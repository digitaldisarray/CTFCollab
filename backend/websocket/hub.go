package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Message struct remains the same
type Message struct {
	Type    string `json:"type"`    // e.g., "ctf_added", "ctf_deleted". "ctf_updated"
	Payload any    `json:"payload"` // The actual data (e.g., the new CTF, the ID of the deleted CTF)
}

// Client represents a single WebSocket connection.
type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan Message // Buffered channel of outbound messages
	roomID string       // ID of the room the client is subscribed to
}

// Hub maintains the set of active clients and broadcasts messages to specific rooms.
type Hub struct {
	// rooms is a map where keys are room IDs and values are maps of clients in that room.
	rooms      map[string]map[*Client]bool
	broadcast  chan broadcastMessage // Inbound messages from the handlers, now includes roomID
	register   chan *Client          // Register requests from clients.
	unregister chan *Client          // Unregister requests from clients.
	mu         sync.Mutex            // Protects the rooms map
}

// broadcastMessage is a helper struct to bundle message with its target room
type broadcastMessage struct {
	roomID  string
	message Message
}

// ClientMessage is what clients send to the server (e.g., to join a room)
type ClientMessage struct {
	Type    string         `json:"type"`
	Payload map[string]any `json:"payload"` // e.g., {"room_id": "your_ctf_room_code"}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:4173" || origin == "http://localhost:3000"
	},
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan broadcastMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[string]map[*Client]bool),
	}
}

// Run starts the hub's event loop.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			// Initial registration; client is not yet in a room.
			// Room assignment will happen when client sends a "join_room" message.
			log.Printf("Client registered (pre-room assignment): %p", client.conn.RemoteAddr())
			// We don't add to any room here yet. Client.readPump will handle room joining.

		case client := <-h.unregister:
			h.mu.Lock()
			if roomClients, ok := h.rooms[client.roomID]; ok {
				if _, clientExists := roomClients[client]; clientExists {
					delete(roomClients, client)
					close(client.send)
					log.Printf("Client %p unregistered from room %s", client.conn.RemoteAddr(), client.roomID)
					if len(roomClients) == 0 {
						delete(h.rooms, client.roomID) // Clean up empty room
						log.Printf("Room %s closed as it's empty", client.roomID)
					}
				}
			}
			h.mu.Unlock()

		case broadcastMsg := <-h.broadcast:
			h.mu.Lock()
			roomClients, roomExists := h.rooms[broadcastMsg.roomID]
			if roomExists {
				log.Printf("Broadcasting message type '%s' to room '%s' (%d clients)",
					broadcastMsg.message.Type, broadcastMsg.roomID, len(roomClients))
				for client := range roomClients {
					select {
					case client.send <- broadcastMsg.message:
						// Message queued successfully
					default:
						log.Printf("Client send channel full in room %s, closing connection: %p",
							client.roomID, client.conn.RemoteAddr())
						close(client.send)
						delete(roomClients, client) // Remove from room
						// If the room becomes empty, you might want to delete it from h.rooms
						if len(roomClients) == 0 {
							delete(h.rooms, client.roomID)
							log.Printf("Room %s closed as it became empty during broadcast", client.roomID)
						}
					}
				}
			} else {
				log.Printf("Attempted to broadcast to non-existent or empty room: %s", broadcastMsg.roomID)
			}
			h.mu.Unlock()
		}
	}
}

// Handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	// Client is created without a roomID initially.
	// It will be set when the client sends a "join_room" message.
	client := &Client{hub: hub, conn: conn, send: make(chan Message, 256), roomID: ""}
	client.hub.register <- client // Basic registration with the hub

	go client.writePump()
	go client.readPump()
}

func (c *Client) joinRoom(roomID string) {
	c.hub.mu.Lock()
	defer c.hub.mu.Unlock()

	// If client was in another room, remove it first (optional, depends on logic)
	if c.roomID != "" && c.roomID != roomID {
		if oldRoomClients, ok := c.hub.rooms[c.roomID]; ok {
			delete(oldRoomClients, c)
			if len(oldRoomClients) == 0 {
				delete(c.hub.rooms, c.roomID)
			}
		}
	}

	c.roomID = roomID
	if _, ok := c.hub.rooms[roomID]; !ok {
		c.hub.rooms[roomID] = make(map[*Client]bool)
		log.Printf("Room %s created", roomID)
	}
	c.hub.rooms[roomID][c] = true
	log.Printf("Client %p joined room %s. Room now has %d clients.", c.conn.RemoteAddr(), roomID, len(c.hub.rooms[roomID]))

}

// readPump pumps messages from the websocket connection to the hub.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
		log.Printf("readPump finished for client: %p, room: %s", c.conn.RemoteAddr(), c.roomID)
	}()

	for {
		var clientMsg ClientMessage
		err := c.conn.ReadJSON(&clientMsg) // Read client's message (e.g., for joining a room)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Read error from client %p (room %s): %v", c.conn.RemoteAddr(), c.roomID, err)
			} else {
				log.Printf("Client %p (room %s) disconnected gracefully or expected error: %v", c.conn.RemoteAddr(), c.roomID, err)
			}
			break
		}

		// Handle client messages, e.g., joining a room
		switch clientMsg.Type {
		case "join_room":
			if roomVal, ok := clientMsg.Payload["room_id"]; ok {
				if roomIDStr, ok := roomVal.(string); ok && roomIDStr != "" {
					c.joinRoom(roomIDStr)
					// Maybe send a confirmation back or current room state
				} else {
					log.Printf("Invalid room_id format in join_room message from %p", c.conn.RemoteAddr())
				}
			} else {
				log.Printf("Missing room_id in join_room message from %p", c.conn.RemoteAddr())
			}
		default:
			log.Printf("Received unhandled message type '%s' from client %p (room %s)", clientMsg.Type, c.conn.RemoteAddr(), c.roomID)
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
		log.Printf("writePump finished for client: %p, room: %s", c.conn.RemoteAddr(), c.roomID)
	}()
	for message := range c.send {
		err := c.conn.WriteJSON(message)
		if err != nil {
			log.Printf("Write error to client %p (room %s): %v", c.conn.RemoteAddr(), c.roomID, err)
			// The hub's unregister logic will handle removing the client on send error (via defer in Run or readPump)
			return
		}
	}
	// If the send channel is closed by the hub, send a close message
	c.conn.WriteMessage(websocket.CloseMessage, []byte{})
}

// BroadcastToRoom sends a message to all clients in a specific room.
func (h *Hub) BroadcastToRoom(roomID string, msg Message) {
	// Ensure roomID is not empty to prevent accidental global broadcasts if logic changes
	if roomID == "" {
		log.Println("Error: BroadcastToRoom called with empty roomID")
		return
	}
	log.Printf("Queueing broadcast for room '%s', message type '%s'", roomID, msg.Type)
	h.broadcast <- broadcastMessage{roomID: roomID, message: msg}
}

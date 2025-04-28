// Create a new file, e.g., websocket/hub.go
package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Message struct defines the structure of messages sent over WebSocket
type Message struct {
	Type    string      `json:"type"`    // e.g., "ctf_added", "ctf_deleted". "ctf_updated"
	Payload interface{} `json:"payload"` // The actual data (e.g., the new CTF, the ID of the deleted CTF)
}

// Client represents a single WebSocket connection.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan Message // Buffered channel of outbound messages
}

// Hub maintains the set of active clients and broadcasts messages.
type Hub struct {
	clients    map[*Client]bool // Registered clients.
	broadcast  chan Message     // Inbound messages from the handlers.
	register   chan *Client     // Register requests from clients.
	unregister chan *Client     // Unregister requests from clients.
	mu         sync.Mutex       // Protects the clients map
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin allows connections from specified origins
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:5173" || origin == "http://localhost:3000" // Allow Svelte dev server
	},
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run starts the hub's event loop.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			log.Printf("Client registered: %p", client.conn.RemoteAddr())
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				log.Printf("Client unregistered: %p", client.conn.RemoteAddr())
			}
			h.mu.Unlock()
		case message := <-h.broadcast:
			h.mu.Lock()
			log.Printf("Broadcasting message: %+v to %d clients", message, len(h.clients))
			for client := range h.clients {
				select {
				case client.send <- message:
					// Message queued successfully
				default:
					log.Printf("Client send channel full, closing connection: %p", client.conn.RemoteAddr())
					close(client.send)
					delete(h.clients, client)
				}
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
	client := &Client{hub: hub, conn: conn, send: make(chan Message, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
		log.Printf("writePump finished for client: %p", c.conn.RemoteAddr())
	}()
	for message := range c.send {
		err := c.conn.WriteJSON(message)
		if err != nil {
			log.Printf("Write error: %v", err)
			return
		}
	}
	// If the send channel is closed by the hub, send a close message
	c.conn.WriteMessage(websocket.CloseMessage, []byte{})
}

// Handles disconnects
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
		log.Printf("readPump finished for client: %p", c.conn.RemoteAddr())
	}()

	for {
		// ReadMessage blocks until a message is received or an error occurs
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Read error: %v", err)
			} else {
				log.Printf("Client disconnected gracefully or expected error: %v", err)
			}
			break
		}
	}
}

// Broadcast sends a message to all connected clients.
func (h *Hub) Broadcast(msg Message) {
	h.broadcast <- msg
}

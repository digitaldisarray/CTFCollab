import { writable } from "svelte/store";
import { challenges, type Challenge } from "../routes/pages/event-room/columns"; // adjust as needed

let manuallyClosed = false;

let ws: WebSocket | null = null;
export let error = writable<string | null>(null);

  export function connectWebSocket(roomcode: string) {
    const wsUrl = "ws://localhost:1337/ws"; // if backend runs elsewhere in the future, change

    console.log(`Attempting to connect WebSocket to: ${wsUrl}`);
    ws = new WebSocket(wsUrl);

    ws.onopen = () => {
      console.log("WebSocket connection established");
      error.set(null);
      if (ws && roomcode) {
        const joinMessage = {
          type: "join_room",
          payload: {
            room_id: roomcode
          }
        };
        ws.send(JSON.stringify(joinMessage));
        console.log(`Sent join_room message for room: ${roomcode}`);
      } else {
        console.error("WebSocket is not open or roomcode is missing, cannot send join_room.")
      }
    };

    ws.onmessage = (event) => {
      try {
        const message = JSON.parse(event.data);
        console.log("WebSocket message received:", message);
        console.log("Received message type:", JSON.stringify(message.type));

        // Handle different message types from the server
        switch (message.type) {
          case 'chal_added':
            const newChallenge = message.payload as Challenge;
            challenges.update((current) => [...current, newChallenge]);
            break;
          case 'chal_deleted':
            const deletedId = message.payload.id.toString();
            challenges.update((current) => current.filter((ch) => ch.id.toString() !== deletedId));
            break;
          case 'chal_flag_updated':
            const { id, flag } = message.payload;
            challenges.update(current =>
                current.map(ch =>
                    ch.id.toString() === id.toString()
                        ? { ...ch, flag, status: flag ? 'complete' : 'pending' }
                        : ch
                )
            );
            break;
          case 'participants_updated':
            console.log("Participants list changed, re-fetching...");
            break;
          // Add cases for challenge updates if needed
          default:
            console.warn("Received unknown WebSocket message type:", message.type);
        }
      } catch (e) {
        console.error("Failed to parse WebSocket message or update store:", e);
      }
    };

    ws.onerror = (event) => {
      console.error("WebSocket error:", event);
      error.set("WebSocket connection error. Trying to reconnect...");
      setTimeout(connectWebSocket, 5000);
    };

    ws.onclose = (event) => {
      console.log("WebSocket connection closed:", event.reason, `Code: ${event.code}`);
      ws = null;
      if (!manuallyClosed && !event.wasClean) {
        error.set("WebSocket connection closed unexpecteldy. Trying to reconnect...");
        setTimeout(connectWebSocket, 5000);
      } else {
        error.set("WebSocket connection closed.");

      }
    };
  }
  export function closeWebSocket() {
    manuallyClosed = true;
    ws?.close();
    ws = null;
}
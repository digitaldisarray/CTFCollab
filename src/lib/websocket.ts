import { challenges, type Challenge } from "../routes/pages/event-room/columns"; // adjust as needed

let ws: WebSocket | null = null;
let manuallyClosed = false;

export function connectWebSocket(backendUrl: string, roomcode: string) {
  const wsUrl = `ws://${backendUrl}/ws`;

  ws = new WebSocket(wsUrl);

  ws.onopen = () => console.log("WebSocket connected");

  ws.onmessage = (event) => {
    try {
      const message = JSON.parse(event.data);

      switch (message.type) {
        case "chal_added":
          challenges.update((current) => [...current, message.payload]);
          break;
        case "chal_deleted":
          challenges.update((current) =>
            current.filter((c) => c.id !== message.payload.id.toString())
          );
          break;
        case "chal_flag_updated":
          const { id, flag } = message.payload;
          challenges.update((current) =>
            current.map((ch) =>
              ch.id === id.toString()
                ? { ...ch, flag, status: flag ? "complete" : "pending" }
                : ch
            )
          );
          break;
        default:
          console.warn("Unknown WebSocket message type:", message.type);
      }
    } catch (err) {
      console.error("WebSocket message parse error:", err);
    }
  };

  ws.onclose = (e) => {
    if (!manuallyClosed) {
      console.warn("WebSocket closed unexpectedly. Reconnecting...");
      setTimeout(() => connectWebSocket(backendUrl, roomcode), 5000);
    }
  };

  ws.onerror = () => {
    console.error("WebSocket error occurred.");
  };
}

export function closeWebSocket() {
  manuallyClosed = true;
  ws?.close();
  ws = null;
}

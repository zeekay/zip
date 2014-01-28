package zip

import (
    "net/http"
    "code.google.com/p/go.net/websocket"
)

type WebSocketHandler func(Conn)

// WebSocket Connection
type Conn struct {
    *websocket.Conn
}

func AddWebSocketHandler(url string, handler WebSocketHandler) {
    http.Handle(url, websocket.Handler(func(ws *websocket.Conn) {
        conn := Conn{ws}
        handler(conn)
    }))

}
func WebSocket(url string, handler WebSocketHandler) {
    AddWebSocketHandler(url, handler)
}

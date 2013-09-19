package client

import (
	"code.google.com/p/go.net/websocket"
)

// Connection - websocket connection
type WebSocketConnection struct {
	*websocket.Conn
}

// Send - write data
func (conn *WebSocketConnection) Send(data interface{}) error {
	return websocket.JSON.Send(conn.Conn, data)
}

// Receive - read data
func (conn *WebSocketConnection) Receive(data interface{}) error {
	return websocket.JSON.Receive(conn.Conn, data)
}

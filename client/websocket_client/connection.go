package websocket_client

import (
	"code.google.com/p/go.net/websocket"
)

// Connection - websocket connection
type Connection struct {
	*websocket.Conn
}

// Send - write data
func (conn *Connection) Send(data interface{}) error {
	return websocket.JSON.Send(conn.Conn, data)
}

// Receive - read data
func (conn *Connection) Receive(data interface{}) error {
	return websocket.JSON.Receive(conn.Conn, data)
}

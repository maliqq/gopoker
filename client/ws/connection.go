package ws

import (
	"code.google.com/p/go.net/websocket"
)

type Connection struct {
	*websocket.Conn
}

func (conn *Connection) Send(data interface{}) error {
	return websocket.JSON.Send(conn.Conn, data)
}

func (conn *Connection) Receive(data interface{}) error {
	return websocket.JSON.Receive(conn.Conn, data)
}

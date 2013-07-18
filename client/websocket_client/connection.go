package websocket_client

import (
	"log"
)

import (
	"code.google.com/p/go.net/websocket"
)

type Connection struct {
	*websocket.Conn
}

func (conn *Connection) Send(data interface{}) error {
	log.Printf("[websocket] sending %+#v", data)
	return websocket.JSON.Send(conn.Conn, data)
}

func (conn *Connection) Receive(data interface{}) error {
	log.Printf("[websocket] receiving")
	return websocket.JSON.Receive(conn.Conn, data)
}

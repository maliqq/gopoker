package ws

import (
	"code.google.com/p/go.net/websocket"
)

import (
	"gopoker/protocol"
)

type Session struct {
	Connection *websocket.Conn
	Receive    chan *protocol.Message
	Send       *chan *protocol.Message
}

func NewSession(connection *websocket.Conn) *Session {
	return &Session{
		Connection: connection,
		Receive:    make(chan *protocol.Message),
	}
}

func (session *Session) Start() {
	go session.receive()
	session.send()
}

func (session *Session) receive() {
	for {
		var message *protocol.Message

		err := websocket.JSON.Receive(session.Connection, message)
		if err != nil {
			break
		}

		*session.Send <- message
	}

	session.Connection.Close()
}

func (session *Session) send() {
	for message := range session.Receive {
		err := websocket.JSON.Send(session.Connection, message)
		if err != nil {
			break
		}
	}

	session.Connection.Close()
}

package client

import (
	"log"
)

import (
	"gopoker/protocol"
)

type Connection interface {
	Close() error
	Receive(interface{}) error
	Send(interface{}) error
}

type Session struct {
	Connection Connection
	Receive    chan *protocol.Message
	Send       *chan *protocol.Message
}

func NewSession(connection Connection) *Session {
	return &Session{
		Connection: connection,
		Receive:    make(chan *protocol.Message),
	}
}

func (session *Session) Start() {
	log.Printf("starting session with connection: %#v", session.Connection)

	go session.receive()

	session.send()
}

func (session *Session) receive() {
	for {
		var message *protocol.Message

		err := session.Connection.Receive(message)
		if err != nil {
			break
		}

		*session.Send <- message
	}

	if err := session.Connection.Close(); err != nil {
		log.Fatalf("Error closing session connection: %s", err)
	}
}

func (session *Session) send() {
	for message := range session.Receive {
		err := session.Connection.Send(message)
		if err != nil {
			break
		}
	}

	session.Connection.Close()
}

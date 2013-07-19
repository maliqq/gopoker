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
	Id string
	Connection Connection
	Receive    protocol.MessageChannel
}

func NewSession(connection Connection) *Session {
	return &Session{
		Connection: connection,
		Receive:    make(chan *protocol.Message),
	}
}

func (session *Session) Start(send *protocol.MessageChannel) {
	log.Printf("starting session with connection: %#v", session.Connection)

	go session.receive(send)

	session.send()
}

func (session *Session) receive(send *protocol.MessageChannel) {
	for {
		var message *protocol.Message

		err := session.Connection.Receive(message)
		if err != nil {
			break
		}

		*send <- message
	}

	session.close()
}

func (session *Session) send() {
	for message := range session.Receive {
		err := session.Connection.Send(message)
		if err != nil {
			break
		}
	}
	session.close()
}

func (session *Session) close() {
	if err := session.Connection.Close(); err != nil {
		log.Fatalf("Error closing session connection: %s", err)
	}
}

package server

import (
	"io"
	"log"
)

import (
	"gopoker/protocol"
	"gopoker/protocol/message"
)

type Connection interface {
	Close() error
	Receive(interface{}) error
	Send(interface{}) error
}

type Session struct {
	Id         string
	Connection Connection
	Recv       protocol.MessageChannel
	Send       *protocol.MessageChannel
}

func NewSession(id string, connection Connection) *Session {
	return &Session{
		Id:         id,
		Recv:       make(protocol.MessageChannel),
		Connection: connection,
	}
}

func (session *Session) Start() {
	log.Printf("[session] start connection: %#v", session.Connection)

	go session.Read()

	session.Write()
}

func (session *Session) Read() {
	for {
		var message message.Message

		err := session.Connection.Receive(&message)
		if err != nil {
			switch err {
			case io.EOF:
				log.Print("[session] EOF")
				break

			default:
				log.Printf("[session] read error: %s", err)
				break
			}

		}

		*session.Send <- &message
	}
}

func (session *Session) Write() {
	for message := range session.Recv {
		err := session.Connection.Send(message)
		if err != nil {
			log.Printf("[session] write error: %s", err)

			break
		}
	}
}

func (session *Session) Close() {
	close(session.Recv)

	if err := session.Connection.Close(); err != nil {
		log.Fatalf("[session] close error: %s", err)
	}
}
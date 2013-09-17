package server

import (
	"io"
	"log"
)

import (
	"gopoker/event"
	"gopoker/model"
)

// Connection - client connection
type Connection interface {
	Close() error
	Receive(interface{}) error
	Send(interface{}) error
}

// Session - client session
type Session struct {
	ID         model.Guid
	Connection Connection
	Recv       event.Channel
	Send       *event.Channel
}

func NewSessionID() model.Guid {
	return model.RandomGuid()
}

// NewSession - create new session
func NewSession(id model.Guid, connection Connection) *Session {
	return &Session{
		ID:         id,
		Recv:       make(event.Channel),
		Connection: connection,
	}
}

// Start - start session
func (session *Session) Start() {
	log.Printf("[session] start connection: %#v", session.Connection)

	go session.read()

	session.write()
}

func (session *Session) read() {
Loop:
	for {
		event := &event.Event{}

		err := session.Connection.Receive(event)
		if err != nil {
			switch err {
			case io.EOF:
				log.Print("[session] EOF")
				break Loop

			default:
				log.Printf("[session] read error: %s", err)
				break Loop
			}

		}

		*session.Send <- event
	}
}

func (session *Session) write() {
	for event := range session.Recv {
		err := session.Connection.Send(event)
		if err != nil {
			log.Printf("[session] write error: %s", err)

			break
		}
	}
}

// Close session
func (session *Session) Close() {
	close(session.Recv)

	if err := session.Connection.Close(); err != nil {
		log.Printf("[session] close error: %s", err)
	}
}

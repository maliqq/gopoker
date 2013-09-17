package server

import (
	"io"
	"log"
)

import (
	"gopoker/event"
	"gopoker/event/message"
	"gopoker/util"
)

// Connection - client connection
type Connection interface {
	Close() error
	Receive(interface{}) error
	Send(interface{}) error
}

// Session - client session
type Session struct {
	ID         string
	Connection Connection
	Recv       event.MessageChannel
	Send       *event.MessageChannel
}

func NewSessionID() string {
	return util.RandomUuid()
}

// NewSession - create new session
func NewSession(id string, connection Connection) *Session {
	return &Session{
		ID:         id,
		Recv:       make(event.MessageChannel),
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
		var message message.Message

		err := session.Connection.Receive(&message)
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

		*session.Send <- &message
	}
}

func (session *Session) write() {
	for message := range session.Recv {
		err := session.Connection.Send(message)
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

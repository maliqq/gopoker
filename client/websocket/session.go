package websocket

import (
  "code.google.com/p/go.net/websocket"
)

type Session struct {
  Connection *websocket.Conn
  Receive chan string
  Send *chan string
}

func NewSession(connection *websocket.Conn) *Session {
  return &Session{
    Connection: connection,
    Receive: make(chan string),
  }
}

func Handler(connection *websocket.Conn) {
  session := NewSession(connection)
  
  go session.receive()
  session.send()
}

func (session *Session) receive() {
  for {
    var message string
  
    err := websocket.Message.Receive(session.Connection, &message)
    if err != nil {
      break
    }
  
    *session.Send <- message
  }
  
  session.Connection.Close()
} 

func (session *Session) send() {
  for message := range session.Receive {
    err := websocket.Message.Send(session.Connection, message)
    if err != nil {
      break
    }
  }

  session.Connection.Close()
}

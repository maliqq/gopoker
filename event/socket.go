package event

import (
  "gopoker/codec"
)

import (
  "code.google.com/p/go.net/websocket"
  zmq "github.com/alecthomas/gozmq"
)

type ZmqSocket struct {
  *zmq.Socket
}

func (s *ZmqSocket) Send(n *Notification) {
  go s.Socket.Send(codec.JSON.Dump(n), 0)
}

type Websocket struct {
  *websocket.Conn
}

func (ws *Websocket) Send(n *Notification) {
  ws.Conn.Write(codec.JSON.Dump(n))
}

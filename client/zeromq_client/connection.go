package zeromq_client

import (
	"log"
	"encoding/json"
)

import (
	zmq "github.com/alecthomas/gozmq"
)

import (
	"gopoker/protocol"
)

type Connection struct {
	Context *zmq.Context
	Socket *zmq.Socket
	Send protocol.MessageChannel
	Recv protocol.MessageChannel
}

func NewConnection(addr string) *Connection {
	context, _ := zmq.NewContext()
  	socket, _ := context.NewSocket(zmq.PAIR)

  	socket.Connect(addr)
  	socket.SetSubscribe("")

	conn := &Connection{
		Context: context,
		Socket: socket,
		Send: make(protocol.MessageChannel),
		Recv: make(protocol.MessageChannel),
	}

	return conn
}

func (conn *Connection) Start() {
	go func() {
		for {
			repl, _ := conn.Socket.Recv(0)
			var msg *protocol.Message
			
			err := json.Unmarshal(repl, &msg)
			if err != nil {
				log.Fatalf("unmarshal error: %s", err)
			}

			conn.Recv <- msg
		}
	}()

	for msg := range conn.Send {
		req, _ := json.Marshal(msg)
		conn.Socket.Send(req, 0)
	}
}

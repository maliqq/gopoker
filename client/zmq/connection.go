package zmq

import (
	"log"
)

import (
	"code.google.com/p/goprotobuf/proto"
	zmq "github.com/alecthomas/gozmq"
)

import (
	"gopoker/protocol"
	"gopoker/protocol/message"
)

// Connection - 0mq connection
type Connection struct {
	Context *zmq.Context
	Socket  *zmq.Socket
	Recv    protocol.MessageChannel
}

// NewConnection - create new connection
func NewConnection(addr string, topic string) *Connection {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)

	socket.Connect(addr)
	socket.SetSubscribe(topic)

	conn := &Connection{
		Context: context,
		Socket:  socket,
		Recv:    make(protocol.MessageChannel),
	}

	return conn
}

// Start - start loop
func (conn *Connection) Start() {
	for {
		repl, _ := conn.Socket.Recv(0)
		var msg *message.Message

		err := proto.Unmarshal(repl, msg)
		if err != nil {
			log.Fatalf("unmarshal error: %s", err)
		}

		conn.Recv <- msg
	}
}

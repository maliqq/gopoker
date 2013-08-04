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
	context *zmq.Context
	socket  *zmq.Socket
	Recv    protocol.MessageChannel
}

// NewConnection - create new connection
func NewConnection(addr string, topic string) *Connection {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)

	if socket.Connect(addr) == nil {
		log.Printf("connected to %s", addr)
	}

	if socket.SetSubscribe(topic) == nil {
		log.Printf("subscribed to %s", topic)
	}

	conn := &Connection{
		context: context,
		socket:  socket,
		Recv:    make(protocol.MessageChannel),
	}

	go conn.receive()

	return conn
}

// Start - start loop
func (conn *Connection) receive() {
	for {
		// receive topic
		topic, _ := conn.socket.Recv(0)
		repl, err := conn.socket.Recv(0)
		if err != nil {
			log.Fatal("receive error:", err)
		}

		log.Printf("received %d bytes for %s", len(repl), topic)

		msg := &message.Message{}
		if err = proto.Unmarshal(repl, msg); err != nil {
			log.Fatalf("unmarshal error: %s", err)
		}

		conn.Recv <- msg
	}
}

package zmq

import (
	"log"
)

import (
	"code.google.com/p/goprotobuf/proto"
	zmq "github.com/alecthomas/gozmq"
)

import (
	"gopoker/exch"
	"gopoker/exch/message"
)

// Connection - 0mq connection
type Connection struct {
	context *zmq.Context
	socket  *zmq.Socket
	Recv    exch.MessageChannel
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
		Recv:    make(exch.MessageChannel),
	}

	go conn.receive()

	return conn
}

// Start - start loop
func (conn *Connection) receive() {
	for {
		parts, err := conn.socket.RecvMultipart(0)
		if err != nil {
			log.Fatal("receive error:", err)
		}

		//topic := parts[0]
		data := parts[1]
		//log.Printf("received %d bytes for %s", len(data), topic)

		msg := &message.Message{}
		if err = proto.Unmarshal(data, msg); err != nil {
			log.Printf("unmarshal error: %s", err)
		} else {
			conn.Recv <- msg
		}
	}
}

func (conn *Connection) Close() {
	conn.context.Close()
	conn.socket.Close()
}

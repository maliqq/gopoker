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
	topic string

	context    *zmq.Context
	sender     *zmq.Socket
	subscriber *zmq.Socket

	Send exch.MessageChannel
	Recv exch.MessageChannel
}

// NewConnection - create new connection
func NewConnection(publisher, receiver string, topic string) *Connection {
	context, _ := zmq.NewContext()

	subscriber, _ := context.NewSocket(zmq.SUB)
	if subscriber.Connect(publisher) == nil {
		log.Printf("connected to %s", publisher)
	}
	if subscriber.SetSubscribe(topic) == nil {
		log.Printf("subscribed to %s", topic)
	}

	sender, _ := context.NewSocket(zmq.PUSH)
	if sender.Bind(receiver) == nil {
		log.Printf("receiver bind success")
	}

	conn := &Connection{
		topic: topic,

		context:    context,
		subscriber: subscriber,
		sender:     sender,

		Send: make(exch.MessageChannel, 100),
		Recv: make(exch.MessageChannel, 100),
	}

	go conn.send()
	go conn.receive()

	return conn
}

func (conn *Connection) send() {
	for {
		select {
		case msg := <-conn.Send:
			data, err := proto.Marshal(msg)
			if err != nil {
				log.Printf("marshal error: %s", err)
			} else {
				conn.sender.SendMultipart([][]byte{[]byte(conn.topic), data}, 0)
			}
		}
	}
}

// Start - start loop
func (conn *Connection) receive() {
	for {
		parts, err := conn.subscriber.RecvMultipart(0)
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
	conn.subscriber.Close()
	conn.sender.Close()
	conn.context.Close()
}

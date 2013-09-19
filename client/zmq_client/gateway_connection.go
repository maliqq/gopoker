package zmq_client

import (
	"log"
)

import (
	zmq "github.com/alecthomas/gozmq"
)

// Connection - 0mq connection
type NodeZMQ struct {
	topic string

	context    *zmq.Context
	sender     *zmq.Socket
	subscriber *zmq.Socket

	Send chan [][]byte
	Recv chan [][]byte
}

// NewConnection - create new connection
func ConnectZmqGateway(publisher, receiver string, topic string) *NodeZMQ {
	context, _ := zmq.NewContext()

	// setup subscriber
	subscriber, _ := context.NewSocket(zmq.SUB)
	if subscriber.Connect(publisher) == nil {
		log.Printf("connected to %s", publisher)
	}
	if subscriber.SetSubscribe(topic) == nil {
		log.Printf("subscribed to %s", topic)
	}

	// setup sender
	sender, _ := context.NewSocket(zmq.PUSH)
	if err := sender.Connect(receiver); err != nil {
		log.Printf("receiver connect error: %s", err)
	}

	conn := &NodeZMQ{
		topic: topic,

		context:    context,
		subscriber: subscriber,
		sender:     sender,

		Send: make(chan [][]byte, 100),
		Recv: make(chan [][]byte, 100),
	}

	go conn.send()
	go conn.receive()

	return conn
}

func (conn *NodeZMQ) send() {
	for {
		select {
		case multipart := <-conn.Send:
			conn.sender.SendMultipart(multipart, zmq.NOBLOCK)
		}
	}
}

// Start - start loop
func (conn *NodeZMQ) receive() {
	for {
		multipart, err := conn.subscriber.RecvMultipart(0)
		if err != nil {
			log.Fatal("receive error:", err)
		}
		conn.Recv <- multipart
	}
}

func (conn *NodeZMQ) Close() {
	conn.subscriber.Close()
	conn.sender.Close()
	conn.context.Close()
}

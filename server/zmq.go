package server

import (
	zmq "github.com/alecthomas/gozmq"
)

type NodeZMQ struct {
	*Node
}

func (n *Node) StartZMQ(addr string) {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	socket.Bind(addr)

	for {
		msg, _ := socket.Recv(0)
		socket.Send(msg, 0)
	}
}

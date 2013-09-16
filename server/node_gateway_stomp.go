package server

import (
	"github.com/jjeffery/stomp"
	stomp_server "github.com/jjeffery/stomp/server"
)

type NodeStomp struct {
	*Node
	server *stomp_server.Server
}

type NodeStompQueue struct {
}

func StartStomp(addr string) *NodeStomp {
	server := stomp_server.Server{
		Addr: addr,
		//QueueStorage: NewNodeStompQueue(),
	}
	go server.ListenAndServe()

	stomp := NodeStomp{
		server: &server,
	}

	return &stomp
}

func NewNodeStompQueue() *NodeStompQueue {
	queue := NodeStompQueue{}
	return &queue
}

func (q *NodeStompQueue) Start() {
}

func (q *NodeStompQueue) Stop() {
}

func (q *NodeStompQueue) Enqueue(queue string, frame *stomp.Frame) error {
	return nil
}

func (q *NodeStompQueue) Requeue(queue string, frame *stomp.Frame) error {
	return nil
}

func (q *NodeStompQueue) Dequeue(queue string) (*stomp.Frame, error) {
	return nil, nil
}

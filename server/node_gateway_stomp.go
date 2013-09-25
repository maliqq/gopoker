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

package server

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type NodeRPC struct {
	Node *Node
}

func (n *Node) StartRPC() {
	nodeRPC := &NodeRPC{n}

	server := rpc.NewServer()
	server.Register(nodeRPC)

	l := listen(n.Rpc.Addr)
	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("[rpc] accept error: %s", c)
			continue
		}

		log.Printf("[rpc] connection started: %v", c.RemoteAddr())
		go server.ServeCodec(jsonrpc.NewServerCodec(c))
	}
}

func listen(addr string) net.Listener {
	log.Printf("[rpc] starting at %s", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("[rpc] listen error:", err)
	}
	return l
}

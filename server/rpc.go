package server

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// NodeRPC - node RPC service
type NodeRPC struct {
	Node *Node
}

// StartRPC - start listening on TCP socket
func (n *Node) StartRPC() {
	nodeRPC := &NodeRPC{n}

	server := rpc.NewServer()
	server.Register(nodeRPC)

	l := listen(n.Config.RPC.Addr)
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

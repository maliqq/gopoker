package server

import (
	"log"
	"net"
	_ "net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type NodeRPC struct {
	Node *Node
}

func (n *Node) StartRPC() {
	nodeRPC := &NodeRPC{n}

	serv := rpc.NewServer()
	serv.Register(nodeRPC)

	log.Printf("[rpc] starting at %s", n.RpcAddr)
	l, err := net.Listen("tcp", n.RpcAddr)
	if err != nil {
		log.Fatal("[rpc] listen error:", err)
	}

	/*
		mux := http.NewServeMux()
		mux.Handle(rpc.DefaultRPCPath, serv)

		http := &http.Server{Handler: mux}
		go http.Serve(l)
	*/

	for {
		c, err := l.Accept()

		if err != nil {
			log.Printf("[rpc] accept error: %s", c)
			continue
		}

		log.Printf("[rpc] connection started: %v", c.RemoteAddr())
		go serv.ServeCodec(jsonrpc.NewServerCodec(c))
	}
}

package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type NodeRPC struct {
	Node *Node
}

func (n *Node) StartRPC() {
	nodeRPC := &NodeRPC{n}

	serv := rpc.NewServer()
	serv.Register(nodeRPC)

	mux := http.NewServeMux()
	mux.Handle(rpc.DefaultRPCPath, serv)

	log.Printf("starting RPC service at %s", n.RpcAddr)
	l, err := net.Listen("tcp", n.RpcAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	http := &http.Server{Handler: mux}
	go http.Serve(l)
}

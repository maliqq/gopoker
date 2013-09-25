package server

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

import (
	"github.com/golang/glog"
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
			glog.Infof("[rpc] accept error: %s", c)
			continue
		}

		glog.Infof("[rpc] connection started: %v", c.RemoteAddr())
		go server.ServeCodec(jsonrpc.NewServerCodec(c))
	}
}

func listen(addr string) net.Listener {
	glog.Infof("[rpc] starting at %s", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		glog.Fatalf("[rpc] listen error:", err)
	}
	return l
}

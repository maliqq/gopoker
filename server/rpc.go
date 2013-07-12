package server

import (
	"net"
	"net/http"
	"net/rpc"
	"log"
)

import (
	"gopoker/server/service"
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

	log.Printf("starting rpc service at %s", n.RpcAddr)
	l, err := net.Listen("tcp", n.RpcAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	http := &http.Server{Handler: mux}
	go http.Serve(l)
}

func (n *NodeRPC) CreateRoom(createRoom *service.CreateRoom, r *service.CallResult) error {
	room := NewRoom(createRoom)
	
	n.Node.AddRoom(room)

	return nil
}

func (n *NodeRPC) DeleteRoom(requestRoom *service.RequestRoom, r *service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)
	
	n.Node.RemoveRoom(room)

	return nil
}

func (n *NodeRPC) StartRoom(requestRoom *service.RequestRoom, r *service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)
	
	room.Start()

	return nil
}

func (n *NodeRPC) PausePlay(requestRoom *service.RequestRoom, r *service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)
	
	room.Pause()

	return nil
}

func (n *NodeRPC) ResumePlay(requestRoom *service.RequestRoom, r *service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)

	room.Resume()

	return nil
}

func (n *NodeRPC) CloseRoom(requestRoom *service.RequestRoom, r *service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)

	room.Close()

	return nil
}

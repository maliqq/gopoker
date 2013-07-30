package server

import (
	"log"
)

import (
	"gopoker/server/rpc_service"
	"net/http"
)

func (n *NodeRPC) CreateRoom(createRoom *rpc_service.CreateRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] CreateRoom args=%+v", createRoom)

	room := NewRoom(createRoom)

	// add to node
	n.Node.AddRoom(room)

	return nil
}

func (n NodeRPC) CreateRoomHTTP(req *http.Request, createRoom *rpc_service.CreateRoom, r *rpc_service.CallResult) error {
	return n.CreateRoom(createRoom, r)
}

func (n *NodeRPC) DeleteRoom(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] DeleteRoom args=%+v", requestRoom)

	room := n.Node.Room(requestRoom.Id)
	n.Node.RemoveRoom(room)

	return nil
}

// send protocol message to room subscribers
func (n *NodeRPC) NotifyRoom(notifyRoom *rpc_service.NotifyRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] NotifyRoom args=%+v", notifyRoom)

	room := n.Node.Room(notifyRoom.Id)
	room.Recv <- notifyRoom.Message

	return nil
}

func (n *NodeRPC) StartRoom(startRoom *rpc_service.StartRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] StartRoom args=%+v", startRoom)

	room := n.Node.Room(startRoom.Id)
	room.Start()

	return nil
}

func (n *NodeRPC) PausePlay(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] PausePlayargs=%+v", requestRoom)

	room := n.Node.Room(requestRoom.Id)
	room.Pause()

	return nil
}

func (n *NodeRPC) ResumePlay(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] ResumePlay args=%+v", requestRoom)

	room := n.Node.Room(requestRoom.Id)
	room.Resume()

	return nil
}

func (n *NodeRPC) CloseRoom(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)
	log.Printf("[rpc] CloseRoom args=%+v", requestRoom)

	room.Close()

	return nil
}

package server

import (
	"gopoker/server/rpc_service"
	"net/http"
)

func (n *NodeRPC) CreateRoom(createRoom *rpc_service.CreateRoom, r *rpc_service.CallResult) error {
	room := NewRoom(createRoom)

	// add to node
	n.Node.AddRoom(room)

	return nil
}

func (n NodeRPC) CreateRoomHTTP(req *http.Request, createRoom *rpc_service.CreateRoom, r *rpc_service.CallResult) error {
	return n.CreateRoom(createRoom, r)
}

func (n *NodeRPC) DeleteRoom(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)
	n.Node.RemoveRoom(room)

	return nil
}

// send protocol message to room subscribers
func (n *NodeRPC) NotifyRoom(notifyRoom *rpc_service.NotifyRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(notifyRoom.Id)
	room.Recv <- notifyRoom.Message

	return nil
}

func (n *NodeRPC) StartRoom(startRoom *rpc_service.StartRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(startRoom.Id)
	room.Start()

	return nil
}

func (n *NodeRPC) PausePlay(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)
	room.Pause()

	return nil
}

func (n *NodeRPC) ResumePlay(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)
	room.Resume()

	return nil
}

func (n *NodeRPC) CloseRoom(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)
	room.Close()

	return nil
}

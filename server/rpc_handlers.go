package server

import (
	"gopoker/server/rpc_service"
)

func (n *NodeRPC) CreateRoom(createRoom *rpc_service.CreateRoom, r *rpc_service.CallResult) error {
	room := NewRoom(createRoom)

	n.Node.AddRoom(room)

	return nil
}

func (n *NodeRPC) DeleteRoom(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(requestRoom.Id)

	n.Node.RemoveRoom(room)

	return nil
}

// send protocol message to room subscribers
func (n *NodeRPC) NotifyRoom(notifyRoom *rpc_service.NotifyRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(notifyRoom.Id)
	msg := notifyRoom.Message

	room.Receive <- msg

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

package server

import (
	"gopoker/server/service"
)

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

// send protocol message to room subscribers
func (n *NodeRPC) NotifyRoom(notifyRoom *service.NotifyRoom, r *service.CallResult) error {
	room := n.Node.Room(notifyRoom.Id)
	msg := notifyRoom.Message

	room.Receive <- msg

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

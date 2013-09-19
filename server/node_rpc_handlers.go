package server

import (
	"log"
)

import (
	rpc_service "gopoker/server/noderpc"
	"net/http"
)

// CreateRoom - create room
func (n *NodeRPC) CreateRoom(createRoom *rpc_service.CreateRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] CreateRoom args=%+v", createRoom)

	room := NewRoom(createRoom)

	// add to node
	n.Node.AddRoom(room)

	return nil
}

// CreateRoomHTTP - create room via HTTP
func (n NodeRPC) CreateRoomHTTP(req *http.Request, createRoom *rpc_service.CreateRoom, r *rpc_service.CallResult) error {
	return n.CreateRoom(createRoom, r)
}

// DeleteRoom - delete room
func (n *NodeRPC) DeleteRoom(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] DeleteRoom args=%+v", requestRoom)

	room := n.Node.Room(requestRoom.Guid)
	n.Node.RemoveRoom(room)

	return nil
}

// StartRoom - start room
func (n *NodeRPC) StartRoom(startRoom *rpc_service.StartRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] StartRoom args=%+v", startRoom)

	room := n.Node.Room(startRoom.Guid)
	room.Start()

	return nil
}

// PausePlay - pause play
func (n *NodeRPC) PausePlay(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] PausePlayargs=%+v", requestRoom)

	room := n.Node.Room(requestRoom.Guid)
	room.Pause()

	return nil
}

// ResumePlay - resume play
func (n *NodeRPC) ResumePlay(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] ResumePlay args=%+v", requestRoom)

	room := n.Node.Room(requestRoom.Guid)
	room.Resume()

	return nil
}

// CloseRoom - close room
func (n *NodeRPC) CloseRoom(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(requestRoom.Guid)
	log.Printf("[rpc] CloseRoom args=%+v", requestRoom)

	room.Close()

	return nil
}

func (n *NodeRPC) Login(login *rpc_service.Login, r *rpc_service.LoginResult) error {
	r.Session, r.Success = n.Node.Login(login.Username, login.Password)
	log.Printf("[prc] Login username=%s", login.Username)

	return nil
}

func (n *NodeRPC) Logout(logout *rpc_service.Logout, r *rpc_service.CallResult) error {
	n.Node.Logout(logout.Session)
	log.Printf("[rpc] Logout args=%+v", logout)

	return nil
}

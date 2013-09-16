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

	room := n.Node.Room(requestRoom.ID)
	n.Node.RemoveRoom(room)

	return nil
}

// NotifyRoom -send protocol message to room subscribers
func (n *NodeRPC) NotifyRoom(notifyRoom *rpc_service.NotifyRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] NotifyRoom args=%+v", notifyRoom)

	room := n.Node.Room(notifyRoom.ID)
	room.Recv <- notifyRoom.Message

	return nil
}

// StartRoom - start room
func (n *NodeRPC) StartRoom(startRoom *rpc_service.StartRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] StartRoom args=%+v", startRoom)

	room := n.Node.Room(startRoom.ID)
	room.Start()

	return nil
}

// PausePlay - pause play
func (n *NodeRPC) PausePlay(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] PausePlayargs=%+v", requestRoom)

	room := n.Node.Room(requestRoom.ID)
	room.Pause()

	return nil
}

// ResumePlay - resume play
func (n *NodeRPC) ResumePlay(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] ResumePlay args=%+v", requestRoom)

	room := n.Node.Room(requestRoom.ID)
	room.Resume()

	return nil
}

// CloseRoom - close room
func (n *NodeRPC) CloseRoom(requestRoom *rpc_service.RequestRoom, r *rpc_service.CallResult) error {
	room := n.Node.Room(requestRoom.ID)
	log.Printf("[rpc] CloseRoom args=%+v", requestRoom)

	room.Close()

	return nil
}

func (n *NodeRPC) Login(login *rpc_service.Login, r *rpc_service.LoginResult) error {
	r.SessionID, r.Success = n.Node.Login(login.Username, login.Password)
	log.Printf("[prc] Login username=%s", login.Username)

	return nil
}

func (n *NodeRPC) Logout(logout *rpc_service.Logout, r *rpc_service.CallResult) error {
	n.Node.Logout(logout.SessionID)
	log.Printf("[rpc] Logout args=%+v", logout)

	return nil
}

// ConnectGateway - connect gateway
func (n *NodeRPC) ConnectGateway(connectReq *rpc_service.ConnectGateway, r *rpc_service.CallResult) error {
	n.Node.ZMQGateway.connect <- *connectReq

	return nil
}

// DisconnectGateway - disconnect gateway
func (n *NodeRPC) DisconnectGateway(disconnectReq *rpc_service.DisconnectGateway, r *rpc_service.CallResult) error {
	n.Node.ZMQGateway.disconnect <- *disconnectReq

	return nil
}

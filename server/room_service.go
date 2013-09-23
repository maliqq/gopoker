package server

type RoomService struct {
	*Room
}

// CreateRoom - create room
func (n *NodeRPC) CreateRoom(createRoom *rpc_service.CreateRoom, r *rpc_service.CallResult) error {
	log.Printf("[rpc] CreateRoom args=%+v", createRoom)

	room := NewRoom(createRoom)

	// add to node
	n.Node.AddRoom(room)

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

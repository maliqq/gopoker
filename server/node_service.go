package server

func (n *Node) CreateRoom() {
	var variation model.Variation
	if createRoom.Mix != nil {
		variation = createRoom.Mix
	} else {
		variation = createRoom.Game
	}
}

func (n *Node) DeleteRoom() {
	room := n.Node.Room(requestRoom.Guid)
	n.Node.RemoveRoom(room)
}

func (n *Node) Login(username, password string) LoginResult {
	n.Node.Login(username, password)
}

func (n *Node) Logout(session model.Guid) {
	n.Node.Logout(session)
}

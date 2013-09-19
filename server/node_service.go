package server

type NodeService struct {
  Node *Node
}

func (n *NodeService) CreateRoom() {
  var variation model.Variation
  if createRoom.Mix != nil {
    variation = createRoom.Mix
  } else {
    variation = createRoom.Game
  }
}

func (n *NodeService) DeleteRoom() {
  room := n.Node.Room(requestRoom.Guid)
  n.Node.RemoveRoom(room)
}

func (service *NodeService) NotifyRoom() {

}

func (n *NodeService) Login(username, password string) LoginResult {
  n.Node.Login(username, password)
}

func (n *NodeRPC) Logout(session model.Guid) {
  n.Node.Logout(session)
}

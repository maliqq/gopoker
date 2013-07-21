package server

type Node struct {
	Name    string
	ApiAddr string
	//apiService
	RpcAddr string
	//rpcService
	Rooms map[string]*Room
}

func NewNode(name string, apiAddr string, rpcAddr string) *Node {
	return &Node{
		Name:    name,
		ApiAddr: apiAddr,
		RpcAddr: rpcAddr,
		Rooms:   map[string]*Room{},
	}
}

func (n *Node) Room(id string) *Room {
	room, _ := n.Rooms[id]

	return room
}

func (n *Node) AddRoom(room *Room) bool {
	n.Rooms[room.Id] = room
	return true
}

func (n *Node) RemoveRoom(room *Room) bool {
	delete(n.Rooms, room.Id)
	return true
}

func (n *Node) Start() {
	go n.StartRPC()
	n.StartHTTP()
}

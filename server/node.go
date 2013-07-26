package server

type Config struct {
	Logdir  string
	ApiAddr string
	RpcAddr string
}

type Node struct {
	Name string

	*Config
	Rooms map[string]*Room
}

func NewNode(name string, config *Config) *Node {
	return &Node{
		Name:   name,
		Config: config,
		Rooms:  map[string]*Room{},
	}
}

func (n *Node) Room(id string) *Room {
	room, _ := n.Rooms[id]

	return room
}

func (n *Node) AddRoom(room *Room) bool {
	n.Rooms[room.Id] = room

	room.createLogger(n.Logdir)

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

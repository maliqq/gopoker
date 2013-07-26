package server

import (
	"time"
)

type HttpConfig struct {
	Addr          string
	ApiPath       string
	RpcPath       string
	WebSocketPath string
}

type RpcConfig struct {
	Addr    string
	Timeout time.Duration
}

type Config struct {
	Logdir string
	Http   *HttpConfig
	Rpc    *RpcConfig
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

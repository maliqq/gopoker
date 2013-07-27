package server

import (
	"time"
	"log"
)

import (
	"gopoker/storage"
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
	Store *storage.StoreConfig
}

type Node struct {
	Name string

	*Config
	Rooms map[string]*Room
	Store *storage.Store
}

func NewNode(name string, config *Config) *Node {
	node := &Node{
		Name:   name,
		Config: config,
		Rooms:  map[string]*Room{},
	}

	node.connectStore()

	return node
}

func (n *Node) connectStore() {
	var err error
	n.Store, err = storage.Open(n.Config.Store)

	if err != nil {
		log.Fatal("Can't open store", err)
	}
	log.Print("[store] connected")
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

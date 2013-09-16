package server

import (
	"log"
)

import (
	"gopoker/storage"
)

// Node - node
type Node struct {
	Name string
	Config     *Config

	Rooms      map[string]*Room
	
	ZMQGateway *NodeZMQ
	
	Store      *storage.Store
	PlayStore  *storage.PlayStore
	SessionStore *storage.SessionStore
}

// NewNode - create new node
func NewNode(name string, config *Config) *Node {
	node := &Node{
		Name:   name,
		Config: config,
		Rooms:  map[string]*Room{},
	}

	if config.Store != nil {
		node.connectStore()
	}
	if config.PlayStore != nil {
		node.connectPlayStore()
	}
	if config.SessionStore != nil {
		node.connectSessionStore()
	}

	return node
}

// Start - start room
func (n *Node) Start() {
	go n.StartRPC()
	go n.StartZMQ()
	n.StartHTTP()
}

func (n *Node) connectStore() {
	var err error
	n.Store, err = storage.OpenStore(n.Config.Store)

	if err != nil {
		log.Fatalf("[store] can't open %s (%s): %s", n.Config.Store.Driver, n.Config.Store.ConnectionString, err)
	}
	log.Printf("[store] connected to %s", n.Config.Store)
}

func (n *Node) connectPlayStore() {
	var err error
	n.PlayStore, err = storage.OpenPlayStore(n.Config.PlayStore)
	if err != nil {
		log.Fatalf("[store] can't open %s: %s", n.Config.PlayStore.Host, err)
	}
	log.Printf("[store] connected to %s", n.Config.PlayStore)
}

func (n *Node) connectSessionStore() {
	var err error
	n.SessionStore, err = storage.OpenSessionStore(n.Config.SessionStore)
	if err != nil {
		log.Fatalf("[store] can't open %#v: %s", n.Config.SessionStore, err)
	}
	log.Printf("[store] connected to %#v", n.Config.SessionStore)
}

// Room - get room
func (n *Node) Room(id string) *Room {
	room, _ := n.Rooms[id]

	return room
}

// AddRoom - add room
func (n *Node) AddRoom(room *Room) bool {
	n.Rooms[room.ID] = room

	if n.Config.Logdir != "" {
		room.createLogger(n.Config.Logdir)
	}
	if n.PlayStore != nil {
		room.createStorage(n.PlayStore)
	}

	return true
}

// RemoveRoom - remove room
func (n *Node) RemoveRoom(room *Room) bool {
	delete(n.Rooms, room.ID)
	return true
}

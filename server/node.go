package server

import (
	"log"
)

import (
	"gopoker/model"
	"gopoker/storage"
)

// Node - node
type Node struct {
	Name   string
	Config *Config

	Rooms map[model.Guid]*Room

	ZMQGateway *NodeZMQ

	Store        *storage.Store
	PlayHistory  *storage.PlayHistory
	SessionStore *storage.SessionStore
}

// NewNode - create new node
func NewNode(name string, config *Config) *Node {
	node := &Node{
		Name:   name,
		Config: config,
		Rooms:  map[model.Guid]*Room{},
	}

	if config.Store != nil {
		node.connectStore()
	}
	if config.PlayHistory != nil {
		node.connectPlayHistory()
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

func (n *Node) connectPlayHistory() {
	var err error
	n.PlayHistory, err = storage.OpenPlayHistory(n.Config.PlayHistory)
	if err != nil {
		log.Fatalf("[store] can't open %s: %s", n.Config.PlayHistory.Host, err)
	}
	log.Printf("[store] connected to %s", n.Config.PlayHistory)
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
func (n *Node) Room(guid model.Guid) *Room {
	room, _ := n.Rooms[guid]

	return room
}

// AddRoom - add room
func (n *Node) AddRoom(room *Room) bool {
	n.Rooms[room.Guid] = room

	if n.Config.Logdir != "" {
		room.createLogger(n.Config.Logdir)
	}
	if n.PlayHistory != nil {
		room.createStorage(n.PlayHistory)
	}

	return true
}

// RemoveRoom - remove room
func (n *Node) RemoveRoom(room *Room) bool {
	delete(n.Rooms, room.Guid)
	return true
}

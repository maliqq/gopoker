package server

import (
	"log"
	"time"
)

import (
	"gopoker/storage"
)

// HTTPConfig - http config
type HTTPConfig struct {
	Addr          string
	APIPath       string
	RPCPath       string
	WebSocketPath string
}

// RPCConfig - RPC config
type RPCConfig struct {
	Addr    string
	Timeout time.Duration
}

// Config - node config
type Config struct {
	Logdir    string
	HTTP      *HTTPConfig
	RPC       *RPCConfig
	ZMQ       string
	Store     *storage.StoreConfig
	PlayStore *storage.PlayStoreConfig
}

// Node - node
type Node struct {
	Name string

	Config     *Config
	Rooms      map[string]*Room
	ZMQGateway *NodeZMQ
	Store      *storage.Store
	PlayStore  *storage.PlayStore
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

	return node
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
	log.Print("[store] connected to %s", n.Config.PlayStore)
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

// Start - start room
func (n *Node) Start() {
	go n.StartRPC()
	go n.StartZMQ()
	n.StartHTTP()
}

// APIPathOr - API path or default
func (httpConfig *HTTPConfig) APIPathOr(defaultPath string) string {
	apiPath := httpConfig.APIPath
	if apiPath == "" {
		return defaultPath
	}
	return apiPath
}

// WebSocketPathOr - websocket path or default
func (httpConfig *HTTPConfig) WebSocketPathOr(defaultPath string) string {
	webSocketPath := httpConfig.WebSocketPath
	if webSocketPath == "" {
		return defaultPath
	}
	return webSocketPath
}

// RPCPathOr - RPC path or default
func (httpConfig *HTTPConfig) RPCPathOr(defaultPath string) string {
	rpcPath := httpConfig.RPCPath
	if rpcPath == "" {
		return defaultPath
	}
	return rpcPath
}

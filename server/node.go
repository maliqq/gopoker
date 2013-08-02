package server

import (
	"log"
	"time"
)

import (
	"gopoker/model"
	"gopoker/storage"
)

// HTTPConfig - http config
type HTTPConfig struct {
	Addr          string
	APIPath       string
	RPCPath       string
	WebSocketPath string
	ZmqAddr       string
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
	Store     *storage.StoreConfig
	PlayStore *storage.PlayStoreConfig
}

// Node - node
type Node struct {
	Name string

	*Config
	Rooms     map[string]*Room
	Store     *storage.Store
	PlayStore *storage.PlayStore
}

// Defaults
const (
	NodeConfigFile = "node.json"
)

// NewNode - create new node
func NewNode(name string, configDir string) *Node {
	var config *Config
	model.ReadConfig(configDir, NodeConfigFile, &config)

	node := &Node{
		Name:   name,
		Config: config,
		Rooms:  map[string]*Room{},
	}

	node.connectStore()
	node.connectPlayStore()

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
	n.Rooms[room.Id] = room

	room.createLogger(n.Logdir)
	room.createStorage(n.PlayStore)

	return true
}

// RemoveRoom - remove room
func (n *Node) RemoveRoom(room *Room) bool {
	delete(n.Rooms, room.Id)
	return true
}

// Start - start room
func (n *Node) Start() {
	go n.StartRPC()
	n.StartHTTP()
}

func (httpConfig *HttpConfig) apiPathOr(defaultPath string) string {
	apiPath := httpConfig.ApiPath
	if apiPath == "" {
		return defaultPath
	}
	return apiPath
}

func (httpConfig *HttpConfig) webSocketPathOr(defaultPath string) string {
	webSocketPath := httpConfig.WebSocketPath
	if webSocketPath == "" {
		return defaultPath
	}
	return webSocketPath
}

func (httpConfig *HttpConfig) rpcPathOr(defaultPath string) string {
	rpcPath := httpConfig.RpcPath
	if rpcPath == "" {
		return defaultPath
	}
	return rpcPath
}

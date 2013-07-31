package server

import (
	"log"
	"time"
)

import (
	"gopoker/model"
	"gopoker/storage"
)

type HttpConfig struct {
	Addr          string
	ApiPath       string
	RpcPath       string
	WebSocketPath string
	ZmqAddr       string
}

type RpcConfig struct {
	Addr    string
	Timeout time.Duration
}

// NodeConfig
type Config struct {
	Logdir    string
	Http      *HttpConfig
	Rpc       *RpcConfig
	Store     *storage.StoreConfig
	PlayStore *storage.PlayStoreConfig
}

type Node struct {
	Name string

	*Config
	Rooms     map[string]*Room
	Store     *storage.Store
	PlayStore *storage.PlayStore
}

const (
	NodeConfigFile = "node.json"
)

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

func (n *Node) Room(id string) *Room {
	room, _ := n.Rooms[id]

	return room
}

func (n *Node) AddRoom(room *Room) bool {
	n.Rooms[room.Id] = room

	room.createLogger(n.Logdir)
	room.createStorage(n.PlayStore)

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

func (httpConfig *HttpConfig) ApiPathOr(defaultPath string) string {
	apiPath := httpConfig.ApiPath
	if apiPath == "" {
		return defaultPath
	}
	return apiPath
}

func (httpConfig *HttpConfig) WebSocketPathOr(defaultPath string) string {
	webSocketPath := httpConfig.WebSocketPath
	if webSocketPath == "" {
		return defaultPath
	}
	return webSocketPath
}

func (httpConfig *HttpConfig) RpcPathOr(defaultPath string) string {
	rpcPath := httpConfig.RpcPath
	if rpcPath == "" {
		return defaultPath
	}
	return rpcPath
}

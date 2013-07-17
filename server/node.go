package server

import (
	"gopoker/client"
	"gopoker/model"
)

const (
	rpcRoot      = "/_rpc"
	rpcDebugRoot = "/_rpc/debug"
)

type Node struct {
	Name    string
	ApiAddr string
	//apiService
	RpcAddr string
	//rpcService
	Rooms    map[model.Id]*Room
	Sessions map[model.Id]*client.Session
}

func NewNode(name string, apiAddr string, rpcAddr string) *Node {
	return &Node{
		Name:    name,
		ApiAddr: apiAddr,
		RpcAddr: rpcAddr,
		Rooms:   map[model.Id]*Room{},
	}
}

func (n *Node) Room(id model.Id) *Room {
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
	n.StartRPC()
	n.StartHTTP()
}

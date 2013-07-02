package server

import (
	"gopoker/play"
	"net/http"
)

type CallResult struct {
	Success bool
}

type RoomParams struct {
	Id   string
	Size int
	Game *play.Game
}

type RoomId struct {
	Id string
}

func (n *Node) CreateRoom(req *http.Request, params *RoomParams, r *CallResult) error {
	room := new(Room)
	room.Id = params.Id
	room.Table = play.CreateTable(params.Size)

	n.Rooms[params.Id] = room
	defer func() { go room.Start() }()

	return nil
}

func (n *Node) CloseRoom(req *http.Request, params *RoomId, r *CallResult) error {
	room := n.Rooms[params.Id]

	defer room.Close()

	return nil
}

func (n *Node) PauseRoom(req *http.Request, params *RoomId, r *CallResult) error {
	room := n.Rooms[params.Id]

	defer room.Pause()

	return nil
}

func (n *Node) StartRoom(req *http.Request, params *RoomId, r *CallResult) error {
	room := n.Rooms[params.Id]

	defer func() { go room.Start() }()

	return nil
}

func (n *Node) DestroyRoom(req *http.Request, params *RoomId, r *CallResult) error {
	delete(n.Rooms, params.Id)
	return nil
}

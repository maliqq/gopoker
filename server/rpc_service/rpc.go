package rpc_service

import (
	"gopoker/model"
	"gopoker/play/mode"
	"gopoker/protocol/message"
)

type CallResult struct {
	Status  string
	Message string
}

type CreateRoom struct {
	Id        string
	TableSize int
	BetSize   float64
	Game      *model.Game
	Mix       *model.Mix
}

type RequestRoom struct {
	Id string
}

type StartRoom struct {
	Id   string
	Mode mode.Type
}

type NotifyRoom struct {
	Id      string
	Message *message.Message
}

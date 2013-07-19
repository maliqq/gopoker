package service

import (
	"gopoker/model"
	"gopoker/protocol"
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

type NotifyRoom struct {
	Id      string
	Message *protocol.Message
}

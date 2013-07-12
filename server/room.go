package server

import (
	"gopoker/model"
	"gopoker/play/context"
	"gopoker/server/service"
)

type Room struct {
	Id model.Id
	*context.Play
}

func NewRoom(createRoom *service.CreateRoom) *Room {
	table := model.NewTable(createRoom.Size)
	stake := model.NewStake(createRoom.BetSize)
	newPlay := context.NewPlay(createRoom.Variation(), stake, table)

	return &Room{
		Id: createRoom.Id,
		Play: newPlay,
	}
}

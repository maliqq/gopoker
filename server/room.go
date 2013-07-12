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
	newTable := model.NewTable(createRoom.Size)
	newPlay := context.NewPlay(createRoom.Variation(), newTable)

	return &Room{
		Id: createRoom.Id,
		Play: newPlay,
	}
}

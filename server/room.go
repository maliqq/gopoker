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
	variation := createRoom.Variation()
	tableSize := createRoom.TableSize

	var maxTableSize int
	if variation.IsMixed() {
		maxTableSize = model.MixedGameMaxTableSize
	} else {
		maxTableSize = variation.(*model.Game).MaxTableSize
	}
	if tableSize == 0 || tableSize > maxTableSize {
		tableSize = maxTableSize
	}

	table := model.NewTable(tableSize)
	stake := model.NewStake(createRoom.BetSize)
	newPlay := context.NewPlay(variation, stake, table)

	return &Room{
		Id:   createRoom.Id,
		Play: newPlay,
	}
}

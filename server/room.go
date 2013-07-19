package server

import (
	"gopoker/model"
	"gopoker/play"
	"gopoker/server/rpc_service"
)

type Room struct {
	Id string
	*play.Play
}

func NewRoom(createRoom *rpc_service.CreateRoom) *Room {
	var variation model.Variation
	if createRoom.Game != nil {
		variation = createRoom.Game.WithDefaults()
	}

	if createRoom.Mix != nil {
		variation = createRoom.Mix.WithDefaults()
	}

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
	newPlay := play.NewPlay(variation, stake, table)

	return &Room{
		Id:   createRoom.Id,
		Play: newPlay,
	}
}

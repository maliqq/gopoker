package service

import (
	"gopoker/model"
)

type CallResult struct {
	Status string
	Message string
}

type CreateRoom struct {
	Id   model.Id
	Size int
	BetSize float64
	Game *model.Game
	Mix *model.Mix
}

func (c CreateRoom) Variation() model.Variation {
	if c.Game != nil {
		return c.Game
	}
	
	if c.Mix != nil {
		return c.Mix
	}

	return nil
}

type RequestRoom struct {
	Id model.Id
}

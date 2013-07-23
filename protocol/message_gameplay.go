package protocol

import (
	"gopoker/model"
	"gopoker/model/game"
)

// error
type Error struct {
	Message string
}

func NewError(err error) *Message {
	return NewMessage(
		Error{
			Message: err.Error(),
		},
	)
}

type CollectPot struct {
	Amount float64
}

func NewCollectPot(pot *model.Pot) *Message {
	return NewMessage(
		CollectPot{
			Amount: pot.Total(),
		},
	)
}

type Chat struct {
	Pos     int
	Message string
}

// deal info
type Deal struct {
}

type ChangeDealState struct {
	State string
}

type ChangeGame struct {
	Type game.LimitedGame
	game.Limit
}

func NewChangeGame(g *model.Game) *Message {
	return NewMessage(
		ChangeGame{
			Type:  g.Type,
			Limit: g.Limit,
		},
	)
}

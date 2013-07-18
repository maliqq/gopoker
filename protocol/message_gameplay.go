package protocol

import (
	"gopoker/model"
	"gopoker/model/game"
)

// error
type Error struct {
	Description string
}

func NewError(err error) *Message {
	return NewMessage(
		Error{
			Description: err.Error(),
		},
	)
}

type PotSummary struct {
	Amount float64
}

func NewPotSummary(pot *model.Pot) *Message {
	return NewMessage(
		PotSummary{
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

// new street
type Street struct {
}

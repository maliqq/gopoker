package protocol

import (
	"gopoker/model"
	"gopoker/model/game"
)

type PlayStart struct {
}

type StreetStart struct {
	Name string
	Pot  float64
	Rake float64
}

type ChangeGame struct {
	Type game.LimitedGame
	game.Limit
}

func NewStreetStart(name string, pot *model.Pot) *Message {
	return NewMessage(StreetStart{
		Name: name,
		Pot:  pot.Total(),
	})
}

func NewPlayStart() *Message {
	return NewMessage(PlayStart{})
}

func NewChangeGame(g *model.Game) *Message {
	return NewMessage(ChangeGame{
		Type:  g.Type,
		Limit: g.Limit,
	})
}

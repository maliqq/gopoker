package protocol

import (
	"gopoker/model"
	"gopoker/model/game"
)

type PlayStart struct {
}

type StreetStart struct {
	Name string
}

type ChangeGame struct {
	Type game.LimitedGame
	game.Limit
}

func NewPlayStart() *Message {
	return NewMessage(PlayStart{})
}

func NewStreetStart(name string) *Message {
	return NewMessage(StreetStart{
		Name: name,
	})
}

func NewChangeGame(g *model.Game) *Message {
	return NewMessage(ChangeGame{
		Type:  g.Type,
		Limit: g.Limit,
	})
}

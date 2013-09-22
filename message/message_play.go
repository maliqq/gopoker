package message

import (
	"gopoker/model"
)

// PlayStart
type PlayStart struct {
	Game  *model.Game
	Stake *model.Stake
	Table *model.Table
}

func (PlayStart) EventMessage() {}

// StreetStart
type StreetStart struct {
	Name string
}

func (StreetStart) EventMessage() {}

// PlayStop
type PlayStop struct {
}

func (PlayStop) EventMessage() {}

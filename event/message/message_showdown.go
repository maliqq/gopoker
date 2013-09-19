package message

import (
	"gopoker/model"
	"gopoker/poker"
)

// ShowHand
type ShowHand struct {
	Pos      int
	Player   model.Player
	Cards    poker.Cards
	Hand     poker.Hand
	HandName string
}

func (ShowHand) EventMessage() {}

// ShowCards
type ShowCards struct {
	Pos    int
	Muck   bool
	Cards  poker.Cards
	Player model.Player
}

func (ShowCards) EventMessage() {}

// Winner
type Winner struct {
	Pos    int
	Player model.Player
	Amount float64
}

func (Winner) EventMessage() {}

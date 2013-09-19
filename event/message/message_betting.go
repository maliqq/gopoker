package message

import (
	"gopoker/model"
	"gopoker/model/bet"
)

// AddBet
type AddBet struct {
	Pos int
	Bet *model.Bet
}

func (AddBet) EventMessage() {}

// RequireBet
type RequireBet struct {
	Pos   int
	Range *bet.Range
}

func (RequireBet) EventMessage() {}

// BettingComplete
type BettingComplete struct {
	Pot float64
}

func (BettingComplete) EventMessage() {}

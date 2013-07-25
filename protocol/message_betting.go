package protocol

import (
	"fmt"
)

import (
	"gopoker/model"
)

type RequireBet struct {
	Pos int
	model.BetRange
}

func (r RequireBet) String() string {
	return fmt.Sprintf("call: %.2f min: %.2f max: %.2f", r.Call, r.Min, r.Max)
}

type AddBet struct {
	Pos int
	Bet model.Bet
	BetString string
}

func NewAddBet(pos int, bet *model.Bet) *Message {
	return NewMessage(AddBet{
		Pos: pos,
		Bet: *bet,
		BetString: bet.PrintString(),
	})
}

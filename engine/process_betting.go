package engine

import (
	"gopoker/engine/context"
	"gopoker/model"
)

type BettingProcess struct {
	Betting *context.Betting
	Recv    chan *model.Bet
}

func NewBettingProcess(g *Gameplay) *BettingProcess {
	g.b = context.NewBetting()
	p := &BettingProcess{
		Betting: g.b,
		Recv:    make(chan *model.Bet),
	}

	return p
}

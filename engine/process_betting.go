package engine

import (
	"log"
	"time"
)

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
		Recv: make(chan *model.Bet),
	}

	return p
}

func (g *Gameplay) bettingRound(exit chan bool) {
BettingRound:
	for {
		done := make(chan bool)
		timeout := time.After(100 * time.Second)

		if !g.requireBetting(done) {
			log.Printf("[betting] done")
			break BettingRound
		}

		select {
		case <-done:
			log.Printf("[betting] none waiting")
			exit <- true
			break BettingRound

		case <-timeout:
			log.Printf("[betting] timeout")
			// process timeout

		case b := <-g.BettingProcess.Recv:
			log.Printf("[betting] got %s", b)
			g.Betting().AddBet(b)
		}
	}
}

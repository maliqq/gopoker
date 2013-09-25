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

func (g *Gameplay) processBetting(exit chan bool) {
BettingRound:
	for {
		done := make(chan bool)
		timeout := time.After(100 * time.Second)

		go g.requireBetting(done)
		
		doExit, doBreak := <-done
		if doExit {
			log.Printf("[betting] none waiting")
			exit <- true
		}
		if doBreak {
			log.Printf("[betting] done")
			break BettingRound
		}

		select {
		case <-timeout:
			log.Printf("[betting] timeout")
			// process timeout
		case b := <-g.BettingProcess.Recv:
			g.Betting().AddBet(b)
		}
	}

	g.completeBetting()
}

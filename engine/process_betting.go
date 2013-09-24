package engine

import (
	"gopoker/engine/context"
	"gopoker/model"
)

type BettingProcess struct {
	Betting *context.Betting

	Bet  chan *model.Bet
	Next chan bool

	stop chan bool
}

func NewBettingProcess(g *Gameplay) *BettingProcess {
	g.b = context.NewBetting()
	p := &BettingProcess{
		Betting: g.b,
	    Bet: make(chan *model.Bet),
	    Next: make(chan bool),
	    stop: make(chan bool),
	  }

  return p
}

func (p *BettingProcess) Start() {
	go p.receive()
}

func (p *BettingProcess) Stop() {
	p.stop <- true
}

func (p *BettingProcess) receive() {
BettingLoop:
	for {
		select {
		case bet := <-p.Bet:
			p.Betting.AddBet(bet)
			p.Next <- true

		case <-p.stop:
			break BettingLoop
		}
	}

	p.Betting.Clear()
}

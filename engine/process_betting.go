package engine

import (
	"gopoker/engine/context"
	"gopoker/model"
)

type BettingProcess struct {
	*context.Betting

	Bet  chan *model.Bet
	Next chan bool

	stop chan bool
}

/*
func NewBetting() *Betting {
  process := Betting{
    Betting: context.NewBetting(),
    Bet: make(chan *model.Bet),
    Next: make(chan int),
    stop: make(chan int),
  }

  process.receive()

  return process
}
*/

func (process *BettingProcess) Start() {
	process.Betting = context.NewBetting()
}

func (process *BettingProcess) Stop() {
	process.stop <- true
}

func (process *BettingProcess) receive() {
BettingLoop:
	for {
		select {
		case bet := <-process.Bet:
			process.Betting.AddBet(bet)
			process.Next <- true

		case <-process.stop:
			break BettingLoop
		}
	}

	process.Betting.Clear()
}

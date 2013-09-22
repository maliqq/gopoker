package process

import (
  "gopoker/engine/context"
  "gopoker/model"
)

type Betting struct {
  *context.Betting
  
  Bet chan *model.Bet
  
  next chan int
  stop chan int
}

func (process *Betting) Start() {
  process.Betting = context.NewBetting()
}

func (process *Betting) Stop() {}

func (process *Betting) receive() {
BettingLoop:
  for {
    select {
    case bet := <-process.Bet:
      process.Betting.AddBet(bet)
    
    case <-process.stop:
      break BettingLoop
    }
  }

  process.Betting.Clear()
}

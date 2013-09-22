package process

import (
  "gopoker/engine/context"
)

type Deal struct {
  *context.Deal
  stop chan int
}

func (process *Deal) Start() {
  process.Deal = context.NewDeal()
}

func (process *Deal) Stop() {
  process.stop <- 1
}

func (process *Deal) receive() {
DealLoop:
  for {
    select {
    case <- process.stop:
      break DealLoop
    }
  }
}

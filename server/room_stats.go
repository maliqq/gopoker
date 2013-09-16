package server

import (
  "github.com/rcrowley/go-metrics"
)

import (
  "time"
)

type RoomStatsWorker struct {
  *Room
  NextDeal chan int
  deals *metrics.StandardMeter
  waiting *metrics.StandardGauge
}

func (w *RoomStatsWorker) Start() {
  w.deals = metrics.NewMeter()
  metrics.Register("deals", w.deals)

  w.waiting = metrics.NewGauge()
  metrics.Register("waiting", w.waiting)

  for {
    select {
    case <- time.After(1 * time.Second):
      // do something

    case <- w.NextDeal:
      w.deals.Mark(1)
    }
  }
}

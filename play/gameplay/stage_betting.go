package gameplay

import (
  "time"
  "fmt"
)

import (
  "gopoker/protocol"
)

const (
  DefaultTimer = 30
)

func (this *GamePlay) StartBettingRound() {
  betting := this.Betting

  for _, pos := range this.Table.Seats.From(betting.Current()).InPlay() {
    seat := this.Table.Seat(pos)

    this.Broadcast.One(seat.Player) <- betting.RequireBet(pos, seat, this.Game, this.Stake)

    select {
    case msg := <-this.Betting.Receive:
      betting.Add(seat, msg)

    case <-time.After(time.Duration(DefaultTimer) * time.Second):
      fmt.Println("timeout!")
    }
  }
}

func (this *GamePlay) ResetBetting() {
  this.Betting.Clear()

  for _, pos := range this.Table.SeatsInPlay() {
    seat := this.Table.Seat(pos)
    seat.Play()
  }

  this.Broadcast.All <- protocol.NewPotSummary(this.Pot)
}

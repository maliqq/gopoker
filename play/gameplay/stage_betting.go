package gameplay

import (
	"gopoker/play/command"
	"gopoker/protocol"
)

const (
	DefaultTimer = 30
)

func (this *GamePlay) NextTurn(current int) bool {
	active := this.Table.Seats.From(current).Playing()
	inPot := this.Table.Seats.From(current).InPot()

	if len(inPot) < 2 {
		this.Control <- command.Showdown
	
	} else if len(active) > 0 {
		pos := active[0]
		seat := this.Table.Seat(pos)
		this.Broadcast.One(seat.Player) <- this.Betting.RequireBet(pos, seat, this.Game, this.Stake)
		return true
	}

	return false
}

func (this *GamePlay) StartBettingRound() {
	pos := make(chan int)
	defer close(pos)

	go this.Betting.Start(&pos)

	for current := range pos {
		if this.NextTurn(current) {
			// ...
		} else {
			this.Betting.Stop()
			break
		}
	}

	this.ResetBetting()
}

func (this *GamePlay) ResetBetting() {
	this.Betting.Clear()

	for _, pos := range this.Table.AllSeats().InPlay() {
		seat := this.Table.Seat(pos)
		seat.Play()
	}

	this.Broadcast.All <- protocol.NewPotSummary(this.Pot)
}

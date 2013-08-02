package gameplay

import (
	seatState "gopoker/model/seat"
	"gopoker/protocol/message"
)

const (
	DefaultTimer = 30
)

func (this *GamePlay) StartBettingRound() Transition {
	//this.Broadcast.All <- message.NewBettingStart(this.Betting)
	nextPos := make(chan int)
	defer close(nextPos)

	go this.Betting.Start(&nextPos)

	var next Transition
	for current := range nextPos {
		for _, pos := range this.Table.AllSeats().InPlay() {
			seat := this.Table.Seat(pos)
			if !seat.Calls(this.Betting.BetRange.Call) {
				seat.State = seatState.Play
			}
		}
		active := this.Table.Seats.From(current).Playing()
		inPot := this.Table.Seats.From(current).InPot()

		if len(inPot) < 2 {
			next = Stop
			break
		} else if len(active) > 0 {
			pos := active[0]
			seat := this.Table.Seat(pos)

			this.Broadcast.One(seat.Player) <- this.Betting.RequireBet(pos, seat.Stack, this.Game, this.Stake)

			continue
		}

		next = Next
		break
	}

	this.Betting.Stop()
	this.ResetBetting()

	return next
}

func (this *GamePlay) ResetBetting() {
	this.Betting.Clear(this.Table.Button)

	for _, pos := range this.Table.AllSeats().InPlay() {
		seat := this.Table.Seat(pos)
		seat.Play()
	}

	total := this.Betting.Pot.Total()
	this.Broadcast.All <- message.NewBettingComplete(total)
}

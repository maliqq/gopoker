package gameplay

import (
	seatState "gopoker/model/seat"
	"gopoker/protocol/message"
)

const (
	// DefaultTimer - default timeout for action
	DefaultTimer = 30
)

// StartBettingRound - start betting round
func (gp *GamePlay) StartBettingRound() Transition {
	//gp.Broadcast.All <- message.NotifyBettingStart(gp.Betting)
	go gp.Betting.Start()

	var next Transition
Loop:
	for {
		select {
		case <-gp.Betting.Next:
			for _, pos := range gp.Table.AllSeats().InPlay() {
				seat := gp.Table.Seat(pos)
				if !seat.Calls(gp.Betting.BetRange.Call) {
					seat.State = seatState.Play
				}
			}

			current := gp.Betting.Pos
			active := gp.Table.Seats.From(current).Playing()
			inPot := gp.Table.Seats.From(current).InPot()

			if len(inPot) < 2 {
				next = Stop
				break Loop
			} else if len(active) > 0 {
				pos := active[0]
				seat := gp.Table.Seat(pos)

				gp.Broadcast.One(seat.Player) <- gp.Betting.RequireBet(pos, seat, gp.Game.Limit, gp.Stake)

				continue Loop
			}
			next = Next
			break Loop
		}
	}

	gp.Betting.Stop()
	gp.ResetBetting()

	return next
}

// ResetBetting - complete betting round
func (gp *GamePlay) ResetBetting() {
	gp.Betting.Clear(gp.Table.Button)

	for _, pos := range gp.Table.AllSeats().InPlay() {
		seat := gp.Table.Seat(pos)
		seat.Play()
	}

	total := gp.Betting.Pot.Total()
	gp.Broadcast.All <- message.NotifyBettingComplete(total)
}

package engine

import (
	"gopoker/message"
	seatState "gopoker/model/seat"
)

const (
	// DefaultTimer - default timeout for action
	DefaultTimer = 30
)

func (g *Gameplay) processBetting(skip chan bool) {
	ring := g.Table.Ring()

	for _, box := range ring.InPlay() {
		if !box.Seat.Calls(g.b.BetRange.Call) {
			box.Seat.State = seatState.Play
		}
	}

	active := ring.Playing()
	inPot := ring.InPot()

	if len(inPot) < 2 || len(active) == 0 {
		skip <- true
		return
	}
	defer close(skip)

	g.b.NewRound(active)

	g.e.Notify(
		g.b.RequireBet(g.Game.Limit, g.Stake),
	).One(g.b.Round.Current().Player)
}

func (g *Gameplay) completeBetting() {
	g.b.Clear()
	ring := g.Table.Ring()

	for _, box := range ring.InPlay() {
		box.Seat.Play()
	}

	total := g.b.Pot.Total()
	g.e.Notify(
		&message.BettingComplete{total},
	).All()
}

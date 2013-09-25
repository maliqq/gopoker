package engine

import (
	"gopoker/message"
	seatState "gopoker/model/seat"
)

const (
	// DefaultTimer - default timeout for action
	DefaultTimer = 30
)

/*
Continue betting normally:
	false, false <- exit

Done betting round:
	false, true <- exit

Go to showdown:
	true, false <- exit
*/
func (g *Gameplay) requireBetting(exit chan bool) {
	ring := g.Table.Ring()

	for _, box := range ring.InPlay() {
		if !box.Seat.Calls(g.b.BetRange.Call) {
			box.Seat.State = seatState.Play
		}
	}

	inPot := ring.InPot()
	if len(inPot) < 2 {
		exit <- true
		return
	}

	active := ring.Playing()
	if len(active) == 0 {
		exit <- false
		return
	}

	defer close(exit)

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

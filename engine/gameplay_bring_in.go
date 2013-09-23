package engine

import (
	"gopoker/model"
	"gopoker/poker"
)

func (g *Gameplay) bringIn() {
	var min model.Box
	var card *poker.Card

	ring := g.Table.Ring()

	for i, box := range ring.Active() {
		pocketCards := g.d.Pocket(box.Seat.Player)

		lastCard := pocketCards[len(pocketCards)-1]
		if i == 0 {
			card = lastCard
		} else {
			if lastCard.Compare(card, poker.AceHigh) > 0 {
				card = lastCard
				min = box
			}
		}
	}

	g.setButton(min.Pos)

	g.b.NewRound(ring.Active())

	g.e.Notify(
		g.b.RequireBet(g.Game.Limit, g.Stake),
	).One(min.Seat.Player)
}

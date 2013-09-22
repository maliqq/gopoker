package engine

import (
	"gopoker/message"
	"gopoker/model/deal"
)

func (g *Gameplay) dealHole(cardsNum int) {
	if cardsNum == 0 {
		cardsNum = g.Game.PocketSize
	}

	ring := g.Table.Ring()

	for _, box := range ring.InPlay() {
		player := box.Seat.Player

		cards := g.d.DealPocket(player, cardsNum)

		g.e.Notify(
			&message.DealCards{box.Pos, cards, deal.Hole},
		).One(player)
	}
}

func (g *Gameplay) dealDoor(cardsNum int) {
	ring := g.Table.Ring()

	for _, box := range ring.InPlay() {
		player := box.Seat.Player

		cards := g.d.DealPocket(player, cardsNum)

		g.e.Notify(
			&message.DealCards{box.Pos, cards, deal.Door},
		).All()
	}
}

func (g *Gameplay) dealBoard(cardsNum int) {
	cards := g.d.DealBoard(cardsNum)

	g.e.Notify(
		&message.DealCards{0, cards, deal.Board},
	).All()
}

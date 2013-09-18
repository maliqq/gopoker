package gameplay

import (
	"gopoker/event/message"
	"gopoker/model/deal"
)

// DealHole - deal hole cards
func (gp *GamePlay) DealHole(cardsNum int) {
	for _, pos := range gp.Table.AllSeats().InPlay() {
		player := gp.Table.Player(pos)

		cards := gp.Deal.DealPocket(player, cardsNum)

		gp.Broadcast.Notify(
			&message.DealCards{pos, cards, deal.Hole},
		).One(player)
	}
}

// DealDoor - deal door cards
func (gp *GamePlay) DealDoor(cardsNum int) {
	for _, pos := range gp.Table.AllSeats().InPlay() {
		player := gp.Table.Player(pos)

		cards := gp.Deal.DealPocket(player, cardsNum)

		gp.Broadcast.Notify(
			&message.DealCards{pos, cards, deal.Door},
		).All()
	}
}

// DealBoard - deal board cards
func (gp *GamePlay) DealBoard(cardsNum int) {
	cards := gp.Deal.DealBoard(cardsNum)

	gp.Broadcast.Notify(
		&message.DealCards{0, cards, deal.Board},
	).All()
}

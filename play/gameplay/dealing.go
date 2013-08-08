package gameplay

import (
	"gopoker/exch/message"
	"gopoker/model/deal"
)

// DealHole - deal hole cards
func (gp *GamePlay) DealHole(cardsNum int) {
	for _, pos := range gp.Table.AllSeats().InPlay() {
		player := gp.Table.Player(pos)

		cards := gp.Deal.DealPocket(player, cardsNum)

		gp.Broadcast.One(player) <- message.NotifyDealPocket(pos, cards.Proto(), deal.Hole)
	}
}

// DealDoor - deal door cards
func (gp *GamePlay) DealDoor(cardsNum int) {
	for _, pos := range gp.Table.AllSeats().InPlay() {
		player := gp.Table.Player(pos)

		cards := gp.Deal.DealPocket(player, cardsNum)

		gp.Broadcast.All <- message.NotifyDealPocket(pos, cards.Proto(), deal.Door)
	}
}

// DealBoard - deal board cards
func (gp *GamePlay) DealBoard(cardsNum int) {
	cards := gp.Deal.DealBoard(cardsNum)

	gp.Broadcast.All <- message.NotifyDealShared(cards.Proto(), deal.Board)
}

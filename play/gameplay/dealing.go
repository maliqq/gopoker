package gameplay

import (
	"gopoker/model/deal"
	"gopoker/protocol/message"
)

func (gp *GamePlay) DealHole(cardsNum int) {
	for _, pos := range gp.Table.AllSeats().InPlay() {
		player := gp.Table.Player(pos)

		cards := gp.Deal.DealPocket(player, cardsNum)

		gp.Broadcast.One(player) <- message.NewDealPocket(pos, cards.Proto(), deal.Hole)
	}
}

func (gp *GamePlay) DealDoor(cardsNum int) {
	for _, pos := range gp.Table.AllSeats().InPlay() {
		player := gp.Table.Player(pos)

		cards := gp.Deal.DealPocket(player, cardsNum)

		gp.Broadcast.All <- message.NewDealPocket(pos, cards.Proto(), deal.Door)
	}
}

func (gp *GamePlay) DealBoard(cardsNum int) {
	cards := gp.Deal.DealBoard(cardsNum)

	gp.Broadcast.All <- message.NewDealShared(cards.Proto(), deal.Board)
}

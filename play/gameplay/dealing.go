package gameplay

import (
	"gopoker/model/deal"
	"gopoker/protocol/message"
)

func (this *GamePlay) DealHole(cardsNum int) {
	for _, pos := range this.Table.AllSeats().InPlay() {
		player := this.Table.Player(pos)

		cards := this.Deal.DealPocket(player, cardsNum)

		this.Broadcast.One(player) <- message.NewDealPocket(pos, cards, deal.Hole)
	}
}

func (this *GamePlay) DealDoor(cardsNum int) {
	for _, pos := range this.Table.AllSeats().InPlay() {
		player := this.Table.Player(pos)

		cards := this.Deal.DealPocket(player, cardsNum)

		this.Broadcast.All <- message.NewDealPocket(pos, cards, deal.Door)
	}
}

func (this *GamePlay) DealBoard(cardsNum int) {
	cards := this.Deal.DealBoard(cardsNum)

	this.Broadcast.All <- message.NewDealShared(cards, deal.Board)
}

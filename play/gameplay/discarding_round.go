package gameplay

import (
	"gopoker/model"
	"gopoker/model/deal"
	"gopoker/poker"
	"gopoker/protocol/message"
)

func (this *GamePlay) StartDiscardingRound() Transition {
	discarding := this.Discarding

	for _, pos := range this.Table.AllSeats().InPlay() {
		seat := this.Table.Seat(pos)

		this.Broadcast.One(seat.Player) <- discarding.RequireDiscard(pos, seat)
	}

	return Next
}

func (this *GamePlay) discard(p model.Player, cards poker.Cards) {
	pos, _ := this.Table.Pos(p)

	cardsNum := len(cards)

	this.Broadcast.All <- message.NewDiscarded(pos, cardsNum)

	if cardsNum > 0 {
		newCards := this.Deal.Discard(p, cards)

		this.Broadcast.One(p) <- message.NewDealPocket(pos, newCards, deal.Discard)
	}
}

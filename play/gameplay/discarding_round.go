package gameplay

import (
	"gopoker/model"
	"gopoker/model/deal"
	"gopoker/poker"
	"gopoker/protocol/message"
)

func (gp *GamePlay) StartDiscardingRound() Transition {
	discarding := gp.Discarding

	for _, pos := range gp.Table.AllSeats().InPlay() {
		seat := gp.Table.Seat(pos)

		gp.Broadcast.One(seat.Player) <- discarding.RequireDiscard(pos, seat)
	}

	return Next
}

func (gp *GamePlay) discard(p model.Player, cards poker.Cards) {
	pos, _ := gp.Table.Pos(p)

	cardsNum := len(cards)

	gp.Broadcast.All <- message.NewDiscarded(pos, cardsNum)

	if cardsNum > 0 {
		newCards := gp.Deal.Discard(p, cards)

		gp.Broadcast.One(p) <- message.NewDealPocket(pos, newCards.Proto(), deal.Discard)
	}
}

package gameplay

import (
	"gopoker/event/message"
	"gopoker/model"
	"gopoker/model/deal"
	"gopoker/poker"
)

// StartDiscardingRound - start discarding round
func (gp *GamePlay) StartDiscardingRound() Transition {
	discarding := gp.Discarding

	for _, pos := range gp.Table.AllSeats().InPlay() {
		seat := gp.Table.Seat(pos)

		gp.Broadcast.Notify(
			discarding.RequireDiscard(pos, seat),
		).One(seat.Player)
	}

	return Next
}

func (gp *GamePlay) discard(p model.Player, cards poker.Cards) {
	pos, _ := gp.Table.Pos(p)

	cardsNum := len(cards)

	gp.Broadcast.Notify(
		message.Discarded{pos, cardsNum},
	).All()

	if cardsNum > 0 {
		newCards := gp.Deal.Discard(p, cards)

		gp.Broadcast.Notify(
			message.DealCards{pos, newCards, deal.Discard},
		).One(p)
	}
}

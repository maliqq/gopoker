package gameplay

import (
	"gopoker/poker"
	"gopoker/protocol"
)

func (this *GamePlay) BringIn() Transition {
	minPos := 0
	var card poker.Card

	for _, pos := range this.Table.AllSeats().Active() {
		s := this.Table.Seat(pos)

		pocketCards := this.Deal.Pocket(s.Player)

		lastCard := pocketCards[len(pocketCards)-1]
		if pos == 0 {
			card = lastCard
		} else {
			if lastCard.Compare(card, poker.AceHigh) > 0 {
				card = lastCard
				minPos = pos
			}
		}
	}

	this.Table.SetButton(minPos)
	this.Broadcast.All <- protocol.NewMoveButton(minPos)

	seat := this.Table.Seat(minPos)
	this.Broadcast.One(seat.Player) <- this.Betting.RequireBet(minPos, seat, this.Game, this.Stake)

	return Next
}

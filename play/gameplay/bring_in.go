package gameplay

import (
	"gopoker/event/message"
	"gopoker/poker"
)

// BringIn - post bring in
func (gp *Gameplay) BringIn() Transition {
	minPos := 0
	var card *poker.Card

	for i, pos := range gp.Table.AllSeats().Active() {
		s := gp.Table.Seat(pos)

		pocketCards := gp.Deal.Pocket(s.Player)

		lastCard := pocketCards[len(pocketCards)-1]
		if i == 0 {
			card = lastCard
		} else {
			if lastCard.Compare(card, poker.AceHigh) > 0 {
				card = lastCard
				minPos = pos
			}
		}
	}

	gp.Table.SetButton(minPos)
	gp.Broadcast.Notify(
		&message.MoveButton{minPos},
	).All()

	seat := gp.Table.Seat(minPos)
	gp.Broadcast.Notify(
		gp.Betting.RequireBet(minPos, seat, gp.Game.Limit, gp.Stake),
	).One(seat.Player)

	return Next
}

package stage

import (
	"gopoker/poker"
)

func (stage *Stage) BringIn() {
	play := stage.Play

	minPos := 0
	var card poker.Card

	for _, pos := range play.Table.SeatsInPlay() {
		s := play.Table.Seat(pos)

		pocketCards := *play.Deal.Pocket(s.Player)

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

	stage.setButton(minPos)

	seat := play.Table.Seat(minPos)

	play.Broadcast.One(seat.Player) <- stage.Betting.RequireBet(minPos, seat, play.Game)
}

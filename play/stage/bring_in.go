package stage

import (
	"log"
)

import (
	"gopoker/play/context"
	"gopoker/poker"
)

func BringIn(play *context.Play) {
	log.Println("[play.stage] bring in")

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

	setButton(play, minPos)

	seat := play.Table.Seat(minPos)

	play.Broadcast.One(seat.Player) <- play.Betting.RequireBet(minPos, seat, play.Game)
}

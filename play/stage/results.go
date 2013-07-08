package stage

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/protocol"
	"gopoker/play/context"
)

func best(sidePot *model.SidePot, hands *showdownHands) (model.Id, *poker.Hand) {
	var winner model.Id
	var best *poker.Hand

	for member, _ := range sidePot.Members {
		hand, hasHand := (*hands)[member]

		if hasHand && (best == nil || hand.Compare(best) > 0) {
			winner = member
			best = hand
		}
	}

	return winner, best
}

func results(play *context.Play, highHands *showdownHands, lowHands *showdownHands) {
	pot := play.Betting.Pot

	hi := highHands != nil
	lo := lowHands != nil
	split := hi && lo

	for _, sidePot := range pot.SidePots() {
		total := sidePot.Total()

		var winnerLow, winnerHigh model.Id
		var bestLow *poker.Hand

		if lo {
			winnerLow, bestLow = best(sidePot, lowHands)
		}

		if hi {
			winnerHigh, _ = best(sidePot, highHands)
		}

		if split && bestLow != nil {
			play.Broadcast.All <- protocol.NewWinner(winnerLow, total/2.)
			play.Broadcast.All <- protocol.NewWinner(winnerHigh, total/2.)
		} else {
			var exclusiveWinner model.Id

			if hi {
				exclusiveWinner = winnerHigh
			} else {
				exclusiveWinner = winnerLow
			}

			play.Broadcast.All <- protocol.NewWinner(exclusiveWinner, total)
		}
	}
}

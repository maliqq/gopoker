package gameplay

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/ranking"
	"gopoker/protocol"
)

type ShowdownHands map[model.Player]*poker.Hand

func (this *GamePlay) ShowHands(ranking ranking.Ranking, withBoard bool) ShowdownHands {
	d := this.Deal

	hands := ShowdownHands{}

	for _, pos := range this.Table.AllSeats().InPot() {
		player := this.Table.Player(pos)
		if pocket, hand := d.Rank(player, ranking, withBoard); hand != nil {
			hands[player] = hand

			this.Broadcast.All <- protocol.NewShowHand(pos, pocket, hand)
		}
	}

	return hands
}

func best(sidePot *model.SidePot, hands ShowdownHands) (model.Player, *poker.Hand) {
	var winner model.Player
	var best *poker.Hand

	for member, _ := range sidePot.Members {
		hand, hasHand := hands[member]

		if hasHand && (best == nil || hand.Compare(best) > 0) {
			winner = member
			best = hand
		}
	}

	return winner, best
}

func (this *GamePlay) Winners(highHands ShowdownHands, lowHands ShowdownHands) {
	hi := highHands != nil
	lo := lowHands != nil

	split := hi && lo

	for _, sidePot := range this.Betting.Pot.SidePots() {
		total := sidePot.Total()

		var winnerLow, winnerHigh model.Player
		var bestLow *poker.Hand

		if lo {
			winnerLow, bestLow = best(sidePot, lowHands)
		}

		if hi {
			winnerHigh, _ = best(sidePot, highHands)
		}

		winners := map[model.Player]float64{}

		if split && bestLow != nil {
			winners[winnerLow] = total / 2.
			winners[winnerHigh] = total / 2.
		} else {
			if hi {
				winners[winnerHigh] = total
			} else {
				winners[winnerLow] = total
			}
		}

		for winner, amount := range winners {
			pos, _ := this.Table.Pos(winner)
			this.Broadcast.All <- protocol.NewWinner(pos, amount)
		}
	}
}

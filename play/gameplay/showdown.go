package gameplay

import (
	"gopoker/event/message"
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/hand"
)

// ShowdownHands - map player and hand
type ShowdownHands map[model.Player]*poker.Hand

// ShowHands - show hands for players in pot
func (gp *Gameplay) ShowHands(ranking hand.Ranking, withBoard bool) ShowdownHands {
	d := gp.Deal

	hands := ShowdownHands{}

	for _, pos := range gp.Table.AllSeats().InPot() {
		player := gp.Table.Player(pos)

		if pocket, hand := d.Rank(player, ranking, withBoard); hand != nil {
			hands[player] = hand

			gp.Broadcast.Notify(
				&message.ShowHand{pos, player, pocket, *hand, hand.PrintString()},
			).All()
		}
	}

	return hands
}

func best(sidePot *model.SidePot, hands ShowdownHands) (model.Player, *poker.Hand) {
	var winner model.Player
	var best *poker.Hand

	for member := range sidePot.Members {
		hand, hasHand := hands[member]

		if hasHand && (best == nil || hand.Compare(best) > 0) {
			winner = member
			best = hand
		}
	}

	return winner, best
}

// Winners - show pot winners
func (gp *Gameplay) Winners(highHands ShowdownHands, lowHands ShowdownHands) {
	hi := highHands != nil
	lo := lowHands != nil

	split := hi && lo

	for _, sidePot := range gp.Betting.Pot.SidePots() {
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
			pos, _ := gp.Table.Pos(winner)
			seat := gp.Table.Seat(pos)
			seat.AdvanceStack(amount)
			gp.Broadcast.Notify(
				&message.Winner{pos, winner, amount},
			).All()
		}
	}
}

// Winner - single winner
func (gp *Gameplay) Winner(pos int) {
	for _, sidePot := range gp.Betting.Pot.SidePots() {
		amount := sidePot.Total()

		seat := gp.Table.Seat(pos)
		seat.AdvanceStack(amount)

		winner := seat.Player
		gp.Broadcast.Notify(
			&message.Winner{pos, winner, amount},
		).All()
	}
}

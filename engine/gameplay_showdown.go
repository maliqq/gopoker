package engine

import (
	"gopoker/message"
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/hand"
)

// showdownHands - map player and hand
type showdownHands map[model.Player]*poker.Hand

// ShowHands - show hands for players in pot
func (g *Gameplay) showHands(ranking hand.Ranking) showdownHands {
	ring := g.Table.Ring()
	hands := showdownHands{}

	for _, box := range ring.InPot() {
		player := box.Seat.Player

		if pocket, hand := g.d.Rank(player, ranking); hand != nil {
			hands[player] = hand

			g.e.Notify(
				&message.ShowHand{box.Pos, player, pocket, *hand, hand.PrintString()},
			).All()
		}
	}

	return hands
}

func best(sidePot *model.SidePot, hands showdownHands) (model.Player, *poker.Hand) {
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
func (g *Gameplay) declareWinners(highHands showdownHands, lowHands showdownHands) {
	hi := highHands != nil
	lo := lowHands != nil

	split := hi && lo

	for _, sidePot := range g.b.Pot.SidePots() {
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
			pos, _ := g.Table.Pos(winner)
			seat := g.Table.Seat(pos)
			seat.AdvanceStack(amount)
			g.e.Notify(
				&message.Winner{pos, winner, amount},
			).All()
		}
	}
}

// Winner - single winner
func (g *Gameplay) declareWinner(pos int) {
	for _, sidePot := range g.b.Pot.SidePots() {
		amount := sidePot.Total()

		seat := g.Table.Seat(pos)
		seat.AdvanceStack(amount)

		winner := seat.Player
		g.e.Notify(
			&message.Winner{pos, winner, amount},
		).All()
	}
}

func (g *Gameplay) showdown() {
	ring := g.Table.Ring()
	inPot := ring.InPot()

	if len(inPot) == 1 {
		// last player left
		g.declareWinner(inPot[0].Pos)

	} else {

		var highHands, lowHands showdownHands

		if g.Game.Lo != "" {
			lowHands = g.showHands(g.Game.Lo)//, g.Game.HasBoard)
		}

		if g.Game.Hi != "" {
			highHands = g.showHands(g.Game.Hi)//, g.Game.HasBoard)
		}

		g.declareWinners(highHands, lowHands)
	}
}

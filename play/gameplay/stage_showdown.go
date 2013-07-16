package gameplay

import (
  "gopoker/poker"
  "gopoker/poker/ranking"
  "gopoker/model"
  "gopoker/protocol"
)

type ShowdownHands map[model.Id]*poker.Hand

func (this *GamePlay) ShowHands(ranking ranking.Ranking, withBoard bool) *ShowdownHands {
  d := this.Deal

  hands := ShowdownHands{}

  for _, pos := range this.Table.SeatsInPot() {
    player := this.Table.Player(pos)
    if pocket, hand := d.Rank(player, ranking, withBoard); hand != nil {
      hands[player.Id] = hand

      this.Broadcast.All <- protocol.NewShowHand(pos, pocket, hand)
    }
  }

  return &hands
}

func best(sidePot *model.SidePot, hands *ShowdownHands) (model.Id, *poker.Hand) {
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

func (this *GamePlay) Winners(highHands *ShowdownHands, lowHands *ShowdownHands) {
  pot := this.Betting.Pot

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
      this.Broadcast.All <- protocol.NewWinner(winnerLow, total/2.)
      this.Broadcast.All <- protocol.NewWinner(winnerHigh, total/2.)
    } else {
      var exclusiveWinner model.Id

      if hi {
        exclusiveWinner = winnerHigh
      } else {
        exclusiveWinner = winnerLow
      }

      this.Broadcast.All <- protocol.NewWinner(exclusiveWinner, total)
    }
  }
}

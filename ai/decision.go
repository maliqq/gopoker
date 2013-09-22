package ai

import (
	"log"
)

import (
	"gopoker/poker/math"
	"gopoker/poker"
)

type decision struct {
	minBet      float64
	maxBet      float64
	raiseChance float64
	allInChance float64
}

func (b *Bot) decidePreflop(cards poker.Cards) decision {
	group := math.SklanskyMalmuthGroup(cards[0], cards[1])

	log.Printf("group=%d", group)

	bb := b.stake.BigBlindAmount()
	switch group {
	case 9:
		return decision{maxBet: 0.}

	case 7, 8:
		return decision{maxBet: bb}

	case 5, 6:
		return decision{
			minBet:      bb,
			maxBet:      bb * 4,
			raiseChance: 0.2,
		}

	case 3, 4:
		return decision{
			maxBet:      b.stack + b.bet,
			raiseChance: 0.5,
			allInChance: 0.1,
		}

	case 1, 2:
		return decision{
			maxBet:      b.stack + b.bet,
			raiseChance: 0.5,
			allInChance: 0.1,
		}
	}

	return decision{}
}

func (b *Bot) decideBoard(cards, board poker.Cards) decision {
	chances := math.ChancesAgainstN{OpponentsNum: b.opponentsNum}.WithBoard(cards, board)

	log.Printf("chances=%s", chances)

	tightness := 0.7
	if chances.Wins() > tightness {
		return decision{
			maxBet:      b.stack + b.bet,
			raiseChance: 0.5,
			allInChance: 0.5,
		}
	} else if chances.Wins() > tightness/2 {
		return decision{
			maxBet:      (b.stack + b.bet) / 3.,
			raiseChance: 0.2,
		}

	} else if chances.Ties() > 0.8 {
		return decision{
			maxBet:      b.stack + b.bet,
			raiseChance: 0.,
			allInChance: 0.,
		}
	}

	return decision{}
}

package util

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/ranking"
)

const (
	AttemptsCount = 10000
)

func Compare(a *poker.Cards, b *poker.Cards) (float64, float64, float64) {
	total := AttemptsCount
	wins, ties, loses := 0, 0, 0
	for i := 0; i <= total; i++ {
		dealer := model.NewDealer()
		dealer.Burn(a)
		dealer.Burn(b)
		board := dealer.Share(5)
		c1 := append(*a, *board...)
		c2 := append(*b, *board...)
		h1, _ := poker.Detect[ranking.High](&c1)
		h2, _ := poker.Detect[ranking.High](&c2)

		switch h1.Compare(h2) {
		case -1:
			loses++
		case 1:
			wins++
		case 0:
			ties++
		}
	}

	return float64(wins) / float64(total), float64(ties) / float64(total), float64(loses) / float64(total)
}

func Equity(cards *poker.Cards) float64 {
	total := AttemptsCount
	wins, ties, loses := 0, 0, 0
	for i := 0; i <= total; i++ {
		dealer := model.NewDealer()
		dealer.Burn(cards)
		other := dealer.Deal(2)
		board := dealer.Share(5)
		c1 := append(*cards, *board...)
		c2 := append(*other, *board...)
		h1, _ := poker.Detect[ranking.High](&c1)
		h2, _ := poker.Detect[ranking.High](&c2)
		switch h1.Compare(h2) {
		case -1:
			loses++
		case 1:
			wins++
		case 0:
			ties++
		}
	}

	return float64(wins) / float64(total)
}

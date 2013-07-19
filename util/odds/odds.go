package odds

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/ranking"
)

const (
	TrialsCount = 1000
)

func Compare(a poker.Cards, b poker.Cards, total int) (float64, float64, float64) {
	if total == 0 {
		total = TrialsCount
	}
	wins, ties, loses := 0, 0, 0
	for i := 0; i <= total; i++ {
		dealer := model.NewDealer()
		dealer.Burn(a)
		dealer.Burn(b)
		board := dealer.Share(5)

		c1 := append(a, board...)
		c2 := append(b, board...)
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

func Equity(cards poker.Cards) float64 {
	total := TrialsCount
	wins, ties, loses := 0, 0, 0

	for i := 0; i <= total; i++ {
		dealer := model.NewDealer()
		dealer.Burn(cards)
		other := dealer.Deal(2)
		board := dealer.Share(5)

		c1 := cards.Append(board)
		c2 := other.Append(board)
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

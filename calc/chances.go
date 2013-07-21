package calc

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/ranking"
	"gopoker/util"
)

const (
	DefaultSamplesCount = 1000
)

func Compare(a poker.Cards, b poker.Cards, total int) (float64, float64, float64) {
	if total == 0 {
		total = DefaultSamplesCount
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

func ChancesAgainstOnePreflop(cards poker.Cards) (int, int, int) {
	total := DefaultSamplesCount
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

	return wins, ties, loses
}

func ChancesAgainstOneWithBoard(cards poker.Cards, board poker.Cards) (int, int, int) {
	switch len(board) {
	case 3:
		return ChancesAgainstOneAtFlop(cards, board)
	case 4:
		return ChancesAgainstOneAtTurn(cards, board)
	case 5:
		return ChancesAgainstOneAtRiver(cards, board)
	}
	return 0, 0, 0
}

func ChancesAgainstOneAtFlop(cards poker.Cards, board poker.Cards) (int, int, int) {
	if len(board) < 3 {
		panic("board length for flop should be 3.")
	}
	if len(board) > 3 {
		board = board[0:3]
	}

	dealer := model.NewDealer()
	dealer.Burn(cards)
	dealer.Burn(board)

	cardsLeft := dealer.Deck

	wins, ties, loses := 0, 0, 0

	for turnCard := 0; turnCard < len(cardsLeft)-1; turnCard++ {
		for riverCard := turnCard + 1; riverCard < len(cardsLeft); riverCard++ {
			otherCombinations := util.Combine(len(cardsLeft), len(cards))
		OtherCombinations:
			for _, otherCombination := range otherCombinations {
				for _, k := range otherCombination {
					if k == turnCard || k == riverCard {
						continue OtherCombinations
					}
				}
				fullBoard := append(board, cardsLeft[turnCard], cardsLeft[riverCard])
				other := poker.Cards{}
				for _, k := range otherCombination {
					other = append(other, cardsLeft[k])
				}
				c1 := cards.Append(fullBoard)
				c2 := other.Append(fullBoard)
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
		}
	}

	return wins, ties, loses
}

func ChancesAgainstOneAtTurn(cards poker.Cards, board poker.Cards) (int, int, int) {
	if len(board) < 4 {
		panic("board length for turn should be 4.")
	}

	if len(board) > 4 {
		board = board[0:4]
	}

	dealer := model.NewDealer()
	dealer.Burn(cards)
	dealer.Burn(board)

	cardsLeft := dealer.Deck

	wins, ties, loses := 0, 0, 0

	for riverCard := 0; riverCard < len(cardsLeft); riverCard++ {
		otherCombinations := util.Combine(len(cardsLeft), len(cards))
	OtherCombinations:
		for _, otherCombination := range otherCombinations {
			for _, k := range otherCombination {
				if k == riverCard {
					continue OtherCombinations
				}
			}
			fullBoard := append(board, cardsLeft[riverCard])
			other := poker.Cards{}
			for _, k := range otherCombination {
				other = append(other, cardsLeft[k])
			}
			c1 := cards.Append(fullBoard)
			c2 := other.Append(fullBoard)
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
	}

	return wins, ties, loses
}

func ChancesAgainstOneAtRiver(cards poker.Cards, board poker.Cards) (int, int, int) {
	if len(board) < 5 {
		panic("board length for river should be 5.")
	}

	if len(board) > 5 {
		board = board[0:5]
	}

	dealer := model.NewDealer()
	dealer.Burn(cards)
	dealer.Burn(board)

	cardsLeft := dealer.Deck

	wins, ties, loses := 0, 0, 0

	otherCombinations := util.Combine(len(cardsLeft), len(cards))

	for _, otherCombination := range otherCombinations {
		other := poker.Cards{}
		for _, k := range otherCombination {
			other = append(other, cardsLeft[k])
		}
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

	return wins, ties, loses
}

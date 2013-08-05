package calc

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/hand"
	"gopoker/util"
)

// Defaults
const (
	DefaultSamplesCount = 1000
	FullBoardLen        = 5
)

// Chances - chances
type Chances struct {
	total int
	wins  int
	ties  int
	loses int
}

func (c Chances) Wins() float64 {
	return float64(c.wins) / float64(c.total)
}

func (c Chances) Ties() float64 {
	return float64(c.ties) / float64(c.total)
}

func (c Chances) Loses() float64 {
	return float64(c.loses) / float64(c.total)
}

// ChancesAgainstOne - chances against one player
type ChancesAgainstOne struct {
	SamplesNum int
}

// ChancesAgainstN - chances against n players
type ChancesAgainstN struct {
	OpponentsNum int
	SamplesNum   int
}

// Compare - compare chances
func (c *Chances) Compare(c1, c2 poker.Cards) {
	h1, _ := poker.Detect[hand.High](&c1)
	h2, _ := poker.Detect[hand.High](&c2)

	switch h1.Compare(h2) {
	case -1:
		c.loses++
	case 1:
		c.wins++
	case 0:
		c.ties++
	}
}

// Preflop - chances preflop
func (c ChancesAgainstOne) Preflop(hole, other poker.Cards) Chances {
	samplesNum := c.SamplesNum
	chances := &Chances{}

	for i := 0; i <= samplesNum; i++ {
		dealer := model.NewDealer()
		dealer.Burn(hole)
		dealer.Burn(other)
		board := dealer.Share(5)

		c1 := append(hole, board...)
		c2 := append(other, board...)

		chances.Compare(c1, c2)
	}

	return *chances
}

// WithBoard - chances with board dealt
func (c ChancesAgainstOne) WithBoard(hole, board poker.Cards) Chances {
	if len(board) > 5 || len(board) == 0 {
		panic("invalid board")
	}

	dealer := model.NewDealer()
	dealer.Burn(hole)
	dealer.Burn(board)

	cardsLeft := dealer.Deck
	holeCardsNum := len(hole)
	cardsNumToCompleteBoard := FullBoardLen - len(board)

	chances := &Chances{}

	for _, boardCombination := range util.Combine(len(cardsLeft), cardsNumToCompleteBoard) {
	OtherCombinations:
		for _, otherCombination := range util.Combine(len(cardsLeft), holeCardsNum) {
			for _, k := range otherCombination {
				for _, b := range boardCombination {
					if k == b {
						continue OtherCombinations // exclude
					}
				}
			}

			fullBoard := board
			for _, b := range boardCombination {
				fullBoard = append(fullBoard, cardsLeft[b])
			}

			other := poker.Cards{}
			for _, k := range otherCombination {
				other = append(other, cardsLeft[k])
			}

			c1 := hole.Append(fullBoard)
			c2 := other.Append(fullBoard)

			chances.Compare(c1, c2)
		}
	}

	return *chances
}

// Equity - hand equity
func (c ChancesAgainstN) Equity(hole, board poker.Cards) float64 {
	var chances Chances
	if len(board) == 0 {
		chances = c.Preflop(hole)
	} else {
		chances = c.WithBoard(hole, board)
	}

	e := float64(chances.wins) / float64(c.SamplesNum)
	e += float64(chances.ties) / float64(c.OpponentsNum)

	return e
}

// Preflop - chances preflop
func (c ChancesAgainstN) Preflop(hole poker.Cards) Chances {
	samplesNum := c.SamplesNum
	chances := &Chances{}

	for i := 0; i < samplesNum; i++ {
		dealer := model.NewDealer()
		dealer.Burn(hole)
		other := dealer.Deal(2)
		board := dealer.Share(5)

		c1 := hole.Append(board)
		c2 := other.Append(board)

		chances.Compare(c1, c2)
	}

	return *chances
}

// WithBoard - chances with board
func (c ChancesAgainstN) WithBoard(hole, board poker.Cards) Chances {
	if len(board) > 5 || len(board) == 0 {
		panic("board invalid")
	}

	opponentsNum := c.OpponentsNum
	samplesNum := c.SamplesNum
	holeCardsNum := len(hole)
	cardsNumToCompleteBoard := FullBoardLen - len(board)

	dealer := model.NewDealer()
	dealer.Burn(hole)
	dealer.Burn(board)

	cardsLeft := dealer.Deck

	chances := &Chances{}

	for i := 0; i < samplesNum; i++ {
		sampleDealer := model.NewDealerWithDeck(cardsLeft.Shuffle())

		fullBoard := board
		if cardsNumToCompleteBoard != 0 {
			fullBoard = append(board, sampleDealer.Deal(cardsNumToCompleteBoard)...)
		}

		for k := 0; k < opponentsNum; k++ {
			other := sampleDealer.Deal(holeCardsNum)
			c1 := hole.Append(fullBoard)
			c2 := other.Append(fullBoard)

			chances.Compare(c1, c2)
		}
	}

	return *chances
}

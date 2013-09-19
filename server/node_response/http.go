package node_response

import (
	"gopoker/poker"
)

// CompareResult - result of comparing cards
type CompareResult struct {
	A      *poker.Hand
	B      *poker.Hand
	Board  poker.Cards
	Result int
}

// OddsResult - result of calculating odds
type OddsResult struct {
	A     poker.Cards
	B     poker.Cards
	Total int
	Wins  float64
	Ties  float64
	Loses float64
}

// PocketHand - cards with hand
type PocketHand struct {
	Pocket poker.Cards
	Hand   *poker.Hand
}

// DealHand - pockets with board
type DealHand struct {
	Board   poker.Cards
	Pockets []PocketHand
}

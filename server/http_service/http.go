package http_service

import (
	"gopoker/poker"
)

type CompareResult struct {
	A      *poker.Hand
	B      *poker.Hand
	Board  *poker.Cards
	Result int
}

type OddsResult struct {
	A     *poker.Cards
	B     *poker.Cards
	Total int
	Wins  float64
	Ties  float64
	Loses float64
}

type PocketHand struct {
	Pocket poker.Cards
	Hand   *poker.Hand
}

type DealHand struct {
	Board   poker.Cards
	Pockets []PocketHand
}

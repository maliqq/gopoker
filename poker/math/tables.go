package math

import (
	"gopoker/poker"
)

var sklanskyMalmuthTable = []int{
	1, 1, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	2, 1, 2, 3, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	3, 4, 1, 3, 4, 5, 7, 9, 9, 9, 9, 9, 9,
	4, 5, 5, 1, 3, 4, 6, 8, 9, 9, 9, 9, 9,
	6, 6, 6, 5, 2, 4, 5, 7, 9, 9, 9, 9, 9,
	8, 8, 8, 7, 7, 3, 4, 5, 8, 9, 9, 9, 9,
	9, 9, 9, 8, 8, 7, 4, 5, 6, 8, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 8, 5, 5, 6, 8, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 8, 5, 6, 7, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 8, 6, 6, 7, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 8, 7, 7, 8,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 7, 8,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 7,
}

var evTable = []float64{
	2.32, 0.77, 0.59, 0.43, 0.33, 0.18, 0.10, 0.08, 0.03, 0.08, 0.06, 0.02, 0.00,
	0.51, 1.67, 0.39, 0.29, 0.20, 0.09, 0.01, 0.00, -0.04, -0.05, -0.05, -0.08, -0.08,
	0.31, 0.16, 1.22, 0.23, 0.17, 0.06, -0.02, -0.06, -0.08, -0.09, -0.10, -0.11, -0.12,
	0.19, 0.07, 0.03, 0.86, 0.15, 0.04, -0.03, -0.07, -0.11, -0.11, -0.11, -0.13, -0.14,
	0.08, 0.01, -0.02, -0.03, 0.58, 0.05, 0.00, -0.05, -0.11, -0.12, -0.13, -0.13, -0.14,
	-0.03, -0.07, -0.08, -0.08, -0.08, 0.38, 0.00, -0.04, -0.09, -0.12, -0.15, -0.14, -0.14,
	-0.07, -0.11, -0.11, -0.10, -0.09, -0.10, 0.25, -0.02, -0.07, -0.11, -0.13, -0.15, -0.14,
	-0.10, -0.11, -0.12, -0.12, -0.10, -0.10, -0.12, 0.16, -0.03, -0.09, -0.11, -0.14, -0.15,
	-0.12, -0.12, -0.13, -0.12, -0.11, -0.12, -0.11, -0.11, 0.07, -0.07, -0.09, -0.11, -0.14,
	-0.12, -0.13, -0.13, -0.13, -0.12, -0.12, -0.11, -0.11, -0.12, 0.02, -0.08, -0.11, -0.14,
	-0.12, -0.13, -0.13, -0.13, -0.12, -0.12, -0.12, -0.12, -0.12, -0.13, -0.03, -0.13, -0.14,
	-0.13, -0.08, -0.13, -0.13, -0.12, -0.12, -0.12, -0.12, -0.12, -0.12, -0.13, -0.07, -0.16,
	-0.15, -0.14, -0.13, -0.13, -0.12, -0.12, -0.12, -0.12, -0.12, -0.12, -0.12, -0.14, -0.09,
}

// SklanskyMalmuthGroup - Sklansky-Malmuth group for 2 cards
func SklanskyMalmuthGroup(card1, card2 *poker.Card) int {
	index1 := 12 - int(card1.Kind())
	index2 := 12 - int(card2.Kind())

	i := index(index1, index2, card1.Suit() == card2.Suit())

	return sklanskyMalmuthTable[i]
}

// RealPlayStatisticsEV - real play statistics EV
func RealPlayStatisticsEV(card1, card2 *poker.Card) float64 {
	index1 := 12 - int(card1.Kind())
	index2 := 12 - int(card2.Kind())

	i := index(index1, index2, card1.Suit() == card2.Suit())

	return evTable[i]
}

func index(index1, index2 int, suited bool) int {
	if suited {
		if index1 > index2 {
			return 13*index2 + index1
		}
		return 13*index1 + index2
	}
	if index1 > index2 {
		return 13*index1 + index2
	}
	return 13*index2 + index1
}
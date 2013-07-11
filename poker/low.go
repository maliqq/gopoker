package poker

import (
	"gopoker/poker/card"
	"gopoker/poker/hand"
)

func isLow(c *ordCards) (*Hand, error) {
	uniq := Cards{}
	for _, cards := range *c.GroupByKind() {
		uniq = append(uniq, cards[0])
	}

	lowCards := uniq.Reverse(c.Ordering)
	lowCards = lowCards[0:5]

	if len(lowCards) == 0 {
		return nil, nil
	}

	max := lowCards.Max(c.Ordering)
	newHand := &Hand{
		Value: lowCards,
		High:  Cards{*max},
	}

	if len(lowCards) == 5 {
		newHand.Rank = hand.CompleteLow
	} else {
		newHand.Rank = hand.IncompleteLow
	}

	return newHand, nil
}

func isGapLow(c *ordCards) (*Hand, error) {
	high, err := isHigh(c.Cards)
	if err != nil {
		return nil, err
	}

	if high.Rank == hand.HighCard {
		return isLow(c)
	}

	return high, nil
}

func isAceFive(c *Cards) (*Hand, error) {
	return isLow(
		NewOrderedCards(c, AceLow),
	)
}

func isAceFive8(c *Cards) (*Hand, error) {
	return isLow(
		NewOrderedCards(c, AceLow).Qualify(card.Eight),
	)
}

func isAceSix(c *Cards) (*Hand, error) {
	return isGapLow(
		NewOrderedCards(c, AceLow),
	)
}

func isDeuceSeven(c *Cards) (*Hand, error) {
	return isGapLow(
		NewOrderedCards(c, AceHigh),
	)
}

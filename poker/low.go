package poker

import (
	"gopoker/poker/card"
	"gopoker/poker/hand"
)

func isLow(c *Cards, ord Ordering) (*Hand, error) {
	helper := cardsHelper{*c, ord, false}

	uniq := Cards{}
	for _, cards := range *helper.GroupByKind() {
		uniq = append(uniq, cards[0])
	}

	lowCards := uniq.Reverse(helper.Ordering)
	lowCards = lowCards[0:5]

	if len(lowCards) == 0 {
		return nil, nil
	}

	max := lowCards.Max(helper.Ordering)
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

func isGapLow(c *Cards, ord Ordering) (*Hand, error) {
	high, err := isHigh(c)
	if err != nil {
		return nil, err
	}

	if high.Rank == hand.HighCard {
		return isLow(c, ord)
	}

	return high, nil
}

func isAceFive(c *Cards) (*Hand, error) {
	return isLow(c, AceLow)
}

func isAceFive8(c *Cards) (*Hand, error) {
	helper := cardsHelper{*c, AceLow, false}
	return isLow(helper.Qualify(card.Eight), AceLow)
}

func isAceSix(c *Cards) (*Hand, error) {
	return isGapLow(c, AceLow)
}

func isDeuceSeven(c *Cards) (*Hand, error) {
	return isGapLow(c, AceHigh)
}

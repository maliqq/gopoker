package poker

import (
	"gopoker/poker/card"
)

func isAceFive(c *Cards) (*Hand, error) {
	helper := cardsHelper{
		Cards:    *c,
		Ordering: AceLow,
		Low:      true,
	}

	return helper.IsLow()
}

func isAceFive8(c *Cards) (*Hand, error) {
	helper := cardsHelper{
		Cards:    *c,
		Ordering: AceLow,
		Low:      true,
	}

	helper.Qualify(card.Eight)

	return helper.IsLow()
}

func isAceSix(c *Cards) (*Hand, error) {
	helper := cardsHelper{
		Cards:    *c,
		Ordering: AceLow,
		Low:      true,
	}

	return helper.IsGapLow()
}

func isDeuceSeven(c *Cards) (*Hand, error) {
	helper := cardsHelper{
		Cards:    *c,
		Ordering: AceHigh,
		Low:      true,
	}

	return helper.IsGapLow()
}

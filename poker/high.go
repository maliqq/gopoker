package poker

import (
	"errors"
)

import (
	"gopoker/poker/hand"
)

var (
	HighRanks = []rankFunc{
		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.StraightFlush, hc.isStraightFlush()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.ThreeKind, hc.isThreeKind()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.TwoPair, hc.isTwoPair()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.OnePair, hc.isOnePair()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.HighCard, hc.isHighCard()
		},
	}

	NoFlushRanks = []rankFunc{
		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.FourKind, hc.isFourKind()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.FullHouse, hc.isFullHouse()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.Straight, hc.isStraight()
		},
	}

	NoStraightRanks = []rankFunc{
		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.FourKind, hc.isFourKind()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.FullHouse, hc.isFullHouse()
		},
	}
)

func (hc *handCards) isStraightFlush() *Hand {
	maybeFlush := hc.isFlush()
	if maybeFlush == nil {
		hand := hc.Detect(NoFlushRanks)

		if hand != nil {
			hand.rank = true
		}

		return hand
	}

	flushCards := maybeFlush.Value

	newhc := NewHandCards(&flushCards, hc.Ordering, false)

	if maybeStraight := newhc.isStraight(); maybeStraight != nil {
		return maybeStraight
	}

	maybeHigher := hc.Detect(NoStraightRanks)

	if maybeHigher != nil {
		maybeHigher.rank = true
		return maybeHigher
	}

	justFlush := *maybeFlush

	justFlush.rank = true
	justFlush.Rank = hand.Flush

	return &justFlush
}

func (hc *handCards) isFourKind() *Hand {
	quads, contains := hc.paired[4]
	if contains == false {
		return nil
	}

	cards := quads[0]

	hand := Hand{
		High:   cards[0:1],
		Value:  cards,
		kicker: true,
	}

	return &hand
}

func (hc *handCards) isFullHouse() *Hand {
	sets, containSets := hc.paired[3]
	if !containSets {
		return nil
	}

	var minor, major Cards

	if len(sets) > 1 {
		sorted := sets.ArrangeByFirst(hc.Ordering)

		major = (*sorted)[0]
		minor = (*sorted)[1]

	} else {
		pairs, containPairs := hc.paired[2]
		if !containPairs {
			return nil
		}

		sortedPairs := pairs.ArrangeByFirst(hc.Ordering)

		major = sets[0]
		minor = (*sortedPairs)[0]
	}

	return &Hand{
		Value: append(major, minor...),
		High:  Cards{major[0], minor[0]},
	}
}

func (hc *handCards) isFlush() *Hand {
	for count, group := range hc.suited {
		if count >= 5 {
			cards := group[0].Arrange(hc.Ordering)

			return &Hand{
				High:  Cards{cards[0]},
				Value: cards[0:5],
			}
		}
	}

	return nil
}

func (hc *handCards) isStraight() *Hand {
	for _, group := range hc.gaps {

		if len(group) >= 5 {
			cards := group.Arrange(hc.Ordering)

			return &Hand{
				Value: cards[0:5],
				High:  Cards{cards[0]},
			}
		}
	}

	return nil
}

func (hc *handCards) isThreeKind() *Hand {
	sets, containSets := hc.paired[3]
	if !containSets || len(sets) != 1 {
		return nil
	}

	return &Hand{
		Value:  sets[0],
		high:   true,
		kicker: true, // 2 kickers
	}
}

func (hc *handCards) isTwoPair() *Hand {
	pairs, containsPairs := hc.paired[2]
	if !containsPairs || len(pairs) < 2 {
		return nil
	}

	cards := pairs.ArrangeByMax(hc.Ordering)
	major, minor := (*cards)[0], (*cards)[1]

	return &Hand{
		Value:  append(major, minor...),
		High:   Cards{major[0], minor[0]},
		kicker: true,
	}
}

func (hc *handCards) isOnePair() *Hand {
	pairs, containsPairs := hc.paired[2]
	if !containsPairs || len(pairs) != 1 {
		return nil
	}

	cards := pairs[0]

	return &Hand{
		Value:  cards,
		high:   true,
		kicker: true,
	}
}

func (hc *handCards) isHighCard() *Hand {
	cards := hc.Arrange()

	return &Hand{
		Value:  cards[0:1],
		high:   true,
		kicker: true,
	}
}

func isHigh(cards *Cards) (*Hand, error) {
	if len(*cards) < 5 {
		return nil, errors.New("5 or more cards required to detect high hand")
	}

	hc := NewHandCards(cards, AceHigh, false)

	hand := hc.Detect(HighRanks)

	return hand, nil
}

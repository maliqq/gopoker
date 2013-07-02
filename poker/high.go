package poker

import (
	"errors"
	"gopoker/poker/hand"
)

func (pocket *PocketCards) isStraightFlush() *Hand {
	maybeFlush := pocket.isFlush()
	if maybeFlush == nil {
		hand := pocket.Detect([]Ranker{
			Ranker{hand.FourKind, func(pocket *PocketCards) *Hand {
				return pocket.isFourKind()
			}},
			Ranker{hand.FullHouse, func(pocket *PocketCards) *Hand {
				return pocket.isFullHouse()
			}},
			Ranker{hand.Straight, func(pocket *PocketCards) *Hand {
				return pocket.isStraight()
			}},
		})
		if hand != nil {
			hand.rank = true
		}
		return hand
	}

	flushCards := maybeFlush.Value
	newPocket := NewPocket(&OrderedCards{&flushCards, pocket.Ordering()})
	maybeStraight := newPocket.isStraight()
	if maybeStraight != nil {
		return maybeStraight
	}

	maybeHigher := pocket.Detect([]Ranker{
		Ranker{hand.FourKind, func(pocket *PocketCards) *Hand {
			return pocket.isFourKind()
		}},
		Ranker{hand.FullHouse, func(pocket *PocketCards) *Hand {
			return pocket.isFullHouse()
		}},
	})

	if maybeHigher != nil {
		maybeHigher.rank = true
		return maybeHigher
	}
	justFlush := *maybeFlush
	justFlush.rank = true
	justFlush.Rank = hand.Flush

	return &justFlush
}

func (pocket *PocketCards) isFourKind() *Hand {
	quads, contains := (*pocket.paired)[4]
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

func (pocket *PocketCards) isFullHouse() *Hand {
	sets, containSets := (*pocket.paired)[3]
	if !containSets {
		return nil
	}

	var minor, major Cards

	if len(sets) > 1 {
		sorted := ArrangeGroupsByFirst(&sets, pocket.Ordering())
		major = (*sorted)[0]
		minor = (*sorted)[1]
	} else {
		pairs, containPairs := (*pocket.paired)[2]
		if !containPairs {
			return nil
		}
		sortedPairs := ArrangeGroupsByFirst(&pairs, pocket.Ordering())
		major = sets[0]
		minor = (*sortedPairs)[0]
	}

	return &Hand{
		Value: append(major, minor...),
		High:  Cards{major[0], minor[0]},
	}
}

func (pocket *PocketCards) isFlush() *Hand {
	for count, group := range *pocket.suited {
		if count >= 5 {
			cards := *ArrangeCards(&group[0], pocket.Ordering())
			return &Hand{
				High:  Cards{cards[0]},
				Value: cards[0:5],
			}
		}
	}
	return nil
}

func (pocket *PocketCards) isStraight() *Hand {
	for _, group := range pocket.gaps {
		if len(group) >= 5 {
			cards := *ArrangeCards(&group, pocket.Ordering())
			// FIXME: wheel straight
			return &Hand{
				Value: cards[0:5],
				High:  Cards{cards[0]},
			}
		}
	}
	return nil
}

func (pocket *PocketCards) isThreeKind() *Hand {
	sets, containSets := (*pocket.paired)[3]
	if !containSets || len(sets) != 1 {
		return nil
	}

	return &Hand{
		Value:  sets[0],
		high:   true,
		kicker: true, // 2 kickers
	}
}

func (pocket *PocketCards) isTwoPair() *Hand {
	pairs, containsPairs := (*pocket.paired)[2]
	if !containsPairs || len(pairs) < 2 {
		return nil
	}

	cards := ArrangeGroupsByMax(&pairs, pocket.Ordering())

	major, minor := (*cards)[0], (*cards)[1]

	return &Hand{
		Value:  append(major, minor...),
		High:   Cards{major[0], minor[0]},
		kicker: true,
	}
}

func (pocket *PocketCards) isOnePair() *Hand {
	pairs, containsPairs := (*pocket.paired)[2]
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

func (pocket *PocketCards) isHighCard() *Hand {
	cards := ArrangeCards(pocket.Cards(), pocket.Ordering())

	return &Hand{
		Value:  (*cards)[0:1],
		high:   true,
		kicker: true,
	}
}

func isHigh(c *Cards) (*Hand, error) {
	if len(*c) < 5 {
		return nil, errors.New("5 or more cards required to detect high hand")
	}
	pocket := NewPocket(&OrderedCards{c, AceHigh})
	hand := pocket.Detect([]Ranker{
		Ranker{hand.StraightFlush, func(pocket *PocketCards) *Hand {
			return pocket.isStraightFlush()
		}},
		Ranker{hand.ThreeKind, func(pocket *PocketCards) *Hand {
			return pocket.isThreeKind()
		}},
		Ranker{hand.TwoPair, func(pocket *PocketCards) *Hand {
			return pocket.isTwoPair()
		}},
		Ranker{hand.OnePair, func(pocket *PocketCards) *Hand {
			return pocket.isOnePair()
		}},
		Ranker{hand.HighCard, func(pocket *PocketCards) *Hand {
			return pocket.isHighCard()
		}},
	})
	return hand, nil
}

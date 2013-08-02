package model

import (
	"gopoker/poker"
	"gopoker/poker/hand"
)

// Deal - wrapper around dealer
type Deal struct {
	dealer *Dealer
	// pocket cards
	Pockets map[Player]poker.Cards `json:"-"`
	// shared cards
	Board poker.Cards
}

// NewDeal initialize deal
func NewDeal() *Deal {
	return &Deal{
		dealer:  NewDealer(),
		Pockets: map[Player]poker.Cards{},
		Board:   poker.Cards{},
	}
}

// Pocket - get dealt pocket cards for player
func (deal *Deal) Pocket(player Player) poker.Cards {
	cards, found := deal.Pockets[player]

	if !found {
		cards = poker.Cards{}
		deal.Pockets[player] = cards
	}

	return cards
}

// DealBoard - deal shared cards
func (deal *Deal) DealBoard(cardsNum int) poker.Cards {
	cards := deal.dealer.Share(cardsNum)

	deal.Board = append(deal.Board, cards...)

	return cards
}

// DealPocket - deal private cards
func (deal *Deal) DealPocket(player Player, cardsNum int) poker.Cards {
	pocket := deal.Pocket(player)

	cards := deal.dealer.Deal(cardsNum)
	deal.Pockets[player] = append(pocket, cards...)

	return cards
}

// Discard - return cards to dealer
func (deal *Deal) Discard(player Player, cards poker.Cards) poker.Cards {
	pocket := deal.Pocket(player)
	newCards := deal.dealer.Discard(cards)

	pocket = append(pocket.Diff(cards), newCards...)

	return newCards
}

// Rank - rank cards for player
func (deal *Deal) Rank(player Player, ranking hand.Ranking, hasBoard bool) (poker.Cards, *poker.Hand) {
	pocket := deal.Pocket(player)

	if !hasBoard {
		hand, _ := poker.Detect[ranking](&pocket)

		return pocket, hand
	}

	var bestHand *poker.Hand

	for _, pair := range pocket.Combine(2) {
		for _, board := range deal.Board.Combine(3) {
			handCards := append(pair, board...)

			hand, _ := poker.Detect[ranking](&handCards)

			if bestHand == nil || hand.Compare(bestHand) > 0 {
				bestHand = hand
			}
		}
	}

	return pocket, bestHand
}

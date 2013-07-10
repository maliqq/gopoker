package model

import (
	"fmt"
)

import (
	"gopoker/poker"
	"gopoker/poker/ranking"
)

type Deal struct {
	dealer  *Dealer
	Pockets map[Id]*poker.Cards
	Board   poker.Cards
}

func NewDeal() *Deal {
	return &Deal{
		dealer:  NewDealer(),
		Pockets: map[Id]*poker.Cards{},
		Board:   poker.Cards{},
	}
}

func (this *Deal) Pocket(player *Player) *poker.Cards {
	cards, found := this.Pockets[player.Id]

	if !found {
		cards = &poker.Cards{}
		this.Pockets[player.Id] = cards
	}

	return cards
}

func (this *Deal) DealBoard(cardsNum int) *poker.Cards {
	cards := this.dealer.Share(cardsNum)

	this.Board = append(this.Board, *cards...)

	return cards
}

func (this *Deal) DealPocket(player *Player, cardsNum int) *poker.Cards {
	pocket := this.Pocket(player)

	cards := this.dealer.Deal(cardsNum)
	*pocket = append(*pocket, *cards...)

	return cards
}

func (this *Deal) Discard(player *Player, cards *poker.Cards) *poker.Cards {
	fmt.Printf("deal=%#v", this)
	pocket := this.Pocket(player)
	newCards := this.dealer.Discard(cards)

	diff := poker.DiffCards(pocket, cards)
	*pocket = append(*diff, *newCards...)

	return newCards
}

func (this *Deal) Rank(cards *poker.Cards, ranking ranking.Type, hasBoard bool) *poker.Hand {
	if !hasBoard {
		hand, _ := poker.Detect[ranking](cards)

		return hand
	}

	var bestHand *poker.Hand
	for _, pair := range cards.CombinePairs() {
		handCards := append(pair, this.Board...)

		hand, _ := poker.Detect[ranking](&handCards)

		if bestHand == nil || hand.Compare(bestHand) > 0 {
			bestHand = hand
		}
	}

	return bestHand
}

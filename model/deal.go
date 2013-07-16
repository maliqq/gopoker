package model

import (
	_ "fmt"
)

import (
	"gopoker/poker"
	"gopoker/poker/ranking"
)

type Deal struct {
	dealer *Dealer
	// pocket cards
	Pockets map[Id]*poker.Cards `json:"-"`
	// shared cards
	Board poker.Cards
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
	pocket := this.Pocket(player)
	newCards := this.dealer.Discard(cards)

	*pocket = append(pocket.Diff(cards), *newCards...)

	return newCards
}

func (this *Deal) Rank(player *Player, ranking ranking.Ranking, hasBoard bool) (*poker.Cards, *poker.Hand) {
	pocket := this.Pocket(player)

	if !hasBoard {
		hand, _ := poker.Detect[ranking](pocket)

		return pocket, hand
	}

	var bestHand *poker.Hand

	for _, pair := range pocket.Combine(2) {
		for _, board := range this.Board.Combine(3) {
			handCards := append(pair, board...)

			hand, _ := poker.Detect[ranking](&handCards)

			if bestHand == nil || hand.Compare(bestHand) > 0 {
				bestHand = hand
			}
		}
	}

	return pocket, bestHand
}

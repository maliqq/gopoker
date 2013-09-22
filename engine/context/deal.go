package context

import (
  "gopoker/model"
  "gopoker/poker"
  "gopoker/poker/hand"
)

// Deal - known & common cards
type Deal struct {
  dealer *model.Dealer
  // pocket cards
  dealt map[model.Player]poker.Cards `json:"-"`
  // shared cards
  Board poker.Cards
}

// NewDeal initialize deal
func NewDeal() *Deal {
  return &Deal{
    dealer:  model.NewDealer(),
    dealt: map[model.Player]poker.Cards{},
    Board:   poker.Cards{},
  }
}

// Pocket - get dealt pocket cards for player
func (ctx *Deal) Pocket(player model.Player) poker.Cards {
  cards, found := ctx.dealt[player]

  if !found {
    cards = poker.Cards{}
    ctx.dealt[player] = cards
  }

  return cards
}

// DealBoard - deal shared cards
func (ctx *Deal) DealBoard(cardsNum int) poker.Cards {
  cards := ctx.dealer.Share(cardsNum)

  ctx.Board = append(ctx.Board, cards...)

  return cards
}

// DealPocket - deal private cards
func (ctx *Deal) DealPocket(player model.Player, cardsNum int) poker.Cards {
  pocket := ctx.dealt[player]

  cards := ctx.dealer.Deal(cardsNum)
  ctx.dealt[player] = append(pocket, cards...)

  return cards
}

// Discard - return cards to dealer
func (ctx *Deal) Discard(player model.Player, cards poker.Cards) poker.Cards {
  pocket := ctx.dealt[player]
  newCards := ctx.dealer.Discard(cards)

  pocket = append(pocket.Diff(cards), newCards...)

  return newCards
}

// Rank - rank cards for player
func (ctx *Deal) Rank(player model.Player, ranking hand.Ranking) (poker.Cards, *poker.Hand) {
  pocket := ctx.dealt[player]

  if len(ctx.Board) == 0 {
    hand, _ := poker.Detect[ranking](&pocket)

    return pocket, hand
  }

  var bestHand *poker.Hand

  for _, pair := range pocket.Combine(2) {
    for _, board := range ctx.Board.Combine(3) {
      handCards := append(pair, board...)

      hand, _ := poker.Detect[ranking](&handCards)

      if bestHand == nil || hand.Compare(bestHand) > 0 {
        bestHand = hand
      }
    }
  }

  return pocket, bestHand
}

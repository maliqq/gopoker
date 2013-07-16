package gameplay

import (
  "gopoker/poker"
  "gopoker/model"
  "gopoker/model/deal"
  "gopoker/protocol"
)

func (this *GamePlay) StartDiscardingRound() {
  discarding := this.Discarding

  for _, pos := range this.Table.SeatsFromButton().InPlay() {
    seat := this.Table.Seat(pos)

    this.Broadcast.One(seat.Player) <- discarding.RequireDiscard(pos, seat)
  }
}

func (this *GamePlay) discard(p *model.Player, cards *poker.Cards) {
  pos, _ := this.Table.Pos(p)

  cardsNum := len(*cards)

  this.Broadcast.All <- protocol.NewDiscarded(pos, cardsNum)

  if cardsNum > 0 {
    newCards := this.Deal.Discard(p, cards)

    this.Broadcast.One(p) <- protocol.NewDealPocket(pos, newCards, deal.Discard)
  }
}

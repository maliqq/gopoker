package gameplay

import (
  "gopoker/protocol"
  "gopoker/model/deal"
)

func (this *GamePlay) DealHole(cardsNum int) {
  for _, pos := range this.Table.SeatsInPlay() {
    player := this.Table.Player(pos)

    cards := this.Deal.DealPocket(player, cardsNum)

    this.Broadcast.One(player) <- protocol.NewDealPocket(pos, cards, deal.Hole)
  }
}

func (this *GamePlay) DealDoor(cardsNum int) {
  for _, pos := range this.Table.SeatsInPlay() {
    player := this.Table.Player(pos)

    cards := this.Deal.DealPocket(player, cardsNum)

    this.Broadcast.All <- protocol.NewDealPocket(pos, cards, deal.Door)
  }
}

func (this *GamePlay) DealBoard(cardsNum int) {
  cards := this.Deal.DealBoard(cardsNum)

  this.Broadcast.All <- protocol.NewDealShared(cards, deal.Board)
}

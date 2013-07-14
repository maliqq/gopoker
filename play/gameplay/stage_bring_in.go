package gameplay

import (
  "gopoker/poker"
  "gopoker/protocol"
)

func (this *GamePlay) SetButton(pos int) {
  this.Table.SetButton(pos)

  this.Broadcast.All <- protocol.NewMoveButton(pos)
}

func (this *GamePlay) BringIn() {
  minPos := 0
  var card poker.Card

  for _, pos := range this.Table.SeatsInPlay() {
    s := this.Table.Seat(pos)

    pocketCards := *this.Deal.Pocket(s.Player)

    lastCard := pocketCards[len(pocketCards)-1]
    if pos == 0 {
      card = lastCard
    } else {
      if lastCard.Compare(card, poker.AceHigh) > 0 {
        card = lastCard
        minPos = pos
      }
    }
  }

  this.SetButton(minPos)

  seat := this.Table.Seat(minPos)

  this.Broadcast.One(seat.Player) <- this.Betting.RequireBet(minPos, seat, this.Game, this.Stake)
}

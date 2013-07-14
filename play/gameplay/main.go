package gameplay

import (
  "gopoker/model"
  "gopoker/protocol"
)

type GamePlay struct {
  // dealt cards context
  Deal *model.Deal

  // mixed or limited game
  Game          *model.Game
  Mix           *model.Mix

  // betting price
  Stake *model.Stake

  // players seating context
  Table *model.Table

  // broadcast protocol messages
  Broadcast *protocol.Broadcast `json:"-"`
}

func (gp *GamePlay) TestF() bool {
  return true
}

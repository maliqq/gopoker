package gameplay

import (
  "gopoker/model"
  "gopoker/protocol"
  "gopoker/play/context"
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
  *context.GameRotation `json:"-"`
  *context.Betting    `json:"-"`
  *context.Discarding `json:"-"`
}

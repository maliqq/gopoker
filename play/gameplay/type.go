package gameplay

import (
  "gopoker/model"
  "gopoker/protocol"
  "gopoker/play/context"
  "gopoker/play/command"
)

type GamePlay struct {
  // dealt cards context
  Deal *model.Deal

  // mixed or limited game
  Game          *model.Game
  Mix           *model.Mix
  *context.GameRotation `json:"-"`

  // betting price
  Stake *model.Stake

  // players seating context
  Table *model.Table

  // broadcast protocol messages
  Broadcast *protocol.Broadcast `json:"-"`
  
  *context.Betting    `json:"-"`
  *context.Discarding `json:"-"`
  // manage play
  Control chan command.Command `json:"-"`
}

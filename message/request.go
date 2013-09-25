package message

import (
  "gopoker/model"
)

type CreateRoom struct {
  Guid model.Guid
  BetSize float64
  Game *model.Game
  Mix *model.Mix
}

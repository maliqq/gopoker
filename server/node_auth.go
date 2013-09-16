package server

import (
  "gopoker/model"
)

func (node *Node) Authorize(key string) (model.Player, bool) {
  result := node.SessionStore.Get(key)
  
  var player model.Player
  
  if result != nil {
    player = result.PlayerID
  }

  return player, player == ""
}

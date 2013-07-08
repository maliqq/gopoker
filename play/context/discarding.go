package context

import (
	"gopoker/protocol"
)

type Discarding struct {
	requireDiscard *protocol.RequireDiscard

  Receive chan *protocol.Message
}

func NewDiscarding() *Discarding {
  return &Discarding{
    Receive: make(chan *protocol.Message),
  }
}

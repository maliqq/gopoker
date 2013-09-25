package event

import (
	"gopoker/message"
)

type Call struct {
	Method  string
	Message *message.Message
}

type CallResult struct {
  Error error
  Type string
  Message *message.Message
}

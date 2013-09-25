package event

import (
	"gopoker/message"
)

type Call struct {
	Method  string
	Message *message.Message
}

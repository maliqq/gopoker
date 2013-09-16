package exch

import (
	"gopoker/exch/message"
)

type MessageHandler interface {
	HandleMessage(*message.Message)
}

package event

import (
	"gopoker/event/message"
)

type MessageHandler interface {
	HandleMessage(*message.Message)
}

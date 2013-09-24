package hub

import (
	"gopoker/event"
)

type Subscriber struct {
	Recv event.Channel
}

func (s Subscriber) Send(message interface{}) {
	n, ok := message.(*event.Notification)
	if ok {
		s.Recv <- n
	}
}

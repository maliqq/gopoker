package hub

import (
	"gopoker/event"
)

type Subscriber struct {
	Recv event.Channel
}

func NewSubscriber(recv event.Channel) Subscriber {
	return Subscriber{recv}
}

func (s Subscriber) Send(message interface{}) {
	n, ok := message.(*event.Notification)
	if ok {
		s.Recv <- n
	}
}

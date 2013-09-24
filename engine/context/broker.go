package context

import (
	"gopoker/hub"
	"gopoker/message"
	"gopoker/event"
)

type Broker struct {
	*hub.Exchange
}

type Notify struct {
	Exchange *hub.Exchange
	Notification  *event.Notification
}

type key interface {
	String() string
}

func NewBroker() *Broker{
	return &Broker{hub.NewExchange()}
}

func (b *Broker) Notify(msg message.Message) Notify {
	return Notify{b.Exchange, event.New(msg)}
}

func (n Notify) All() {
	n.Exchange.Dispatch(
		hub.Route{
			All: true,
		},
		n.Notification,
	)
}

func (n Notify) One(key key) {
	n.Exchange.Dispatch(
		hub.Route{
			One: key.String(),
		},
		n.Notification,
	)
}

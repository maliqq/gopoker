package protocol

import (
	"gopoker/protocol/message"
)

type Notification struct {
	*message.Message
	*Notify
}

type Broadcast struct {
	*Broker
	All    MessageChannel
	System MessageChannel
	Route  chan *Notification
}

type Actor interface {
	RouteKey() string
}

func NewBroadcast() *Broadcast {
	broadcast := &Broadcast{
		Broker: NewBroker(),
		All:    make(MessageChannel),
		System: make(MessageChannel),
		Route:  make(chan *Notification),
	}

	go broadcast.receive()

	return broadcast
}

func (bcast *Broadcast) receive() {
	for {
		select {
		case msg := <-bcast.All:
			bcast.Broker.Dispatch(&Notify{All: true}, msg)

		case msg := <-bcast.System:
			bcast.Broker.Dispatch(&Notify{None: true}, msg)

		case notification := <-bcast.Route:
			bcast.Broker.Dispatch(notification.Notify, notification.Message)
		}
	}
}

func (bcast *Broadcast) route(notify *Notify) MessageChannel {
	channel := make(MessageChannel)

	go func() {
		msg := <-channel
		bcast.Route <- &Notification{msg, notify}
	}()

	return channel
}

func (bcast *Broadcast) One(actor Actor) MessageChannel {
	notify := &Notify{One: actor.RouteKey()}

	return bcast.route(notify)
}

func (bcast *Broadcast) Except(actors ...Actor) MessageChannel {
	keys := make([]string, len(actors))
	for i, a := range actors {
		keys[i] = a.RouteKey()
	}

	notify := &Notify{Except: keys}

	return bcast.route(notify)
}

func (bcast *Broadcast) Only(actors ...Actor) MessageChannel {
	keys := make([]string, len(actors))
	for i, a := range actors {
		keys[i] = a.RouteKey()
	}

	notify := &Notify{Only: keys}

	return bcast.route(notify)
}

func (bcast *Broadcast) Bind(actor Actor, ch *MessageChannel) {
	bcast.Broker.BindUser(actor.RouteKey(), ch)
}

func (bcast *Broadcast) Unbind(actor Actor) {
	bcast.Broker.UnbindUser(actor.RouteKey())
}

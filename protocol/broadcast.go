package protocol

import (
	"gopoker/protocol/message"
)

// Route - message with route
type Route struct {
	*message.Message
	*Notify
}

// Broadcast - broadcast hub
type Broadcast struct {
	*Broker
	All   MessageChannel
	Route chan *Route
}

// Actor - message receiver
type Actor interface {
	RouteKey() string
}

// NewBroadcast - create new broadcast hub
func NewBroadcast() *Broadcast {
	broadcast := &Broadcast{
		Broker: NewBroker(),
		All:    make(MessageChannel),
		Route:  make(chan *Route),
	}

	go broadcast.receive()

	return broadcast
}

func (bcast *Broadcast) receive() {
	for {
		select {
		case msg := <-bcast.All:
			bcast.Broker.Dispatch(&Notify{All: true}, msg)

		case route := <-bcast.Route:
			bcast.Broker.Dispatch(route.Notify, route.Message)
		}
	}
}

func (bcast *Broadcast) route(notify *Notify) MessageChannel {
	channel := make(MessageChannel)

	go func() {
		msg := <-channel
		bcast.Route <- &Route{msg, notify}
	}()

	return channel
}

// One - route to one receiver
func (bcast *Broadcast) One(actor Actor) MessageChannel {
	notify := &Notify{One: actor.RouteKey()}

	return bcast.route(notify)
}

// Except - route to all except
func (bcast *Broadcast) Except(actors ...Actor) MessageChannel {
	keys := make([]string, len(actors))
	for i, a := range actors {
		keys[i] = a.RouteKey()
	}

	notify := &Notify{Except: keys}

	return bcast.route(notify)
}

// Only - route to only
func (bcast *Broadcast) Only(actors ...Actor) MessageChannel {
	keys := make([]string, len(actors))
	for i, a := range actors {
		keys[i] = a.RouteKey()
	}

	notify := &Notify{Only: keys}

	return bcast.route(notify)
}

// Bind - bind receiver to hub
func (bcast *Broadcast) Bind(group Group, actor Actor, ch *MessageChannel) {
	bcast.Broker.Bind(group, actor.RouteKey(), ch)
}

// Unbind - unbind receiver from hub
func (bcast *Broadcast) Unbind(group Group, actor Actor) {
	bcast.Broker.Unbind(group, actor.RouteKey())
}

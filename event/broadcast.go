package event

import (
	"gopoker/event/message"
)

// Route - message with route
type Route struct {
	*Event
	Notify
}

type messageChannel chan message.Message
type key interface {
	String() string
}

// Broadcast - broadcast hub
type Broadcast struct {
	*Broker
	All   messageChannel
	Route chan Route
}

// NewBroadcast - create new broadcast hub
func NewBroadcast() *Broadcast {
	broadcast := &Broadcast{
		Broker: NewBroker(),
		All:    make(chan message.Message),
		Route:  make(chan Route),
	}

	go broadcast.receive()

	return broadcast
}

func (bcast *Broadcast) Pass(event *Event) {
	bcast.Broker.Dispatch(Notify{All: true}, event)
}

func (bcast *Broadcast) receive() {
	for {
		select {
		case msg := <-bcast.All:
			bcast.Pass(NewEvent(msg))

		case route := <-bcast.Route:
			bcast.Broker.Dispatch(route.Notify, route.Event)
		}
	}
}

func (bcast *Broadcast) route(notify Notify) messageChannel {
	channel := make(messageChannel)

	go func() {
		msg := <-channel
		event := NewEvent(msg)
		bcast.Route <- Route{event, notify}
	}()

	return channel
}

// One - route to one receiver
func (bcast *Broadcast) One(key key) messageChannel {
	return bcast.route(Notify{
		One: key.String(),
	})
}

// Except - route to all except
func (bcast *Broadcast) Except(keys ...key) messageChannel {
	except := make([]string, len(keys))
	for i, a := range keys {
		except[i] = a.String()
	}

	return bcast.route(Notify{
		Except: except,
	})
}

// Only - route to only
func (bcast *Broadcast) Only(keys ...key) messageChannel {
	only := make([]string, len(keys))
	for i, a := range keys {
		only[i] = a.String()
	}

	return bcast.route(Notify{
		Only: only,
	})
}

// Bind - bind receiver to hub
func (bcast *Broadcast) Bind(key key, channel *Channel) {
	subscriber := Subscriber{
		Role:    User,
		Key:     key.String(),
		Channel: channel,
	}

	bcast.Broker.Bind(Default, subscriber)
}

// Unbind - unbind receiver from hub
func (bcast *Broadcast) Unbind(key key) {
	bcast.Broker.Unbind(Default, key.String())
}

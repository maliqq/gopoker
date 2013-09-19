package event

import (
	"gopoker/event/message"
)

type Notification struct {
	Broker *Broker
	Event  *Event
}

type messageChannel chan message.Message
type key interface {
	String() string
}

// Broadcast - broadcast hub
type Broadcast struct {
	*Broker
}

// NewBroadcast - create new broadcast hub
func NewBroadcast() *Broadcast {
	broadcast := Broadcast{
		Broker: NewBroker(),
	}
	return &broadcast
}

func (bcast *Broadcast) Notify(msg message.Message) *Notification {
	return &Notification{
		Event:  New(msg),
		Broker: bcast.Broker,
	}
}

func (bcast *Broadcast) Pass(event *Event) *Notification {
	return &Notification{
		Event:  event,
		Broker: bcast.Broker,
	}
}

func (n *Notification) route(notify Notify) {
	notify.Topic = Default // FIXME
	n.Broker.Dispatch(notify, n.Event)
}

func (n *Notification) All() {
	n.route(Notify{All: true})
}

// One - route to one receiver
func (n *Notification) One(key key) {
	n.route(Notify{
		One: key.String(),
	})
}

// Except - route to all except
func (n *Notification) Except(keys ...key) {
	except := make([]string, len(keys))
	for i, a := range keys {
		except[i] = a.String()
	}

	n.route(Notify{
		Except: except,
	})
}

// Only - route to only
func (n *Notification) Only(keys ...key) {
	only := make([]string, len(keys))
	for i, a := range keys {
		only[i] = a.String()
	}

	n.route(Notify{
		Only: only,
	})
}

// Bind - bind receiver to hub
func (bcast *Broadcast) Bind(key key, channel *Channel) {
	bcast.Broker.Bind(Default, SubscribeUser(key.String(), channel))
}

// Unbind - unbind receiver from hub
func (bcast *Broadcast) Unbind(key key) {
	bcast.Broker.Unbind(Default, key.String())
}

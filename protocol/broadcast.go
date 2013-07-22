package protocol

type Notification struct {
	*Message
	*Notify
}

type Broadcast struct {
	*Broker
	All   MessageChannel
	Route chan *Notification
}

type Actor interface {
	RouteKey() string
}

func NewBroadcast() *Broadcast {
	broadcast := &Broadcast{
		Broker: NewBroker(),
		All:    make(MessageChannel),
		Route:  make(chan *Notification),
	}

	go broadcast.receive()

	return broadcast
}

func (broadcast *Broadcast) receive() {
	for {
		select {
		case msg := <-broadcast.All:
			broadcast.Broker.Dispatch(&Notify{All: true}, msg)

		case notification := <-broadcast.Route:
			broadcast.Broker.Dispatch(notification.Notify, notification.Message)
		}
	}
}

func (broadcast *Broadcast) route(notify *Notify) MessageChannel {
	channel := make(MessageChannel)

	go func() {
		msg := <-channel
		broadcast.Route <- &Notification{msg, notify}
	}()

	return channel
}

func (broadcast *Broadcast) One(actor Actor) MessageChannel {
	notify := &Notify{One: actor.RouteKey()}

	return broadcast.route(notify)
}

func (broadcast *Broadcast) Except(actors ...Actor) MessageChannel {
	keys := make([]string, len(actors))
	for i, a := range actors {
		keys[i] = a.RouteKey()
	}

	notify := &Notify{Except: keys}

	return broadcast.route(notify)
}

func (broadcast *Broadcast) Only(actors ...Actor) MessageChannel {
	keys := make([]string, len(actors))
	for i, a := range actors {
		keys[i] = a.RouteKey()
	}

	notify := &Notify{Only: keys}

	return broadcast.route(notify)
}

func (broadcast *Broadcast) Bind(actor Actor, ch *MessageChannel) {
	broadcast.Broker.Bind(actor.RouteKey(), ch)
}

func (broadcast *Broadcast) For(actor Actor) *MessageChannel {
	return broadcast.Broker.For(actor.RouteKey())
}

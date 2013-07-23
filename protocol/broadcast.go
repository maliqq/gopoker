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

func (this *Broadcast) receive() {
	for {
		select {
		case msg := <-this.All:
			this.Broker.Dispatch(&Notify{All: true}, msg)

		case notification := <-this.Route:
			this.Broker.Dispatch(notification.Notify, notification.Message)
		}
	}
}

func (this *Broadcast) route(notify *Notify) MessageChannel {
	channel := make(MessageChannel)

	go func() {
		msg := <-channel
		this.Route <- &Notification{msg, notify}
	}()

	return channel
}

func (this *Broadcast) One(actor Actor) MessageChannel {
	notify := &Notify{One: actor.RouteKey()}

	return this.route(notify)
}

func (this *Broadcast) Except(actors ...Actor) MessageChannel {
	keys := make([]string, len(actors))
	for i, a := range actors {
		keys[i] = a.RouteKey()
	}

	notify := &Notify{Except: keys}

	return this.route(notify)
}

func (this *Broadcast) Only(actors ...Actor) MessageChannel {
	keys := make([]string, len(actors))
	for i, a := range actors {
		keys[i] = a.RouteKey()
	}

	notify := &Notify{Only: keys}

	return this.route(notify)
}

func (this *Broadcast) Bind(actor Actor, ch *MessageChannel) {
	this.Broker.Bind(actor.RouteKey(), ch)
}

func (this *Broadcast) For(actor Actor) *MessageChannel {
	return this.Broker.For(actor.RouteKey())
}

package event

type Handler interface {
	HandleEvent(*Event)
}

type Observer struct {
	Recv    Channel
	Handler Handler
}

func NewObserver(handler Handler) *Observer {
	observer := Observer{Handler: handler}
	observer.Start()
	return &observer
}

func (observer *Observer) Start() {
	observer.Recv = make(Channel)
	go observer.receive()
}

func (observer *Observer) receive() {
	for {
		msg := <-observer.Recv
		observer.Handler.HandleEvent(msg)
	}
}

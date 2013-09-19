package event

type Handler interface {
	HandleEvent(*Event)
}

type Observer struct {
	Recv    Channel
	Handler Handler
}

func NewObserver(handler Handler) *Observer {
	observer := Observer{
		Handler: handler,
		Recv:    make(Channel),
	}
	go observer.receive()
	return &observer
}

func (observer *Observer) receive() {
	for {
		event := <-observer.Recv
		observer.Handler.HandleEvent(event)
	}
}

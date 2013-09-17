package event

type MessageObserver struct {
	Recv       MessageChannel
	Observable MessageHandler
}

func NewMessageObserver(observable MessageHandler) *MessageObserver {
	observer := MessageObserver{Observable: observable}
	observer.Start()
	return &observer
}

func (observer *MessageObserver) Start() {
	observer.Recv = make(MessageChannel)
	go observer.receive()
}

func (observer *MessageObserver) receive() {
	for {
		msg := <-observer.Recv
		observer.Observable.HandleMessage(msg)
	}
}

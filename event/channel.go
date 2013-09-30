package event

type Channel chan *Notification

func (channel Channel) Send(message interface{}) {
	n, ok := message.(*Notification)
	if !ok {
		return
	}
	channel <- n
}

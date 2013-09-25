package event

type Subscriber struct {
  recv Channel
}

func NewSubscriber(recv Channel) *Subscriber {
  return &Subscriber{recv}
}

func (subscriber *Subscriber) Send(n *Notification) {
  subscriber.recv <- n
}

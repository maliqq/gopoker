package event

type Channel chan *Notification

func (channel Channel) Send(n *Notification) {
  channel <- n
}

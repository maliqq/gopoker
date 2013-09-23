package hub

type Subscriber struct {
	Key  EndpointKey
	Recv *chan interface{}
}

func (s *Subscriber) Send(message interface{}) {
	*s.Recv <- message
}

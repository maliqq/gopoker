package hub

type Notification struct {
	Route   ExchangeRoute
	Message interface{}
}

type Notify struct {
	Exchange *Exchange
	Message  interface{}
}

type key interface {
	String() string
}

func (n Notify) All() {
	n.Exchange.Dispatch(Notification{
		Route:   ExchangeRoute{All: true},
		Message: n.Message,
	})
}

func (n Notify) One(key key) {
	n.Exchange.Dispatch(Notification{
		Route:   ExchangeRoute{One: EndpointKey(key.String())},
		Message: n.Message,
	})
}

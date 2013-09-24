package hub

type Endpoint interface {
	Send(interface{})
}

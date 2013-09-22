package hub

type Broker struct {
  *Exchange
}

func NewBroker() *Broker{
  broker := Broker{
    Exchange: NewExchange(),
  }
  return &broker
}

func (broker *Broker) Notify(message interface{}) Notify {
  return Notify{
    Exchange: broker.Exchange,
    Message: message,
  }
}

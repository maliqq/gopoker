package hub

type ExchangeRoute struct {
  None bool
  All bool
  One EndpointKey
  Only []EndpointKey
  Except []EndpointKey
  Where func(Endpoint) bool
}

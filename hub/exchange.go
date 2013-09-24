package hub

import (
	_ "log"
)

import (
	"gopoker/event"
	_ "gopoker/util"
)

type Exchange struct {
	endpoints map[string]Endpoint
}

func NewExchange() *Exchange {
	return &Exchange{
		endpoints: map[string]Endpoint{},
	}
}

func (exchange *Exchange) Subscribe(key string, channel event.Channel) {
	exchange.endpoints[key] = Subscriber{channel}
}

func (exchange *Exchange) Dispatch(route Route, message interface{}) {
	//log.Printf(util.Colorf(util.Cyan, "[dispatch] %s %#v", route, message))

	if route.None {
		return
	}

	if route.One != "" {
		if endpoint, found := exchange.endpoints[route.One]; found {
			endpoint.Send(message)
		}
		return
	}

	for key, endpoint := range exchange.endpoints {
		if !route.All {
			var skip bool

			if len(route.Only) != 0 {
				skip = true
				for _, Only := range route.Only {
					if Only == key {
						skip = false
						break
					}
				}
			}

			if len(route.Except) != 0 {
				skip = false
				for _, Except := range route.Except {
					if Except == key {
						skip = true
						break
					}
				}
			}

			if skip {
				continue
			}
		}

		if route.Where != nil && !route.Where(endpoint) {
			continue
		}

		endpoint.Send(message)
	}
}

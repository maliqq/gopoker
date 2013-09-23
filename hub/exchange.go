package hub

import (
	"fmt"
)

type Exchange struct {
	endpoints map[EndpointKey]Endpoint
}

func NewExchange() *Exchange {
	exchange := Exchange{
		endpoints: map[EndpointKey]Endpoint{},
	}

	return &exchange
}

func (exchange *Exchange) Dispatch(notification Notification) {
	route := notification.Route

	if route.None {
		return
	}

	if route.One != "" {
		if endpoint, found := exchange.endpoints[route.One]; found {
			fmt.Printf("%s", endpoint)
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

		//
	}
}

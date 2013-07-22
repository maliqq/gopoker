package protocol

import (
	"fmt"
	"log"
)

import (
	"gopoker/util/console"
)

type MessageChannel chan *Message

// pubsub
type Broker struct {
	Send map[string]*MessageChannel
}

func NewBroker() *Broker {
	return &Broker{
		Send: map[string]*MessageChannel{},
	}
}

func (broker *Broker) Bind(key string, channel *MessageChannel) {
	broker.Send[key] = channel
}

func (broker *Broker) For(key string) *MessageChannel {
	ch, found := broker.Send[key]

	if !found {
		// ...
	}

	return ch
}

func (broker *Broker) Unbind(key string) {
	delete(broker.Send, key)
}

// route
type Notify struct {
	All    bool
	None   bool
	One    string
	Only   []string
	Except []string
}

func (n *Notify) RouteType() string {
	if n.All {
		return "all"
	}
	if n.None {
		return "none"
	}
	if n.One != "" {
		return "one"
	}
	if len(n.Except) != 0 {
		return "except"
	}
	if len(n.Only) != 0 {
		return "only"
	}

	return ""
}

func (n *Notify) String() string {
	s := fmt.Sprintf("[notify] %s", n.RouteType())

	if n.One != "" {
		s += fmt.Sprintf(": %v", n.One)
	}

	if len(n.Except) != 0 {
		s += fmt.Sprintf(": %v", n.Except)
	}

	if len(n.Only) != 0 {
		s += fmt.Sprintf(": %v", n.Only)
	}

	return s
}

func (broker *Broker) send(key string, msg *Message) {
	ch, found := broker.Send[key]
	if found {
		*ch <- msg
	}
}

func (broker *Broker) Dispatch(n *Notify, msg *Message) {
	log.Println(console.Color(console.CYAN, fmt.Sprintf("%s %s", n, msg)))

	if n.None {
		return
	}

	if n.One != "" {
		broker.send(n.One, msg)

		return
	}

	for key, _ := range broker.Send {
		if !n.All {
			var skip bool

			if len(n.Only) != 0 {
				skip = true
				for _, Only := range n.Only {
					if Only == key {
						skip = false
						break
					}
				}
			}

			if len(n.Except) != 0 {
				skip = false
				for _, Except := range n.Except {
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

		broker.send(key, msg)
	}
}

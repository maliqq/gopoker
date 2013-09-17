package event

import (
	"fmt"
	"log"
	"time"
)

import (
	_ "code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/event/message"
	"gopoker/util/console"
)

// MessageChannel - channel of messages
type MessageChannel chan *message.Message

type Group string

const (
	Observer Group = "observer"
	Watcher  Group = "watcher"
	Player   Group = "player"
)

// Broker - pubsub broker
type Broker struct {
	exchange map[Group]map[string]*MessageChannel
}

// Notify - route
type Notify struct {
	All    bool
	None   bool
	Group  Group
	One    string
	Only   []string
	Except []string
}

// NewBroker - create new broker
func NewBroker() *Broker {
	return &Broker{
		exchange: map[Group]map[string]*MessageChannel{
			Observer: map[string]*MessageChannel{},
			Watcher:  map[string]*MessageChannel{},
			Player:   map[string]*MessageChannel{},
		},
	}
}

// Bind - bind receiver
func (broker *Broker) Bind(group Group, key string, receiver *MessageChannel) {
	broker.exchange[group][key] = receiver
}

// Unbind - unbind receiver
func (broker *Broker) Unbind(group Group, key string) {
	delete(broker.exchange[group], key)
}

// RouteType - route type to string
func (n *Notify) RouteType() string {
	if n.None {
		return "none"
	}
	if n.One != "" {
		return "one"
	}
	if n.Group != "" {
		return string(n.Group)
	}
	if len(n.Except) != 0 {
		return "except"
	}
	if len(n.Only) != 0 {
		return "only"
	}
	if n.All {
		return "all"
	}

	return ""
}

// String - route to string
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

func (broker *Broker) send(group Group, key string, msg *message.Message) {
	if receiver, ok := broker.exchange[group][key]; ok {
		// sign message with timestamp
		if msg.Timestamp == 0 {
			msg.Timestamp = time.Now().Unix()
		}
		//
		*receiver <- msg
	}
}

// Dispatch - dispatch message
func (broker *Broker) Dispatch(n *Notify, msg *message.Message) {
	log.Println(console.Color(console.Cyan, fmt.Sprintf("%s %s", n, msg)))

	// notify observers
	defer broker.dispatchGroup(Observer, n, msg)
	if n.None {
		return
	}

	// notify one from group
	if n.One != "" && n.Group != "" {
		broker.send(n.Group, n.One, msg)
		return
	}

	// notify all group
	if n.Group != "" {
		broker.dispatchGroup(n.Group, n, msg)

		return
	}

	// notify all
	if n.All {
		broker.dispatchGroup(Player, n, msg)
		broker.dispatchGroup(Watcher, n, msg)
	}
}

func (broker *Broker) dispatchGroup(group Group, n *Notify, msg *message.Message) {
	for key := range broker.exchange[group] {
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

		broker.send(group, key, msg)
	}
}

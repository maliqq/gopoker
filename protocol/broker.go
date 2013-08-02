package protocol

import (
	"fmt"
	"log"
	"time"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/protocol/message"
	"gopoker/util/console"
)

type MessageChannel chan *message.Message

// pubsub
type Broker struct {
	User   map[string]*MessageChannel
	System map[string]*MessageChannel
}

// route
type Notify struct {
	All    bool
	None   bool
	One    string
	Only   []string
	Except []string
}

type System string

func (s System) RouteKey() string {
	return string(s)
}

func NewBroker() *Broker {
	return &Broker{
		User:   map[string]*MessageChannel{},
		System: map[string]*MessageChannel{},
	}
}

func (broker *Broker) BindUser(key string, channel *MessageChannel) {
	broker.User[key] = channel
}

func (broker *Broker) UnbindUser(key string) {
	delete(broker.User, key)
}

func (broker *Broker) BindSystem(key string, channel *MessageChannel) {
	broker.System[key] = channel
}

func (broker *Broker) UnbindSystem(key string) {
	delete(broker.System, key)
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

func (broker *Broker) sendUser(key string, msg *message.Message) {
	// sign message with timestamp
	if msg.GetTimestamp() == 0 {
		msg.Timestamp = proto.Int64(time.Now().Unix())
	}

	user, found := broker.User[key]
	if found {
		*user <- msg
	}
}

func (broker *Broker) sendSystem(msg *message.Message) {
	for _, system := range broker.System {
		*system <- msg
	}
}

func (broker *Broker) Dispatch(n *Notify, msg *message.Message) {
	log.Println(console.Color(console.Cyan, fmt.Sprintf("%s %s", n, msg)))

	defer broker.sendSystem(msg)

	if n.None {
		return
	}

	if n.One != "" {
		broker.sendUser(n.One, msg)

		return
	}

	for key, _ := range broker.User {
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

		broker.sendUser(key, msg)
	}
}

package event

import (
	"fmt"
	"log"
)

import (
	"gopoker/util"
)

// Channel - channel of messages

type Topic string

const (
	Default Topic = "default"
)

type Role string

const (
	System Role = "system"
	User   Role = "user"
)

type Subscriber struct {
	Key     string
	Role    Role
	Channel *Channel
}

// Broker - pubsub broker
type Broker struct {
	exchange map[Topic]map[string]Subscriber
}

// Notify - route
type Notify struct {
	All    bool
	None   bool
	Topic  Topic
	Role   Role
	One    string
	Only   []string
	Except []string
}

// NewBroker - create new broker
func NewBroker() *Broker {
	exchange := map[Topic]map[string]Subscriber{
		Default: map[string]Subscriber{},
	}
	return &Broker{exchange}
}

// Bind - bind receiver
func (broker *Broker) Bind(topic Topic, subscriber Subscriber) {
	broker.exchange[topic][subscriber.Key] = subscriber
}

// Unbind - unbind receiver
func (broker *Broker) Unbind(topic Topic, key string) {
	delete(broker.exchange[topic], key)
}

// RouteType - route type to string
func (n *Notify) RouteType() string {
	if n.None {
		return "none"
	}
	if n.One != "" {
		return "one"
	}
	if n.Topic != "" {
		return string(n.Topic)
	}
	if n.Role != "" {
		return string(n.Role)
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

func (broker *Broker) send(topic Topic, key string, event *Event) {
	if subscriber, ok := broker.exchange[topic][key]; ok {
		*subscriber.Channel <- event
	}
}

// Dispatch - dispatch message
func (broker *Broker) Dispatch(n Notify, event *Event) {
	log.Println(util.Color(util.Cyan, fmt.Sprintf("%s %s", n, event)))

	if n.None {
		return
	}

	// notify one from topic
	if n.One != "" && n.Topic != "" {
		broker.send(n.Topic, n.One, event)
		return
	}

	// notify all topic
	if n.Topic != "" {
		broker.dispatchTopic(n.Topic, n, event)
		return
	}

	// notify all
	if n.All {
		for topic, _ := range broker.exchange {
			broker.dispatchTopic(topic, n, event)
		}
	}
}

func (broker *Broker) dispatchTopic(topic Topic, n Notify, event *Event) {
	for key, subscriber := range broker.exchange[topic] {
		role := subscriber.Role

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

		if n.Role != "" && n.Role != role {
			continue
		}

		broker.send(topic, key, event)
	}
}

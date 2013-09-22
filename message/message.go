package message

import (
	"fmt"
	"reflect"
)

// Event Message
type Message interface {
	EventMessage()
}

// type register
var types = map[string]Message{}

func register(name string, instance Message) {
	types[name] = instance
}

func TypeFor(instance Message) string {
	typeName := reflect.Indirect(reflect.ValueOf(instance)).Type().Name()

	if typeName == "" {
		fmt.Printf("msg: %#v", instance)
		panic("unknown message type")
	}

	return typeName
}

func InstanceFor(name string) Message {
	if instance, found := types[name]; found {
		return instance
	}

	fmt.Printf("type: %s", name)
	panic("unknown message type")
}

func init() {
	register("AddBet", &AddBet{})
	register("RequireBet", &RequireBet{})
	register("BettingComplete", &BettingComplete{})

	register("DealCards", &DealCards{})
	register("RequireDiscard", &RequireDiscard{})
	register("Discarded", &Discarded{})
	register("DiscardCards", &DiscardCards{})

	register("PlayStart", &PlayStart{})
	register("StreetStart", &StreetStart{})
	register("PlayStop", &PlayStop{})

	register("ShowHand", &ShowHand{})
	register("ShowCards", &ShowCards{})
	register("Winner", &Winner{})

	register("MoveButton", &MoveButton{})
	register("Join", &Join{})
	register("SitOut", &SitOut{})
	register("ComeBack", &ComeBack{})
	register("Leave", &Leave{})

	register("ErrorMessage", &ErrorMessage{})
	register("ChatMessage", &ChatMessage{})
}

// ErrorMessage
type ErrorMessage struct {
	Error error
}

func (ErrorMessage) EventMessage() {}

// ChatMessage
type ChatMessage struct {
	Body string
}

func (ChatMessage) EventMessage() {}

package message

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

const (
	UseIndent = false
)

type Payload interface{}

func NewMessage(payload Payload) *Message {
	payloadType := reflect.TypeOf(payload)
	typeName := payloadType.Name()

	if typeName == "" {
		fmt.Printf("payload: %#v", payload)

		panic("unknown message type")
	}

	envelope := Envelope{}
	/*
		value := reflect.ValueOf(envelope)
		field := value.FieldByName(typeName)
		field.SetPointer(&payload)
	*/
	switch v := payload.(type) {
	case ErrorMessage:
		envelope.ErrorMessage = &v

	case PlayStart:
		envelope.PlayStart = &v

	case PlayStop:
		envelope.PlayStop = &v

	case StreetStart:
		envelope.StreetStart = &v

	case BettingComplete:
		envelope.BettingComplete = &v

	case JoinTable:
		envelope.JoinTable = &v

	case LeaveTable:
		envelope.LeaveTable = &v

	case SitOut:
		envelope.SitOut = &v

	case ComeBack:
		envelope.ComeBack = &v

	case MoveButton:
		envelope.MoveButton = &v

	case RequireBet:
		envelope.RequireBet = &v

	case AddBet:
		envelope.AddBet = &v

	case DealCards:
		envelope.DealCards = &v

	case RequireDiscard:
		envelope.RequireDiscard = &v

	case Discarded:
		envelope.Discarded = &v

	case DiscardCards:
		envelope.DiscardCards = &v

	case ShowHand:
		envelope.ShowHand = &v

	case ShowCards:
		envelope.ShowCards = &v

	case Winner:
		envelope.Winner = &v

	default:
		log.Fatalf("[protocol] can't handle %#v", payload)
	}

	return &Message{
		Type:      proto.String(typeName),
		Timestamp: proto.Int64(time.Now().Unix()),
		Envelope:  &envelope,
	}
}

func (msg *Message) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	data["Type"] = msg.GetType()
	data["Timestamp"] = msg.GetTimestamp()
	// cleanup fields
	value := reflect.ValueOf(*msg.Envelope)
	field := value.FieldByName(msg.GetType())
	data["Envelope"] = map[string]interface{}{
		*msg.Type: field.Interface(),
	}
	return json.Marshal(data)
}

func (msg *Message) Payload() Payload {
	value := reflect.ValueOf(msg.Envelope)
	method := value.MethodByName("Get" + msg.GetType())

	if method.IsValid() {
		result := method.Call([]reflect.Value{})
		return result[0].Interface()
	}

	log.Printf("[protocol] Got nil value on %#v", msg)

	return nil
}

func (msg *Message) PrintString() string {
	var err error
	var s []byte
	if UseIndent {
		s, err = json.MarshalIndent(msg, "", "\t")
	} else {
		s, err = json.Marshal(msg)
	}

	if err != nil {
		log.Printf("Message: %#v\n", msg)
		log.Printf("Error: %s\n", err)

		panic("error marshaling message")
	}

	return string(s)
}

func NewErrorMessage(err error) *Message {
	return NewMessage(ErrorMessage{
		Message: proto.String(err.Error()),
	})
}

func NewChatMessage(body string) *Message {
	return NewMessage(ChatMessage{
		Message: proto.String(body),
	})
}

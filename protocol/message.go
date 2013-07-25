package protocol

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

const (
	UseIndent = false
)

type Payload interface{}

type Envelope struct {
	// notice
	ErrorMessage  *ErrorMessage
	ChatMessage   *ChatMessage
	DealerMessage *DealerMessage

	// gameplay
	PlayStart   *PlayStart
	StreetStart *StreetStart
	ChangeGame  *ChangeGame

	// table
	JoinTable  *JoinTable
	LeaveTable *LeaveTable
	SitOut     *SitOut
	ComeBack   *ComeBack
	MoveButton *MoveButton

	// betting
	RequireBet *RequireBet
	AddBet     *AddBet

	// dealing
	DealCards      *DealCards
	RequireDiscard *RequireDiscard
	Discarded      *Discarded
	DiscardCards   *DiscardCards

	// showdown
	ShowHand  *ShowHand
	ShowCards *ShowCards
	Winner    *Winner
}

type Message struct {
	Type      string
	Timestamp int64
	Notify    Notify
	Envelope  Envelope
}

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
	case PlayStart:
		envelope.PlayStart = &v

	case StreetStart:
		envelope.StreetStart = &v

	case ChangeGame:
		envelope.ChangeGame = &v

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
		Type:      typeName,
		Timestamp: time.Now().Unix(),
		Envelope:  envelope,
	}
}

func (msg *Message) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	data["Type"] = msg.Type
	data["Timestamp"] = msg.Timestamp
	data["Notify"] = msg.Notify // FIXME
	// cleanup fields
	value := reflect.ValueOf(msg.Envelope)
	field := value.FieldByName(msg.Type)
	data["Envelope"] = map[string]interface{}{
		msg.Type: field.Interface(),
	}
	return json.Marshal(data)
}

func (msg *Message) Payload() Payload {
	value := reflect.ValueOf(msg.Envelope)
	field := value.FieldByName(msg.Type)
	return reflect.Indirect(field).Interface()
}

func (msg *Message) String() string {
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

// error
type ErrorMessage struct {
	Message string
}

// chat
type ChatMessage struct {
	Pos     int
	Message string
}

type DealerMessage struct {
	Message string
}

func NewErrorMessage(err error) *Message {
	return NewMessage(ErrorMessage{
		Message: err.Error(),
	})
}

func NewChatMessage(body string) *Message {
	return NewMessage(ChatMessage{
		Message: body,
	})
}

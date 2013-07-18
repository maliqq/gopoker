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

type Message struct {
	Type      string
	Timestamp int64
	Payload   Payload
}

func NewMessage(payload Payload) *Message {
	typeName := reflect.TypeOf(payload).Name()

	if typeName == "" {
		fmt.Printf("payload: %#v", payload)

		panic("unknown message type")
	}

	return &Message{
		Type:      typeName,
		Timestamp: time.Now().Unix(),
		Payload:   payload,
	}
}

func (msg *Message) UnmarshalJSON(data []byte) error {
	var raw map[string]*json.RawMessage
	err := json.Unmarshal(data, &raw)

	// Type
	var typeName string
	err = json.Unmarshal(*raw["Type"], &typeName)
	msg.Type = typeName

	// Timestamp
	var timestamp int64
	err = json.Unmarshal(*raw["Timestamp"], &timestamp)
	msg.Timestamp = timestamp

	// Payload
	switch msg.Type {
	case "JoinTable":
		var join JoinTable
		err = json.Unmarshal(*raw["Payload"], &join)
		msg.Payload = join
	}

	return err
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

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
	name := reflect.TypeOf(payload).Name()

	if name == "" {
		fmt.Printf("payload: %#v", payload)

		panic("unknown message type")
	}

	return &Message{
		Type:      name,
		Timestamp: time.Now().Unix(),
		Payload:   payload,
	}
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

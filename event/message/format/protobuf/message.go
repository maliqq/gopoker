package protobuf

/*
import (
	"fmt"
	"log"
	"reflect"
	"time"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

NewMessage - create new message with payload
func NewMessage(payload interface{}) *Message {
	payloadType := reflect.TypeOf(payload)
	typeName := payloadType.Name()

	if typeName == "" {
		fmt.Printf("payload: %#v", payload)

		panic("unknown message type")
	}

	return &Message{
		Event:      proto.String(typeName),
		Timestamp: proto.Int64(time.Now().Unix()),
		Payload:  &payload,
	}
}

// Payload - get message payload
func (msg *Message) Payload2() Payload {
	value := reflect.ValueOf(msg.Payload)
	method := value.MethodByName("Get" + msg.GetType())

	if method.IsValid() {
		result := method.Call([]reflect.Value{})
		return result[0].Interface()
	}

	log.Printf("[protocol] Got nil value on %#v", msg)

	return nil
}
*/

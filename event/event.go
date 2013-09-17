package event

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

import (
	"gopoker/event/message"
	"gopoker/event/message/format/protobuf"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

type Event struct {
	Type      string          `json:_type`
	Timestamp int64           `json:timestamp`
	Message   message.Message `json:message`
}

type Channel chan *Event

func getType(msg message.Message) string {
	msgType := reflect.TypeOf(msg)
	typeName := msgType.Name()

	if typeName == "" {
		fmt.Printf("msg: %#v", msg)
		panic("unknown message type")
	}

	return typeName
}

func NewEvent(msg message.Message) *Event {
	event := Event{
		Type:      getType(msg),
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}

	return &event
}

func (event *Event) UnmarshalJSON(data []byte) error {
	var raw map[string]*json.RawMessage

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	json.Unmarshal(*raw["_type"], &event.Type)
	json.Unmarshal(*raw["timestamp"], &event.Timestamp)
	/*
		msg := message.Type(msgType)
		json.Unmarshal(*raw["message"], &msg)
	*/
	return nil
}

func (event *Event) UnmarshalProto(data []byte) error {
	return nil
}

func (e *Event) String() string {
	return fmt.Sprintf("%#v", e)
}

func (e *Event) Proto() *protobuf.Event {
	return &protobuf.Event{
		Type:      proto.String(e.Type),
		Timestamp: proto.Int64(e.Timestamp),
		Message:   e.Message.Proto(),
	}
}

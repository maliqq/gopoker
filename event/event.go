package event

import (
	"encoding/json"
	"fmt"
	"time"
)

import (
	"gopoker/event/message"
	"gopoker/event/message/protobuf"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

type Event struct {
	Type      string          `json:_type bson:_type`
	Timestamp int64           `json:timestamp bson:timestamp`
	Message   message.Message `json:message bson:message`
}

type Channel chan *Event

func NewEvent(msg message.Message) *Event {
	event := Event{
		Type:      message.TypeFor(msg),
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

	event.Message = message.InstanceFor(event.Type)
	json.Unmarshal(*raw["message"], &event.Message)

	return nil
}

func (event *Event) Unproto(data []byte) error {
	raw := &protobuf.Event{}
	err := proto.Unmarshal(data, raw)
	if err != nil {
		return err
	}

	event.Timestamp = raw.GetTimestamp()
	event.Type = raw.GetType()
	event.Message = message.InstanceFor(event.Type)
	event.Message.Unproto(raw.Message)

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

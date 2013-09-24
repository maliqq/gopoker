package event

import (
	"encoding/json"
	"fmt"
	"time"
)

import (
	"gopoker/message"
)

type Notification struct {
	Type      string          `json:_type bson:_type`
	Timestamp int64           `json:timestamp bson:timestamp`
	Message   message.Message `json:message bson:message`
}

func New(msg message.Message) *Notification {
	n := &Notification{
		Type:      message.TypeFor(msg),
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}
	
	return n
}

func (n *Notification) UnmarshalJSON(data []byte) error {
	var raw map[string]*json.RawMessage

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	json.Unmarshal(*raw["_type"], &n.Type)
	json.Unmarshal(*raw["timestamp"], &n.Timestamp)

	n.Message = message.InstanceFor(n.Type)
	json.Unmarshal(*raw["message"], &n.Message)

	return nil
}

func (n *Notification) String() string {
	return fmt.Sprintf("[%s %d %#v]", n.Type, n.Timestamp, n.Message)
}

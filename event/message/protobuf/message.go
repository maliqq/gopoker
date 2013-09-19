package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

func NewErrorMessage(err error) *Message {
	return &Message{
		ErrorMessage: &ErrorMessage{
			Error: proto.String(err.Error()),
		},
	}
}

func NewChatMessage(body string) *Message {
	return &Message{
		ChatMessage: &ChatMessage{
			Body: proto.String(body),
		},
	}
}

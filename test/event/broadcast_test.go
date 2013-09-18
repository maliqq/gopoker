package event

import (
	"testing"
)

import (
	"gopoker/event/message"
)

func TestBroadcast(t *testing.T) {
	bcast := NewBroadcast()
	bcast.Notify(message.MoveButton{1}).All()
}

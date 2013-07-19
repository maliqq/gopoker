package protocol

import (
	"testing"
)

func TestMessage(t *testing.T) {
	payload := JoinTable{Pos: 1, Amount: 30.0}
	msg := NewMessage(payload)

	t.Logf("msg=%s", msg)
}

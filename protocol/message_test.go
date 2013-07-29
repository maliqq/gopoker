package protocol

import (
	"encoding/json"
	"testing"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

func TestMessage(t *testing.T) {
	payload := JoinTable{Pos: proto.Int32(1), Amount: proto.Float64(30.0)}
	msg := NewMessage(payload)

	s, err := json.Marshal(msg)
	t.Logf("msg=%s err=%s", s, err)
}

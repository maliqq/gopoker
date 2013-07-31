package bet

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/protocol/message"
)

type Range struct {
	Call float64
	Min  float64
	Max  float64
}

func (r Range) Proto() *message.BetRange {
	return &message.BetRange{
		Call: proto.Float64(r.Call),
		Min:  proto.Float64(r.Min),
		Max:  proto.Float64(r.Max),
	}
}

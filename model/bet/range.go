package bet

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/protocol/message"
)

// Range - bet range, how much to call and how much to raise
type Range struct {
	Call float64
	Min  float64
	Max  float64
}

// Reset - reset call and raise range
func (r *Range) Reset() {
	r.Call = 0.
	r.ResetRaise()
}

// ResetRaise - set min and max to 0.0
func (r *Range) ResetRaise() {
	r.Min, r.Max = 0., 0.
}

// SetRaise - set raise range according to available stack
func (r *Range) SetRaise(available float64, min float64, max float64) {
	minRaise, maxRaise := r.Call+min, r.Call+max

	if available < maxRaise {
		if available < r.Call {
			minRaise, maxRaise = 0., 0.
		} else if available < minRaise {
			minRaise, maxRaise = available, available
		} else {
			maxRaise = available
		}
	}

	r.Min, r.Max = minRaise, maxRaise
}

// Proto - range to protobuf
func (r *Range) Proto() *message.BetRange {
	return &message.BetRange{
		Call: proto.Float64(r.Call),
		Min:  proto.Float64(r.Min),
		Max:  proto.Float64(r.Max),
	}
}

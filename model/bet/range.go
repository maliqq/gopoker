package bet

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/event/message/protobuf"
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

func (r *Range) SetRaise(min, max float64) {
	r.Min, r.Max = r.Call+min, r.Call+max
}

// AdjustByAvailable - set raise range according to available stack
func (r *Range) AdjustByAvailable(available float64) {
	minRaise, maxRaise := r.Min, r.Max

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
func (r *Range) Proto() *protobuf.BetRange {
	return &protobuf.BetRange{
		Call: proto.Float64(r.Call),
		Min:  proto.Float64(r.Min),
		Max:  proto.Float64(r.Max),
	}
}

func (r *Range) UnmarshalProto(protoRange *protobuf.BetRange) {
	newRange := Range{
		Call: protoRange.GetCall(),
		Min:  protoRange.GetMin(),
		Max:  protoRange.GetMax(),
	}
	*r = newRange
}

package message

import (
	"code.google.com/p/goprotobuf/proto"
)

// NewAddBet - create new add bet
func NewAddBet(pos int, bet *Bet) *Message {
	return NewMessage(AddBet{
		Pos: proto.Int32(int32(pos)),
		Bet: bet,
	})
}

// NewRequireBet - create new require bet
func NewRequireBet(pos int, betRange *BetRange) *Message {
	return NewMessage(RequireBet{
		Pos:      proto.Int32(int32(pos)),
		BetRange: betRange,
	})
}

// NewBettingComplete - notify betting complete
func NewBettingComplete(total float64) *Message {
	return NewMessage(BettingComplete{
		Pot: proto.Float64(total),
	})
}

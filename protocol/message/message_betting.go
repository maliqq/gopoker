package message

import (
	"code.google.com/p/goprotobuf/proto"
)

// NotifyAddBet - create new add bet
func NotifyAddBet(pos int, bet *Bet) *Message {
	return NewMessage(AddBet{
		Pos: proto.Int32(int32(pos)),
		Bet: bet,
	})
}

// NotifyRequireBet - create new require bet
func NotifyRequireBet(pos int, betRange *BetRange) *Message {
	return NewMessage(RequireBet{
		Pos:      proto.Int32(int32(pos)),
		BetRange: betRange,
	})
}

// NotifyBettingComplete - notify betting complete
func NotifyBettingComplete(total float64) *Message {
	return NewMessage(BettingComplete{
		Pot: proto.Float64(total),
	})
}

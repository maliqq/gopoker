package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

// NotifyAddBet - create new add bet
func NewAddBet(pos int, bet *Bet) *Message {
	return &Message{
		Payload: &Payload{
			AddBet: &AddBet{
				Pos: proto.Int32(int32(pos)),
				Bet: bet,
			},
		},
	}
}

// NotifyRequireBet - create new require bet
func NewRequireBet(pos int, betRange *BetRange) *Message {
	return &Message{
		Payload: &Payload{
			RequireBet: &RequireBet{
				Pos:      proto.Int32(int32(pos)),
				BetRange: betRange,
			},
		},
	}
}

// NotifyBettingComplete - notify betting complete
func NewBettingComplete(total float64) *Message {
	return &Message{
		Payload: &Payload{
			BettingComplete: &BettingComplete{
				Pot: proto.Float64(total),
			},
		},
	}
}

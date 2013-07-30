package message

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model"
)

func NewAddBet(pos int, bet *model.Bet) *Message {
	return NewMessage(AddBet{
		Pos: proto.Int32(int32(pos)),
		Bet: &Bet{
			Type:   BetType(BetType_value[string(bet.Type)]).Enum(),
			Amount: proto.Float64(bet.Amount),
		},
	})
}

func NewRequireBet(pos int, betRange model.BetRange) *Message {
	return NewMessage(RequireBet{
		Pos: proto.Int32(int32(pos)),
		BetRange: &BetRange{
			Call: proto.Float64(betRange.Call),
			Min:  proto.Float64(betRange.Min),
			Max:  proto.Float64(betRange.Max),
		},
	})
}

func NewBettingComplete(pot *model.Pot) *Message {
	return NewMessage(BettingComplete{
		Pot: proto.Float64(pot.Total()),
	})
}

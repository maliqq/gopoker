package protocol

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

func NewBettingComplete(pot *model.Pot) *Message {
	return NewMessage(BettingComplete{
		Pot: proto.Float64(pot.Total()),
	})
}

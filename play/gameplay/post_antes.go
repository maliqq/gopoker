package gameplay

import (
	"gopoker/model/bet"
	"gopoker/protocol/message"
)

func (gp *GamePlay) PostAntes() {
	for _, pos := range gp.Table.AllSeats().Active() {
		seat := gp.Table.Seat(pos)

		newBet := gp.Betting.ForceBet(pos, bet.Ante, gp.Stake)

		gp.Betting.AddBet(seat, newBet)

		gp.Broadcast.All <- message.NewAddBet(pos, newBet.Proto())
	}
}

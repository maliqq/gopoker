package gameplay

import (
	"gopoker/model/bet"
	"gopoker/protocol/message"
)

// PostAntes - post antes
func (gp *GamePlay) PostAntes() {
	for _, pos := range gp.Table.AllSeats().Active() {
		seat := gp.Table.Seat(pos)

		newBet := gp.Betting.ForceBet(pos, seat, bet.Ante, gp.Stake)

		gp.Betting.AddBet(newBet)

		gp.Broadcast.All <- message.NewAddBet(pos, newBet.Proto())
	}
}

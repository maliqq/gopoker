package gameplay

import (
	"gopoker/event/message"
	"gopoker/model/bet"
)

// PostAntes - post antes
func (gp *GamePlay) PostAntes() {
	for _, pos := range gp.Table.AllSeats().Active() {
		seat := gp.Table.Seat(pos)

		newBet := gp.Betting.ForceBet(pos, seat, bet.Ante, gp.Stake)

		gp.Betting.AddBet(newBet)

		gp.Broadcast.All <- message.AddBet{pos, newBet}
	}
}

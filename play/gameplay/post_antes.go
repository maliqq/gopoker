package gameplay

import (
	"gopoker/model/bet"
	"gopoker/protocol/message"
)

func (this *GamePlay) PostAntes() {
	for _, pos := range this.Table.AllSeats().Active() {
		seat := this.Table.Seat(pos)

		newBet := this.Betting.ForceBet(pos, bet.Ante, this.Stake)

		this.Betting.AddBet(seat, newBet)

		this.Broadcast.All <- message.NewAddBet(pos, newBet.Proto())
	}
}

package gameplay

import (
	"gopoker/model/bet"
	"gopoker/protocol"
)

func (this *GamePlay) PostAntes() {
	for _, pos := range this.Table.SeatsInPlay() {
		seat := this.Table.Seat(pos)

		newBet := this.ForceBet(pos, bet.Ante, this.Stake)

		this.AddBet(seat, newBet)

		this.Broadcast.All <- protocol.NewAddBet(pos, newBet)
	}
}

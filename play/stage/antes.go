package stage

import (
	"gopoker/model/bet"
	"gopoker/protocol"
)

func (stage *Stage) postAntes() {
	play := stage.Play

	stake := play.Game.Stake

	for _, pos := range play.Table.SeatsInPlay() {
		seat := play.Table.Seat(pos)

		newBet := stage.Betting.ForceBet(pos, bet.Ante, stake)

		stage.Betting.AddBet(seat, newBet)
		
		play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
	}
}

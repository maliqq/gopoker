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

		newBet := &bet.Bet{
			Type:   bet.Ante,
			Amount: stake.AnteAmount(),
		}

		play.Broadcast.All <- protocol.NewAddBet(pos, newBet)

		stage.Betting.AddBet(seat, newBet)
	}
}

package stage

import (
	"gopoker/model/bet"
	"gopoker/play/context"
	"gopoker/protocol"
)

func postAntes(play *context.Play) {
	stake := play.Game.Stake

	for _, pos := range play.Table.SeatsInPlay() {
		seat := play.Table.Seat(pos)

		newBet := play.Betting.ForceBet(pos, bet.Ante, stake)

		play.Betting.AddBet(seat, newBet)

		play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
	}

	resetBetting(play)
}

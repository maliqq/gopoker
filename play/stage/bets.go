package stage

import (
	"fmt"
	"time"
)

import (
	"gopoker/protocol"
)

const (
	DefaultTimer = 30
)

func (stage *Stage) resetBetting() {
	play := stage.Play
	betting := stage.Betting

	betting.Reset()

	for _, pos := range play.Table.SeatsInPlay() {
		seat := play.Table.Seat(pos)

		seat.SetPlaying()
	}

	play.Broadcast.All <- protocol.NewPotSummary(betting.Pot)
}

func (stage *Stage) BettingRound() {
	play := stage.Play
	betting := stage.Betting

	for _, pos := range play.Table.Seats.From(betting.Current()).InPlay() {
		seat := play.Table.Seat(pos)

		play.Broadcast.One(seat.Player) <- betting.RequireBet(pos, seat, play.Game)

		select {
		case msg := <-play.Betting:
			betting.Add(seat, msg)

		case <-time.After(time.Duration(DefaultTimer) * time.Second):
			fmt.Printf("timeout!")
		}
	}

	stage.resetBetting()
}

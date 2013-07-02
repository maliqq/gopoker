package stage

import (
	"fmt"
	"time"
)

import (
	"gopoker/protocol"
)

const (
	DefaultTimer = 30 * time.Second
)

func (stage *Stage) resetBetting() {
	play := stage.Play
	betting := stage.Betting

	betting.Reset()
}

func (stage *Stage) BettingRound() {
	play := stage.Play
	betting := stage.Betting

	for _, pos := range play.Table.SeatsInPot() {
		seat := play.Table.Seat(pos)

		play.Broadcast.One(seat.Player) <- betting.RequireBet(pos, seat, play.Game)

		select {
		case msg := <-play.NextTurn:
			err := betting.Add(seat, msg)

			if err != nil {
				play.Broadcast.One(seat.Player) <- protocol.NewError(err)
			}

		case <-time.After(DefaultTimer):
			fmt.Printf("timeout!")
		}
	}

	stage.resetBetting()
}

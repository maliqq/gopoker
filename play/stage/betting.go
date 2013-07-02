package stage

import (
	"fmt"
	"time"
)

const (
	DefaultTimer = 30 * time.Second
)

func (stage *Stage) BettingRound() {
	play := stage.Play
	betting := stage.Betting

	betting.Reset()

	for _, pos := range play.Table.SeatsInPot() {
		seat := play.Table.Seat(pos)

		play.Broadcast.One(seat.Player) <- betting.RequireBet(pos, seat, play.Game)

		select {
		case msg := <-play.NextTurn:
			betting.Add(seat, msg)

		case <-time.After(DefaultTimer):
			fmt.Printf("timeout!")
		}
	}
}

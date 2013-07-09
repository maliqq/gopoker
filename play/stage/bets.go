package stage

import (
	"fmt"
	"log"
	"time"
)

import (
	"gopoker/play/context"
	"gopoker/protocol"
)

const (
	DefaultTimer = 30
)

func resetBetting(play *context.Play) {
	betting := play.Betting

	betting.Reset()

	for _, pos := range play.Table.SeatsInPlay() {
		seat := play.Table.Seat(pos)

		seat.SetPlaying()
	}

	play.Broadcast.All <- protocol.NewPotSummary(betting.Pot)
}

func BettingRound(play *context.Play) {
	log.Println("[play.stage] betting")

	betting := play.Betting

	for _, pos := range play.Table.Seats.From(betting.Current()).InPlay() {
		seat := play.Table.Seat(pos)

		play.Broadcast.One(seat.Player) <- betting.RequireBet(pos, seat, play.Game)

		select {
		case msg := <-betting.Receive:
			betting.Add(seat, msg)

		case <-time.After(time.Duration(DefaultTimer) * time.Second):
			fmt.Printf("timeout!")
		}
	}

	resetBetting(play)
}

func BigBets(play *context.Play) {
	log.Println("[play.stage] big bets")

	play.Betting.BigBets = true
}

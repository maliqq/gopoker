package stage

import (
	"log"
)

import (
	"gopoker/model/bet"
	"gopoker/play/context"
	"gopoker/protocol"
)

func postSmallBlind(play *context.Play, pos int) {
	stake := play.Game.Stake

	t := play.Table
	newBet := play.Betting.ForceBet(pos, bet.SmallBlind, stake)

	err := play.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", pos, err)
	}

	play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func postBigBlind(play *context.Play, pos int) {
	stake := play.Game.Stake

	t := play.Table
	newBet := play.Betting.ForceBet(pos, bet.BigBlind, stake)

	err := play.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding big blind for %d: %s", pos, err)
	}

	play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func postBlinds(play *context.Play) {
	moveButton(play)

	t := play.Table

	active := t.Seats.From(t.Button).Active()
	waiting := t.Seats.From(t.Button).Waiting()

	if len(active)+len(waiting) < 2 {
		log.Println("[play.stage.blinds] none waiting")

		return
	}

	//headsUp := len(active) == 2 && len(waiting) == 0 || len(active) == 1 && len(waiting) == 1

	sb := active[0]
	postSmallBlind(play, sb)

	bb := active[1]
	postBigBlind(play, bb)
}

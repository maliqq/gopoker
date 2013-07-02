package stage

import (
	"log"
)

import (
	"gopoker/model/bet"
	"gopoker/protocol"
)

func (stage *Stage) postSmallBlind(pos int) {
	play := stage.Play
	stake := play.Game.Stake

	t := play.Table
	newBet := stage.Betting.ForceBet(pos, bet.SmallBlind, stake)

	err := stage.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", pos, err)
	}

	play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (stage *Stage) postBigBlind(pos int) {
	play := stage.Play
	stake := play.Game.Stake

	t := play.Table
	newBet := stage.Betting.ForceBet(pos, bet.BigBlind, stake)

	err := stage.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding big blind for %d: %s", pos, err)
	}

	play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (stage *Stage) postBlinds() {
	stage.moveButton()

	play := stage.Play
	t := play.Table

	active := t.Seats.From(t.Button).Active()
	waiting := t.Seats.From(t.Button).Waiting()

	if len(active)+len(waiting) < 2 {
		log.Println("[play.stage.blinds] none waiting")

		return
	}

	//headsUp := len(active) == 2 && len(waiting) == 0 || len(active) == 1 && len(waiting) == 1

	sb := active[0]
	stage.postSmallBlind(sb)

	bb := active[1]
	stage.postBigBlind(bb)
}

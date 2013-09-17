package gameplay

import (
	"log"
)

import (
	"gopoker/event/message"
	"gopoker/model/bet"
)

// MoveButton - move table button
func (gp *GamePlay) MoveButton() {
	gp.Table.MoveButton()
	gp.Betting.Pos = gp.Table.Button

	gp.Broadcast.All <- message.MoveButton{gp.Table.Button}
}

func (gp *GamePlay) postSmallBlind(pos int) {
	t := gp.Table
	newBet := gp.Betting.ForceBet(pos, t.Seat(pos), bet.SmallBlind, gp.Stake)

	err := gp.Betting.AddBet(newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", pos, err)
	}

	gp.Broadcast.All <- message.AddBet{pos, newBet}
}

func (gp *GamePlay) postBigBlind(pos int) {
	t := gp.Table
	newBet := gp.Betting.ForceBet(pos, t.Seat(pos), bet.BigBlind, gp.Stake)

	err := gp.Betting.AddBet(newBet)
	if err != nil {
		log.Fatalf("Error adding big blind for %d: %s", pos, err)
	}

	gp.Broadcast.All <- message.AddBet{pos, newBet}
}

// PostBlinds - post blinds
func (gp *GamePlay) PostBlinds() {
	t := gp.Table

	active := t.AllSeats().Active()
	waiting := t.AllSeats().Waiting()

	if len(active)+len(waiting) < 2 {
		log.Println("[gp.stage.blinds] none waiting")

		return
	}
	//headsUp := len(active) == 2 && len(waiting) == 0 || len(active) == 1 && len(waiting) == 1

	sb := active[0]
	gp.postSmallBlind(sb)

	bb := active[1]
	gp.postBigBlind(bb)
}

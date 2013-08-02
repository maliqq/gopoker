package gameplay

import (
	"log"
)

import (
	"gopoker/model/bet"
	"gopoker/protocol/message"
)

func (gp *GamePlay) MoveButton() {
	gp.Table.MoveButton()

	gp.Broadcast.All <- message.NewMoveButton(gp.Table.Button)
}

func (gp *GamePlay) postSmallBlind(pos int) {
	t := gp.Table
	newBet := gp.Betting.ForceBet(pos, bet.SmallBlind, gp.Stake)

	err := gp.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", pos, err)
	}

	gp.Broadcast.All <- message.NewAddBet(pos, newBet.Proto())
}

func (gp *GamePlay) postBigBlind(pos int) {
	t := gp.Table
	newBet := gp.Betting.ForceBet(pos, bet.BigBlind, gp.Stake)

	err := gp.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding big blind for %d: %s", pos, err)
	}

	gp.Broadcast.All <- message.NewAddBet(pos, newBet.Proto())
}

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

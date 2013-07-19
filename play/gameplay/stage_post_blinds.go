package gameplay

import (
	"log"
)

import (
	"gopoker/model/bet"
	"gopoker/protocol"
)

func (this *GamePlay) MoveButton() {
	this.Table.MoveButton()

	this.Broadcast.All <- protocol.NewMoveButton(this.Table.Button)
}

func (this *GamePlay) postSmallBlind(pos int) {
	t := this.Table
	newBet := this.Betting.ForceBet(pos, t.Seat(pos), bet.SmallBlind, this.Stake)

	err := this.Betting.AddBet(newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", pos, err)
	}

	this.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (this *GamePlay) postBigBlind(pos int) {
	t := this.Table
	newBet := this.Betting.ForceBet(pos, t.Seat(pos), bet.BigBlind, this.Stake)

	err := this.Betting.AddBet(newBet)
	if err != nil {
		log.Fatalf("Error adding big blind for %d: %s", pos, err)
	}

	this.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (this *GamePlay) PostBlinds() {
	t := this.Table

	active := t.AllSeats().Active()
	waiting := t.AllSeats().Waiting()

	if len(active)+len(waiting) < 2 {
		log.Println("[this.stage.blinds] none waiting")

		return
	}
	//headsUp := len(active) == 2 && len(waiting) == 0 || len(active) == 1 && len(waiting) == 1

	sb := active[0]
	this.postSmallBlind(sb)

	bb := active[1]
	this.postBigBlind(bb)
}

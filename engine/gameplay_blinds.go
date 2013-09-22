package engine

import (
	"log"
)

import (
	"gopoker/message"
	"gopoker/model/bet"
)

func (g *Gameplay) moveButton() {
	g.Table.MoveButton()
	g.e.Notify(
		&message.MoveButton{g.Table.Button},
	).All()
}

func (g *Gameplay) setButton(pos int) {
	g.Table.SetButton(pos)
	g.e.Notify(
		&message.MoveButton{g.Table.Button},
	).All()
}

func (g *Gameplay) postSmallBlind() {
	newBet := g.b.ForceBet(bet.SmallBlind, g.Stake)

	err := g.b.AddBet(newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", g.b.Round.Pos(), err)
	}

	g.e.Notify(
		&message.AddBet{g.b.Round.Pos(), newBet},
	).All()
	g.b.Round.Move()
}

func (g *Gameplay) postBigBlind() {
	newBet := g.b.ForceBet(bet.BigBlind, g.Stake)

	err := g.b.AddBet(newBet)
	if err != nil {
		log.Fatalf("Error adding big blind for %d: %s", g.b.Round.Pos(), err)
	}

	g.e.Notify(
		&message.AddBet{g.b.Round.Pos(), newBet},
	).All()
	g.b.Round.Move()
}

func (g *Gameplay) postBlinds() {
	g.moveButton()

	ring := g.Table.Ring()

	active := ring.Active()
	waiting := ring.Waiting()

	if len(active)+len(waiting) < 2 {
		g.Stage.Stop() // stop stage
		g.d.Stop() // stop deal

		return
	}
	//headsUp := len(active) == 2 && len(waiting) == 0 || len(active) == 1 && len(waiting) == 1
	g.b.NewRound(active)

	g.postSmallBlind()
	g.postBigBlind()
}

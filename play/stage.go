package play

import (
	"log"
)

import (
	"gopoker/model/deal"
	"gopoker/play/gameplay"
)

// Stage - strategy invoking
type Stage struct {
	Name   string
	Invoke func(*Play) gameplay.Transition
}

var (
	// BringIn - bring in stage
	BringIn = Stage{"bring-in", func(play *Play) gameplay.Transition {
		log.Println("[play] bring in")

		return play.Gameplay.BringIn()
	}}

	// Betting - betting stage
	Betting = Stage{"betting", func(play *Play) gameplay.Transition {
		//log.Println("[play] betting")

		return play.Gameplay.StartBettingRound()
	}}

	// Discarding - discarding stage
	Discarding = Stage{"discarding", func(play *Play) gameplay.Transition {
		log.Println("[play] discarding")

		return play.Gameplay.StartDiscardingRound()
	}}

	// BigBets - big bets stage
	BigBets = Stage{"big-bets", func(play *Play) gameplay.Transition {
		log.Println("[play] big bets")

		return play.Gameplay.BigBets()
	}}
)

type dealing struct {
	deal.Type
	n int
}

func (d dealing) stage(play *Play) gameplay.Transition {
	n := d.n
	if d.n == 0 && d.Type == deal.Hole {
		n = play.Game.PocketSize
	}
	log.Printf("[play] dealing %s %d cards\n", d.Type, n)

	switch d.Type {
	case deal.Hole:
		play.Gameplay.DealHole(n)

	case deal.Door:
		play.Gameplay.DealDoor(n)

	case deal.Board:
		play.Gameplay.DealBoard(n)
	}

	return gameplay.Next
}

// Dealing - dealing stage
func Dealing(t deal.Type, n int) Stage {
	return Stage{"dealing", dealing{t, n}.stage}
}

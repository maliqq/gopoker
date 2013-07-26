package play

import (
	"log"
)

import (
	"gopoker/model/deal"
	"gopoker/play/gameplay"
)

/*
Strategy invoking
*/
type Stage struct {
	Name   string
	Invoke func(*Play) gameplay.Transition
}

var (
	BringIn = Stage{"bring-in", func(play *Play) gameplay.Transition {
		log.Println("[play] bring in")

		return play.GamePlay.BringIn()
	}}

	Betting = Stage{"betting", func(play *Play) gameplay.Transition {
		//log.Println("[play] betting")

		return play.GamePlay.StartBettingRound()
	}}

	Discarding = Stage{"discarding", func(play *Play) gameplay.Transition {
		log.Println("[play] discarding")

		return play.GamePlay.StartDiscardingRound()
	}}

	BigBets = Stage{"big-bets", func(play *Play) gameplay.Transition {
		log.Println("[play] big bets")

		return play.GamePlay.BigBets()
	}}
)

type dealing struct {
	deal.Type
	n int
}

func (d dealing) Stage(play *Play) gameplay.Transition {
	n := d.n
	if d.n == 0 && d.Type == deal.Hole {
		n = play.Game.PocketSize
	}
	log.Printf("[play] dealing %s %d cards\n", d.Type, n)

	switch d.Type {
	case deal.Hole:
		play.GamePlay.DealHole(n)

	case deal.Door:
		play.GamePlay.DealDoor(n)

	case deal.Board:
		play.GamePlay.DealBoard(n)
	}

	return gameplay.Next
}

func Dealing(t deal.Type, n int) Stage {
	return Stage{"dealing", dealing{t, n}.Stage}
}

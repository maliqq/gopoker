package play

import (
	"log"
)

import (
	"gopoker/model/deal"
	"gopoker/play/gameplay"
	"gopoker/play/street"
)

/*
Strategy invoking
*/
type Stage func(*Play)

func DealStart(play *Play) {
	play.StartNewDeal()
	play.ResetSeats()
	play.RotateGame()
}

func PostAntes(play *Play) {
	if play.Game.HasAnte || play.Stake.HasAnte() {
		log.Println("[play] post antes")

		play.PostAntes()
		play.ResetBetting()
	}
}

func PostBlinds(play *Play) {
	if play.Game.HasBlinds {
		log.Println("[play] post blinds")

		play.MoveButton()
		play.PostBlinds()
	}
}

func StartStreets(play *Play) {
	for _, street := range street.Get(play.Game.Group) {
		play.Street = street
		play.RunStreet()
	}
}

func BringIn(play *Play) {
	log.Println("[play] bring in")

	play.BringIn()
}

type dealing struct {
	deal.Type
	n int
}

func (d dealing) Stage(play *Play) {
	n := d.n
	if d.n == 0 && d.Type == deal.Hole {
		n = play.Game.PocketSize
	}
	log.Printf("[play] dealing %s %d cards\n", d.Type, n)

	switch d.Type {
	case deal.Hole:
		play.DealHole(n)

	case deal.Door:
		play.DealDoor(n)

	case deal.Board:
		play.DealBoard(n)
	}
}

func Dealing(t deal.Type, n int) Stage {
	return dealing{t, n}.Stage
}

func Betting(play *Play) {
	log.Println("[play] betting")

	play.StartBettingRound()
}

func Discarding(play *Play) {
	log.Println("[play] discarding")

	play.StartDiscardingRound()
}

func BigBets(play *Play) {
	log.Println("[play] big bets")

	play.Betting.BigBets = true
}

func Showdown(play *Play) {
	log.Println("[play] showdown")

	var highHands, lowHands gameplay.ShowdownHands

	if play.Game.Lo != "" {
		lowHands = play.ShowHands(play.Game.Lo, play.Game.HasBoard)
	}

	if play.Game.Hi != "" {
		highHands = play.ShowHands(play.Game.Hi, play.Game.HasBoard)
	}

	play.Winners(highHands, lowHands)
}

func DealStop(play *Play) {
	log.Println("[play] deal stop")
	
	play.ScheduleNextDeal()
}

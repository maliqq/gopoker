package play

import (
	"log"
)

import (
	"gopoker/model/deal"
	"gopoker/play/context"
	"gopoker/play/street"
)

/*
Strategy invoking
*/
type Stage func(*context.Play)

func Init(play *context.Play) {
	play.StartNewDeal()
	play.ResetSeats()
	play.RotateGame()
}

func PostAntes(play *context.Play) {
	gameOptions := play.Game.Options
	stake := play.Game.Stake

	if gameOptions.HasAnte || stake.HasAnte() {
		log.Println("[play] post antes")

		play.PostAntes()
		play.ResetBetting()
	}
}

func PostBlinds(play *context.Play) {
	gameOptions := play.Game.Options
	if gameOptions.HasBlinds {
		log.Println("[play] post blinds")

		play.MoveButton()
		play.PostBlinds()
	}
}

func StartStreets(play *context.Play) {
	for _, street := range street.Get(play.Game.Options.Group) {
		log.Printf("[play] %s\n", street)

		ByStreet[street].Proceed(play)
	}
}

func BringIn(play *context.Play) {
	log.Println("[play] bring in")

	play.BringIn()
}

type dealing struct {
	deal.Type
	n int
}

func (d dealing) Stage(play *context.Play) {
	n := d.n
	if d.n == 0 && d.Type == deal.Hole {
		n = play.Game.Options.Pocket
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

func Betting(play *context.Play) {
	log.Println("[play] betting")

	play.StartBettingRound()
	play.ResetBetting()
}

func Discarding(play *context.Play) {
	log.Println("[play] discarding")

	play.StartDiscardingRound()
}

func BigBets(play *context.Play) {
	log.Println("[play] big bets")

	play.Betting.BigBets = true
}

func Showdown(play *context.Play) {
	log.Println("[play] showdown")

	gameOptions := play.Game.Options

	var highHands, lowHands *context.ShowdownHands

	if gameOptions.Lo != "" {
		lowHands = play.ShowHands(gameOptions.Lo, gameOptions.HasBoard)
	}

	if gameOptions.Hi != "" {
		highHands = play.ShowHands(gameOptions.Hi, gameOptions.HasBoard)
	}

	play.Winners(highHands, lowHands)

	play.ScheduleNextDeal()
}

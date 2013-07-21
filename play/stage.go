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
type Stage struct {
	Name   string
	Invoke func(*Play)
}

var (
	DealStart = Stage{"deal-start", func(play *Play) {
		play.StartNewDeal()
		play.RotateGame()
		play.ResetSeats()
	}}

	PostAntes = Stage{"post-antes", func(play *Play) {
		if play.Game.HasAnte || play.Stake.HasAnte() {
			log.Println("[play] post antes")

			play.GamePlay.PostAntes()
			play.GamePlay.ResetBetting()
		}
	}}

	PostBlinds = Stage{"post-blinds", func(play *Play) {
		if play.Game.HasBlinds {
			log.Println("[play] post blinds")

			play.GamePlay.MoveButton()
			play.GamePlay.PostBlinds()
		}
	}}

	Streets = Stage{"streets", func(play *Play) {
		for _, street := range street.Get(play.Game.Group) {
			play.Street = street
			play.RunStreet()
		}
		play.Street = ""
	}}

	BringIn = Stage{"bring-in", func(play *Play) {
		log.Println("[play] bring in")

		play.GamePlay.BringIn()
	}}

	Betting = Stage{"betting", func(play *Play) {
		log.Println("[play] betting")

		play.GamePlay.StartBettingRound()
	}}

	Discarding = Stage{"discarding", func(play *Play) {
		log.Println("[play] discarding")

		play.GamePlay.StartDiscardingRound()
	}}

	BigBets = Stage{"big-bets", func(play *Play) {
		log.Println("[play] big bets")

		play.Betting.BigBets = true
	}}

	Showdown = Stage{"showdown", func(play *Play) {
		log.Println("[play] showdown")

		var highHands, lowHands gameplay.ShowdownHands

		if play.Game.Lo != "" {
			lowHands = play.GamePlay.ShowHands(play.Game.Lo, play.Game.HasBoard)
		}

		if play.Game.Hi != "" {
			highHands = play.GamePlay.ShowHands(play.Game.Hi, play.Game.HasBoard)
		}

		play.GamePlay.Winners(highHands, lowHands)
	}}

	DealStop = Stage{"deal-stop", func(play *Play) {
		log.Println("[play] deal stop")

		play.ScheduleNextDeal()
	}}
)

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
		play.GamePlay.DealHole(n)

	case deal.Door:
		play.GamePlay.DealDoor(n)

	case deal.Board:
		play.GamePlay.DealBoard(n)
	}
}

func Dealing(t deal.Type, n int) Stage {
	return Stage{"dealing", dealing{t, n}.Stage}
}

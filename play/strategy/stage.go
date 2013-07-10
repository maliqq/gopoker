package strategy

import (
	"log"
)

import (
	"gopoker/model/deal"
	"gopoker/play/context"
)

type Stage func(*context.Play)

func ResetSeats(play *context.Play) {
	play.ResetSeats()
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
	streets, _ := Streets[play.Game.Options.Group]

	for _, street := range streets {
		log.Printf("[play] %s\n", street)

		StreetStrategies[street].Proceed(play)
	}
}

func BringIn(play *context.Play) {
	log.Println("[play] bring in")

	play.BringIn()
}

func Dealing(dealingType deal.Type, cardsNum int) Stage {
	return func(play *context.Play) {
		if cardsNum == 0 && dealingType == deal.Hole {
			cardsNum = play.Game.Options.Pocket
		}

		log.Printf("[play] dealing %s %d cards\n", dealingType, cardsNum)

		switch dealingType {
		case deal.Hole:
			play.DealHole(cardsNum)

		case deal.Door:
			play.DealDoor(cardsNum)

		case deal.Board:
			play.DealBoard(cardsNum)
		}
	}
}

func Betting(play *context.Play) {
	log.Println("[play] betting")

	play.StartBettingRound()
	play.ResetBetting()
}

func Discarding(play *context.Play) {
	log.Println("[play] discarding")

	log.Fatalf("not implemented")
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

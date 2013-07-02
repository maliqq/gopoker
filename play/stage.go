package play

import (
	"log"
)

import (
	"gopoker/model/deal"
	"gopoker/play/stage"
)

type Stage func(play *stage.Stage)

var (
	BringIn Stage = func(stage *stage.Stage) {
		log.Println("[play.stage] bring in")

		stage.BringIn()
	}

	BigBets Stage = func(stage *stage.Stage) {
		log.Println("[play.stage] big bets")

		stage.Betting.BigBets = true
	}

	BettingRound Stage = func(stage *stage.Stage) {
		log.Println("[play.stage] betting")

		stage.BettingRound()
	}

	DiscardingRound Stage = func(stage *stage.Stage) {
		log.Println("[play.stage] discarding")

		stage.DiscardingRound()
	}
)

func Dealing(dealType deal.Type, dealNum int) Stage {
	return func(stage *stage.Stage) {
		if dealNum == 0 && dealType == deal.Hole {
			game := stage.Play.Game

			dealNum = game.Options.Pocket
		}

		log.Printf("[play.stage] dealing %s %d cards\n", dealType, dealNum)

		stage.Dealing(dealType, dealNum)
	}
}

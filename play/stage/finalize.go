package stage

import (
	"log"
)

import (
	"gopoker/play/context"
)

func Finalize(play *context.Play) {
	log.Println("[play] finalize")

	gameOptions := play.Game.Options

	var highHands, lowHands *showdownHands

	if gameOptions.Lo != "" {
		lowHands = showdown(play, gameOptions.Lo, gameOptions.HasBoard)
	}

	if gameOptions.Hi != "" {
		highHands = showdown(play, gameOptions.Hi, gameOptions.HasBoard)
	}

	results(play, highHands, lowHands)
}

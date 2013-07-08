package stage

import (
	"log"
)

import (
	"gopoker/play/context"
)

var Finalize = func(play *context.Play) {
	log.Println("[play] finalize")

	finalize(play)
}

func finalize(play *context.Play) {
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

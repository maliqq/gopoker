package stage

func (stage *Stage) Finalize() {
	play := stage.Play

	gameOptions := play.Game.Options

	var highHands, lowHands *showdownHands

	if gameOptions.Lo != "" {
		lowHands = stage.showdown(gameOptions.Lo, gameOptions.HasBoard)
	}

	if gameOptions.Hi != "" {
		highHands = stage.showdown(gameOptions.Hi, gameOptions.HasBoard)
	}

	stage.results(highHands, lowHands)
}

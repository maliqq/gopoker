package strategy

import (
	"gopoker/play/context"
)

type Strategy []Stage

func (strategy Strategy) Proceed(play *context.Play) {
	for _, context := range strategy {
		context(play)
	}
}

var (
	Default = Strategy{
		//DealStart,
		ResetSeats, // FIXME
		//MoveButton,
		PostAntes,
		PostBlinds,
		StartStreets,
		Showdown,
		//DealStop,
	}
)

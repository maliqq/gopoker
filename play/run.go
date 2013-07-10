package play

import (
	"gopoker/play/command"
	"gopoker/play/context"
)

func Run(play *context.Play) {
Loop:
	for {
		select {
		case cmd := <-play.Control:
			switch cmd {
			case command.NextDeal:
				go DefaultStrategy.Proceed(play)

			case command.Exit:
				break Loop
			}
		}
	}
}

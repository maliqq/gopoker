package play

import (
	_ "fmt"
)

import (
	"gopoker/play/command"
	"gopoker/play/context"
	"gopoker/play/strategy"
)

func Run(play *context.Play) {
Loop:
	for {
		select {
		case cmd := <-play.Control:
			switch cmd {
			case command.NextDeal:
				play.NextDeal()
				strategy.Default.Proceed(play)

			case command.Exit:
				break Loop
			}
		}
	}
}

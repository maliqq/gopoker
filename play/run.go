package play

import (
	"log"
)

import (
	"gopoker/play/command"
	_ "gopoker/play/mode"
)

/*

Run in different modes:
- *cash* mode
	Don't force blinds and ante when players sits out.
- *tournament* mode
	Force blinds and ante on sit out.
	Increase stake on new level.
- *random* mode
	Redirect to other table on new deal.

*/

func (play *Play) RunLoop() {
Loop:
	for {
		select {
		case cmd := <-play.Control:
			switch cmd {
			case command.NextDeal:
				go play.RunMode()

			case command.Showdown:

			case command.Exit:
				break Loop
			}
		}
	}
}

func (play *Play) RunMode() {
	log.Printf("[run] mode %s\n", play.Mode)
	ByMode[play.Mode].Proceed(play)
}

func (play *Play) RunStreet() {
	log.Printf("[run] street %s\n", play.Street)
	ByStreet[play.Street].Proceed(play)
}

package play

import (
	"gopoker/play/command"
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
func (play *Play) Run() {
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

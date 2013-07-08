package play

import (
  "gopoker/play/context"
  "gopoker/play/strategy"
  "gopoker/play/command"
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

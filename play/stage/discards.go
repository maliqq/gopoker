package stage

import (
	"log"
)

import (
	"gopoker/model"
	"gopoker/model/deal"
	"gopoker/poker"
	"gopoker/protocol"
	"gopoker/play/context"
)

var DiscardingRound = func(play *context.Play) {
  log.Println("[play.stage] discarding")
  log.Fatalf("not implemented")
}

func discard(play *context.Play, p *model.Player, cards *poker.Cards) {
	pos, _ := play.Table.Pos(p)

	cardsNum := len(*cards)

	play.Broadcast.All <- protocol.NewDiscarded(pos, cardsNum)

	if cardsNum > 0 {
		newCards := play.Deal.Discard(p, cards)

		play.Broadcast.One(p) <- protocol.NewDealPocket(pos, newCards, deal.Discard)
	}
}

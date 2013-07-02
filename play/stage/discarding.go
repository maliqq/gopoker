package stage

import (
	"gopoker/model"
	"gopoker/model/deal"
	"gopoker/poker"
	"gopoker/protocol"
)

func (stage *Stage) discard(p *model.Player, cards *poker.Cards) {
	play := stage.Play

	pos, _ := play.Table.Pos(p)

	cardsNum := len(*cards)

	play.Broadcast.All <- protocol.NewDiscardCards(pos, cardsNum)

	if cardsNum > 0 {
		newCards := play.Deal.Discard(p, cards)

		play.Broadcast.One(p) <- protocol.NewDealPocket(pos, newCards, deal.Discard)
	}
}

func (stage *Stage) DiscardingRound() {

}

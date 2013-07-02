package stage

import (
	"gopoker/model/deal"
	"gopoker/protocol"
)

func (stage *Stage) dealHole(cardsNum int) {
	play := stage.Play

	for _, pos := range play.Table.SeatsInPlay() {
		player := play.Table.Player(pos)

		cards := play.Deal.DealPocket(player, cardsNum)

		play.Broadcast.One(player) <- protocol.NewDealPocket(pos, cards, deal.Hole)
	}
}

func (stage *Stage) dealDoor(cardsNum int) {
	play := stage.Play

	for _, pos := range play.Table.SeatsInPlay() {
		player := play.Table.Player(pos)

		cards := play.Deal.DealPocket(player, cardsNum)

		play.Broadcast.All <- protocol.NewDealPocket(pos, cards, deal.Door)
	}
}

func (stage *Stage) dealBoard(cardsNum int) {
	play := stage.Play

	cards := play.Deal.DealBoard(cardsNum)

	play.Broadcast.All <- protocol.NewDealShared(cards, deal.Board)
}

func (stage *Stage) Dealing(dealingType deal.Type, cardsNum int) {
	switch dealingType {
	case deal.Hole:
		stage.dealHole(cardsNum)

	case deal.Door:
		stage.dealDoor(cardsNum)

	case deal.Board:
		stage.dealBoard(cardsNum)
	}
}

package stage

import (
	"log"
)

import (
	"gopoker/model/deal"
	"gopoker/play/context"
	"gopoker/protocol"
)

func dealHole(play *context.Play, cardsNum int) {
	for _, pos := range play.Table.SeatsInPlay() {
		player := play.Table.Player(pos)

		cards := play.Deal.DealPocket(player, cardsNum)

		play.Broadcast.One(player) <- protocol.NewDealPocket(pos, cards, deal.Hole)
	}
}

func dealDoor(play *context.Play, cardsNum int) {
	for _, pos := range play.Table.SeatsInPlay() {
		player := play.Table.Player(pos)

		cards := play.Deal.DealPocket(player, cardsNum)

		play.Broadcast.All <- protocol.NewDealPocket(pos, cards, deal.Door)
	}
}

func dealBoard(play *context.Play, cardsNum int) {
	cards := play.Deal.DealBoard(cardsNum)

	play.Broadcast.All <- protocol.NewDealShared(cards, deal.Board)
}

func dealing(play *context.Play, dealingType deal.Type, cardsNum int) {
	switch dealingType {
	case deal.Hole:
		dealHole(play, cardsNum)

	case deal.Door:
		dealDoor(play, cardsNum)

	case deal.Board:
		dealBoard(play, cardsNum)
	}
}

func Dealing(dealType deal.Type, dealNum int) func(*context.Play) {
	return func(play *context.Play) {
		if dealNum == 0 && dealType == deal.Hole {
			dealNum = play.Game.Options.Pocket
		}

		log.Printf("[play.stage] dealing %s %d cards\n", dealType, dealNum)

		dealing(play, dealType, dealNum)
	}
}

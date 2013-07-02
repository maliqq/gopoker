package stage

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/protocol"
)

type showdownHands map[model.Id]*poker.Hand

func (stage *Stage) showdown(ranking poker.Ranking, withBoard bool) *showdownHands {
	play := stage.Play

	d := play.Deal

	hands := showdownHands{}

	for _, pos := range play.Table.SeatsInPot() {
		player := play.Table.Player(pos)

		pocket := d.Pocket(player)

		if hand := d.Rank(pocket, ranking, withBoard); hand != nil {
			hands[player.Id] = hand

			play.Broadcast.All <- protocol.NewShowHand(pos, pocket, hand)
		}
	}

	return &hands
}

package engine

import (
	"github.com/golang/glog"
)

import (
	"gopoker/engine/stage"
	"gopoker/engine/street"
	"gopoker/model/deal"
	"gopoker/model/game"
)

func (g *Gameplay) processStreets(skip chan bool) {
	for _, street := range buildStreets(g) {
		glog.Infof("[street] %s", street)
		if !street.Run() {
			skip <- true
			return
		}
	}
	close(skip)
}

func buildStreets(g *Gameplay) StreetStrategy {
	betting := StageExit{
		Stage: Stage{
			Type: stage.Betting,
		},
		Do: g.processBetting,
	}

	discarding := StageDo{
		Stage: Stage{
			Type: stage.Discarding,
		},
		Do: func() {}, // FIXME
	}

	bigBets := StageDo{
		Stage: Stage{
			Type: stage.BigBets,
		},
		Do: g.turnOnBigBets,
	}

	bringIn := StageDo{
		Stage: Stage{
			Type: stage.BringIn,
		},
		Do: g.bringIn,
	}

	dealing := func(dealType deal.Type, cardsNum int) StageDo {
		n := cardsNum

		var f func()
		switch dealType {
		case deal.Hole:
			f = func() { g.dealHole(n) }
		case deal.Board:
			f = func() { g.dealBoard(n) }
		case deal.Door:
			f = func() { g.dealDoor(n) }
		}

		return StageDo{
			Stage: Stage{
				Type: stage.Dealing,
			},
			Do: f,
		}
	}

	switch g.Game.Group {
	case game.Holdem:

		return StreetStrategy{

			Street{
				street.Preflop,
				StageStrategy{
					dealing(deal.Hole, 0),
					betting,
				},
			},

			Street{
				street.Flop,
				StageStrategy{
					dealing(deal.Board, 3),
					betting,
				},
			},

			Street{
				street.Turn,
				StageStrategy{
					dealing(deal.Board, 1),
					bigBets,
					betting,
				},
			},

			Street{
				street.River,
				StageStrategy{
					dealing(deal.Board, 1),
					betting,
				},
			},
		}

	case game.SevenCard:

		return StreetStrategy{

			Street{
				street.Second,
				StageStrategy{
					dealing(deal.Hole, 2),
				},
			},

			Street{
				street.Third,
				StageStrategy{
					dealing(deal.Door, 1),
					bringIn,
					betting,
				},
			},

			Street{
				street.Fourth,
				StageStrategy{
					dealing(deal.Door, 1),
					betting,
				},
			},

			Street{
				street.Fifth,
				StageStrategy{
					dealing(deal.Door, 1),
					bigBets,
					betting,
				},
			},

			Street{
				street.Sixth,
				StageStrategy{
					dealing(deal.Door, 1),
					betting,
				},
			},

			Street{
				street.Seventh,
				StageStrategy{
					dealing(deal.Hole, 1),
					betting,
				},
			},
		}

	case game.SingleDraw:

		return StreetStrategy{

			Street{
				street.Predraw,
				StageStrategy{
					dealing(deal.Hole, 5),
					betting,
					discarding,
				},
			},

			Street{
				street.Draw,
				StageStrategy{
					bigBets,
					betting,
					discarding,
				},
			},
		}

	case game.TripleDraw:

		return StreetStrategy{

			Street{
				street.FirstDraw,
				StageStrategy{
					betting,
					discarding,
				},
			},

			Street{
				street.SecondDraw,
				StageStrategy{
					betting,
					discarding,
				},
			},

			Street{
				street.ThirdDraw,
				StageStrategy{
					betting,
					discarding,
				},
			},
		}
	default:
		return StreetStrategy{}
	}

}

package engine

import (
	"gopoker/engine/stage"
	"gopoker/engine/street"
	"gopoker/model/deal"
	"gopoker/model/game"
)

func buildStreets(g *Gameplay) []Street {
	betting := Stage{
		Type: stage.Betting,
		Do:   func() {},
	}

	discarding := Stage{
		Type: stage.Discarding,
		Do:   func() {},
	}

	bigBets := Stage{
		Type: stage.BigBets,
		Do:   g.turnOnBigBets,
	}

	bringIn := Stage{
		Type: stage.BringIn,
		Do:   g.bringIn,
	}

	dealing := func(dealType deal.Type, cardsNum int) Stage {
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

		return Stage{
			Type: stage.Dealing,
			Do:   f,
		}
	}

	switch g.Game.Group {
	case game.Holdem:

		return []Street{

			Street{
				street.Preflop,
				Stages{
					dealing(deal.Hole, 0),
					betting,
				},
			},

			Street{
				street.Flop,
				Stages{
					dealing(deal.Board, 3),
					betting,
				},
			},

			Street{
				street.Turn,
				Stages{
					dealing(deal.Board, 1),
					bigBets,
					betting,
				},
			},

			Street{
				street.River,
				Stages{
					dealing(deal.Board, 1),
					betting,
				},
			},
		}

	case game.SevenCard:

		return []Street{

			Street{
				street.Second,
				Stages{
					dealing(deal.Hole, 2),
				},
			},

			Street{
				street.Third,
				Stages{
					dealing(deal.Door, 1),
					bringIn,
					betting,
				},
			},

			Street{
				street.Fourth,
				Stages{
					dealing(deal.Door, 1),
					betting,
				},
			},

			Street{
				street.Fifth,
				Stages{
					dealing(deal.Door, 1),
					bigBets,
					betting,
				},
			},

			Street{
				street.Sixth,
				Stages{
					dealing(deal.Door, 1),
					betting,
				},
			},

			Street{
				street.Seventh,
				Stages{
					dealing(deal.Hole, 1),
					betting,
				},
			},
		}

	case game.SingleDraw:

		return []Street{

			Street{
				street.Predraw,
				Stages{
					dealing(deal.Hole, 5),
					betting,
					discarding,
				},
			},

			Street{
				street.Draw,
				Stages{
					bigBets,
					betting,
					discarding,
				},
			},
		}

	case game.TripleDraw:

		return []Street{

			Street{
				street.FirstDraw,
				Stages{
					betting,
					discarding,
				},
			},

			Street{
				street.SecondDraw,
				Stages{
					betting,
					discarding,
				},
			},

			Street{
				street.ThirdDraw,
				Stages{
					betting,
					discarding,
				},
			},
		}
	default:
		return []Street{}
	}

}

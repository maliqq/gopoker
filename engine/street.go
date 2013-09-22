package engine

import (
	"gopoker/model/deal"
	"gopoker/model/game"
	"gopoker/engine/street"
	"gopoker/engine/process"
)

func (i *Instance) buildStreets() []process.Street {
	betting := process.Stage{
		Name: "betting-round",
		Run:  nil,//i.Gameplay.startBettingRound,
	}

	discarding := process.Stage{
		Name: "discarding-round",
		Run:  nil,//i.startDiscardingRound,
	}

	bigBets := process.Stage{
		Name: "big-bets",
		Run:  i.Gameplay.turnOnBigBets,
	}

	bringIn := process.Stage{
		Name: "bring-in",
		Run:  i.Gameplay.bringIn,
	}

	dealing := func(dealType deal.Type, cardsNum int) process.Stage {
		return process.Stage{
			Name: "dealing",
			Run: func() {
				switch dealType {
				case deal.Hole:
					i.dealHole(cardsNum)
				case deal.Board:
					i.dealBoard(cardsNum)
				case deal.Door:
					i.dealDoor(cardsNum)
				}
			},
		}
	}

	switch i.Game.Group {
	case game.Holdem:

		return []process.Street{

			process.Street{
				street.Preflop,
				process.Stages{
					dealing(deal.Hole, 0),
					betting,
				},
			},

			process.Street{
				street.Flop,
				process.Stages{
					dealing(deal.Board, 3),
					betting,
				},
			},

			process.Street{
				street.Turn,
				process.Stages{
					dealing(deal.Board, 1),
					bigBets,
					betting,
				},
			},

			process.Street{
				street.River,
				process.Stages{
					dealing(deal.Board, 1),
					betting,
				},
			},
		}

	case game.SevenCard:

		return []process.Street{

			process.Street{
				street.Second,
				process.Stages{
					dealing(deal.Hole, 2),
				},
			},

			process.Street{
				street.Third,
				process.Stages{
					dealing(deal.Door, 1),
					bringIn,
					betting,
				},
			},

			process.Street{
				street.Fourth,
				process.Stages{
					dealing(deal.Door, 1),
					betting,
				},
			},

			process.Street{
				street.Fifth,
				process.Stages{
					dealing(deal.Door, 1),
					bigBets,
					betting,
				},
			},

			process.Street{
				street.Sixth,
				process.Stages{
					dealing(deal.Door, 1),
					betting,
				},
			},

			process.Street{
				street.Seventh,
				process.Stages{
					dealing(deal.Hole, 1),
					betting,
				},
			},
		}

	case game.SingleDraw:

		return []process.Street{

			process.Street{
				street.Predraw,
				process.Stages{
					dealing(deal.Hole, 5),
					betting,
					discarding,
				},
			},

			process.Street{
				street.Draw,
				process.Stages{
					bigBets,
					betting,
					discarding,
				},
			},
		}

	case game.TripleDraw:

		return []process.Street{

			process.Street{
				street.FirstDraw,
				process.Stages{
					betting,
					discarding,
				},
			},

			process.Street{
				street.SecondDraw,
				process.Stages{
					betting,
					discarding,
				},
			},

			process.Street{
				street.ThirdDraw,
				process.Stages{
					betting,
					discarding,
				},
			},
		}
	default:
		return []process.Street{}
	}

}

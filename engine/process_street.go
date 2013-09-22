package engine

import (
	"gopoker/model/deal"
	"gopoker/model/game"
	"gopoker/engine/street"
)

func (i *Instance) processStreets() {
	for _, streetProcess := range i.buildStreets() {
		i.Street = streetProcess.Street
		streetProcess.Stages.Process()
	}
}

func (i *Instance) buildStreets() []StreetProcess {
	betting := StageProcess{
		Name: "betting-round",
		Run:  i.Gameplay.startBettingRound,
	}

	discarding := StageProcess{
		Name: "discarding-round",
		Run:  nil,//i.startDiscardingRound,
	}

	bigBets := StageProcess{
		Name: "big-bets",
		Run:  i.Gameplay.turnOnBigBets,
	}

	bringIn := StageProcess{
		Name: "bring-in",
		Run:  i.Gameplay.bringIn,
	}

	dealing := func(dealType deal.Type, cardsNum int) StageProcess {
		return StageProcess{
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

		return []StreetProcess{

			StreetProcess{
				Street: street.Preflop,
				Stages: StageProcesses{
					dealing(deal.Hole, 0),
					betting,
				},
			},

			StreetProcess{
				Street: street.Flop,
				Stages: StageProcesses{
					dealing(deal.Board, 3),
					betting,
				},
			},

			StreetProcess{
				Street: street.Turn,
				Stages: StageProcesses{
					dealing(deal.Board, 1),
					bigBets,
					betting,
				},
			},

			StreetProcess{
				Street: street.River,
				Stages: StageProcesses{
					dealing(deal.Board, 1),
					betting,
				},
			},
		}

	case game.SevenCard:

		return []StreetProcess{

			StreetProcess{
				Street: street.Second,
				Stages: StageProcesses{
					dealing(deal.Hole, 2),
				},
			},

			StreetProcess{
				Street: street.Third,
				Stages: StageProcesses{
					dealing(deal.Door, 1),
					bringIn,
					betting,
				},
			},

			StreetProcess{
				Street: street.Fourth,
				Stages: StageProcesses{
					dealing(deal.Door, 1),
					betting,
				},
			},

			StreetProcess{
				Street: street.Fifth,
				Stages: StageProcesses{
					dealing(deal.Door, 1),
					bigBets,
					betting,
				},
			},

			StreetProcess{
				Street: street.Sixth,
				Stages: StageProcesses{
					dealing(deal.Door, 1),
					betting,
				},
			},

			StreetProcess{
				Street: street.Seventh,
				Stages: StageProcesses{
					dealing(deal.Hole, 1),
					betting,
				},
			},
		}

	case game.SingleDraw:

		return []StreetProcess{

			StreetProcess{
				Street: street.Predraw,
				Stages: StageProcesses{
					dealing(deal.Hole, 5),
					betting,
					discarding,
				},
			},

			StreetProcess{
				Street: street.Draw,
				Stages: StageProcesses{
					bigBets,
					betting,
					discarding,
				},
			},
		}

	case game.TripleDraw:

		return []StreetProcess{

			StreetProcess{
				Street: street.FirstDraw,
				Stages: StageProcesses{
					betting,
					discarding,
				},
			},

			StreetProcess{
				Street: street.SecondDraw,
				Stages: StageProcesses{
					betting,
					discarding,
				},
			},

			StreetProcess{
				Street: street.ThirdDraw,
				Stages: StageProcesses{
					betting,
					discarding,
				},
			},
		}
	default:
		return []StreetProcess{}
	}

}

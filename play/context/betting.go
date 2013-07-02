package context

import (
	"fmt"
	"log"
	"reflect"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/seat"
	"gopoker/protocol"
)

type Betting struct {
	raiseCount int
	requireBet *protocol.RequireBet

	BigBets bool

	*model.Pot
	Log []*protocol.Message
}

func NewBetting() *Betting {
	return &Betting{
		requireBet: &protocol.RequireBet{},

		Pot: model.NewPot(),
		Log: []*protocol.Message{},
	}
}

func (betting *Betting) String() string {
	return fmt.Sprintf("Require %s %s raiseCount: %d bigBets: %t pot total: %.2f",
		betting.requireBet,
		betting.raiseCount,
		betting.BigBets,
		betting.Pot.Total(),
	)
}

func (betting *Betting) Reset() {
	req := betting.requireBet
	
	req.Call, req.Min, req.Max, betting.raiseCount = 0., 0., 0., 0
}

func (betting *Betting) RequireBet(pos int, s *model.Seat, game *model.Game) *protocol.Message {
	req := betting.requireBet

	req.Pos = pos
	req.Min, req.Max = game.Limit.RaiseRange(game.Stake, s.Stack+s.Bet, betting.Pot.Total(), betting.BigBets)
	req.Call -= s.Bet

	return protocol.NewRequireBet(req)
}

func (betting *Betting) AddBet(s *model.Seat, newBet *bet.Bet) error {
	require := betting.requireBet

	switch newBet.Type {
	case bet.Fold:
		s.Fold()

	case bet.Check:
		if require.Call != 0. {
			return fmt.Errorf("Can't check, need call: %.2f", require.Call)
		}

	default:
		if newBet.Amount != 0. {
			amount := newBet.Amount

			switch reflect.TypeOf(newBet).Name() {
			// raise, call
			case "ActiveBet":
				if newBet.Amount > s.Stack {
					return fmt.Errorf("Amount is greater than available stack: amount=%.2f stack=%.2f", amount, s.Stack)
				}

				if newBet.Type == bet.Raise {
					valid := true
					if require.Max == 0. || require.Max < amount {
						valid = false
					}
					if require.Min > amount && amount+require.Call != s.Stack {
						valid = false
					}

					if !valid {
						return fmt.Errorf("Raise is invalid: amount=%.2f stack=%.2f", amount, s.Stack)
					}

					betting.raiseCount++
					require.Call += amount
				}

			// ante, blinds
			case "ForcedBet":
				require.Call = amount
			}

			s.SetBet(amount)

			betting.Pot.Add(s.Player.Id, amount, s.State == seat.AllIn)
		}
	}

	return nil
}

func (betting *Betting) Add(seat *model.Seat, msg *protocol.Message) error {
	newBet := msg.Payload.(protocol.AddBet).Bet

	log.Printf("Player %s %s\n", seat.Player, newBet.String())

	err := betting.AddBet(seat, &newBet)

	if err == nil {
		betting.log(msg)
	}

	return err
}

func (betting *Betting) log(msg *protocol.Message) {
	betting.Log = append(betting.Log, msg)
}

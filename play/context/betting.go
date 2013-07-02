package context

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/game"
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

func (betting *Betting) ForceBet(pos int, betType bet.Type, stake *game.Stake) *bet.Bet {
	req := betting.requireBet

	amount := stake.Amount(betType)

	req.Pos = pos
	req.Call = amount

	return &bet.Bet{
		Type:   betType,
		Amount: amount,
	}
}

func (betting *Betting) RequireBet(pos int, seat *model.Seat, game *model.Game) *protocol.Message {
	req := betting.requireBet

	req.Pos = pos
	req.Min, req.Max = game.Limit.RaiseRange(game.Stake, seat.Stack+seat.Bet, betting.Pot.Total(), betting.BigBets)
	req.Call -= seat.Bet

	return protocol.NewRequireBet(req)
}

func ValidateBet(require *protocol.RequireBet, seat *model.Seat, newBet *bet.Bet) error {
	switch newBet.Type {
	case bet.Check:
		if require.Call != 0. {
			return fmt.Errorf("Can't check, need to call: %.2f", require.Call)
		}

	case bet.Call, bet.Raise:
		amount := newBet.Amount
		all_in := amount == seat.Stack

		if amount > seat.Stack {
			return fmt.Errorf("Bet amount is greater than available stack: amount=%.2f stack=%.2f", amount, seat.Stack)
		}

		if newBet.Type == bet.Call && amount != require.Call {
			return fmt.Errorf("Call mismatch: amount=%.2f call=%.2f", amount, require.Call)
		}

		if newBet.Type == bet.Raise {
			if require.Max == 0. {
				return fmt.Errorf("Raise not allowed in current betting: amount=%.2f", amount)
			}

			raiseAmount := require.Call - amount

			if raiseAmount > require.Max {
				return fmt.Errorf("Raise invalid: amount=%.2f max=%.2f", amount, require.Max)
			}

			if raiseAmount < require.Min && !all_in {
				return fmt.Errorf("Raise invalid: amount=%.2f min=%.2f", amount, require.Min)
			}
		}
	}

	return nil
}

func (betting *Betting) AddBet(seat *model.Seat, newBet *bet.Bet) error {
	require := betting.requireBet

	switch newBet.Type {
	case bet.Fold:
		seat.Fold()

	default:
		err := ValidateBet(require, seat, newBet)

		if err != nil {
			seat.Fold() // force fold

		} else {
			amount := newBet.Amount
			all_in := amount == seat.Stack

			if newBet.IsForced() {
				// ante, blinds
				require.Call = amount

				seat.SetBet(amount)

			} else if newBet.IsActive() {
				// raise, call
				if newBet.Type == bet.Raise {
					betting.raiseCount++
					require.Call += amount
				}

				seat.PutBet(amount)

				betting.Pot.Add(seat.Player.Id, amount, all_in)
			}
		}

		return err
	}

	return nil
}

func (betting *Betting) Add(seat *model.Seat, msg *protocol.Message) error {
	newBet := msg.Payload.(protocol.AddBet).Bet

	log.Printf("Player %s %s\n", seat.Player, newBet.String())

	err := betting.AddBet(seat, &newBet)

	if err != nil {
		log.Printf("[betting.error] %s", err)
	} else {
		betting.log(msg)
	}

	return err
}

func (betting *Betting) log(msg *protocol.Message) {
	betting.Log = append(betting.Log, msg)
}

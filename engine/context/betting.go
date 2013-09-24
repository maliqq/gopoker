package context

import (
	"container/ring"
	_ "fmt"
	"log"
)

import (
	"gopoker/message"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/game"
)

const (
	// MaxRaises - maximum raise count
	DefaultMaxRaises = 8 // TODO into game options
)

// Betting - betting context
type Betting struct {
	raiseCount int  // current raise count
	bigBets    bool // big bets mode

	Round    *Round
	Pot      *model.Pot
	BetRange *bet.Range
}

// NewBetting - create new betting context
func NewBetting() *Betting {
	return &Betting{
		Pot:      model.NewPot(),
		BetRange: &bet.Range{},
	}
}

func (ctx *Betting) NewRound(boxes []model.Box) {
	r := ring.New(len(boxes))
	for _, box := range boxes {
		r.Value = box
		r = r.Next()
	}
	ctx.Round = &Round{r}
}

// BigBets - increase bets
func (ctx *Betting) BigBets() {
	ctx.bigBets = true
}

// Clear - clear betting context
func (ctx *Betting) Clear() {
	ctx.raiseCount = 0
	ctx.BetRange.Reset()
}

// RaiseRange - bet range for seat
func (ctx *Betting) RaiseRange(limit game.Limit, stake *model.Stake) (float64, float64) {
	bb := stake.BigBlindAmount()
	seat := ctx.Round.Current()

	var min, max float64
	switch limit {
	case game.NoLimit:
		min, max = bb, seat.Stack

	case game.PotLimit:
		min, max = bb, ctx.Pot.Total()

	case game.FixedLimit:
		if ctx.bigBets {
			min, max = bb*2, bb*2
		} else {
			min, max = bb, bb
		}
	}

	return min, max
}

// ForceBet - force action
func (ctx *Betting) ForceBet(betType bet.Type, stake *model.Stake) *model.Bet {
	amount := stake.Amount(betType)

	ctx.BetRange.Call = amount

	return model.NewBet(betType, amount)
}

// RequireBet - require action
func (ctx *Betting) RequireBet(limit game.Limit, stake *model.Stake) *message.RequireBet {
	box := ctx.Round.Box()

	if ctx.raiseCount >= DefaultMaxRaises {
		ctx.BetRange.DisableRaise()
	} else {
		min, max := ctx.RaiseRange(limit, stake)
		// FIXME
		ctx.BetRange.SetRaise(min, max)
		ctx.BetRange.SetAvailable(box.Seat.Stack)
	}

	return &message.RequireBet{box.Pos, ctx.BetRange}
}

// AddBet - add action
func (ctx *Betting) AddBet(newBet *model.Bet) error {
	seat := ctx.Round.Current()

	log.Printf("[betting] %s %s\n", seat, newBet.String())

	err := newBet.Validate(seat, ctx.BetRange)

	if err != nil {
		seat.Fold()
	} else {
		putAmount, isAllIn := seat.AddBet(newBet)

		amount := newBet.Amount
		if amount > 0 {
			if newBet.Type != bet.Call {
				ctx.BetRange.Call = amount
			}

			if newBet.Type == bet.Raise {
				ctx.raiseCount++
			}

			ctx.Pot.Add(seat.Player, putAmount, isAllIn)
		}
	}

	return err
}

/*/ String - betting to string
func (ctx *Betting) String() string {
	return fmt.Sprintf("Pos %d BetRange %s raiseCount: %d bigBets: %t pot total: %.2f",
		ctx.Pos,
		ctx.BetRange,
		ctx.raiseCount,
		ctx.bigBets,
		ctx.Pot.Total(),
	)
}
*/

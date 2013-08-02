package context

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/game"
	"gopoker/protocol/message"
)

const (
	// MaxRaises - maximum raise count
	MaxRaises = 8 // TODO into game options
)

// Betting - betting context
type Betting struct {
	raiseCount int
	bigBets    bool

	Pot *model.Pot `json:"-"`

	active bool
	*Required

	Bet  chan *Action `json:"-"`
	Next chan int     `json:"-"`
	stop chan int     `json:"-"`
}

// Action - seat bet
type Action struct {
	Seat *model.Seat
	Bet  *model.Bet
}

// Required - action required
type Required struct {
	Pos      int
	BetRange bet.Range
}

// NewBetting - create new betting context
func NewBetting() *Betting {
	return &Betting{
		Pot: model.NewPot(),

		Required: &Required{},

		Bet: make(chan *Action),

		stop: make(chan int),
	}
}

// Clear - clear betting context
func (betting *Betting) Clear(startPos int) {
	betting.Pos = startPos // start from button
	betting.BetRange.Call, betting.BetRange.Min, betting.BetRange.Max, betting.raiseCount = 0., 0., 0., 0
}

// BigBets - increase bets
func (betting *Betting) BigBets() {
	betting.bigBets = true
}

// String - betting to string
func (betting *Betting) String() string {
	return fmt.Sprintf("Required %s raiseCount: %d bigBets: %t pot total: %.2f",
		betting.Required,
		betting.raiseCount,
		betting.bigBets,
		betting.Pot.Total(),
	)
}

// Start - start betting
func (betting *Betting) Start(pos *chan int) {
	log.Println("[betting] start")

	betting.active = true

	*pos <- betting.Pos

Loop:
	for {
		select {
		case <-betting.stop:
			log.Println("[betting] stop")
			betting.active = false
			break Loop

		case action := <-betting.Bet:
			err := betting.AddBet(action.Seat, action.Bet)

			if err != nil {
				log.Printf("[betting] %s", err)
			}

			*pos <- betting.Pos
		}
	}
}

// IsActive - check if betting is active
func (betting *Betting) IsActive() bool {
	return betting.active
}

// Stop - stop betting
func (betting *Betting) Stop() {
	betting.stop <- 1
}

// RaiseRange - bet range for seat
func (betting *Betting) RaiseRange(stackAvailable float64, g *model.Game, stake *model.Stake) (float64, float64) {
	_, bb := stake.Blinds()

	switch g.Limit {
	case game.NoLimit:
		return bb, stackAvailable

	case game.PotLimit:
		return bb, betting.Pot.Total()

	case game.FixedLimit:
		if betting.bigBets {
			return bb * 2, bb * 2
		}
		return bb, bb
	}

	return 0., 0.
}

// ForceBet - force action
func (betting *Betting) ForceBet(pos int, betType bet.Type, stake *model.Stake) *model.Bet {
	amount := stake.Amount(betType)

	betting.Pos = pos
	betting.BetRange.Call = amount

	return model.NewBet(betType, amount)
}

// RequireBet - require action
func (betting *Betting) RequireBet(pos int, stackAvailable float64, game *model.Game, stake *model.Stake) *message.Message {
	betting.Pos = pos

	if betting.raiseCount >= MaxRaises {
		betting.BetRange.Min, betting.BetRange.Max = 0., 0.
	} else {
		call := betting.BetRange.Call
		min, max := betting.RaiseRange(stackAvailable, game, stake)
		minRaise, maxRaise := call+min, call+max

		// FIXME
		//log.Printf("------\nstackAvailable=%.2f; call=%.2f; minRaise=%.2f; maxRaise=%.2f\n", stackAvailable, call, minRaise, maxRaise)
		if stackAvailable < maxRaise {
			if stackAvailable < call {
				minRaise, maxRaise = 0., 0.
			} else if stackAvailable < minRaise {
				minRaise, maxRaise = stackAvailable, stackAvailable
			} else {
				maxRaise = stackAvailable
			}
		}

		betting.BetRange.Min, betting.BetRange.Max = minRaise, maxRaise
	}

	return message.NewRequireBet(betting.Pos, betting.BetRange.Proto())
}

// AddBet - add action
func (betting *Betting) AddBet(seat *model.Seat, newBet *model.Bet) error {
	log.Printf("[betting] Player %s %s\n", seat.Player, newBet.String())

	err := newBet.Validate(seat, betting.BetRange)

	if err != nil {
		seat.Fold()
	} else {
		putAmount, isAllIn := seat.AddBet(newBet)

		amount := newBet.Amount
		if amount > 0 {
			if newBet.Type != bet.Call {
				betting.BetRange.Call = amount
			}

			if newBet.Type == bet.Raise {
				betting.raiseCount++
			}

			betting.Pot.Add(seat.Player, putAmount, isAllIn)
		}
	}

	return err
}

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

	Seat *model.Seat `json:"-"`
	Pot  *model.Pot  `json:"-"`

	active   bool
	Pos      int
	BetRange *bet.Range

	Bet  chan *model.Bet `json:"-"`
	Next chan int        `json:"-"`
	stop chan int        `json:"-"`
}

// NewBetting - create new betting context
func NewBetting() *Betting {
	return &Betting{
		Pot: model.NewPot(),

		BetRange: &bet.Range{},

		Bet:  make(chan *model.Bet),
		Next: make(chan int),
		stop: make(chan int),
	}
}

// Clear - clear betting context
func (betting *Betting) Clear(button int) {
	betting.Pos = button // start from button
	betting.raiseCount = 0.
	betting.BetRange.Reset()
}

// BigBets - increase bets
func (betting *Betting) BigBets() {
	betting.bigBets = true
}

// String - betting to string
func (betting *Betting) String() string {
	return fmt.Sprintf("Pos %d BetRange %s raiseCount: %d bigBets: %t pot total: %.2f",
		betting.Pos,
		betting.BetRange,
		betting.raiseCount,
		betting.bigBets,
		betting.Pot.Total(),
	)
}

// Start - start betting
func (betting *Betting) Start() {
	log.Println("[betting] start")

	betting.active = true
	betting.Next <- 1

Loop:
	for {
		select {
		case <-betting.stop:
			log.Println("[betting] stop")
			betting.active = false
			break Loop

		case newBet := <-betting.Bet:
			err := betting.AddBet(newBet)

			if err != nil {
				log.Printf("[betting] %s", err)
			}

			betting.Next <- 1
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
func (betting *Betting) RaiseRange(limit game.Limit, stake *model.Stake) (float64, float64) {
	bb := stake.BigBlindAmount()

	var min, max float64
	switch limit {
	case game.NoLimit:
		min, max = bb, betting.Seat.Stack

	case game.PotLimit:
		min, max = bb, betting.Pot.Total()

	case game.FixedLimit:
		if betting.bigBets {
			min, max = bb*2, bb*2
		} else {
			min, max = bb, bb
		}
	}

	return min, max
}

// ForceBet - force action
func (betting *Betting) ForceBet(pos int, seat *model.Seat, betType bet.Type, stake *model.Stake) *model.Bet {
	amount := stake.Amount(betType)

	betting.Pos = pos
	betting.Seat = seat
	betting.BetRange.Call = amount

	return model.NewBet(betType, amount)
}

// RequireBet - require action
func (betting *Betting) RequireBet(pos int, seat *model.Seat, limit game.Limit, stake *model.Stake) *message.Message {
	betting.Pos = pos
	betting.Seat = seat

	if betting.raiseCount >= MaxRaises {
		betting.BetRange.ResetRaise()
	} else {
		min, max := betting.RaiseRange(limit, stake)
		betting.BetRange.SetRaise(seat.Stack, min, max)
	}

	return message.NotifyRequireBet(betting.Pos, betting.BetRange.Proto())
}

// AddBet - add action
func (betting *Betting) AddBet(newBet *model.Bet) error {
	seat := betting.Seat
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

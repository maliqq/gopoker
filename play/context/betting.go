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
	MaxRaises = 8 // TODO into game options
)

// betting context
type Betting struct {
	raiseCount int `json:"-"`
	BigBets    bool

	Pot *model.Pot

	Required *Required

	Bet chan *Action `json:"-"`

	Next chan int `json:"-"`
	stop chan int `json:"-"`
}

type Action struct {
	Seat *model.Seat
	Bet *model.Bet
}

type Required struct {
	Pos int
	bet.Range
}

func NewBetting() *Betting {
	return &Betting{
		Pot: model.NewPot(),

		Required: &Required{},

		Bet: make(chan *Action),

		stop: make(chan int),
	}
}

func (this *Betting) Clear() {
	this.Required.Call, this.Required.Min, this.Required.Max, this.raiseCount = 0., 0., 0., 0
}

func (this *Betting) String() string {
	return fmt.Sprintf("Required %s raiseCount: %d bigBets: %t pot total: %.2f",
		this.Required,
		this.raiseCount,
		this.BigBets,
		this.Pot.Total(),
	)
}

func (this *Betting) Start(pos *chan int) {
	log.Println("[betting] start")

	*pos <- this.Required.Pos

Loop:
	for {
		select {
		case <-this.stop:
			log.Println("[betting] stop")
			break Loop

		case action := <-this.Bet:
			err := this.AddBet(action.Seat, action.Bet)

			if err != nil {
				log.Printf("[betting] %s", err)
			}

			*pos <- this.Required.Pos
		}
	}
}

func (this *Betting) Stop() {
	this.stop <- 1
}

func (this *Betting) RaiseRange(stackAvailable float64, g *model.Game, stake *model.Stake) (float64, float64) {
	_, bb := stake.Blinds()

	switch g.Limit {
	case game.NoLimit:
		return bb, stackAvailable

	case game.PotLimit:
		return bb, this.Pot.Total()

	case game.FixedLimit:
		if this.BigBets {
			return bb * 2, bb * 2
		}
		return bb, bb
	}

	return 0., 0.
}

func (this *Betting) ForceBet(pos int, betType bet.Type, stake *model.Stake) *model.Bet {
	amount := stake.Amount(betType)

	this.Required.Pos = pos
	this.Required.Call = amount

	return model.NewBet(betType, amount)
}

func (this *Betting) RequireBet(pos int, stackAvailable float64, game *model.Game, stake *model.Stake) *message.Message {
	this.Required.Pos = pos

	if this.raiseCount >= MaxRaises {
		this.Required.Min, this.Required.Max = 0., 0.
	} else {
		min, max := this.RaiseRange(stackAvailable, game, stake)
		call := this.Required.Call
		this.Required.Min, this.Required.Max = call+min, call+max
	}

	return message.NewRequireBet(this.Required.Pos, this.Required.Range.Proto())
}

func (this *Betting) AddBet(seat *model.Seat, newBet *model.Bet) error {
	log.Printf("[betting] Player %s %s\n", seat.Player, newBet.String())

	err := newBet.Validate(seat, this.Required.Range)

	if err != nil {
		seat.Fold()
	} else {
		putAmount, isAllIn := seat.AddBet(newBet)

		amount := newBet.Amount
		if amount > 0 {
			if newBet.Type != bet.Call {
				this.Required.Call = amount
			}

			if newBet.Type == bet.Raise {
				this.raiseCount++
			}

			this.Pot.Add(seat.Player, putAmount, isAllIn)
		}
	}

	return err
}

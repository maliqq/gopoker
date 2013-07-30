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

	Seat     *model.Seat
	Required *Required

	Bet chan *message.Message `json:"-"`

	Next chan int `json:"-"`
	stop chan int `json:"-"`
}

type Required struct {
	Pos int
	model.BetRange
}

func NewBetting() *Betting {
	return &Betting{
		Pot: model.NewPot(),

		Required: &Required{},

		Bet: make(chan *message.Message),

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

		case msg := <-this.Bet:
			protoBet := msg.Envelope.AddBet.Bet
			betType := protoBet.GetType().String()
			amount := protoBet.GetAmount()

			newBet := model.NewBet(bet.Type(betType), amount)
			err := this.AddBet(newBet)

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

func (this *Betting) RaiseRange(seat *model.Seat, g *model.Game, stake *model.Stake) (float64, float64) {
	_, bb := stake.Blinds()

	switch g.Limit {
	case game.NoLimit:
		return bb, seat.Stack

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

func (this *Betting) ForceBet(pos int, seat *model.Seat, betType bet.Type, stake *model.Stake) *model.Bet {
	amount := stake.Amount(betType)

	this.Seat = seat
	this.Required.Pos = pos
	this.Required.Call = amount

	return model.NewBet(betType, amount)
}

func (this *Betting) RequireBet(pos int, seat *model.Seat, game *model.Game, stake *model.Stake) *message.Message {
	this.Seat = seat
	this.Required.Pos = pos

	if this.raiseCount >= MaxRaises {
		this.Required.Min, this.Required.Max = 0., 0.
	} else {
		min, max := this.RaiseRange(seat, game, stake)
		call := this.Required.Call
		this.Required.Min, this.Required.Max = call+min, call+max
	}

	return message.NewRequireBet(this.Required.Pos, this.Required.BetRange)
}

func (this *Betting) AddBet(newBet *model.Bet) error {
	log.Printf("[betting] Player %s %s\n", this.Seat.Player, newBet.String())

	err := newBet.Validate(this.Seat, this.Required.BetRange)

	if err != nil {
		this.Seat.Fold()
	} else {
		putAmount, isAllIn := this.Seat.AddBet(newBet)

		amount := newBet.Amount
		if amount > 0 {
			if newBet.Type != bet.Call {
				this.Required.Call = amount
			}

			if newBet.Type == bet.Raise {
				this.raiseCount++
			}

			this.Pot.Add(this.Seat.Player, putAmount, isAllIn)
		}
	}

	return err
}

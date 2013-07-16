package context

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/game"
	"gopoker/model/seat"
	"gopoker/protocol"
)

const (
	MaxRaises = 8 // TODO into game options
)

// betting context
type Betting struct {
	raiseCount int
	BigBets    bool

	Pot *model.Pot

	Seat *model.Seat
	Required *protocol.RequireBet
	
	Bet chan *protocol.Message `json:"-"`
	
	Next chan int `json:"-"`
	stop chan int `json:"-"`
}

func NewBetting() *Betting {
	return &Betting{
		Pot: model.NewPot(),

		Required: &protocol.RequireBet{},

		Bet: make(chan *protocol.Message),

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

	for {
		select {
		case <-this.stop:
			log.Println("[betting] stop")
			break

		case msg := <-this.Bet:
			newBet := msg.Payload.(protocol.AddBet).Bet

			seat := this.Seat

			log.Printf("[betting] Player %s %s\n", seat.Player, newBet.String())

			err := this.AddBet(seat, &newBet)

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

func (this *Betting) ForceBet(pos int, betType bet.Type, stake *model.Stake) *model.Bet {
	amount := stake.Amount(betType)

	this.Required.Pos = pos
	this.Required.Call = amount

	return model.NewBet(betType, amount)
}

func (this *Betting) RequireBet(pos int, seat *model.Seat, game *model.Game, stake *model.Stake) *protocol.Message {
	this.Seat = seat
	this.Required.Pos = pos

	if this.raiseCount >= MaxRaises {
		this.Required.Min, this.Required.Max = 0., 0.
	} else {
		this.Required.Min, this.Required.Max = this.RaiseRange(seat, game, stake)
	}

	return protocol.NewRequireBet(seat, this.Required)
}

func (this *Betting) AddBet(s *model.Seat, newBet *model.Bet) error {
	if newBet.Type == bet.Fold {
		s.Fold()

		return nil
	}

	err := newBet.Validate(s, this.Required.RequireBet)

	if err != nil {
		s.Fold() // force fold
	} else {
		amount := newBet.Amount
		s.SetBet(amount)

		if newBet.IsForced() {
			this.Required.Call = amount
		} else {
			this.Required.Call += amount
		}

		this.Pot.Add(s.Player.Id, amount, amount == s.Stack)

		if newBet.Type == bet.Raise {
			this.raiseCount++
			s.State = seat.Bet
		}
	}

	return err
}

package model

import (
	"fmt"
)

import (
	"gopoker/model/bet"
	"gopoker/model/seat"
)

type Seat struct {
	Player Player

	State seat.State
	Stack float64

	Bet float64
}

func NewSeat() *Seat {
	return &Seat{State: seat.Empty}
}

func (this *Seat) String() string {
	return fmt.Sprintf("state=%s player=%s stack=%.2f bet=%.2f",
		this.State,
		this.Player,
		this.Stack,
		this.Bet,
	)
}

func (this *Seat) Clear() {
	this.State = seat.Empty
	this.Player = ""
	this.Stack = 0.
	this.Bet = 0.
}

func (this *Seat) Play() {
	this.Bet = 0.
	this.State = seat.Play
}

func (this *Seat) Check() {
	this.State = seat.Bet
}

func (this *Seat) Fold() {
	this.Bet = 0.
	this.State = seat.Fold
}

func (this *Seat) SetBet(amount float64) {
	this.Stack -= (amount - this.Bet)
	this.Bet = amount

	if this.Stack == 0. {
		this.State = seat.AllIn
	} else {
		this.State = seat.Bet
	}
}

func (this *Seat) ForceBet(amount float64) {
	this.Stack -= amount
	this.Bet = amount

	if this.Stack == 0. {
		this.State = seat.AllIn
	} else {
		this.State = seat.Play
	}
}

func (this *Seat) AddBet(b *Bet) (float64, bool) {
	var put float64

	switch b.Type {
	case bet.Fold:
		put = 0
		this.Fold()

	case bet.Check:
		put = 0
		this.Check()

	case bet.Call, bet.Raise:
		put = b.Amount - this.Bet
		this.SetBet(b.Amount)

	default:
		put = b.Amount - this.Bet
		if b.IsForced() {
			this.ForceBet(b.Amount)
		}
	}

	return put, this.State == seat.AllIn
}

func (this *Seat) SetPlayer(player Player) error {
	if this.State != seat.Empty {
		return fmt.Errorf("Seat is not empty.")
	}

	this.Player = player
	this.State = seat.Taken

	return nil
}

func (this *Seat) SetStack(amount float64) {
	this.Stack = amount

	if this.State == seat.Taken {
		this.State = seat.Ready
	}
}

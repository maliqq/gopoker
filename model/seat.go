package model

import (
	"fmt"
)

import (
	"gopoker/model/seat"
)

type Seat struct {
	Player *Player

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
	this.Player = nil
	this.Stack = 0.
	this.Bet = 0.
}

func (this *Seat) Play() {
	this.Bet = 0.
	this.State = seat.Play
}

func (this *Seat) Fold() {
	this.Bet = 0.
	this.State = seat.Fold
}

func (this *Seat) SetBet(amount float64) {
	this.Stack -= (amount - this.Bet)
	this.Bet = amount
}

func (this *Seat) PutBet(amount float64) {
	this.SetBet(amount)

	if this.Stack == 0. {
		this.State = seat.AllIn
	} else {
		this.State = seat.Bet
	}
}

func (this *Seat) SetPlayer(player *Player) error {
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

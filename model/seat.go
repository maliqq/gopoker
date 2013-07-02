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

func (this *Seat) Clear() {
	this.State = seat.Empty
	this.Player = nil
	this.Stack = 0.
	this.Bet = 0.
}

func (this *Seat) SetPlaying() {
	this.State = seat.Play

	this.Bet = 0.
}

func (this *Seat) Fold() {
	this.State = seat.Fold

	this.Bet = 0.
}

func (this *Seat) SetBet(amount float64) {
	this.State = seat.Bet

	this.Bet += this.AdvanceStack(-amount)
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

func (this *Seat) AdvanceStack(amount float64) float64 {
	this.Stack += amount

	return -amount
}

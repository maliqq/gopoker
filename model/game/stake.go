package game

import (
	"fmt"
)

type Stake struct {
	Size float64

	WithAnte    bool
	WithBringIn bool

	BigBlind   float64
	SmallBlind float64

	Ante float64

	BringIn float64
}

const (
	BigBetAmount     = 2.
	BigBlindAmount   = 1.
	SmallBlindAmount = 0.5
	AnteAmount       = 0.25
	BringInAmount    = 0.125
)

func NewStake(size float64) *Stake {
	return &Stake{Size: size}
}

func (stake *Stake) Amount(n float64) float64 {
	return stake.Size * n
}

func (stake *Stake) BigBetAmount() float64 {
	return stake.Amount(BigBetAmount)
}

func (stake *Stake) Blinds() (float64, float64) {
	return stake.SmallBlindAmount(), stake.BigBlindAmount()
}

func (stake *Stake) BigBlindAmount() float64 {
	if stake.BigBlind == 0. {
		return stake.Amount(BigBlindAmount)
	}

	return stake.BigBlind
}

func (stake *Stake) SmallBlindAmount() float64 {
	if stake.SmallBlind == 0. {
		return stake.Amount(SmallBlindAmount)
	}

	return stake.SmallBlind
}

func (stake *Stake) AnteAmount() float64 {
	if stake.Ante == 0. {
		return stake.Amount(AnteAmount)
	}

	return stake.Ante
}

func (stake *Stake) BringInAmount() float64 {
	if stake.BringIn == 0. {
		return stake.Amount(BringInAmount)
	}

	return stake.BringIn
}

func (stake *Stake) HasAnte() bool {
	return stake.WithAnte || stake.Ante != 0.
}

func (stake *Stake) HasBringIn() bool {
	return stake.WithBringIn || stake.BringIn != 0.
}

func (stake *Stake) String() string {
	s := fmt.Sprintf("%.2f/%.2f", stake.SmallBlindAmount(), stake.BigBlindAmount())

	if stake.HasAnte() {
		s += fmt.Sprintf(" ante: %.2f", stake.AnteAmount())
	}

	if stake.HasBringIn() {
		s += fmt.Sprintf(" bring in: %.2f", stake.BringInAmount())
	}

	return s
}

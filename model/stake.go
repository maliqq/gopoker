package model

import (
	"encoding/json"
	"fmt"
)

import (
	"gopoker/model/bet"
)

type Stake struct {
	Size float64

	WithAnte    bool `json:"-"`
	WithBringIn bool `json:"-"`

	BringIn float64
	Ante    float64

	BigBlind   float64
	SmallBlind float64
}

const (
	BigBetAmount     = 2.
	BigBlindAmount   = 1.
	SmallBlindAmount = 0.5
	AnteAmount       = 0.25
	BringInAmount    = 0.125
)

var betAmounts = map[bet.Type]float64{
	bet.Ante:       AnteAmount,
	bet.BringIn:    BringInAmount,
	bet.SmallBlind: SmallBlindAmount,
	bet.BigBlind:   BigBlindAmount,
	bet.GuestBlind: BigBlindAmount,
	bet.Straddle:   BigBetAmount,
}

func NewStake(size float64) *Stake {
	return &Stake{Size: size}
}

func (stake *Stake) Amount(betType bet.Type) float64 {
	switch betType {
	case bet.Ante:
		if stake.Ante != 0. {
			return stake.Ante
		}

	case bet.BringIn:
		if stake.BringIn != 0. {
			return stake.BringIn
		}

	case bet.SmallBlind:
		if stake.SmallBlind != 0. {
			return stake.SmallBlind
		}

	case bet.BigBlind:
		if stake.BigBlind != 0. {
			return stake.BigBlind
		}
	}

	k := betAmounts[betType]

	return k * stake.Size
}

func (stake *Stake) BigBetAmount() float64 {
	return stake.Size * BigBetAmount
}

func (stake *Stake) Blinds() (float64, float64) {
	return stake.SmallBlindAmount(), stake.BigBlindAmount()
}

func (stake *Stake) BigBlindAmount() float64 {
	return stake.Amount(bet.BigBlind)
}

func (stake *Stake) SmallBlindAmount() float64 {
	return stake.Amount(bet.SmallBlind)
}

func (stake *Stake) AnteAmount() float64 {
	return stake.Amount(bet.Ante)
}

func (stake *Stake) BringInAmount() float64 {
	return stake.Amount(bet.BringIn)
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

func (stake *Stake) MarshalJSON() ([]byte, error) {
	data := map[string]float64{
		"SmallBlind": stake.SmallBlindAmount(),
		"BigBlind":   stake.BigBlindAmount(),
	}
	if stake.HasAnte() {
		data["Ante"] = stake.AnteAmount()
	}
	if stake.HasBringIn() {
		data["BringIn"] = stake.BringInAmount()
	}
	return json.Marshal(data)
}

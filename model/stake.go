package model

import (
	"encoding/json"
	"fmt"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/bet"
	"gopoker/protocol/message"
)

// Stake - bet sizing
type Stake struct {
	Size float64

	WithAnte    bool `json:"-"`
	WithBringIn bool `json:"-"`

	BringIn float64
	Ante    float64

	BigBlind   float64
	SmallBlind float64
}

// Amount rates
const (
	BigBetAmount = 2. // BigBetAmount - big bet amount rate

	BigBlindAmount = 1. // BigBlindAmount - big blind amount rate

	SmallBlindAmount = 0.5 // SmallBlindAmount - small blind amount rate

	AnteAmount = 0.25 // AnteAmount - ante amount rate

	BringInAmount = 0.125 // BringInAmount - bring in amount rate
)

var betAmounts = map[bet.Type]float64{
	bet.Ante:       AnteAmount,
	bet.BringIn:    BringInAmount,
	bet.SmallBlind: SmallBlindAmount,
	bet.BigBlind:   BigBlindAmount,
	bet.GuestBlind: BigBlindAmount,
	bet.Straddle:   BigBetAmount,
}

// NewStake - create new stake
func NewStake(size float64) *Stake {
	return &Stake{Size: size}
}

// Amount - bet sizing for bet type
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

// BigBetAmount - amount for big bets
func (stake *Stake) BigBetAmount() float64 {
	return stake.Size * BigBetAmount
}

// Blinds - return small blind and big blind amounts
func (stake *Stake) Blinds() (float64, float64) {
	return stake.SmallBlindAmount(), stake.BigBlindAmount()
}

// BigBlindAmount - return big blind amount
func (stake *Stake) BigBlindAmount() float64 {
	return stake.Amount(bet.BigBlind)
}

// SmallBlindAmount - return small blind amount
func (stake *Stake) SmallBlindAmount() float64 {
	return stake.Amount(bet.SmallBlind)
}

// AnteAmount - return ante amount
func (stake *Stake) AnteAmount() float64 {
	return stake.Amount(bet.Ante)
}

// BringInAmount - return bring in amount
func (stake *Stake) BringInAmount() float64 {
	return stake.Amount(bet.BringIn)
}

// HasAnte - check for ante
func (stake *Stake) HasAnte() bool {
	return stake.WithAnte || stake.Ante != 0.
}

// HasBringIn - check for bring in
func (stake *Stake) HasBringIn() bool {
	return stake.WithBringIn || stake.BringIn != 0.
}

// String - stake to string
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

// MarshalJSON - stake to JSON
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

// Proto - stake to protobuf
func (stake *Stake) Proto() *message.Stake {
	return &message.Stake{
		BigBlind:   proto.Float64(stake.BigBlindAmount()),
		SmallBlind: proto.Float64(stake.SmallBlindAmount()),
		Ante:       proto.Float64(stake.AnteAmount()),
		BringIn:    proto.Float64(stake.BringInAmount()),
	}
}

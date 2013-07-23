package model

import (
	"math"
)

type Currency struct {
	Name      string // "US Dollar"
	IsoCode   string // "USD"
	Precision int    // 2
}

type Money struct {
	Value int64
	*Currency
}

func (m Money) Float64() float64 {
	return float64(m.Value) / math.Pow10(m.Precision)
}

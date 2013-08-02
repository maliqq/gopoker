package model

import (
	"math"
)

// Currency - currency
type Currency struct {
	Name      string // "US Dollar"
	IsoCode   string // "USD"
	Precision int    // 2
}

// Money - money (decimal)
type Money struct {
	Value int64
	*Currency
}

// Float64 - money to float64
func (m Money) Float64() float64 {
	return float64(m.Value) / math.Pow10(m.Precision)
}

package model

type Currency struct {
	Name string
	IsoCode string
	Precision int
}

type Money struct {
	Value int64
	*Currency
}

func (m Money) Float64() float64 {
	return float64(m.Value) / float64(m.Precision)
}

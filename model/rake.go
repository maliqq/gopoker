package model

// Rake - rake pot
type Rake struct {
	barrier float64
	amount  float64
}

func NewRake(barrier float64) *Rake {
	return &Rake{
		barrier: barrier,
	}
}

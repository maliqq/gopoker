package bet

// Range - bet range, how much to call and how much to raise
type Range struct {
	Call float64
	Min  float64
	Max  float64
}

// Reset - reset call and raise range
func (r *Range) Reset() {
	r.Call = 0.
	r.Min, r.Max = 0., 0.
}

// ResetRaise - set min and max to 0.0
func (r *Range) DisableRaise() {
	r.Min, r.Max = 0., 0.
}

func (r *Range) SetRaise(min, max float64) {
	r.Min, r.Max = r.Call+min, r.Call+max
}

// AdjustByAvailable - set raise range according to available stack
func (r *Range) SetAvailable(available float64) {
	minRaise, maxRaise := r.Min, r.Max

	if available < maxRaise {
		if available < r.Call {
			minRaise, maxRaise = 0., 0.
		} else if available < minRaise {
			minRaise, maxRaise = available, available
		} else {
			maxRaise = available
		}
	}

	r.Min, r.Max = minRaise, maxRaise
}

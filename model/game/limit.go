package game

type Limit string

const (
	FixedLimit Limit = "fixed-limit"
	PotLimit   Limit = "pot-limit"
	NoLimit    Limit = "no-limit"
)

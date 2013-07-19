package game

type Limit string

const (
	FixedLimit Limit = "fixed-limit"
	PotLimit   Limit = "pot-limit"
	NoLimit    Limit = "no-limit"
)

var (
	limitNames = map[Limit]string{
		FixedLimit: "Fixed Limit",
		PotLimit:   "Pot Limit",
		NoLimit:    "No Limit",
	}
)

func (l Limit) HumanString() string {
	return limitNames[l]
}

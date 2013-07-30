package game

type Limit string

const (
	FixedLimit Limit = "FixedLimit"
	PotLimit   Limit = "PotLimit"
	NoLimit    Limit = "NoLimit"
)

var (
	limitNames = map[Limit]string{
		FixedLimit: "Fixed Limit",
		PotLimit:   "Pot Limit",
		NoLimit:    "No Limit",
	}
)

func (l Limit) PrintString() string {
	return limitNames[l]
}

package game

// Limit - game limit type
type Limit string

const (
	// FixedLimit - fixed limit
	FixedLimit Limit = "FixedLimit"
	// PotLimit - pot limit
	PotLimit Limit = "PotLimit"
	// NoLimit - no limit
	NoLimit Limit = "NoLimit"
)

var (
	limitNames = map[Limit]string{
		FixedLimit: "Fixed Limit",
		PotLimit:   "Pot Limit",
		NoLimit:    "No Limit",
	}
)

// PrintString - limit to print string
func (l Limit) PrintString() string {
	return limitNames[l]
}

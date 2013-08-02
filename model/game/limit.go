package game

// Limit - game limit type
type Limit string

// Limits
const (
	FixedLimit Limit = "FixedLimit" // FixedLimit - fixed limit
	PotLimit   Limit = "PotLimit"   // PotLimit - pot limit
	NoLimit    Limit = "NoLimit"    // NoLimit - no limit
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

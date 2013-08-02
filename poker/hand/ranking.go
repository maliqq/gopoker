package hand

// Ranking - ranking type
type Ranking string

const (
	// High - ranking type high
	High Ranking = "High"
	// Badugi - ranking type badugi
	Badugi Ranking = "Badugi"
	// AceFive - ranking type Ace to Five low
	AceFive Ranking = "AceFive"
	// AceFive8 - ranking type Ace to Five low with Eight qualifier
	AceFive8 Ranking = "AceFive8"
	// AceSix - ranking type Ace to Six low
	AceSix Ranking = "AceSix"
	// DeuceSeven - ranking type Deuce to Seven low
	DeuceSeven Ranking = "DeuceSeven"
)

package hand

// Ranking - ranking type
type Ranking string

// Rankings
const (
	High       Ranking = "High"       // High - ranking type high
	Badugi     Ranking = "Badugi"     // Badugi - ranking type badugi
	AceFive    Ranking = "AceFive"    // AceFive - ranking type Ace to Five low
	AceFive8   Ranking = "AceFive8"   // AceFive8 - ranking type Ace to Five low with Eight qualifier
	AceSix     Ranking = "AceSix"     // AceSix - ranking type Ace to Six low
	DeuceSeven Ranking = "DeuceSeven" // DeuceSeven - ranking type Deuce to Seven low
)

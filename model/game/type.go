package game

// LimitedGame - game type with limit
type LimitedGame string

// MixedGame - game type with variations
type MixedGame string

// Game - limited or mixed game
type Game string

// Group - game group
type Group string

// Type - game type interface
type Type interface {
	Game() Game
}

// Game - limited game to game
func (g LimitedGame) Game() Game {
	return Game(g)
}

// Game - mixed game to game
func (g MixedGame) Game() Game {
	return Game(g)
}

// limited games
const (
	Texas  LimitedGame = "Texas"  // Texas - texas holdem
	Omaha  LimitedGame = "Omaha"  // Omaha - omaha holdem
	Omaha8 LimitedGame = "Omaha8" // Omaha8 - omaha holdem hi/lo

	Stud   LimitedGame = "Stud"   // Stud - seven card stud
	Stud8  LimitedGame = "Stud8"  // Stud8 - seven card stud hi/lo
	Razz   LimitedGame = "Razz"   // Razz - Ace to Five lowball stud
	London LimitedGame = "London" // London - Ace to Six lowball stud

	FiveCard LimitedGame = "FiveCard" // FiveCard - five card draw poker
	Single27 LimitedGame = "Single27" // Single27 - 2-7 lowball draw poker
	Triple27 LimitedGame = "Triple27" // Triple27 - 2-7 lowball poker with 3 draws
	Badugi   LimitedGame = "Badugi"   // Badugi - badugi poker
)

// mixed games
const (
	Horse MixedGame = "Horse" // Horse - H.O.R.S.E.
	Eight MixedGame = "Eight" // Eight - 8-game
)

// groups
const (
	Holdem     Group = "Holdem"     // Holdem - holdem poker
	SevenCard  Group = "SevenCard"  // SevenCard - 7-card poker
	SingleDraw Group = "SingleDraw" // SingleDraw - single draw poker
	TripleDraw Group = "TripleDraw" // TripleDraw - triple draw poker
)

var (
	limitedGameNames = map[LimitedGame]string{
		Texas:    "Texas",
		Omaha:    "Omaha",
		Omaha8:   "Omaha Hi/Lo",
		Stud:     "Stud",
		Stud8:    "Stud Hi/Lo",
		Razz:     "Razz",
		London:   "London",
		FiveCard: "5-card",
		Single27: "Single 2-7",
		Triple27: "Triple 2-7",
		Badugi:   "Badugi",
	}

	mixedGameNames = map[MixedGame]string{
		Horse: "HORSE",
		Eight: "8-game",
	}

	groupNames = map[Group]string{
		Holdem:     "Holdem",
		SevenCard:  "Seven Card",
		SingleDraw: "Single Draw",
		TripleDraw: "Triple Draw",
	}
)

// PrintString - limited game to print string
func (g LimitedGame) PrintString() string {
	return limitedGameNames[g]
}

// PrintString - mixed game to print string
func (g MixedGame) PrintString() string {
	return mixedGameNames[g]
}

// PrintString - group to print string
func (g Group) PrintString() string {
	return groupNames[g]
}

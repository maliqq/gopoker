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

const (
	// Texas - texas holdem
	Texas LimitedGame = "Texas"
	// Omaha - omaha holdem
	Omaha LimitedGame = "Omaha"
	// Omaha8 - omaha holdem hi/lo
	Omaha8 LimitedGame = "Omaha8"
	// Stud - seven card stud
	Stud LimitedGame = "Stud"
	// Stud8 - seven card stud hi/lo
	Stud8 LimitedGame = "Stud8"
	// Razz - Ace to Five lowball stud
	Razz LimitedGame = "Razz"
	// London - Ace to Six lowball stud
	London LimitedGame = "London"
	// FiveCard - five card draw poker
	FiveCard LimitedGame = "FiveCard"
	// Single27 - 2-7 lowball draw poker
	Single27 LimitedGame = "Single27"
	// Triple27 - 2-7 lowball poker with 3 draws
	Triple27 LimitedGame = "Triple27"
	// Badugi - badugi poker
	Badugi LimitedGame = "Badugi"

	// mixes

	// Horse - H.O.R.S.E.
	Horse MixedGame = "Horse"
	// Eight - 8-game
	Eight MixedGame = "Eight"

	// groups

	// Holdem - holdem poker
	Holdem Group = "Holdem"
	// SevenCard - 7-card poker
	SevenCard Group = "SevenCard"
	// SingleDraw - single draw poker
	SingleDraw Group = "SingleDraw"
	// TripleDraw - triple draw poker
	TripleDraw Group = "TripleDraw"
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

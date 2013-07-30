package game

type LimitedGame string
type MixedGame string
type Game string
type Group string

type Type interface {
	Game() Game
}

func (g LimitedGame) Game() Game {
	return Game(g)
}

func (g MixedGame) Game() Game {
	return Game(g)
}

const (
	Texas    LimitedGame = "Texas"
	Omaha    LimitedGame = "Omaha"
	Omaha8   LimitedGame = "Omaha8"
	Stud     LimitedGame = "Stud"
	Stud8    LimitedGame = "Stud8"
	Razz     LimitedGame = "Razz"
	London   LimitedGame = "London"
	FiveCard LimitedGame = "FiveCard"
	Single27 LimitedGame = "Single27"
	Triple27 LimitedGame = "Triple27"
	Badugi   LimitedGame = "Badugi"

	// mixes
	Horse MixedGame = "Horse"
	Eight MixedGame = "Eight"

	// groups
	Holdem     Group = "Holdem"
	SevenCard  Group = "SevenCard"
	SingleDraw Group = "SingleDraw"
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

func (l LimitedGame) PrintString() string {
	return limitedGameNames[l]
}

func (m MixedGame) PrintString() string {
	return mixedGameNames[m]
}

func (g Group) PrintString() string {
	return groupNames[g]
}

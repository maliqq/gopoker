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
	Texas    LimitedGame = "texas"
	Omaha    LimitedGame = "omaha"
	Omaha8   LimitedGame = "omaha8"
	Stud     LimitedGame = "stud"
	Stud8    LimitedGame = "stud8"
	Razz     LimitedGame = "razz"
	London   LimitedGame = "london"
	FiveCard LimitedGame = "five-card"
	Single27 LimitedGame = "single27"
	Triple27 LimitedGame = "triple27"
	Badugi   LimitedGame = "badugi"

	// mixes
	Horse MixedGame = "horse"
	Eight MixedGame = "eight"

	// groups
	Holdem     Group = "holdem"
	SevenCard  Group = "seven-card"
	SingleDraw Group = "single-draw"
	TripleDraw Group = "triple-draw"
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

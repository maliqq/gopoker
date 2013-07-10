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

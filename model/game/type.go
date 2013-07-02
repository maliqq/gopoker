package game

type Type string

type Group string

const (
	// simple
	Texas    Type = "texas"
	Omaha    Type = "omaha"
	Omaha8   Type = "omaha8"
	Stud     Type = "stud"
	Stud8    Type = "stud8"
	Razz     Type = "razz"
	London   Type = "london"
	FiveCard Type = "five-card"
	Single27 Type = "single27"
	Triple27 Type = "triple27"
	Badugi   Type = "badugi"
	// mixed
	Horse Type = "horse"
	Eight Type = "eight"
)

const (
	Holdem     Group = "holdem"
	SevenCard  Group = "seven-card"
	SingleDraw Group = "single-draw"
	TripleDraw Group = "triple-draw"
)

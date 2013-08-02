package position

// Position - position
type Position string

const (
	// SB - small blind
	SB Position = "SB"
	// BB - big blind
	BB Position = "BB"
	// BU - button
	BU Position = "BU"
	// CO - cutoff
	CO Position = "CO"
	// MP1 - middle player 1
	MP1 Position = "MP1"
	// MP2 - middle player 2
	MP2 Position = "MP2"
	// MP3 - middle player 3
	MP3 Position = "MP3"
	// UTG1 - under the gun 1
	UTG1 Position = "UTG1"
	// UTG2 - under the gun 2
	UTG2 Position = "UTG2"
	// UTG3 - under the gun 3
	UTG3 Position = "UTG3"
)

type pair struct {
	pos Position
	max int
}

var pairs = []pair{
	{BU, 2}, {SB, 1}, {BB, 1}, {UTG1, 3}, {UTG2, 7}, {UTG3, 9}, {MP1, 5}, {MP2, 6}, {MP3, 8}, {CO, 4},
}

// Draw - display positions
func Draw(n int) []Position {
	var r = make([]Position, n)

	i := 0
	for _, pair := range pairs {
		if n > pair.max {
			r[i] = pair.pos
			i++
		}
	}

	return r
}

package deal

// Type - deal type
type Type string

const (
	// Hole - deal hole cards
	Hole Type = "Hole"
	// Door - deal door cards
	Door Type = "Door"
	// Board - deal board cards
	Board Type = "Board"
	// Discard - deal discarded cards
	Discard Type = "Discard"
)

// IsBoard - check we're dealing board cards
func (t Type) IsBoard() bool {
	return t == Board
}

// IsShared - check we're dealing shared cards (door or board)
func (t Type) IsShared() bool {
	return t == Board || t == Hole
}

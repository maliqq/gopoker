package deal

// Type - deal type
type Type string

// Deal types
const (
	Hole Type = "Hole" // Hole - deal hole cards

	Door Type = "Door" // Door - deal door cards

	Board Type = "Board" // Board - deal board cards

	Discard Type = "Discard" // Discard - deal discarded cards
)

// IsBoard - check we're dealing board cards
func (t Type) IsBoard() bool {
	return t == Board
}

// IsShared - check we're dealing shared cards (door or board)
func (t Type) IsShared() bool {
	return t == Board || t == Hole
}

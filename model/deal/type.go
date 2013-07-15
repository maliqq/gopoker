package deal

type Type string

const (
	Hole    Type = "hole"
	Door    Type = "door"
	Board   Type = "board"
	Discard Type = "discard"
)

func (t Type) IsBoard() bool {
	return t == Board
}

func (t Type) IsShared() bool {
	return t == Board || t == Hole
}

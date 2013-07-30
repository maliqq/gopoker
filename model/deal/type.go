package deal

type Type string

const (
	Hole    Type = "Hole"
	Door    Type = "Door"
	Board   Type = "Board"
	Discard Type = "Discard"
)

func (t Type) IsBoard() bool {
	return t == Board
}

func (t Type) IsShared() bool {
	return t == Board || t == Hole
}

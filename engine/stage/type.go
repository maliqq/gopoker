package stage

type Type string

const (
	PrepareSeats Type = "prepare-seats"
	RotateGame   Type = "rotate-game"
	PostBlinds   Type = "post-blinds"
	PostAntes    Type = "post-antes"
	Streets      Type = "streets"
	Showdown     Type = "showdown"
	BigBets      Type = "big-bets"
	BringIn      Type = "bring-in"
	Betting      Type = "betting"
	Discarding   Type = "discarding"
	Dealing      Type = "dealing"
)

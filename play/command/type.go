package command

type Command string

const (
	NextTurn Command = "next-turn"
	NextDeal Command = "next-deal"
	Showdown Command = "showdown"
	Exit     Command = "exit"
)

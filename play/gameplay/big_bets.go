package gameplay

func (gp *GamePlay) BigBets() Transition {
	gp.Betting.BigBets()

	return Next
}

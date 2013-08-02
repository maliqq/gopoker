package gameplay

// BigBets - switch to big bets mode
func (gp *GamePlay) BigBets() Transition {
	gp.Betting.BigBets()

	return Next
}

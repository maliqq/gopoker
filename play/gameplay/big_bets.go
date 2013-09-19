package gameplay

// BigBets - switch to big bets mode
func (gp *Gameplay) BigBets() Transition {
	gp.Betting.BigBets()

	return Next
}

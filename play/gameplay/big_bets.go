package gameplay

func (this *GamePlay) BigBets() Transition {
	this.Betting.BigBets()

	return Next
}

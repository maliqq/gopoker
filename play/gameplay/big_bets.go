package gameplay

func (this *GamePlay) BigBets() Transition {
	this.Betting.BigBets = true

	return Next
}

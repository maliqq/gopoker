package gameplay

const (
	DefaultTimer = 30
)

func (this *GamePlay) StartBettingRound() Transition {
	//this.Broadcast.All <- protocol.NewBettingStart(this.Betting)
	pos := make(chan int)
	defer close(pos)

	go this.Betting.Start(&pos)

	var next Transition
	for current := range pos {
		active := this.Table.Seats.From(current).Playing()
		inPot := this.Table.Seats.From(current).InPot()

		if len(inPot) < 2 {
			next = Stop
			break
		} else if len(active) > 0 {
			pos := active[0]
			seat := this.Table.Seat(pos)

			this.Broadcast.One(seat.Player) <- this.Betting.RequireBet(pos, seat, this.Game, this.Stake)

			continue
		}

		next = Next
		break
	}

	this.Betting.Stop()
	this.ResetBetting()

	return next
}

func (this *GamePlay) ResetBetting() {
	this.Betting.Clear()

	for _, pos := range this.Table.AllSeats().InPlay() {
		seat := this.Table.Seat(pos)
		seat.Play()
	}

	//this.Broadcast.All <- protocol.NewBettingStop(this.Betting)
}

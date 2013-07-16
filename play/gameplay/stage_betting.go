package gameplay

import (
	"fmt"
	_ "time"
)

import (
	"gopoker/play/command"
	"gopoker/protocol"
)

const (
	DefaultTimer = 30
)

func (this *GamePlay) NextTurn(current int) {
	fmt.Printf("current pos=%d\n", current)

	active := this.Table.Seats.From(current).InPlay()
	inPot := this.Table.Seats.From(current).InPot()

	if len(inPot) < 2 {
		fmt.Println("goto showdown")
		this.Control <- command.Showdown
		this.Betting.Stop()
	} else if len(active) == 0 {
		fmt.Println("reset betting")
		this.Betting.Stop()
	} else {
		fmt.Printf("seats active=%#v\n\n", active)
		pos := active[0]
		seat := this.Table.Seat(pos)
		this.Broadcast.One(seat.Player) <- this.Betting.RequireBet(pos, seat, this.Game, this.Stake)
	}
}

func (this *GamePlay) StartBettingRound() {
	pos := make(chan int)
	defer close(pos)

	go this.Betting.Start(&pos)

	for current := range pos {
		this.NextTurn(current)
	}

	this.ResetBetting()
}

func (this *GamePlay) ResetBetting() {
	this.Betting.Clear()

	for _, pos := range this.Table.SeatsInPlay() {
		seat := this.Table.Seat(pos)
		seat.Play()
	}

	this.Broadcast.All <- protocol.NewPotSummary(this.Pot)
}

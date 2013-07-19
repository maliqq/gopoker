package play

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
	"gopoker/model/seat"
	"gopoker/play/command"
	"gopoker/play/context"
	"gopoker/play/gameplay"
	"gopoker/play/mode"
	"gopoker/play/street"
	"gopoker/protocol"
	"gopoker/util/console"
)

type Play struct {
	// players action context
	*gameplay.GamePlay

	// current mode
	Mode mode.Type
	// current street
	Street street.Type
	// receive protocol messages
	Receive chan *protocol.Message `json:"-"`
}

func (this *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", this.Game, this.Table)
}

func (this *Play) RotateGame() {
	if this.Mix == nil {
		return
	}

	if nextGame := this.GameRotation.Next(); nextGame != nil {
		this.Game = nextGame
		this.Broadcast.All <- protocol.NewChangeGame(nextGame)
	}
}

func NewPlay(variation model.Variation, stake *model.Stake, table *model.Table) *Play {
	play := &Play{
		GamePlay: &gameplay.GamePlay{
			Table:     table,
			Stake:     stake,
			Broadcast: protocol.NewBroadcast(),
			Control:   make(chan command.Command),
		},
		Receive: make(chan *protocol.Message),
	}

	if variation.IsMixed() {
		mix := variation.(*model.Mix)
		play.Mix = mix
		play.GameRotation = context.NewGameRotation(mix, 1)
		play.Game = play.GameRotation.Current()
	} else {
		play.Game = variation.(*model.Game)
	}

	go play.receive()

	return play
}

func (this *Play) receive() {
	for {
		msg := <-this.Receive
		log.Printf(console.Color(console.YELLOW, msg.String()))

		switch msg.Payload().(type) {
		case protocol.JoinTable:
			join := msg.Envelope.JoinTable
			_, err := this.Table.AddPlayer(join.Player, join.Pos, join.Amount)
			if err == nil {
				// start next deal
			} else {
				log.Printf("[protocol] error: %s", err)
			}
			//this.Broadcast.Except(join.Player) <- join

		case protocol.LeaveTable:
			leave := msg.Envelope.LeaveTable
			this.Table.RemovePlayer(leave.Player)

			//this.Broadcast.Except(join.Player) <- leave

			// TODO: fold & autoplay

		case protocol.SitOut:
			sitOut := msg.Envelope.SitOut
			this.Table.Seat(sitOut.Pos).State = seat.Idle

			// TODO: fold

		case protocol.ComeBack:
			comeBack := msg.Envelope.ComeBack
			this.Table.Seat(comeBack.Pos).State = seat.Ready

		case protocol.AddBet:
			this.Betting.Bet <- msg

		case protocol.DiscardCards:
			this.Discarding.Discard <- msg
		}
	}
}

func (this *Play) Start() {

}

func (this *Play) Pause() {

}

func (this *Play) Resume() {

}

func (this *Play) Close() {

}

func (this *Play) ResetSeats() {
	for _, s := range this.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play:
			s.Play()
		}
	}
}

func (this *Play) StartNewDeal() {
	this.Deal = model.NewDeal()
	this.Betting = context.NewBetting()
	if this.Game.Discards {
		this.Discarding = context.NewDiscarding(this.Deal)
	}
}

func (this *Play) ScheduleNextDeal() {
	this.GamePlay.Control <- command.NextDeal
}

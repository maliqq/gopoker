package context

import (
	"fmt"
	"log"
	"time"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/model/seat"
	"gopoker/play/command"
	g "gopoker/play/gameplay"
	"gopoker/poker"
	"gopoker/poker/ranking"
	"gopoker/protocol"
	"gopoker/util/console"
)

/*********************************
* Context
*********************************/
type Play struct {
	// players action context
	*g.GamePlay
  *GameRotation `json:"-"`
	*Betting    `json:"-"`
	*Discarding `json:"-"`

	// receive protocol messages
	Receive chan *protocol.Message `json:"-"`

	// manage play
	Control chan command.Type `json:"-"`
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
		GamePlay: &g.GamePlay{
			Table:     table,
			Stake:     stake,
			Broadcast: protocol.NewBroadcast(),
		},
		Receive:   make(chan *protocol.Message),
		Control:   make(chan command.Type),
	}

	if variation.IsMixed() {
		mix := variation.(*model.Mix)
		play.Mix = mix
		play.GameRotation = NewGameRotation(mix, 1)
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

		switch msg.Payload.(type) {
		case protocol.JoinTable:
			join := msg.Payload.(protocol.JoinTable)
			this.Table.AddPlayer(join.Player, join.Pos, join.Amount)

			//this.Broadcast.Except(join.Player) <- join

		case protocol.LeaveTable:
			leave := msg.Payload.(protocol.LeaveTable)
			this.Table.RemovePlayer(leave.Player)

			//this.Broadcast.Except(join.Player) <- leave

			// TODO: fold & autoplay

		case protocol.SitOut:
			sitOut := msg.Payload.(protocol.SitOut)
			this.Table.Seat(sitOut.Pos).State = seat.Idle

			// TODO: fold

		case protocol.ComeBack:
			comeBack := msg.Payload.(protocol.ComeBack)
			this.Table.Seat(comeBack.Pos).State = seat.Ready

		case protocol.AddBet:
			this.Betting.Receive <- msg

		case protocol.DiscardCards:
			this.Discarding.Receive <- msg
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

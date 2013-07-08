package context

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
	"gopoker/play/command"
	"gopoker/protocol"
	"gopoker/util/console"
)

type Play struct {
	*model.Deal
	*model.Table
	*model.Game

	*protocol.Broadcast
	Receive chan *protocol.Message
	Control chan command.Type

	*Betting
	*Discarding
}

func (this *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", this.Game, this.Table)
}

func NewPlay(game *model.Game, table *model.Table) *Play {
	play := &Play{
		Game:      game,
		Table:     table,
		Broadcast: protocol.NewBroadcast(),
		Receive:   make(chan *protocol.Message),
		Control:   make(chan command.Type),
	}

	go play.receive()

	return play
}

func (this *Play) NextDeal() {
	this.Deal = model.NewDeal()
	this.Betting = NewBetting()
	if this.Game.Options.Discards {
		this.Discarding = NewDiscarding()
	}
}

func (this *Play) receive() {
	for {
		msg := <-this.Receive
		log.Printf(console.Color(console.YELLOW, msg.String()))

		switch msg.Payload.(type) {
		case protocol.JoinTable:
			join := msg.Payload.(protocol.JoinTable)

			this.Table.AddPlayer(join.Player, join.Pos, join.Amount)

		case protocol.LeaveTable:
			leave := msg.Payload.(protocol.LeaveTable)

			this.Table.RemovePlayer(leave.Player)

		case protocol.AddBet:
			this.Betting.Receive <- msg

		case protocol.DiscardCards:
			this.Discarding.Receive <- msg
		}
	}
}

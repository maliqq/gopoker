package context

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
	"gopoker/protocol"
	"gopoker/util/console"
	"gopoker/play/command"
)

type State string

type Play struct {
	*model.Deal
	*model.Table
	*model.Game

	*protocol.Broadcast
	Receive  chan *protocol.Message
	Betting chan *protocol.Message
	Discarding chan *protocol.Message

	Control chan command.Type
}

func (play *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", play.Game, play.Table)
}

func NewPlay(game *model.Game, table *model.Table) *Play {
	play := &Play{
		State:      Waiting,
		Game:       game,
		Table:      table,
		Broadcast:  protocol.NewBroadcast(),
		Receive:    make(chan *protocol.Message),
	}
	play.receive()

	return play
}

func (play *Play) receive() {
	for {
		msg := <-play.Receive

		log.Printf(console.Color(console.YELLOW, msg.String()))

		switch msg.Payload.(type) {
		case protocol.JoinPlayer:
			join := msg.Payload.(protocol.JoinPlayer)
			play.Table.AddPlayer(join.Player, join.Pos, join.Amount)

		case protocol.LeaveTable:
			leave := msg.Payload.(protocol.LeaveTable)
			play.Table.RemovePlayer(leave.Player)

		case protocol.DiscardCards:
			play.Discarding <- msg

		case protocol.AddBet:
			play.Betting <- msg
		}
	}
}

func (play *Play) Start() {
	play.Deal = model.NewDeal()
	play.Betting = make(chan *protocol.Message)

	if play.Game.Options.Discards {
		play.Discarding = make(chan *protocol.Message)
	}
	
	play.Control <- command.Start
}

func (play *Play) Pause() {
	play.Control <- command.Pause
}

func (play *Play) Resume() {
	play.Control <- command.Resume
}

func (play *Play) Stop() {
	play.Control <- command.Stop
}

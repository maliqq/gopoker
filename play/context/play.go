package context

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
	"gopoker/protocol"
	"gopoker/util/console"
)

type State string

const (
	Active = "active"
	Waiting = "waiting"
	Paused = "paused"
	Closed = "closed"
)

type Play struct {
	State
	*model.Deal
	*model.Table
	*model.Game

	*protocol.Broadcast
	Receive  chan *protocol.Message
	Betting chan *protocol.Message
	Discarding chan *protocol.Message

	NextDeal chan int
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
		Betting:    make(chan *protocol.Message),
		NextDeal:   make(chan int),
	}
	
	if game.Options.Discards {
		play.Discarding = make(chan *protocol.Message)
	}

	play.receive()

	return play
}

func (play *Play) Start() {
	play.receive()
}

func (play *Play) Close() {
	play.State = Closed
}

func (play *Play) Pause() {
	play.State = Paused
}

func (play *Play) Resume() {
	play.State = Active
}

func (play *Play) nextDeal() {
	play.Deal = model.NewDeal()
	play.State = Active
}

func (play *Play) receive() {
	for {
		select {
		case msg := <-play.Receive:
			log.Printf(console.Color(console.YELLOW, msg.String()))

			switch msg.Payload.(type) {
			case protocol.JoinPlayer:
				join := msg.Payload.(protocol.JoinPlayer)
				play.Table.AddPlayer(join.Player, join.Pos, join.Amount)

			case protocol.LeaveTable:
				leave := msg.Payload.(protocol.LeaveTable)
				play.Table.RemovePlayer(leave.Player)

			case protocol.ChangeSeatState:
				change := msg.Payload.(protocol.ChangeSeatState)
				seat := play.Table.Seat(change.Pos)
				seat.State = change.State
			
			case protocol.DiscardCards:
				play.Discarding <- msg

			case protocol.AddBet:
				play.Betting <- msg
			}

		case <-play.NextDeal:
			play.nextDeal()
		}
	}
}

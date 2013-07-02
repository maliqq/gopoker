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

type Play struct {
	*model.Deal
	*model.Table
	*model.Game

	*protocol.Broadcast
	Receive  chan *protocol.Message
	NextTurn chan *protocol.Message
	NextDeal chan int
}

func (play *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", play.Game, play.Table)
}

func NewPlay(game *model.Game, table *model.Table) *Play {
	play := &Play{
		Game:      game,
		Table:     table,
		Broadcast: protocol.NewBroadcast(),
		Receive:   make(chan *protocol.Message),
		NextTurn:  make(chan *protocol.Message),
		NextDeal:  make(chan int),
	}

	go play.receive()

	return play
}

func (play *Play) receive() {
	for {
		select {
		case msg := <-play.Receive:
			log.Printf(console.Color(console.YELLOW, msg.String()))

			switch msg.Payload.(type) {
			case protocol.AddBet:
				play.NextTurn <- msg
			}

		case <-play.NextDeal:
			log.Println("starting new deal")
		}
	}
}

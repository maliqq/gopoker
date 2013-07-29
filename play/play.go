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
	"gopoker/protocol"
	"gopoker/util/console"
)

type Play struct {
	// players action context
	*gameplay.GamePlay

	// finite state machine
	FSM
	// receive protocol messages
	Recv protocol.MessageChannel `json:"-"`
}

func NewPlay(variation model.Variation, stake *model.Stake, table *model.Table) *Play {
	play := &Play{
		GamePlay: &gameplay.GamePlay{
			Table:     table,
			Stake:     stake,
			Broadcast: protocol.NewBroadcast(),
			Control:   make(chan command.Command),
		},
		FSM: FSM{
			State:       Waiting,
			stateChange: make(chan State),
			Mode:        mode.Cash, // FIXME
		},
		Recv: make(chan *protocol.Message),
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

func (this *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", this.Game, this.Table)
}

func (this *Play) receive() {
	for {
		select {
		case msg := <-this.Recv:
			this.processMessage(msg)
		case newState := <-this.stateChange:
			this.processStateChange(newState)
		}
	}
}

func (this *Play) processMessage(msg *protocol.Message) {
	log.Printf(console.Color(console.YELLOW, msg.String()))

	switch msg.Payload().(type) {
	case protocol.JoinTable:
		join := msg.Envelope.JoinTable
		player := model.Player(*join.Player)

		_, err := this.Table.AddPlayer(player, int(*join.Pos), *join.Amount)
		if err == nil {
			// start next deal
		} else {
			log.Printf("[protocol] error: %s", err)
		}
		// retranslate
		this.Broadcast.All <- msg

	case protocol.LeaveTable:
		leave := msg.Envelope.LeaveTable
		player := model.Player(*leave.Player)

		this.Table.RemovePlayer(player)
		this.Broadcast.All <- msg
		// TODO: fold & autoplay

	case protocol.SitOut:
		sitOut := msg.Envelope.SitOut
		this.Table.Seat(int(*sitOut.Pos)).State = seat.Idle
		// TODO: fold

	case protocol.ComeBack:
		comeBack := msg.Envelope.ComeBack
		this.Table.Seat(int(*comeBack.Pos)).State = seat.Ready

	case protocol.ChatMessage:
		this.Broadcast.All <- msg

	case protocol.AddBet:
		this.Betting.Bet <- msg
		this.Broadcast.All <- msg

	case protocol.DiscardCards:
		this.Discarding.Discard <- msg
	}
}

func (this *Play) processStateChange(newState State) {
	log.Printf("[play] state %s", newState)

	oldState := this.State
	if oldState != newState {
		this.stateLock.Lock()
		defer this.stateLock.Unlock()

		this.State = newState
		if newState == Active {
			go this.Run()
			this.Control <- command.NextDeal
		}
	}
}

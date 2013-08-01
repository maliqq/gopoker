package play

import (
	"fmt"
	"log"
	"time"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/seat"
	"gopoker/play/context"
	"gopoker/play/gameplay"
	"gopoker/play/mode"
	"gopoker/protocol"
	"gopoker/protocol/message"
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
	gp := gameplay.NewGamePlay()
	gp.Table = table
	gp.Stake = stake

	play := &Play{
		GamePlay: gp,
		FSM: FSM{
			State:       Waiting,
			stateChange: make(chan State),
			Mode:        mode.Cash, // FIXME
		},
		Recv: make(chan *message.Message),
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
			this.handleMessage(msg)

		case newState := <-this.stateChange:
			this.handleStateChange(newState)
		}
	}
}

func (this *Play) handleMessage(msg *message.Message) {
	log.Printf(console.Color(console.YELLOW, msg.String()))

	payload := msg.Payload() // FIXME
	if payload == nil {
		return
	}

	switch payload.(type) {
	case *message.JoinTable:
		join := msg.Envelope.JoinTable
		player := model.Player(join.GetPlayer())
		pos := int(join.GetPos())
		amount := join.GetAmount()

		_, err := this.Table.AddPlayer(player, pos, amount)
		if err == nil {
			// start next deal
		} else {
			log.Printf("[protocol] error: %s", err)
		}
		// retranslate
		this.Broadcast.All <- msg

	case *message.LeaveTable:
		leave := msg.Envelope.LeaveTable
		player := model.Player(leave.GetPlayer())

		this.Table.RemovePlayer(player)
		this.Broadcast.All <- msg
		// TODO: fold & autoplay

	case *message.SitOut:
		sitOut := msg.Envelope.SitOut
		pos := int(sitOut.GetPos())

		this.Table.Seat(pos).State = seat.Idle
		// TODO: fold

	case *message.ComeBack:
		comeBack := msg.Envelope.ComeBack
		pos := int(comeBack.GetPos())

		this.Table.Seat(pos).State = seat.Ready

	case *message.ChatMessage:
		this.Broadcast.All <- msg

	case *message.AddBet:
		if !this.Betting.IsActive() {
			return
		}

		add := msg.Envelope.AddBet

		pos := int(add.GetPos())
		seat := this.Table.Seat(pos)

		betType := bet.Type(add.Bet.GetType().String())
		amount := add.Bet.GetAmount()
		newBet := model.NewBet(betType, amount)

		this.Betting.Bet <- &context.Action{Seat: seat, Bet: newBet}
		this.Broadcast.All <- msg

	case *message.DiscardCards:
		this.Discarding.Discard <- msg

	default:
		log.Printf("Unknown message: %#v", msg.Payload())
	}
}

func (this *Play) handleStateChange(newState State) {
	log.Printf("[play] state %s", newState)

	oldState := this.State
	if oldState != newState {
		this.stateLock.Lock()
		defer this.stateLock.Unlock()

		this.State = newState
		if newState == Active {
			go this.Run()
			this.scheduleNextDeal()
		}
	}
}

func (this *Play) scheduleNextDeal() {
	log.Printf("[play] scheduling next deal in %d seconds", 5)

	this.NextDeal <- time.After(5 * time.Second)
}

func (this *Play) Proto() *message.Play {
	return &message.Play{
		Table: this.Table.Proto(),
		Game:  this.Game.Proto(),
		Stake: this.Stake.Proto(),
	}
}

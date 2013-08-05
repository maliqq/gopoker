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

// Play - play
type Play struct {
	// players action context
	*gameplay.GamePlay

	// finite state machine
	FSM
	// receive protocol messages
	Recv protocol.MessageChannel `json:"-"`
}

// NewPlay - create new play
func NewPlay(variation model.Variation, stake *model.Stake) *Play {
	gp := gameplay.NewGamePlay()
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

	// game
	if variation.IsMixed() {
		mix := variation.(*model.Mix)
		play.Mix = mix
		play.GameRotation = context.NewGameRotation(mix, 1)
		play.Game = play.GameRotation.Current()
	} else {
		play.Game = variation.(*model.Game)
	}
	// create table
	play.Table = model.NewTable(play.Game.TableSize)

	go play.receive()

	return play
}

// String - play to string
func (play *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", play.Game, play.Table)
}

func (play *Play) receive() {
	for {
		select {
		case msg := <-play.Recv:
			play.handleMessage(msg)

		case newState := <-play.stateChange:
			play.handleStateChange(newState)
		}
	}
}

func (play *Play) handleMessage(msg *message.Message) {
	log.Printf(console.Color(console.Yellow, msg.String()))

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

		_, err := play.Table.AddPlayer(player, pos, amount)
		if err == nil {
			// start next deal
		} else {
			log.Printf("[protocol] error: %s", err)
		}
		// retranslate
		play.Broadcast.All <- msg

	case *message.LeaveTable:
		leave := msg.Envelope.LeaveTable
		player := model.Player(leave.GetPlayer())

		play.Table.RemovePlayer(player)
		play.Broadcast.All <- msg
		// TODO: fold & autoplay

	case *message.SitOut:
		sitOut := msg.Envelope.SitOut
		pos := int(sitOut.GetPos())

		play.Table.Seat(pos).State = seat.Idle
		// TODO: fold

	case *message.ComeBack:
		comeBack := msg.Envelope.ComeBack
		pos := int(comeBack.GetPos())

		play.Table.Seat(pos).State = seat.Ready

	case *message.ChatMessage:
		play.Broadcast.All <- msg

	case *message.AddBet:
		if !play.Betting.IsActive() {
			return
		}

		add := msg.Envelope.AddBet

		pos := int(add.GetPos())
		if pos != play.Betting.Pos {
			return
		}
		//seat := play.Table.Seat(pos)

		betType := bet.Type(add.Bet.GetType().String())
		amount := add.Bet.GetAmount()

		play.Betting.Bet <- model.NewBet(betType, amount)
		play.Broadcast.All <- msg

	case *message.DiscardCards:
		play.Discarding.Discard <- msg

	default:
		log.Printf("Unknown message: %#v", msg.Payload())
	}
}

func (play *Play) handleStateChange(newState State) {
	log.Printf("[play] state %s", newState)

	oldState := play.State
	if oldState != newState {
		play.stateLock.Lock()
		defer play.stateLock.Unlock()

		play.State = newState
		if newState == Active {
			go play.Run()
			play.scheduleNextDeal()
		}
	}
}

func (play *Play) scheduleNextDeal() {
	log.Printf("[play] scheduling next deal in %d seconds", 5)

	play.NextDeal <- time.After(5 * time.Second)
}

// Proto - play to protobuf
func (play *Play) Proto() *message.Play {
	return &message.Play{
		Table: play.Table.Proto(),
		Game:  play.Game.Proto(),
		Stake: play.Stake.Proto(),
	}
}

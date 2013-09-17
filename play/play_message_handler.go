package play

import (
	"log"
	"time"
)

import (
	"gopoker/event/message"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/seat"
	"gopoker/util/console"
)

func (play *Play) HandleMessage(msg *message.Message) {
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

func (play *Play) HandleStateChange(newState State) {
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

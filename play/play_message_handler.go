package play

import (
	"log"
	"time"
)

import (
	"gopoker/event"
	"gopoker/event/message"
	"gopoker/model/seat"
	"gopoker/util"
)

func (play *Play) HandleEvent(event *event.Event) {
	log.Printf(util.Color(util.Yellow, event.String()))

	switch msg := event.Message.(type) {
	case message.JoinTable:

		player, pos, amount := msg.Player, msg.Pos, msg.Amount
		_, err := play.Table.AddPlayer(player, pos, amount)
		if err == nil {
			// start next deal
		} else {
			log.Printf("[protocol] error: %s", err)
		}
		// retranslate
		play.Broadcast.Pass(event).All()

	case message.LeaveTable:

		player := msg.Player
		play.Table.RemovePlayer(player)
		play.Broadcast.Pass(event).All()
		// TODO: fold & autoplay

	case message.SitOut:

		pos := msg.Pos
		play.Table.Seat(pos).State = seat.Idle
		// TODO: fold

	case message.ComeBack:

		pos := msg.Pos

		play.Table.Seat(pos).State = seat.Ready

	case message.ChatMessage:

		play.Broadcast.Pass(event).All()

	case message.AddBet:

		if !play.Betting.IsActive() {
			return
		}

		pos := msg.Pos
		if pos != play.Betting.Pos {
			return
		}
		//seat := play.Table.Seat(pos)

		play.Betting.Bet <- msg.Bet
		play.Broadcast.Pass(event).All()

	case message.DiscardCards:

		play.Discarding.Discard <- msg

	default:
		log.Printf("Unknown message: %#v", msg)
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

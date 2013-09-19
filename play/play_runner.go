package play

//
// game main loop
//

import (
	"log"
)

import (
	"gopoker/event/message"
	"gopoker/model"
	"gopoker/model/seat"
	"gopoker/play/context"
	"gopoker/play/gameplay"
	"gopoker/play/street"
)

// Run - main run loop
func (play *Play) Run() {
	log.Printf("started run loop")
Loop:
	for {
		select {
		case timeout := <-play.NextDeal:
			<-timeout
			go play.run()
		case <-play.Exit:
			break Loop
		}
	}
}

func (play *Play) run() {
	// prepare seats
	log.Println("[play] prepare seats")

	for _, s := range play.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play, seat.Fold:
			s.Play()
		}
	}

	// start new deal
	log.Println("[play] start new deal")

	play.Deal = model.NewDeal()
	play.Betting = context.NewBetting()
	if play.Game.Discards {
		play.Discarding = context.NewDiscarding(play.Deal)
	}

	// notify about play start
	play.Broadcast.Notify(&message.PlayStart{
		Game:  play.Game,
		Stake: play.Stake,
		Table: play.Table,
	}).All()

	// rotate game
	if play.Mix != nil {
		log.Println("[play] rotate game")

		if nextGame := play.GameRotation.Next(); nextGame != nil {
			play.Game = nextGame
		}
	}

	// post antes
	if play.Game.HasAnte || play.Stake.HasAnte() {
		log.Println("[play] post antes")

		play.Gameplay.PostAntes()
		play.Gameplay.ResetBetting()
	}

	// post blinds
	if play.Game.HasBlinds {
		log.Println("[play] post blinds")

		play.Gameplay.MoveButton()
		play.Gameplay.PostBlinds()
	}

	// run streets
Streets:
	for _, street := range street.Get(play.Game.Group) {

		log.Printf("[play] street %s\n", street)
		play.Broadcast.Notify(&message.StreetStart{string(street)}).All()

		play.Street = street

		for _, stage := range ByStreet[street] {
			play.Stage = stage.Name

			log.Printf("[play] stage %s\n", stage.Name)

			switch stage.Invoke(play) {
			case gameplay.Next:
				continue
			case gameplay.Stop:
				break Streets
			}
		}
	}

	inPot := play.Table.AllSeats().InPot()
	if len(inPot) == 1 {
		// last player left
		play.Gameplay.Winner(inPot[0])
	} else {
		// showdown
		log.Println("[play] showdown")

		var highHands, lowHands gameplay.ShowdownHands

		if play.Game.Lo != "" {
			lowHands = play.Gameplay.ShowHands(play.Game.Lo, play.Game.HasBoard)
		}

		if play.Game.Hi != "" {
			highHands = play.Gameplay.ShowHands(play.Game.Hi, play.Game.HasBoard)
		}

		play.Gameplay.Winners(highHands, lowHands)
	}

	// deal stop
	log.Println("[play] deal stop")
	play.Broadcast.Notify(&message.PlayStop{}).All()

	play.scheduleNextDeal()
}

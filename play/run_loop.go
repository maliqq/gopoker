package play

//
// game main loop
//

import (
	"log"
)

import (
	"gopoker/model"
	"gopoker/model/seat"
	"gopoker/play/context"
	"gopoker/play/gameplay"
	"gopoker/play/street"
	"gopoker/protocol/message"
)

func (this *Play) Run() {
	log.Printf("started run loop")
Loop:
	for {
		select {
		case timeout := <-this.NextDeal:
			<-timeout
			go this.run()
		case <-this.Exit:
			break Loop
		}
	}
}

func (this *Play) run() {
	// prepare seats
	log.Println("[play] prepare seats")

	for _, s := range this.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play, seat.Fold:
			s.Play()
		}
	}

	// start new deal
	log.Println("[play] start new deal")

	this.Deal = model.NewDeal()
	this.Betting = context.NewBetting()
	if this.Game.Discards {
		this.Discarding = context.NewDiscarding(this.Deal)
	}

	// notify about play start
	this.Broadcast.System <- message.NewMessage(this.dumpProto())

	// rotate game
	if this.Mix != nil {
		log.Println("[play] rotate game")

		if nextGame := this.GameRotation.Next(); nextGame != nil {
			this.Game = nextGame
		}
	}

	// post antes
	if this.Game.HasAnte || this.Stake.HasAnte() {
		log.Println("[play] post antes")

		this.GamePlay.PostAntes()
		this.GamePlay.ResetBetting()
	}

	// post blinds
	if this.Game.HasBlinds {
		log.Println("[play] post blinds")

		this.GamePlay.MoveButton()
		this.GamePlay.PostBlinds()
	}

	// run streets
	for _, street := range street.Get(this.Game.Group) {
		log.Printf("[play] street %s\n", street)
		this.Broadcast.All <- message.NewStreetStart(string(street))

		this.Street = street

		for _, stage := range ByStreet[street] {
			this.Stage = stage.Name

			log.Printf("[play] stage %s\n", stage.Name)

			switch stage.Invoke(this) {
			case gameplay.Next:
				continue
			case gameplay.Stop:
				break
			}
		}
	}

	// showdown
	log.Println("[play] showdown")

	var highHands, lowHands gameplay.ShowdownHands

	if this.Game.Lo != "" {
		lowHands = this.GamePlay.ShowHands(this.Game.Lo, this.Game.HasBoard)
	}

	if this.Game.Hi != "" {
		highHands = this.GamePlay.ShowHands(this.Game.Hi, this.Game.HasBoard)
	}

	this.GamePlay.Winners(highHands, lowHands)

	// deal stop
	log.Println("[play] deal stop")
	this.Broadcast.All <- message.NewMessage(message.PlayStop{})

	this.scheduleNextDeal()
}
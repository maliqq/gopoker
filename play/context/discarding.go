package context

import (
	"log"
)

import (
	"gopoker/model"
	"gopoker/protocol"
)

type Discarding struct {
	*model.Deal
	Seat     *model.Seat
	Required *protocol.RequireDiscard

	Discard chan *protocol.Message
}

func NewDiscarding(d *model.Deal) *Discarding {
	return &Discarding{
		Required: &protocol.RequireDiscard{},
		Discard:  make(chan *protocol.Message),
	}
}
func (this *Discarding) RequireDiscard(pos int, seat *model.Seat) *protocol.Message {
	this.Seat = seat
	this.Required.Pos = pos
	return protocol.NewRequireDiscard(this.Required)
}

func (this *Discarding) Start() {
	for {
		select {
		case msg := <-this.Discard:
			payload := msg.Payload.(protocol.DiscardCards)
			seat := this.Seat
			cards := payload.Cards

			if len(cards) == 0 {
				log.Printf("[discarding] Player %s stands pat", seat.Player)
			} else {
				log.Printf("[discarding] Player %s discards %s", seat.Player, cards)
			}
		}
	}
}

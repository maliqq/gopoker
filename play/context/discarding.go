package context

import (
	"log"
)

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/protocol"
)

type Discarding struct {
	*model.Deal
	Required *protocol.RequireDiscard

	Receive chan *protocol.Message
}

func NewDiscarding(d *model.Deal) *Discarding {
	return &Discarding{
		Required: &protocol.RequireDiscard{},
		Receive:  make(chan *protocol.Message),
	}
}

func (this *Discarding) RequireDiscard(pos int) *protocol.Message {
	this.Required.Pos = pos
	return protocol.NewRequireDiscard(this.Required)
}

func (this *Discarding) Add(seat *model.Seat, msg *protocol.Message) (*model.Player, *poker.Cards) {
	payload := msg.Payload.(protocol.DiscardCards)
	cards := payload.Cards

	if len(cards) == 0 {
		log.Printf("[discarding] Player %s stands pat", seat.Player)
	} else {
		log.Printf("[discarding] Player %s discards %s", seat.Player, cards)
	}

	return seat.Player, &cards
}

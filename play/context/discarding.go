package context

import (
	"log"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model"
	"gopoker/protocol/message"
)

type Discarding struct {
	*model.Deal
	Seat     *model.Seat
	Required *message.RequireDiscard

	Discard chan *message.Message
}

func NewDiscarding(d *model.Deal) *Discarding {
	return &Discarding{
		Required: &message.RequireDiscard{},
		Discard:  make(chan *message.Message),
	}
}

func (this *Discarding) RequireDiscard(pos int, seat *model.Seat) *message.Message {
	this.Seat = seat
	this.Required.Pos = proto.Int32(int32(pos))
	return message.NewRequireDiscard(this.Required)
}

func (this *Discarding) Start() {
	for {
		select {
		case msg := <-this.Discard:
			payload := msg.Envelope.DiscardCards
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

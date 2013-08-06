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

// Discarding - discarding context
type Discarding struct {
	*model.Deal
	Seat     *model.Seat
	Required *message.RequireDiscard

	Discard chan *message.Message
}

// NewDiscarding - create new discarding context
func NewDiscarding(d *model.Deal) *Discarding {
	return &Discarding{
		Required: &message.RequireDiscard{},
		Discard:  make(chan *message.Message),
	}
}

// RequireDiscard - require discard
func (discarding *Discarding) RequireDiscard(pos int, seat *model.Seat) *message.Message {
	discarding.Seat = seat
	discarding.Required.Pos = proto.Int32(int32(pos))
	return message.NotifyRequireDiscard(discarding.Required)
}

// Start - start discarding
func (discarding *Discarding) Start() {
	for {
		select {
		case msg := <-discarding.Discard:
			payload := msg.Envelope.DiscardCards
			seat := discarding.Seat
			cards := payload.Cards

			if len(cards) == 0 {
				log.Printf("[discarding] Player %s stands pat", seat.Player)
			} else {
				log.Printf("[discarding] Player %s discards %s", seat.Player, cards)
			}
		}
	}
}

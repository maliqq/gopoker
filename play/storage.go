package play

import (
	"log"
	"time"
)

import (
	"gopoker/protocol"
	"gopoker/protocol/message"
	"gopoker/storage"
)

type Storage struct {
	*storage.PlayStore
	Current *storage.Play
	Recv    protocol.MessageChannel
}

func NewStorage(ps *storage.PlayStore) *Storage {
	storage := &Storage{
		PlayStore: ps,
		Recv:      make(protocol.MessageChannel),
	}

	go storage.receive()

	return storage
}

func (this *Storage) receive() {
	for {
		msg := <-this.Recv
		this.handle(msg)
	}
}

func (this *Storage) handle(msg *message.Message) {
	switch msg.Payload().(type) {
	case *message.PlayStart:
		this.Current = this.NewPlay()
		this.Current.Play = msg.Envelope.PlayStart.Play

	case *message.PlayStop:
		this.Current.Stop = time.Now()

		log.Printf("[storage] saving %+v", this.Current)

		this.PlayStore.Collection("plays").Insert(this.Current)

	case *message.AddBet:
		this.Current.Log = append(this.Current.Log, msg)

	case *message.ShowHand:
		show := msg.Envelope.ShowHand
		this.Current.KnownCards[show.GetPlayer()] = show.GetCards()

	case *message.Winner:
		winner := msg.Envelope.Winner
		this.Current.Winners[winner.GetPlayer()] = winner.GetAmount()

	default:
		log.Printf("[storage] got %s", msg)
	}
}

func (this *Storage) NewPlay() *storage.Play {
	return &storage.Play{
		Id:         storage.NewObjectId(),
		Start:      time.Now(),
		Winners:    map[string]float64{},
		KnownCards: map[string]message.Cards{},
	}
}

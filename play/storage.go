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

func (stor *Storage) receive() {
	for {
		msg := <-stor.Recv
		stor.handle(msg)
	}
}

func (stor *Storage) handle(msg *message.Message) {
	switch msg.Payload().(type) {
	case *message.PlayStart:
		stor.Current = stor.NewPlay()
		stor.Current.Play = msg.Envelope.PlayStart.Play

	case *message.PlayStop:
		stor.Current.Stop = time.Now()

		log.Printf("[storage] saving %+v", stor.Current)

		stor.PlayStore.Collection("plays").Insert(stor.Current)

	case *message.ShowHand:
		show := msg.Envelope.ShowHand
		stor.Current.KnownCards[show.GetPlayer()] = show.GetCards()

	case *message.Winner:
		winner := msg.Envelope.Winner
		stor.Current.Winners[winner.GetPlayer()] = winner.GetAmount()

	default:
		log.Printf("[storage] got %s", msg)
	}
}

func (stor *Storage) NewPlay() *storage.Play {
	return &storage.Play{
		Id:         storage.NewObjectId(),
		Start:      time.Now(),
		Winners:    map[string]float64{},
		KnownCards: map[string]message.Cards{},
	}
}

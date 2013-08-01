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
		Current:   &storage.Play{},
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
		this.Current.Id = storage.NewObjectId()
		this.Current.Start = time.Now()
		this.Current.Play = msg.Envelope.PlayStart.Play

	case *message.PlayStop:
		this.Current.Stop = time.Now()

		log.Printf("[storage] saving %+v", this.Current)

		this.PlayStore.Collection("plays").Insert(this.Current)

	default:
		log.Printf("[storage] got %s", msg)
	}
}

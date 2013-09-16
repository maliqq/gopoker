package play

import (
	"log"
	"time"
)

import (
	"gopoker/exch/message"
	"gopoker/storage"
)

func (stor *Storage) HandleMessage(msg *message.Message) {
	switch msg.Payload().(type) {
	case *message.PlayStart:
		stor.Current = storage.NewPlayHistoryEntry()
		stor.Current.Play = msg.Envelope.PlayStart.Play

	case *message.PlayStop:
		stor.Current.Stop = time.Now()

		log.Printf("[storage] saving %+v", stor.Current)

		stor.History.Store(stor.Current)

	case *message.AddBet:
		stor.Current.Log = append(stor.Current.Log, msg)

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

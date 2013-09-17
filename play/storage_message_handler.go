package play

import (
	"log"
	"time"
)

import (
	"gopoker/event"
	"gopoker/event/message"
	"gopoker/storage"
)

func (stor *Storage) HandleEvent(event *event.Event) {
	switch msg := event.Message.(type) {
	case message.PlayStart:
		stor.Current = storage.NewPlayHistoryEntry()
		// FIXME
		//stor.Current.Play = event.Message.(message.PlayStart.Play)

	case message.PlayStop:
		stor.Current.Stop = time.Now()

		log.Printf("[storage] saving %+v", stor.Current)

		stor.History.Store(stor.Current)

	case message.AddBet:
		stor.Current.Log = append(stor.Current.Log, msg)

	case message.ShowHand:
		player, cards := msg.Player, msg.Cards
		stor.Current.KnownCards[player] = cards

	case message.Winner:
		player, amount := msg.Player, msg.Amount
		stor.Current.Winners[player] = amount

	default:
		log.Printf("[storage] got %#v", msg)
	}
}

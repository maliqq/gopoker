package server

import (
	"log"
	"os"
	"path"
)

import (
	"gopoker/model"
	"gopoker/play"
	"gopoker/server/rpc_service"
	"gopoker/storage"
)

type Room struct {
	Id string
	*play.Play
	Waiting  map[string]*Session
	Watchers map[string]*Session
}

func NewRoom(createRoom *rpc_service.CreateRoom) *Room {
	var variation model.Variation
	if createRoom.Game != nil {
		variation = createRoom.Game.WithDefaults()
	}

	if createRoom.Mix != nil {
		variation = createRoom.Mix.WithDefaults()
	}

	tableSize := createRoom.TableSize

	var maxTableSize int
	if variation.IsMixed() {
		maxTableSize = model.MixedGameMaxTableSize
	} else {
		maxTableSize = variation.(*model.Game).MaxTableSize
	}
	if tableSize == 0 || tableSize > maxTableSize {
		tableSize = maxTableSize
	}

	table := model.NewTable(tableSize)
	stake := model.NewStake(createRoom.BetSize)
	newPlay := play.NewPlay(variation, stake, table)

	room := &Room{
		Id:       createRoom.Id,
		Play:     newPlay,
		Waiting:  map[string]*Session{},
		Watchers: map[string]*Session{},
	}

	return room
}

func (r *Room) createLogger(dir string) {
	f, err := os.OpenFile(path.Join(dir, r.Id+".log"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal("cant open logger file", err)
	}
	logger := play.NewLogger(f)
	r.Broadcast.Broker.BindSystem("logger", &logger.Recv)
}

func (r *Room) createStorage(ps *storage.PlayStore) {
	storage := play.NewStorage(ps)
	r.Broadcast.Broker.BindSystem("storage", &storage.Recv)
}

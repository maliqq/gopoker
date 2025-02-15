package storage

import (
	"log"
	"os"
	"time"
)

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

import (
	"gopoker/event/message"
	"gopoker/model"
	"gopoker/poker"
)

const DefaultCollectionName = "play_history"

// PlayHistoryConfig - MongoDB store config
type PlayHistoryConfig struct {
	Host           string
	Database       string
	CollectionName string
}

// PlayHistory - MongoDB store
type PlayHistory struct {
	Session *mgo.Session
	Config  *PlayHistoryConfig
}

// Play - play data dump
type PlayHistoryEntry struct {
	ID         bson.ObjectId `bson:"_id"`
	Start      time.Time
	Stop       time.Time
	Play       *message.PlayStart `bson:"play"`
	Winners    map[model.Player]float64
	KnownCards map[model.Player]poker.Cards
	Pot        float64
	Rake       float64
	Log        []message.Message
}

// NewPlay - create new play
func NewPlayHistoryEntry() *PlayHistoryEntry {
	return &PlayHistoryEntry{
		ID:         NewObjectID(),
		Start:      time.Now(),
		Winners:    map[model.Player]float64{},
		KnownCards: map[model.Player]poker.Cards{},
	}
}

// NewObjectID - create new object id
func NewObjectID() bson.ObjectId {
	return bson.NewObjectId()
}

// Debug - debug mode
var Debug = false

// OpenPlayHistory - open MongoDB
func OpenPlayHistory(config *PlayHistoryConfig) (*PlayHistory, error) {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	if Debug {
		mgo.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
		mgo.SetDebug(true)
	}

	store := &PlayHistory{
		Session: session,
		Config:  config,
	}

	return store, nil
}

// Close - close MongoDB
func (history *PlayHistory) Close() {
	history.Session.Close()
}

// Database - get database
func (history *PlayHistory) database() *mgo.Database {
	return history.Session.DB(history.Config.Database)
}

// Collection - get collection by name
func (history *PlayHistory) collection() *mgo.Collection {
	collectionName := history.Config.CollectionName
	if collectionName == "" {
		collectionName = DefaultCollectionName
	}
	return history.database().C(collectionName)
}

// FindPlayByID - find play data by id
func (history *PlayHistory) Find(id string) (*PlayHistoryEntry, error) {
	var document PlayHistoryEntry

	query := history.collection().Find(bson.M{"_id": bson.ObjectIdHex(id)})
	err := query.One(&document)

	return &document, err
}

func (history *PlayHistory) Store(document *PlayHistoryEntry) {
	history.collection().Insert(document)
}

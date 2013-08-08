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
	"gopoker/exch/message"
)

// PlayStoreConfig - MongoDB store config
type PlayStoreConfig struct {
	Host     string
	Database string
}

// PlayStore - MongoDB store
type PlayStore struct {
	Session *mgo.Session
	Config  *PlayStoreConfig
}

// Play - play data dump
type Play struct {
	ID         bson.ObjectId `bson:"_id"`
	Start      time.Time
	Stop       time.Time
	Play       *message.Play `bson:"play"`
	Winners    map[string]float64
	KnownCards map[string]message.Cards
	Pot        float64
	Rake       float64
	Log        []*message.Message
}

// NewPlay - create new play
func NewPlay() *Play {
	return &Play{
		ID:         NewObjectID(),
		Start:      time.Now(),
		Winners:    map[string]float64{},
		KnownCards: map[string]message.Cards{},
	}
}

// NewObjectID - create new object id
func NewObjectID() bson.ObjectId {
	return bson.NewObjectId()
}

// Debug - debug mode
var Debug = false

// OpenPlayStore - open MongoDB
func OpenPlayStore(config *PlayStoreConfig) (*PlayStore, error) {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	if Debug {
		mgo.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
		mgo.SetDebug(true)
	}

	store := &PlayStore{
		Session: session,
		Config:  config,
	}

	return store, nil
}

// Close - close MongoDB
func (ps *PlayStore) Close() {
	ps.Session.Close()
}

// Database - get database
func (ps *PlayStore) Database() *mgo.Database {
	return ps.Session.DB(ps.Config.Database)
}

// Collection - get collection by name
func (ps *PlayStore) Collection(collection string) *mgo.Collection {
	return ps.Database().C(collection)
}

// FindPlayByID - find play data by id
func (ps *PlayStore) FindPlayByID(id string) (*Play, error) {
	var play Play

	plays := ps.Collection("plays")
	query := plays.Find(bson.M{"_id": bson.ObjectIdHex(id)})
	err := query.One(&play)

	return &play, err
}

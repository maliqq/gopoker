package storage

import (
	"database/sql"
	_ "github.com/bmizerany/pq" // import pq sql
)

// StoreConfig - SQL store config
type StoreConfig struct {
	Driver           string
	ConnectionString string
}

// Store - SQL store
type Store struct {
	*sql.DB
}

const (
	// DefaultDriver - default SQL driver
	DefaultDriver = "postgres"
)

// OpenStore - open SQL store
func OpenStore(config *StoreConfig) (*Store, error) {
	store := &Store{}

	var err error
	store.DB, err = sql.Open(config.Driver, config.ConnectionString)

	return store, err
}

// Close - close SQL store
func (store *Store) Close() {
	store.DB.Close()
}

package storage

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
)

type StoreConfig struct {
	Driver           string
	ConnectionString string
}

type Store struct {
	*sql.DB
}

const (
	DefaultDriver = "postgres"
)

func Open(config *StoreConfig) (*Store, error) {
	store := &Store{}

	var err error
	store.DB, err = sql.Open(config.Driver, config.ConnectionString)

	return store, err
}

func (store *Store) Close() {
	store.DB.Close()
}

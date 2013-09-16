package storage

import (
	"database/sql"
	_ "github.com/bmizerany/pq" // import pq sql
	"log"
)

import (
	"gopoker/util"
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

type User struct {
	PlayerID          string
	Username          string
	PasswordEncrypted string
	PasswordSalt      string
}

func (u *User) MatchPassword(password string) bool {
	return util.HashPasswordWithSalt(u.PasswordEncrypted, u.PasswordSalt) == password
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

func (store *Store) FindUserByUsername(username string) *User {
	query := `SELECT players.uuid, users.encrypted_password, users.password_salt
		FROM users
		LEFT OUTER JOIN players ON players.user_id = users.id
		WHERE users.username=$1`
	row := store.DB.QueryRow(query, username)

	user := User{
		Username: username,
	}

	var passwordEncrypted, passwordSalt sql.NullString
	if err := row.Scan(&user.PlayerID, &passwordEncrypted, &passwordSalt); err != nil {
		if err != sql.ErrNoRows {
			log.Printf("[store] error querying user: %s", err)
		}
		return nil
	}

	if passwordEncrypted.Valid {
		user.PasswordEncrypted = passwordEncrypted.String
	}
	if passwordSalt.Valid {
		user.PasswordSalt = passwordSalt.String
	}

	return &user
}

package storage

import (
  "github.com/hoisie/redis"
)

import (
  "fmt"
  "encoding/json"
)

import (
  "gopoker/model"
)

type SessionStoreConfig struct {
  Address string
  Host string
  Port int
  KeyPrefix string
}

type SessionStore struct {
  Config *SessionStoreConfig
  redisClient *redis.Client
}

type SessionData struct {
  PlayerID model.Player `json:player_id`
}

func OpenSessionStore(config *SessionStoreConfig) (*SessionStore, error) {
  var redisClient redis.Client

  addr := ""
  if config.Address != "" {
    addr = config.Address
  } else if config.Host != "" && config.Port != 0 {
    addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
  }
  if addr != "" {
    redisClient.Addr = addr
  }

  store := SessionStore{
    Config: config,
    redisClient: &redisClient,
  }

  return &store, nil
}

func (store *SessionStore) Get(key string) *SessionData {
  data, found := store.redisClient.Get(store.Config.KeyPrefix + key)

  if found == nil {
    return nil
  }

  var sessionData SessionData
  json.Unmarshal([]byte(data), &sessionData)

  return &sessionData
}

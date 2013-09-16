package server

import (
	"gopoker/model"
	"gopoker/storage"
)

func (node *Node) Auth(key string) (model.Player, bool) {
	result := node.SessionStore.Get(key)

	var player model.Player

	if result != nil {
		player = result.PlayerID
	}

	return player, player == ""
}

func (node *Node) Login(username string, password string) (string, bool) {
	user := node.Store.FindUserByUsername(username)
	var sessionID string
	if user != nil {
		if user.MatchPassword(password) {
			sessionID = NewSessionID()
			node.SessionStore.Set(sessionID, storage.SessionData{
				PlayerID: model.Player(user.PlayerID),
			})
		}
	}
	return sessionID, sessionID == ""
}

func (node *Node) Logout(sessionID string) {
	node.SessionStore.Del(sessionID)
}

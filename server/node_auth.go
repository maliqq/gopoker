package server

import (
	"gopoker/model"
	"gopoker/storage"
)

func (node *Node) Auth(key string) (model.Player, bool) {
	result := node.SessionStore.Get(key)

	var player model.Player

	if result != nil {
		player = result.Player
	}

	return player, player == ""
}

func (node *Node) Login(username string, password string) (model.Guid, bool) {
	user := node.Store.FindUserByUsername(username)
	var sessionID model.Guid
	if user != nil {
		if user.MatchPassword(password) {
			sessionID = NewSessionID()
			node.SessionStore.Set(string(sessionID), storage.SessionData{
				Player: model.Player(user.Player),
			})
		}
	}
	return sessionID, sessionID == ""
}

func (node *Node) Logout(session model.Guid) {
	node.SessionStore.Del(string(session))
}

package noderpc

import (
	"gopoker/model"
	"gopoker/play/mode"
	"gopoker/protocol/message"
)

// CallResult - RPC call result
type CallResult struct {
	Status  string
	Message string
}

// CreateRoom - create room request
type CreateRoom struct {
	ID        string
	BetSize   float64
	Game      *model.Game
	Mix       *model.Mix
}

// RequestRoom - get room by id
type RequestRoom struct {
	ID string
}

// StartRoom - start room by id
type StartRoom struct {
	ID   string
	Mode mode.Type
}

// NotifyRoom - send protocol message
type NotifyRoom struct {
	ID      string
	Message *message.Message
}

// ConnectGateway - connect node ZMQ gateway
type ConnectGateway struct {
	RoomID   string
	PlayerID string
}

// DisconnectGateway - disconnect node ZMQ gateway
type DisconnectGateway struct {
	PlayerID string
}

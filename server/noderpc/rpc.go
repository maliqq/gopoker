package noderpc

import (
	"gopoker/event"
	"gopoker/model"
	"gopoker/play/mode"
)

// CallResult - RPC call result
type CallResult struct {
	Status  string
	Message string
}

// CreateRoom - create room request
type CreateRoom struct {
	Guid    model.Guid
	BetSize float64
	Game    *model.Game
	Mix     *model.Mix
}

// RequestRoom - get room by id
type RequestRoom struct {
	Guid model.Guid
}

// StartRoom - start room by id
type StartRoom struct {
	Guid model.Guid
	Mode mode.Type
}

// NotifyRoom - send protocol message
type NotifyRoom struct {
	Guid  model.Guid
	Event *event.Event
}

// ConnectGateway - connect node ZMQ gateway
type ConnectGateway struct {
	Room   model.Guid
	Player model.Player
}

// DisconnectGateway - disconnect node ZMQ gateway
type DisconnectGateway struct {
	Player model.Player
}

type Login struct {
	Username string
	Password string
}

type LoginResult struct {
	Session model.Guid
	Success bool
}

type Logout struct {
	Session model.Guid
}

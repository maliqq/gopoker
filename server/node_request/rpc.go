package node_request

import (
	"gopoker/model"
	"gopoker/play/mode"
)
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

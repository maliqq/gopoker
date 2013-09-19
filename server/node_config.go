package server

import (
	"time"
)

import (
	"gopoker/storage"
)

// Defaults
const (
	DefaultAPIPath       = "/_api"
	DefaultRPCPath       = "/_rpc"
	DefaultWebSocketPath = "/_ws"
)

// HTTPConfig - http config
type HTTPConfig struct {
	Addr          string
	APIPath       string
	RPCPath       string
	WebSocketPath string
}

// RPCConfig - RPC config
type RPCConfig struct {
	Addr    string
	Timeout time.Duration
}

type ZMQConfig struct {
	Publisher string
	Receiver  string
}

// Config - node config
type Config struct {
	Logdir string

	HTTP *HTTPConfig
	RPC  *RPCConfig

	ZMQ   *ZMQConfig
	Stomp string

	Store        *storage.StoreConfig
	PlayHistory  *storage.PlayHistoryConfig
	SessionStore *storage.SessionStoreConfig
}

// APIPathOr - API path or default
func (httpConfig *HTTPConfig) GetAPIPath() string {
	apiPath := httpConfig.APIPath
	if apiPath == "" {
		return DefaultAPIPath
	}
	return apiPath
}

// WebSocketPathOr - websocket path or default
func (httpConfig *HTTPConfig) GetWebSocketPath() string {
	webSocketPath := httpConfig.WebSocketPath
	if webSocketPath == "" {
		return DefaultWebSocketPath
	}
	return webSocketPath
}

// RPCPathOr - RPC path or default
func (httpConfig *HTTPConfig) GetRPCPath() string {
	rpcPath := httpConfig.RPCPath
	if rpcPath == "" {
		return DefaultRPCPath
	}
	return rpcPath
}

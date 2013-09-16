package server

import (
	"time"
)

import (
	"gopoker/storage"
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

// Config - node config
type Config struct {
	Logdir string

	HTTP *HTTPConfig
	RPC  *RPCConfig

	ZMQ   string
	Stomp string

	Store        *storage.StoreConfig
	PlayHistory  *storage.PlayHistoryConfig
	SessionStore *storage.SessionStoreConfig
}

// APIPathOr - API path or default
func (httpConfig *HTTPConfig) APIPathOr(defaultPath string) string {
	apiPath := httpConfig.APIPath
	if apiPath == "" {
		return defaultPath
	}
	return apiPath
}

// WebSocketPathOr - websocket path or default
func (httpConfig *HTTPConfig) WebSocketPathOr(defaultPath string) string {
	webSocketPath := httpConfig.WebSocketPath
	if webSocketPath == "" {
		return defaultPath
	}
	return webSocketPath
}

// RPCPathOr - RPC path or default
func (httpConfig *HTTPConfig) RPCPathOr(defaultPath string) string {
	rpcPath := httpConfig.RPCPath
	if rpcPath == "" {
		return defaultPath
	}
	return rpcPath
}

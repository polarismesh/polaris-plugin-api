package discoverevent

import (
	"time"

	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"
)

// ConfigEntry 单个插件配置
type ConfigEntry struct {
	Name   string                 `yaml:"name"`
	Option map[string]interface{} `yaml:"option"`
}

// InstanceEventType 探测事件类型
type InstanceEventType string

const (
	// EventDiscoverNone empty discover event
	EventDiscoverNone InstanceEventType = "EventDiscoverNone"
	// EventInstanceOnline instance becoming online
	EventInstanceOnline InstanceEventType = "InstanceOnline"
	// EventInstanceTurnUnHealth Instance becomes unhealthy
	EventInstanceTurnUnHealth InstanceEventType = "InstanceTurnUnHealth"
	// EventInstanceTurnHealth Instance becomes healthy
	EventInstanceTurnHealth InstanceEventType = "InstanceTurnHealth"
	// EventInstanceOpenIsolate Instance is in isolation
	EventInstanceOpenIsolate InstanceEventType = "InstanceOpenIsolate"
	// EventInstanceCloseIsolate Instance shutdown isolation state
	EventInstanceCloseIsolate InstanceEventType = "InstanceCloseIsolate"
	// EventInstanceOffline Instance offline
	EventInstanceOffline InstanceEventType = "InstanceOffline"
	// EventInstanceSendHeartbeat Instance send heartbeat package to server
	EventInstanceSendHeartbeat InstanceEventType = "InstanceSendHeartbeat"
	// EventInstanceUpdate Instance metadata and info update event
	EventInstanceUpdate InstanceEventType = "InstanceUpdate"
	// EventClientOffline .
	EventClientOffline InstanceEventType = "ClientOffline"
)

// InstanceEvent 服务实例事件
type InstanceEvent struct {
	Id         string
	SvcId      string
	Namespace  string
	Service    string
	Instance   *apiservice.Instance
	EType      InstanceEventType
	CreateTime time.Time
	MetaData   map[string]string
}

// DiscoverChannel is used to receive discover events from the agent
type DiscoverChannel interface {
	// Name .
	Name() string
	// Initialize .
	Initialize(c *ConfigEntry) error
	// Destroy .
	Destroy() error
	// PublishEvent Release a service event
	PublishEvent(event *InstanceEvent)
}

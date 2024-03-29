/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package healthcheck

import (
	"context"
	"fmt"
)

var (
	slots = make(map[string]HealthChecker)
)

// Register 注册插件
func Register(name string, plugin HealthChecker) {
	if _, exist := slots[name]; exist {
		panic(fmt.Sprintf("existed plugin: name=%v", name))
	}
	slots[name] = plugin
}

func Get(name string) (HealthChecker, bool) {
	server, exist := slots[name]
	return server, exist
}

// ReportRequest report heartbeat request
type ReportRequest struct {
	QueryRequest
	LocalHost  string
	CurTimeSec int64
	Count      int64
}

// CheckRequest check heartbeat request
type CheckRequest struct {
	QueryRequest
	ExpireDurationSec uint32
	CurTimeSec        func() int64
}

// CheckResponse check heartbeat response
type CheckResponse struct {
	Healthy              bool
	LastHeartbeatTimeSec int64
	StayUnchanged        bool
	Regular              bool
}

// QueryRequest query heartbeat request
type QueryRequest struct {
	InstanceId string
	Host       string
	Port       uint32
	Healthy    bool
}

// BatchQueryRequest batch query heartbeat request
type BatchQueryRequest struct {
	Requests []*QueryRequest
}

// QueryResponse query heartbeat response
type QueryResponse struct {
	Server           string
	Exists           bool
	LastHeartbeatSec int64
	Count            int64
}

// BatchQueryResponse batch query heartbeat response
type BatchQueryResponse struct {
	Responses []*QueryResponse
}

// AddCheckRequest add check request
type AddCheckRequest struct {
	Instances []string
	LocalHost string
}

// HealthCheckType health check type
type HealthCheckType int32

const (
	HealthCheckerHeartbeat HealthCheckType = iota + 1
)

// ConfigEntry 单个插件配置
type ConfigEntry struct {
	Name   string                 `yaml:"name"`
	Option map[string]interface{} `yaml:"option"`
}

// HealthChecker health checker plugin interface
type HealthChecker interface {
	// Name .
	Name() string
	// Initialize .
	Initialize(c *ConfigEntry) error
	// Destroy .
	Destroy() error
	// Type for health check plugin, only one same type plugin is allowed
	Type() HealthCheckType
	// Report process heartbeat info report
	Report(ctx context.Context, request *ReportRequest) error
	// Check process the instance check
	Check(request *CheckRequest) (*CheckResponse, error)
	// Query queries the heartbeat time
	Query(ctx context.Context, request *QueryRequest) (*QueryResponse, error)
	// BatchQuery batch queries the heartbeat time
	BatchQuery(ctx context.Context, request *BatchQueryRequest) (*BatchQueryResponse, error)
	// Suspend health checker for entire expired duration manually
	Suspend()
	// SuspendTimeSec get the suspend time in seconds
	SuspendTimeSec() int64
	// Delete delete the id
	Delete(ctx context.Context, id string) error
}

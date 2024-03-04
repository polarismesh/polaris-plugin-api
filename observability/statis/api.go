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

package statis

import (
	"fmt"
	"strconv"
	"time"
)

var (
	slots = make(map[string]Statis)
)

// Register 注册插件
func Register(name string, plugin Statis) {
	if _, exist := slots[name]; exist {
		panic(fmt.Sprintf("existed plugin: name=%v", name))
	}
	slots[name] = plugin
}

func Get(name string) (Statis, bool) {
	server, exist := slots[name]
	return server, exist
}

// ConfigEntry 单个插件配置
type ConfigEntry struct {
	Name   string                 `yaml:"name"`
	Option map[string]interface{} `yaml:"option"`
}

// Statis Statistical plugin interface
type Statis interface {
	// Name .
	Name() string
	// Initialize .
	Initialize(c *ConfigEntry) error
	// Destroy .
	Destroy() error
	// ReportCallMetrics report call metrics info
	ReportCallMetrics(metric CallMetric)
	// ReportDiscoveryMetrics report discovery metrics
	ReportDiscoveryMetrics(metric ...DiscoveryMetric)
	// ReportConfigMetrics report config_center metrics
	ReportConfigMetrics(metric ...ConfigMetrics)
	// ReportDiscoverCall report discover service times
	ReportDiscoverCall(metric ClientDiscoverMetric)
}

type TrafficDirection string

const (
	// TrafficDirectionInBound .
	TrafficDirectionInBound TrafficDirection = "INBOUND"
	// TrafficDirectionOutBound .
	TrafficDirectionOutBound TrafficDirection = "OUTBOUND"
)

const (
	LabelApi      = "api"
	LabelProtocol = "protocol"
	LabelErrCode  = "err_code"
)

// CallMetricType .
type CallMetricType string

type CallMetric struct {
	Type             CallMetricType
	API              string
	Protocol         string
	Code             int
	Times            int
	Success          bool
	Duration         time.Duration
	Labels           map[string]string
	TrafficDirection TrafficDirection
}

func (m CallMetric) GetLabels() map[string]string {
	if len(m.Labels) == 0 {
		m.Labels = map[string]string{}
	}
	m.Labels[LabelApi] = m.API
	m.Labels[LabelProtocol] = m.Protocol
	m.Labels[LabelErrCode] = strconv.FormatInt(int64(m.Code), 10)
	return m.Labels
}

type DiscoveryMetricType string

const (
	ClientMetrics   DiscoveryMetricType = "client"
	ServiceMetrics  DiscoveryMetricType = "service"
	InstanceMetrics DiscoveryMetricType = "instance"
)

type DiscoveryMetric struct {
	Type     DiscoveryMetricType
	Total    int64
	Abnormal int64
	Offline  int64
	Online   int64
	Isolate  int64
	Labels   map[string]string
}

type ClientDiscoverMetric struct {
	ClientIP  string
	Action    string
	Namespace string
	Resource  string
	Revision  string
	Timestamp int64
	CostTime  int64
	Success   bool
}

func (c ClientDiscoverMetric) String() string {
	revision := c.Revision
	if revision == "" {
		revision = "-"
	}
	return fmt.Sprintf("%s|%s|%s|%s|%s|%s|%dms|%+v", c.ClientIP, c.Action, c.Namespace, c.Resource,
		revision, time.Unix(c.Timestamp/1000, 0).Format(time.DateTime), c.CostTime, c.Success)
}

type ConfigMetricType string

const (
	ConfigGroupMetric ConfigMetricType = "config_group"
	FileMetric        ConfigMetricType = "file"
	ReleaseFileMetric ConfigMetricType = "release_file"
)

type ConfigMetrics struct {
	Type    ConfigMetricType
	Total   int64
	Release int64
	Labels  map[string]string
}

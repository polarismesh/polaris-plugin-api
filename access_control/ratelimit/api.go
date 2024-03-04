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

package ratelimit

import (
	"fmt"
)

var (
	slots = make(map[string]Ratelimit)
)

// Register 注册插件
func Register(name string, plugin Ratelimit) {
	if _, exist := slots[name]; exist {
		panic(fmt.Sprintf("existed plugin: name=%v", name))
	}
	slots[name] = plugin
}

func Get(name string) (Ratelimit, bool) {
	plugin, exist := slots[name]
	return plugin, exist
}

// RatelimitType rate limit type
type RatelimitType int

const (
	// IPRatelimit Based on IP flow control
	IPRatelimit RatelimitType = iota + 1

	// APIRatelimit Based on interface-level flow control
	APIRatelimit

	// ServiceRatelimit Based on Service flow control
	ServiceRatelimit

	// InstanceRatelimit Based on Instance flow control
	InstanceRatelimit
)

// RatelimitStr rate limit string map
var RatelimitStr = map[RatelimitType]string{
	IPRatelimit:       "ip-limit",
	APIRatelimit:      "api-limit",
	ServiceRatelimit:  "service-limit",
	InstanceRatelimit: "instance-limit",
}

// ConfigEntry 单个插件配置
type ConfigEntry struct {
	Name   string                 `yaml:"name"`
	Option map[string]interface{} `yaml:"option"`
}

// Ratelimit Ratelimit plugin interface
type Ratelimit interface {
	// Name .
	Name() string
	// Initialize .
	Initialize(c *ConfigEntry) error
	// Destroy .
	Destroy() error
	// Allow Whether to allow access, true: allow, FALSE: not allowing Todo
	// - Parameter ratingype is the type of current limits, and the ID is the key that limits the current
	// - If RateType is Ratelimitip, the ID is IP, RateType is Ratelimitservice, and the ID is
	//  IP_NAMESPACE_SERVICE or IP_SERVICEID
	Allow(typ RatelimitType, resource string) bool
}

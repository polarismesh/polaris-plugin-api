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

package history

import (
	"fmt"
	"time"
)

var (
	slots = make(map[string]History)
)

// Register 注册插件
func Register(name string, plugin History) {
	if _, exist := slots[name]; exist {
		panic(fmt.Sprintf("existed plugin: name=%v", name))
	}
	slots[name] = plugin
}

func Get(name string) (History, bool) {
	server, exist := slots[name]
	return server, exist
}

// ConfigEntry 单个插件配置
type ConfigEntry struct {
	Name   string                 `yaml:"name"`
	Option map[string]interface{} `yaml:"option"`
}

// History 历史记录插件
type History interface {
	// Name .
	Name() string
	// Initialize .
	Initialize(c *ConfigEntry) error
	// Destroy .
	Destroy() error
	// Record .
	Record(entry *RecordEntry)
}

// OperationType Operating type
type OperationType string

// Resource Operating resources
type Resource string

// RecordEntry Operation records
type RecordEntry struct {
	ResourceType  Resource
	ResourceName  string
	Namespace     string
	Operator      string
	OperationType OperationType
	Detail        string
	Server        string
	HappenTime    time.Time
}

func (r *RecordEntry) String(format func(time.Time) string) string {
	return fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s",
		format(r.HappenTime),
		r.ResourceType,
		r.ResourceName,
		r.Namespace,
		r.OperationType,
		r.Operator,
		r.Detail,
		r.Server,
	)
}

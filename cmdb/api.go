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

package plugin

import (
	apimodel "github.com/polarismesh/specification/source/go/api/v1/model"
)

// Location cmdb信息，对应内存结构体
type Location struct {
	Proto    *apimodel.Location
	RegionID uint32
	ZoneID   uint32
	CampusID uint32
	Valid    bool
}

// ConfigEntry 单个插件配置
type ConfigEntry struct {
	Name   string                 `yaml:"name"`
	Option map[string]interface{} `yaml:"option"`
}

// CMDB CMDB插件接口
type CMDB interface {
	// Name .
	Name() string
	// Initialize .
	Initialize(c *ConfigEntry) error
	// Destroy .
	Destroy() error
	// GetLocation 在CMDB中没有找到Host，返回error为nil，location为nil
	// 插件内部出现错误，返回error不为nil，忽略location
	GetLocation(host string) (*Location, error)
	// Range 提供一个Range接口，遍历所有的数据
	// 遍历失败，通过Range返回值error可以额捕获
	// 参数为一个回调函数
	// 返回值：bool，是否继续遍历
	// 返回值：error，回调函数处理结果，error不为nil，则停止遍历过程，并且通过Range返回error
	Range(handler func(host string, location *Location) (bool, error)) error
	// Size 获取当前CMDB存储的entry个数
	Size() int32
}

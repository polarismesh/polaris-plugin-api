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

package model

import (
	"time"

	apimodel "github.com/polarismesh/specification/source/go/api/v1/model"
	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"
)

// Location cmdb信息，对应内存结构体
type Location struct {
	Proto    *apimodel.Location
	RegionID uint32
	ZoneID   uint32
	CampusID uint32
	Valid    bool
}

type ServicePort struct {
	Port     uint32
	Protocol string
}

// Service 服务数据
type Service struct {
	ID           string
	Name         string
	Namespace    string
	Business     string
	Ports        string
	Meta         map[string]string
	Comment      string
	Department   string
	CmdbMod1     string
	CmdbMod2     string
	CmdbMod3     string
	Token        string
	Owner        string
	Revision     string
	Reference    string
	ReferFilter  string
	PlatformID   string
	Valid        bool
	CreateTime   time.Time
	ModifyTime   time.Time
	Mtime        int64
	Ctime        int64
	ServicePorts []*ServicePort
	// ExportTo 服务可见性暴露设置
	ExportTo map[string]struct{}
}

// ServiceAlias 服务别名结构体
type ServiceAlias struct {
	ID             string
	Alias          string
	AliasNamespace string
	ServiceID      string
	Service        string
	Namespace      string
	Owner          string
	Comment        string
	CreateTime     time.Time
	ModifyTime     time.Time
	ExportTo       map[string]struct{}
}

// LocationStore 地域信息，对应数据库字段
type LocationStore struct {
	IP         string
	Region     string
	Zone       string
	Campus     string
	RegionID   uint32
	ZoneID     uint32
	CampusID   uint32
	Flag       int
	ModifyTime int64
}

// Instance 组合了api的Instance对象
type Instance struct {
	Proto             *apiservice.Instance
	ServiceID         string
	ServicePlatformID string
	// Valid Whether it is deleted by logic
	Valid bool
	// ModifyTime Update time of instance
	ModifyTime time.Time
}

// InstanceArgs 用于通过服务实例查询服务的参数
type InstanceArgs struct {
	Hosts []string
	Ports []uint32
	Meta  map[string]string
}

type InstanceMetadataRequest struct {
	InstanceID string
	Revision   string
	Keys       []string
	Metadata   map[string]string
}

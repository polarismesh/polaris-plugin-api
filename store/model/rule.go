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

	apifault "github.com/polarismesh/specification/source/go/api/v1/fault_tolerance"
	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"
	apitraffic "github.com/polarismesh/specification/source/go/api/v1/traffic_manage"
)

// CircuitBreaker 熔断规则
type CircuitBreaker struct {
	ID         string
	Version    string
	Name       string
	Namespace  string
	Business   string
	Department string
	Comment    string
	Inbounds   string
	Outbounds  string
	Token      string
	Owner      string
	Revision   string
	Valid      bool
	CreateTime time.Time
	ModifyTime time.Time
}

// CircuitBreakerRelation 熔断规则绑定关系
type CircuitBreakerRelation struct {
	ServiceID   string
	RuleID      string
	RuleVersion string
	Valid       bool
	CreateTime  time.Time
	ModifyTime  time.Time
}

// CircuitBreakerRule 熔断规则
type CircuitBreakerRule struct {
	Proto        *apifault.CircuitBreakerRule
	ID           string
	Name         string
	Namespace    string
	Description  string
	Level        int
	SrcService   string
	SrcNamespace string
	DstService   string
	DstNamespace string
	DstMethod    string
	Rule         string
	Revision     string
	Enable       bool
	Valid        bool
	CreateTime   time.Time
	ModifyTime   time.Time
	EnableTime   time.Time
}

// FaultDetectRule 故障探测规则
type FaultDetectRule struct {
	Proto        *apifault.FaultDetectRule
	ID           string
	Name         string
	Namespace    string
	Description  string
	DstService   string
	DstNamespace string
	DstMethod    string
	Rule         string
	Revision     string
	Valid        bool
	CreateTime   time.Time
	ModifyTime   time.Time
}

type RoutingConfig struct {
	ID            string
	ServiceName   string
	NamespaceName string
	InBounds      string
	OutBounds     string
	Revision      string
	Valid         bool
	CreateTime    time.Time
	ModifyTime    time.Time
}

// RouterConfig Routing rules
type RouterConfig struct {
	// ID The unique id of the rules
	ID string `json:"id"`
	// namespace router config owner namespace
	Namespace string `json:"namespace"`
	// name router config name
	Name string `json:"name"`
	// policy Rules
	Policy string `json:"policy"`
	// config Specific routing rules content
	Config string `json:"config"`
	// enable Whether the routing rules are enabled
	Enable bool `json:"enable"`
	// priority Rules priority
	Priority uint32 `json:"priority"`
	// revision Edition information of routing rules
	Revision string `json:"revision"`
	// Description Simple description of rules
	Description string `json:"description"`
	// valid Whether the routing rules are valid and have not been deleted by logic
	Valid bool `json:"flag"`
	// createtime Rules creation time
	CreateTime time.Time `json:"ctime"`
	// modifytime Rules modify time
	ModifyTime time.Time `json:"mtime"`
	// enabletime The last time the rules enabled
	EnableTime time.Time `json:"etime"`
}

// RateLimit 限流规则
type RateLimit struct {
	Proto         *apitraffic.Rule
	ID            string
	ServiceID     string
	ServiceName   string
	NamespaceName string
	Name          string
	Method        string
	// Labels for old compatible, will be removed later
	Labels     string
	Priority   uint32
	Rule       string
	Revision   string
	Disable    bool
	Valid      bool
	CreateTime time.Time
	ModifyTime time.Time
	EnableTime time.Time
}

type ServiceContract struct {
	ID string
	// 所属命名空间
	Namespace string
	// 所属服务名称
	Service string
	// 契约名称
	Name string
	// 协议，http/grpc/dubbo/thrift
	Protocol string
	// 契约版本
	Version string
	// 信息摘要
	Revision string
	// 额外描述
	Content string
	// 创建时间
	CreateTime time.Time
	// 更新时间
	ModifyTime time.Time
	// 是否有效
	Valid bool
	// ClientInterfaces 客户端主动上报的接口定义
	ClientInterfaces map[string]*InterfaceDescriptor
	// ManualInterfaces 通过 OpenAPI 上报的接口定义
	ManualInterfaces map[string]*InterfaceDescriptor
}

type InterfaceDescriptor struct {
	// ID
	ID string
	// Name 接口名称
	Name string
	// ContractID
	ContractID string
	// 方法名称，对应 http method/ dubbo interface func/grpc service func
	Method string
	// 接口名称，http path/dubbo interface/grpc service
	Path string
	// 接口描述信息
	Content string
	// 接口信息摘要
	Revision string
	// 创建来源
	Source apiservice.InterfaceDescriptor_Source
	// 创建时间
	CreateTime time.Time
	// 更新时间
	ModifyTime time.Time
	// Valid
	Valid bool
}

// GrayRule 灰度资源
type GrayResource struct {
	Name       string
	MatchRule  string
	CreateTime time.Time
	ModifyTime time.Time
	CreateBy   string
	ModifyBy   string
	Valid      bool
}

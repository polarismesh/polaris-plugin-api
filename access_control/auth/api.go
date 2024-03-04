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

package auth

import (
	"context"
	"errors"
	"fmt"
	"sync"

	apisecurity "github.com/polarismesh/specification/source/go/api/v1/security"
	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"
)

// AuthChecker 权限管理通用接口定义
type AuthChecker interface {
	// Initialize 执行初始化动作
	Initialize(options *Config) error
	// VerifyCredential 验证令牌
	VerifyCredential(preCtx *AcquireContext) error
	// CheckClientPermission 执行检查客户端动作判断是否有权限，并且对 RequestContext 注入操作者数据
	CheckClientPermission(preCtx *AcquireContext) (bool, error)
	// CheckConsolePermission 执行检查控制台动作判断是否有权限，并且对 RequestContext 注入操作者数据
	CheckConsolePermission(preCtx *AcquireContext) (bool, error)
	// IsOpenConsoleAuth 返回是否开启了操作鉴权，可以用于前端查询
	IsOpenConsoleAuth() bool
	// IsOpenClientAuth
	IsOpenClientAuth() bool
}

// UserServer 用户数据管理 server
type UserServer interface {
	// Initialize 初始化
	Initialize(authOpt *Config) error
	// Name 用户数据管理server名称
	Name() string
	// CreateUsers 批量创建用户
	CreateUsers(ctx context.Context, users []*apisecurity.User) *apiservice.BatchWriteResponse
	// UpdateUser 更新用户信息
	UpdateUser(ctx context.Context, user *apisecurity.User) *apiservice.Response
	// UpdateUserPassword 更新用户密码
	UpdateUserPassword(ctx context.Context, req *apisecurity.ModifyUserPassword) *apiservice.Response
	// DeleteUsers 批量删除用户
	DeleteUsers(ctx context.Context, users []*apisecurity.User) *apiservice.BatchWriteResponse
	// GetUsers 查询用户列表
	GetUsers(ctx context.Context, query map[string]string) *apiservice.BatchQueryResponse
	// GetUserToken 获取用户的 token
	GetUserToken(ctx context.Context, user *apisecurity.User) *apiservice.Response
	// UpdateUserToken 禁止用户的token使用
	UpdateUserToken(ctx context.Context, user *apisecurity.User) *apiservice.Response
	// ResetUserToken 重置用户的token
	ResetUserToken(ctx context.Context, user *apisecurity.User) *apiservice.Response
	// Login 登录动作
	Login(req *apisecurity.LoginRequest) *apiservice.Response
	GroupOperator
}

// GroupOperator 用户组相关操作
type GroupOperator interface {
	// CreateGroup 创建用户组
	CreateGroup(ctx context.Context, group *apisecurity.UserGroup) *apiservice.Response
	// UpdateGroups 更新用户组
	UpdateGroups(ctx context.Context, groups []*apisecurity.ModifyUserGroup) *apiservice.BatchWriteResponse
	// DeleteGroups 批量删除用户组
	DeleteGroups(ctx context.Context, group []*apisecurity.UserGroup) *apiservice.BatchWriteResponse
	// GetGroups 查询用户组列表（不带用户详细信息）
	GetGroups(ctx context.Context, query map[string]string) *apiservice.BatchQueryResponse
	// GetGroup 根据用户组信息，查询该用户组下的用户相信
	GetGroup(ctx context.Context, req *apisecurity.UserGroup) *apiservice.Response
	// GetGroupToken 获取用户组的 token
	GetGroupToken(ctx context.Context, group *apisecurity.UserGroup) *apiservice.Response
	// UpdateGroupToken 取消用户组的 token 使用
	UpdateGroupToken(ctx context.Context, group *apisecurity.UserGroup) *apiservice.Response
	// ResetGroupToken 重置用户组的 token
	ResetGroupToken(ctx context.Context, group *apisecurity.UserGroup) *apiservice.Response
}

// StrategyServer 策略相关操作
type StrategyServer interface {
	// Initialize 初始化
	Initialize(authOpt *Config) error
	// Name 策略管理server名称
	Name() string
	// CreateStrategy 创建策略
	CreateStrategy(ctx context.Context, strategy *apisecurity.AuthStrategy) *apiservice.Response
	// UpdateStrategies 批量更新策略
	UpdateStrategies(ctx context.Context, reqs []*apisecurity.ModifyAuthStrategy) *apiservice.BatchWriteResponse
	// DeleteStrategies 删除策略
	DeleteStrategies(ctx context.Context, reqs []*apisecurity.AuthStrategy) *apiservice.BatchWriteResponse
	// GetStrategies 获取资源列表
	// support 1. 支持按照 principal-id + principal-role 进行查询
	// support 2. 支持普通的鉴权策略查询
	GetStrategies(ctx context.Context, query map[string]string) *apiservice.BatchQueryResponse
	// GetStrategy 获取策略详细
	GetStrategy(ctx context.Context, strategy *apisecurity.AuthStrategy) *apiservice.Response
	// GetPrincipalResources 获取某个 principal 的所有可操作资源列表
	GetPrincipalResources(ctx context.Context, query map[string]string) *apiservice.Response
	// GetAuthChecker 获取鉴权检查器
	GetAuthChecker() AuthChecker
	// AfterResourceOperation 操作完资源的后置处理逻辑
	AfterResourceOperation(afterCtx *AcquireContext) error
}

const (
	// DefaultUserMgnPluginName default user server name
	DefaultUserMgnPluginName = "defaultUser"
	// DefaultStrategyMgnPluginName default strategy server name
	DefaultStrategyMgnPluginName = "defaultStrategy"
)

// Config 鉴权能力的相关配置参数
type Config struct {
	// Name 原AuthServer名称，已废弃
	Name string
	// Option 原AuthServer的option，已废弃
	// Deprecated
	Option map[string]interface{}
	// User UserOperator的相关配置
	User *UserConfig `yaml:"user"`
	// Strategy StrategyOperator的相关配置
	Strategy *StrategyConfig `yaml:"strategy"`
}

func (c *Config) SetDefault() {
	if c.User == nil {
		c.User = &UserConfig{
			Name:   DefaultUserMgnPluginName,
			Option: map[string]interface{}{},
		}
	}
	if c.Strategy == nil {
		c.Strategy = &StrategyConfig{
			Name:   DefaultStrategyMgnPluginName,
			Option: map[string]interface{}{},
		}
	}
}

// UserConfig UserOperator的相关配置
type UserConfig struct {
	// Name UserOperator的名称
	Name string `yaml:"name"`
	// Option UserOperator的option
	Option map[string]interface{} `yaml:"option"`
}

// StrategyConfig StrategyOperator的相关配置
type StrategyConfig struct {
	// Name StrategyOperator的名称
	Name string `yaml:"name"`
	// Option StrategyOperator的option
	Option map[string]interface{} `yaml:"option"`
}

var (
	// userMgnSlots 保存用户管理manager slot
	userMgnSlots = map[string]UserServer{}
	// strategyMgnSlots 保存策略管理manager slot
	strategyMgnSlots = map[string]StrategyServer{}
	once             sync.Once
	userMgn          UserServer
	strategyMgn      StrategyServer
	finishInit       bool
)

// RegisterUserServer 注册一个新的 UserServer
func RegisterUserServer(s UserServer) error {
	name := s.Name()
	if _, ok := userMgnSlots[name]; ok {
		return fmt.Errorf("UserServer=[%s] exist", name)
	}

	userMgnSlots[name] = s
	return nil
}

// GetUserServer 获取一个 UserServer
func GetUserServer() (UserServer, error) {
	if !finishInit {
		return nil, errors.New("UserServer has not done Initialize")
	}
	return userMgn, nil
}

// RegisterStrategyServer 注册一个新的 StrategyServer
func RegisterStrategyServer(s StrategyServer) error {
	name := s.Name()
	if _, ok := strategyMgnSlots[name]; ok {
		return fmt.Errorf("StrategyServer=[%s] exist", name)
	}

	strategyMgnSlots[name] = s
	return nil
}

// GetStrategyServer 获取一个 StrategyServer
func GetStrategyServer() (StrategyServer, error) {
	if !finishInit {
		return nil, errors.New("StrategyServer has not done Initialize")
	}
	return strategyMgn, nil
}

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
)

type ReleaseType string

const (
	// ReleaseTypeFull 全量类型
	ReleaseTypeFull = ""
	// ReleaseTypeGray 灰度类型
	ReleaseTypeGray = "gray"
)

/** ----------- DataObject ------------- */

// ConfigFileGroup 配置文件组数据持久化对象
type ConfigFileGroup struct {
	Id         uint64
	Name       string
	Namespace  string
	Comment    string
	Owner      string
	Business   string
	Department string
	Metadata   map[string]string
	CreateTime time.Time
	ModifyTime time.Time
	CreateBy   string
	ModifyBy   string
	Valid      bool
	Revision   string
}

type ConfigFileKey struct {
	Name      string
	Namespace string
	Group     string
}

func (c ConfigFileKey) String() string {
	return c.Namespace + "@" + c.Group + "@" + c.Name
}

// ConfigFile 配置文件数据持久化对象
type ConfigFile struct {
	Id        uint64
	Name      string
	Namespace string
	Group     string
	// OriginContent 最原始的配置文件内容数据
	OriginContent string
	Content       string
	Comment       string
	Format        string
	Flag          int
	Valid         bool
	Metadata      map[string]string
	Encrypt       bool
	EncryptAlgo   string
	Status        string
	CreateBy      string
	ModifyBy      string
	ReleaseBy     string
	CreateTime    time.Time
	ModifyTime    time.Time
	ReleaseTime   time.Time
}

func NewConfigFileRelease() *ConfigFileRelease {
	return &ConfigFileRelease{
		SimpleConfigFileRelease: &SimpleConfigFileRelease{
			ConfigFileReleaseKey: &ConfigFileReleaseKey{},
		},
	}
}

// ConfigFileRelease 配置文件发布数据持久化对象
type ConfigFileRelease struct {
	*SimpleConfigFileRelease
	Content string
}

type ConfigFileReleaseKey struct {
	Id          uint64
	Name        string
	Namespace   string
	Group       string
	FileName    string
	ReleaseType ReleaseType
}

// SimpleConfigFileRelease 配置文件发布数据持久化对象
type SimpleConfigFileRelease struct {
	*ConfigFileReleaseKey
	Version            uint64
	Comment            string
	Md5                string
	Flag               int
	Active             bool
	Valid              bool
	Format             string
	Metadata           map[string]string
	CreateTime         time.Time
	CreateBy           string
	ModifyTime         time.Time
	ModifyBy           string
	ReleaseDescription string
	BetaLabels         []*apimodel.ClientLabel
}

// ConfigFileReleaseHistory 配置文件发布历史记录数据持久化对象
type ConfigFileReleaseHistory struct {
	Id                 uint64
	Name               string
	Namespace          string
	Group              string
	FileName           string
	Format             string
	Metadata           map[string]string
	Content            string
	Comment            string
	Version            uint64
	Md5                string
	Type               string
	Status             string
	CreateTime         time.Time
	CreateBy           string
	ModifyTime         time.Time
	ModifyBy           string
	Valid              bool
	Reason             string
	ReleaseDescription string
}

// ConfigFileTag 配置文件标签数据持久化对象
type ConfigFileTag struct {
	Id         uint64
	Key        string
	Value      string
	Namespace  string
	Group      string
	FileName   string
	CreateTime time.Time
	CreateBy   string
	ModifyTime time.Time
	ModifyBy   string
	Valid      bool
}

// ConfigFileTemplate config file template data object
type ConfigFileTemplate struct {
	Id         uint64
	Name       string
	Content    string
	Comment    string
	Format     string
	CreateTime time.Time
	CreateBy   string
	ModifyTime time.Time
	ModifyBy   string
}

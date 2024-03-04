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

package crypto

import "fmt"

var (
	slots = make(map[string]Crypto)
)

// Register 注册插件
func RegisterCrypto(name string, plugin Crypto) {
	if _, exist := slots[name]; exist {
		panic(fmt.Sprintf("existed plugin: name=%v", name))
	}
	slots[name] = plugin
}

func GetCrypto(name string) (Crypto, bool) {
	server, exist := slots[name]
	return server, exist
}

// ConfigEntry 单个插件配置
type ConfigEntry struct {
	Name   string                 `yaml:"name"`
	Option map[string]interface{} `yaml:"option"`
}

// Crypto Crypto interface
type Crypto interface {
	// Name .
	Name() string
	// Initialize .
	Initialize(c *ConfigEntry) error
	// Destroy .
	Destroy() error
	// GenerateKey .
	GenerateKey() ([]byte, error)
	// Encrypt .
	Encrypt(plaintext string, key []byte) (cryptotext string, err error)
	// Decrypt .
	Decrypt(cryptotext string, key []byte) (string, error)
}

var (
	_pwdslots = make(map[string]ParsePassword)
)

// Register 注册插件
func RegisterParsePassword(name string, plugin ParsePassword) {
	if _, exist := _pwdslots[name]; exist {
		panic(fmt.Sprintf("existed plugin: name=%v", name))
	}
	_pwdslots[name] = plugin
}

func GetParsePassword(name string) (ParsePassword, bool) {
	server, exist := _pwdslots[name]
	return server, exist
}

// ParsePassword Password plug -in
type ParsePassword interface {
	// Name .
	Name() string
	// Initialize .
	Initialize(c *ConfigEntry) error
	// Destroy .
	Destroy() error
	// ParsePassword .
	ParsePassword(cipher string) (string, error)
}

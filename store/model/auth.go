package model

import "time"

// User 用户
type User struct {
	ID          string
	Name        string
	Password    string
	Owner       string
	Source      string
	Mobile      string
	Email       string
	Type        string
	Token       string
	TokenEnable bool
	Valid       bool
	Comment     string
	CreateTime  time.Time
	ModifyTime  time.Time
}

// UserGroup 用户组
type UserGroup struct {
	ID          string
	Name        string
	Owner       string
	Token       string
	TokenEnable bool
	Valid       bool
	Comment     string
	UserIds     []string
	CreateTime  time.Time
	ModifyTime  time.Time
}

// StrategyDetail 鉴权策略详细
type StrategyDetail struct {
	ID         string
	Name       string
	Action     string
	Comment    string
	Principals []Principal
	Default    bool
	Owner      string
	Resources  []StrategyResource
	Valid      bool
	Revision   string
	CreateTime time.Time
	ModifyTime time.Time
}

// Strategy 策略main信息
type Strategy struct {
	ID         string
	Name       string
	Principal  string
	Action     string
	Comment    string
	Owner      string
	Default    bool
	Valid      bool
	CreateTime time.Time
	ModifyTime time.Time
}

// StrategyResource 策略资源
type StrategyResource struct {
	StrategyID string
	ResType    int32
	ResID      string
}

// Principal 策略相关人
type Principal struct {
	StrategyID    string
	PrincipalID   string
	PrincipalRole string
}

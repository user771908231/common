package sichuan

import (
	"casino_common/api/majiang/api"
	"github.com/name5566/leaf/module"
)

//四川麻将相关的都可以放在这里
type SCMJDesk struct {
	*api.MJDeskCore
}

//四川麻将相关的核心类都在这里
func NewSCMJDesk(s *module.Skeleton) *SCMJDesk {
	return &SCMJDesk{
		MJDeskCore: api.NewMJDeskCore(s),
	}
}

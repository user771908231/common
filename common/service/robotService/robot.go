package robotService

import (
	"casino_common/proto/ddproto"
)

//机器人
type Robots struct {
	*ddproto.User
	available bool
}

func NewRobots(u *ddproto.User) *Robots {
	return &Robots{
		User:u,
	}
}

//机器人是否可用
func (r *Robots) IsAvailable() bool {
	return r.available
}

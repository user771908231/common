package robotService

import (
	"casino_common/proto/ddproto"
)

//机器人
type Robot struct {
	*ddproto.User
	available bool
}

func NewRobots(u *ddproto.User) *Robot {
	return &Robot{
		User:u,
	}
}

//机器人是否可用
func (r *Robot) IsAvailable() bool {
	return r.available
}

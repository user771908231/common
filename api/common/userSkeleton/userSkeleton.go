package userSkeleton

import (
	"github.com/golang/protobuf/proto"
	"time"
)

type UserSkeleton struct {
	ApplyDissolve int32 //申请离线的状态
	DissolveTimer *time.Timer
}

func (*UserSkeleton) WriteMsg(p proto.Message) {
	panic("implement me")
}

func (*UserSkeleton) GetUserId() uint32 {
	panic("implement me")
}

func (*UserSkeleton) GetBaseUserSkeleton() interface{} {
	panic("implement me")
}

func (*UserSkeleton) IsBreak() bool {
	panic("implement me")
}

func (*UserSkeleton) IsLeave() bool {
	panic("implement me")
}

func (u *UserSkeleton) GetSkeleton() *UserSkeleton {
	return u
}

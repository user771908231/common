package userSkeleton

import "github.com/golang/protobuf/proto"

//user的骨架
type UserSkeletonApi interface {
	WriteMsg(p proto.Message)
	GetUserId() uint32
	GetBaseUserSkeleton() interface{}
	IsBreak() bool
	IsLeave() bool
}

package deskSkeleton

import (
	"casino_common/api/common/roomSkeleton"
	"casino_common/api/common/userSkeleton"
	"github.com/golang/protobuf/proto"
)

//这个是骨架的desk，所有的游戏都可以组合这个骨架
type DeskSkeleton struct {
	Id         int32
	GameNumber int32
	ApplyDis   bool
	ListUsers  func() []userSkeleton.UserSkeletonApi
	GetRoom    func() roomSkeleton.RoomSkeletonApi
}

func (d *DeskSkeleton) BroadCastProto(msg proto.Message) {
	for _, u := range d.ListUsers() {
		if u != nil {
			u.WriteMsg(msg)
		}
	}
}

func (d *DeskSkeleton) DlogDes() string {
	return "desk 描述"
}

func (d *DeskSkeleton) GetBaseUserSkeleton(userId uint32) *userSkeleton.UserSkeleton {
	for _, u := range d.ListUsers() {
		if u != nil && u.GetUserId() == userId {
			return u.GetBaseUserSkeleton().(*userSkeleton.UserSkeleton)
		}
	}
	return nil
}

func (d *DeskSkeleton) ListBaseUserSkeletons() []*userSkeleton.UserSkeleton {
	var ret []*userSkeleton.UserSkeleton
	for _, u := range d.ListUsers() {
		if u != nil {
			ret = append(ret, u.GetBaseUserSkeleton().(*userSkeleton.UserSkeleton))
		}
	}
	return ret
}

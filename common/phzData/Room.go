package phzData

import (
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/util"
)

type Room struct {
	S        *module.Skeleton //leaf骨架
	RoomId   int32            //唯一ID
	RoomType int32            //room类型
	Desks    *util.Map        //room下的Desk
}
package api

import (
	"github.com/name5566/leaf/util"
	"github.com/name5566/leaf/module"
)

type MJRoom interface {
	CreateDesk(interface{}) (MJDesk, error) //创建房间
	GetDesk(interface{}) (MJDesk, error)    //得到一个房间
	CreateFee(...interface{}) int64         //房费
}

//room 基本的操作
type MJRoomCore struct {
	S     *module.Skeleton
	desks *util.Map //所有的desk
}

func NewMJRoomCore(s *module.Skeleton) *MJRoomCore {
	return &MJRoomCore{
		desks: new(util.Map), //所有的desk
		S:     s,             //组合leaf的骨架
	}
}

//增加一个desk
func (rs *MJRoomCore) AddDesk(desk MJDesk) {
	rs.desks.Set(desk.GetDeskId(), desk)
}

//得到一个desk
func (rs *MJRoomCore) GetDesk(deskId int32) MJDesk {
	ret := rs.desks.Get(deskId)
	if ret != nil {
		return ret.(MJDesk)
	}
	return nil
}

//通过key得到desk
func (rs *MJRoomCore) GetDeskByPass(key string) MJDesk {
	var ret MJDesk
	rs.desks.UnsafeRange(func(k interface{}, v interface{}) {
		if v != nil {
			d := v.(MJDesk)
			if d.GetPassword() == key {
				ret = d
			}
		}
	})
	return ret
}

//删除一个desk
func (rs *MJRoomCore) RmDesk(deskId int32) error {
	rs.desks.Del(deskId)
	return nil
}

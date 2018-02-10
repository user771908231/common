package api

import (
	"casino_common/common/consts"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"casino_common/utils/chessUtils"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"fmt"
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/util"
)

type MJRoom interface {
	CreateDesk(interface{}) (MJDesk, error)   //创建房间
	CreateDeskV2(interface{}) (MJDesk, error) //创建房间v2
	GetEnterDesk(interface{}) (MJDesk, error) //得到一个房间
	GetDeskById(int32) MJDesk                 //得到一个desk
	GetRoomId() int32                         //得到id
	ListDesk() *util.Map                      //获取所有桌子
	GetRoomType() int32                       //房间类型
	DissolveDesk(...interface{}) error
	CreateFee(...interface{}) int64 //房费
}

const (
	ROOM_FIREND int32 = int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) //朋友桌
	ROOM_COIN   int32 = int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN)   //金币场
)

//room 基本的操作
type MJRoomCore struct {
	S        *module.Skeleton
	RoomId   int32
	RoomType int32
	desks    *util.Map //所有的desk
}

func NewMJRoomCore(s *module.Skeleton) *MJRoomCore {
	return &MJRoomCore{
		desks: new(util.Map), //所有的desk
		S:     s,             //组合leaf的骨架
	}
}

func (r *MJRoomCore) CreateDeskV2(interface{}) (MJDesk, error) {
	log.W("MJRoomCore CreateDeskV2 nothing to do.")
	return nil, nil
}

func (r *MJRoomCore) CreateFee(args ...interface{}) int64 {
	return 0
}

func (r *MJRoomCore) GetRoomId() int32 {
	return r.RoomId
}

func (r *MJRoomCore) GetRoomType() int32 {
	return r.RoomType
}

//随机一个房间号码
func (r *MJRoomCore) RandRoomKey(gid int32) string {
	//金币场没有房间号码
	roomKey := chessUtils.GetRoomPass(gid)
	//1,判断roomKey是否已经存在
	if r.IsRoomKeyExist(roomKey) {
		//log.E("房间密钥[%v]已经存在,创建房间失败,重新创建", roomKey)
		return r.RandRoomKey(gid)
	} else {
		//log.T("最终得到的密钥是[%v]", roomKey)
		return roomKey
	}
	return ""
}

//随机一个房间号码
func (r *MJRoomCore) RandRoomKeyV2(gid int32) string {
	//金币场没有房间号码
	roomKey := chessUtils.GetRoomPassV3(gid)
	//1,判断roomKey是否已经存在
	if r.IsRoomKeyExistV2(roomKey) {
		//log.E("房间密钥[%v]已经存在,创建房间失败,重新创建", roomKey)
		return r.RandRoomKeyV2(gid)
	}
	redisUtils.Set(fmt.Sprintf("%s_%s", consts.RKEY_ROOMPWD_GAMEID, roomKey), fmt.Sprintf("%v", gid))
	log.T("最终得到的密钥是[%v]", roomKey)
	return roomKey
}

//判断roomkey是否已经存在了
func (r *MJRoomCore) IsRoomKeyExist(roomkey string) bool {
	ret := false
	r.ListDesk().RLockRange(func(k interface{}, v interface{}) {
		if v != nil {
			d := v.(MJDesk)
			if d.GetPassword() == roomkey {
				ret = true
			}
		}
	})
	return ret
}

//判断roomkey是否已经存在了
func (r *MJRoomCore) IsRoomKeyExistV2(roomkey string) bool {
	//先从内存找
	memIsExist := false
	r.ListDesk().RLockRange(func(k interface{}, v interface{}) {
		if v != nil {
			d := v.(MJDesk)
			if d.GetPassword() == roomkey {
				memIsExist = true
			}
		}
	})

	//再从redis中找
	redisIsExist := true
	gidStr := redisUtils.Get(fmt.Sprintf("%s_%s", consts.RKEY_ROOMPWD_GAMEID, roomkey))
	if gidStr == "" {
		redisIsExist = false
	} else {
		gid := numUtils.String2Int(gidStr)
		if gid <= 0 {
			redisIsExist = false
		}
	}
	return memIsExist || redisIsExist
}

//增加一个desk
func (rs *MJRoomCore) AddDesk(desk MJDesk) {
	rs.desks.Set(desk.GetDeskId(), desk)
}

//得到一个desk
func (rs *MJRoomCore) GetDeskById(deskId int32) MJDesk {
	ret := rs.desks.Get(deskId)
	if ret != nil {
		return ret.(MJDesk)
	}
	return nil
}

func (rs *MJRoomCore) ListDesk() *util.Map {
	return rs.desks
}

//通过key得到desk
func (rs *MJRoomCore) GetDeskByPass(key string) MJDesk {
	var ret MJDesk
	rs.desks.RLockRange(func(k interface{}, v interface{}) {
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
	desk := rs.desks.Get(deskId)
	if desk == nil {
		return nil
	}

	//需要置空desk指针
	desk = nil
	rs.desks.Del(deskId)
	return nil
}

//删除一个desk
func (rs *MJRoomCore) RmDeskV2(d MJDesk) error {
	//删除desk 清理房间号
	redisUtils.Del(fmt.Sprintf("%s_%s", consts.RKEY_ROOMPWD_GAMEID, d.GetPassword()))

	desk := rs.desks.Get(d.GetDeskId())
	if desk == nil {
		return nil
	}

	//需要置空desk指针
	desk = nil
	rs.desks.Del(d.GetDeskId())
	return nil
}

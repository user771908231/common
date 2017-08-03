package sessionService

import (
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
	"github.com/name5566/leaf/util"
	"casino_common/common/log"
	"github.com/golang/protobuf/proto"
)

// 各游戏服务器内存session相关的...
/**
	目前的session每个人将会保存两个：朋友桌和金币场

 */

var FSession *util.Map //朋友桌玩家id 对应deskId
var CSession *util.Map //金币场玩家id 对应deskId

func init() {
	FSession = new(util.Map)
	CSession = new(util.Map)
}

func GetSessionv2(userId uint32, gid, roomType int32) *ddproto.GameSession {
	log.T("开始从内存中获取玩家session userId[%v] gid[%v] roomType[%v]", userId, gid, roomType)
	gets := FSession.Get(userId)
	str := "朋友桌"
	if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
		gets = CSession.Get(userId)
		str = "金币场"
	}

	if gets == nil {
		log.W("无法从内存中获取玩家[%v]在游戏为[%v]的%v的session,找到的session为nil", userId, gid, str)
		return nil
	}
	session := gets.(*ddproto.GameSession)
	if session == nil {
		log.W("无法从内存中获取玩家[%v]在游戏为[%v]的%vsession,找到的session为nil", userId, gid, str)
		return nil
	}

	if session.GetDeskId() == 0 {
		log.W("从内存中获取玩家[%v]在游戏为[%v]的%vsession错误,桌子id错误,找到的session[%v]", userId, gid, str, session)
		return nil
	}

	if session.GetGameId() != gid {
		log.W("从内存中获取玩家[%v]在游戏为[%v]的%vsession错误,游戏id不匹配,找到的session[%v]", userId, gid, str, session)
		return nil
	}
	return session
}

//得到朋友桌的session
func GetFriendSessionv2(userId uint32, gid int32) *ddproto.GameSession {
	return GetSessionv2(userId, gid, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND))
}

//得到金币场的session
func GetCoinSessionv2(userId uint32, gid int32) *ddproto.GameSession {
	return GetSessionv2(userId, gid, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN))
}

//更新用户的session信息，具体更新什么信息待定
func UpdateSessionv2(userId uint32, gameStatus int32, gameId int32, gameNumber int32, roomId int32, deskId int32, gameCustomStatus int32, isBreak bool, isLeave bool, roomType int32, roomPass string) (*ddproto.GameSession, error) {
	session := GetSession(userId, roomType)
	if session == nil {
		log.W("没有找到user[%v]内存中的session,需要重新申请一个并保存...", userId)
		session = commonNewPorot.NewGameSession()
	}
	*session.UserId = userId
	*session.DeskId = deskId
	*session.RoomId = roomId
	*session.GameStatus = gameStatus
	*session.GameId = gameId
	*session.GameNumber = gameNumber
	*session.GameCustomStatus = gameCustomStatus
	*session.IsBreak = isBreak
	*session.IsLeave = isLeave
	session.RoomType = proto.Int32(roomType)
	session.RoomPassword = proto.String(roomPass)

	//保存session
	log.T("保存到内存中的session【%v】", session)
	if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
		CSession.Set(userId, session)
		return session, nil
	}

	FSession.Set(userId, session)
	return session, nil
}

//删除玩家session 增加对空的判断
func delSessionv2(s *ddproto.GameSession) {
	if s == nil {
		log.W("无法删除内存中的session, session为nil")
		return
	}
	log.T("开始删除玩家[%v]的roomType[%v]内存中的session[%v] ", s.GetUserId(), s.GetRoomType(), s)
	if s.GetRoomType() == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
		//删除金币场
		CSession.Del(s.GetUserId())
		return
	}
	//默认删除朋友桌
	FSession.Del(s.GetUserId())
	return
}

//通过suerId roomType 删除玩家的session
func DelSessionByKeyv2(userId uint32, gid, roomType, deskId int32) {
	s := GetSessionv2(userId, gid, roomType)
	if s == nil {
		log.W("删除玩家[%v]内存中的session时出错, 当前session为nil, gid[%v] roomType[%v] deskId[%v]", userId, gid, roomType, deskId)
		return
	}

	if s.GetDeskId() != deskId {
		log.E("删除玩家[%v]roomType[%v],gid[%v],dekid[%v] 内存中session的时候出错, 桌子id不匹配 session[%v],",
			userId, roomType, gid, deskId, s)
		return
	}

	delSessionv2(s)
	return
}

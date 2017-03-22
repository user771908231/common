package sessionService

import (
	"strings"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"casino_common/common/log"
	"casino_common/proto/funcsInit"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"casino_common/common/cfg"
)

// session相关的...
/**
	目前的session每个人将会保存两个：朋友桌和金币场

 */

func getSessionKey(userId uint32, gameId int32) string {
	idstr, _ := numUtils.Uint2String(userId)
	gameIdSrt, _ := numUtils.Int2String(gameId)
	ret := strings.Join([]string{cfg.RKEY_MJSESSION_KEY_PRE, idstr, gameIdSrt}, "_")
	return ret
}

//
func GetSession(userId uint32, roomType int32) *ddproto.GameSession {
	s := redisUtils.GetObj(getSessionKey(userId, roomType), &ddproto.GameSession{})
	if s != nil {
		return s.(*ddproto.GameSession)
	} else {
		return nil
	}
}

//自动查找session
func GetSessionAuto(userId uint32) *ddproto.GameSession {
	session := GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND))
	if session == nil || session.GetDeskId() == 0 || session.GetGameStatus() == int32(ddproto.COMMON_ENUM_GAMESTATUS_NOGAME) {
		//2,第二步获取朋友桌的session
		session = GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN))
	}
	return session
}

//得到朋友桌的session
func GetFriendSession(userId uint32, gid int32) *ddproto.GameSession {
	session := GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND))
	if session == nil || session.GetDeskId() == 0 || session.GetGameId() != gid {
		return nil
	} else {
		return session
	}

}

//得到金币场的session
func GetCoinSession(userId uint32, gid int32) *ddproto.GameSession {
	session := GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN))
	if session == nil || session.GetDeskId() == 0 || session.GetGameId() != gid {
		return nil
	} else {
		return session
	}
}

//更新用户的session信息，具体更新什么信息待定
func UpdateSession(userId uint32, gameStatus int32, gameId int32, gameNumber int32, roomId int32, deskId int32, gameCustomStatus int32, isBreak bool, isLeave bool, roomType int32, roomPass string) (*ddproto.GameSession, error) {
	session := GetSession(userId, roomType)
	if session == nil {
		log.W("没有找到user[%v]的session,需要重新申请一个并保存...", userId)
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
	log.T("保存到redis的session【%v】", session)
	redisUtils.SetObj(getSessionKey(userId, roomType), session)
	return session, nil
}

//删除玩家session 增加对空的判断
func DelSession(s *ddproto.GameSession) {
	if s == nil {
		return
	} else {
		log.T("开始删除玩家[%v]的roomType[%v] session[%v] ", s.GetUserId(), s.GetRoomType(), s)
		redisUtils.Del(getSessionKey(s.GetUserId(), s.GetRoomType()))
	}
}

//通过suerId roomType 删除玩家的session
func DelSessionByKey(userId uint32, roomType int32) {
	s := GetSession(userId, roomType)
	DelSession(s)
}

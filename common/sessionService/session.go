package sessionService

import (
	"casino_common/common/Error"
	"casino_common/common/consts"
	"casino_common/common/log"
	"casino_common/common/service/statisticsService"
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
	"casino_common/utils"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strings"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
)

// session相关的...
/**
	目前的session每个人将会保存两个：朋友桌和金币场

 */

func getSessionKey(userId uint32, gameId int32) string {
	idstr, _ := numUtils.Uint2String(userId)
	gameIdSrt, _ := numUtils.Int2String(gameId)
	ret := strings.Join([]string{consts.RKEY_MJSESSION_KEY_PRE, idstr, gameIdSrt}, "_")
	return ret
}

//
func GetSession(userId uint32, roomType int32) *ddproto.GameSession {
	s := redisUtils.GetObj(getSessionKey(userId, roomType), &ddproto.GameSession{})
	if s != nil {
		return s.(*ddproto.GameSession)
	}
	return nil
}

//自动查找session
func GetSessionAuto(userId uint32) *ddproto.GameSession {
	session := GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND))
	log.T("开始自动查找玩家[%v]朋友sessin[%v]", userId, session)
	if session == nil || session.GetDeskId() == 0 || session.GetGameStatus() == int32(ddproto.COMMON_ENUM_GAMESTATUS_NOGAME) {
		//2,第二步获取金币场的session
		session = GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN))
		log.T("开始自动查找玩家[%v]金币sessin[%v]", userId, session)

	}
	return session
}

//得到朋友桌的session
func GetFriendSession(userId uint32, gid int32) *ddproto.GameSession {
	session := GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND))
	if session == nil {
		log.W("无法获取玩家[%v]在游戏为[%v]的朋友桌session,找到的session为nil", userId, gid)
		return nil
	}

	if session.GetDeskId() == 0 {
		log.W("获取玩家[%v]在游戏为[%v]的朋友桌session错误,桌子id错误,找到的session[%v]", userId, gid, session)
		return nil
	}

	if session.GetGameId() != gid {
		log.W("获取玩家[%v]在游戏为[%v]的朋友桌session错误,游戏id不匹配,找到的session[%v]", userId, gid, session)
		return nil
	}
	return session
}

//得到金币场的session
func GetCoinSession(userId uint32, gid int32) *ddproto.GameSession {
	session := GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN))
	if session == nil {
		log.W("无法获取玩家[%v]在游戏为[%v]的金币场session,找到的session为nil", userId, gid)
		return nil
	}

	if session.GetDeskId() == 0 {
		log.W("获取玩家[%v]在游戏为[%v]的金币场session错误,桌子id错误,找到的session[%v]", userId, gid, session)
		return nil
	}

	if session.GetGameId() != gid {
		log.W("获取玩家[%v]在游戏为[%v]的金币场session错误,游戏id不匹配,找到的session[%v]", userId, gid, session)
		return nil
	}
	return session
}

//更新用户的session信息，具体更新什么信息待定
func UpdateSession(userId uint32, gameStatus int32, gameId int32, gameNumber int32, roomId int32, deskId int32, gameCustomStatus int32, isBreak bool, isLeave bool, roomType int32, roomPass string) (*ddproto.GameSession, error) {
	//同步更新内存中的session
	UpdateSessionv2(userId, gameStatus, gameId, gameNumber, roomId, deskId, gameCustomStatus, isBreak, isLeave, roomType, roomPass)

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

	//这里需要在redis-set中保存在线用户的数据
	statisticsService.IncrOnline(userId, gameId)
	//异步更新session表
	go db.C(tableName.DBT_GAME_SESSION).Upsert(bson.M{"userid": session.GetUserId(), "gameid": session.GetGameId()}, session)
	return session, nil
}

//删除玩家session 增加对空的判断
func delSession(s *ddproto.GameSession) {
	if s == nil {
		return
	} else {
		log.T("开始删除玩家[%v]的roomType[%v] session[%v] ", s.GetUserId(), s.GetRoomType(), s)
		redisUtils.Del(getSessionKey(s.GetUserId(), s.GetRoomType()))
		statisticsService.DecrOnline(s.GetUserId(), s.GetGameId())
		//异步删除玩家session表
		go db.C(tableName.DBT_GAME_SESSION).Remove(bson.M{"userid": s.GetUserId(), "gameid": s.GetGameId()})
	}
}

//检测是否能够进入游戏房间
func CanEnter(userId uint32, gid int32, roomType int32, roomLevel int32) (bool, error) {
	log.T("开始根据session判断玩家[%v]能否进入游戏id为[%v],roomType[%v],roomLevel[%v]", userId, gid, roomType, roomLevel)
	var getError = func(gid, roomType int32) error {
		msg := fmt.Sprintf("已经在游戏%v中了，结束之后再来吧!", utils.GetGameName(gid, roomType))
		return Error.NewError(consts.ACK_RESULT_ERROR, msg)
	}

	s := GetSession(userId, roomType)
	if s == nil {
		//没有session 可以进
		log.T("找不到玩家[%v]的session, 可以进入", userId)
		return true, nil
	}
	log.T("判断玩家能否进入游戏时，找到的玩家session[%v]", s)
	if s.GetGameStatus() != int32(ddproto.COMMON_ENUM_GAMESTATUS_GAMING) {
		//没在游戏中 可以进
		log.T("找到的玩家[%v]session没在游戏中, 可以进入", userId)
		return true, nil
	}

	if s.GetGameId() != gid {
		//根据roomType找到的session是在其他游戏中 不能进
		return false, getError(s.GetGameId(), s.GetRoomType())
	}
	//其他条件可以进入房间
	log.T("根据session判断玩家[%v]可以进入游戏", userId)
	return true, nil
}

//通过suerId roomType 删除玩家的session
func DelSessionByKey(userId uint32, roomType int32, gid int32, deskId int32) {
	//同步删除内存中session
	DelSessionByKeyv2(userId, gid, roomType, deskId)

	s := GetSession(userId, roomType)

	if s == nil {
		log.E("删除玩家[%v]roomType[%v],gid[%v],dekid[%v] 的时候出错 session[%v],",
			userId, roomType, gid, deskId, s)
		return
	}

	if s.GetGameId() != gid {
		log.E("删除玩家[%v]roomType[%v],gid[%v],dekid[%v] 的时候出错 session[%v],",
			userId, roomType, gid, deskId, s)
		return
	}

	if deskId > 0 && s.GetDeskId() != deskId {
		log.E("删除玩家[%v]roomType[%v],gid[%v],dekid[%v] 的时候出错 session[%v],",
			userId, roomType, gid, deskId, s)
		return
	}
	delSession(s)
}

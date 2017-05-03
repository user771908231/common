package sessionService

import (
	"strings"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"casino_common/common/log"
	"casino_common/proto/funcsInit"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"casino_common/common/Error"
	"casino_common/common/consts"
	"casino_common/utils"
	"fmt"
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
	} else {
		return nil
	}
}

//自动查找session
func GetSessionAuto(userId uint32) *ddproto.GameSession {
	session := GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND))
	log.T("开始自动查找玩家[%v]朋友sessin[%v]", userId, session)
	if session == nil || session.GetDeskId() == 0 || session.GetGameStatus() == int32(ddproto.COMMON_ENUM_GAMESTATUS_NOGAME) {
		//2,第二步获取朋友桌的session
		session = GetSession(userId, int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN))
		log.T("开始自动查找玩家[%v]金币sessin[%v]", userId, session)

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
func delSession(s *ddproto.GameSession) {
	if s == nil {
		return
	} else {
		log.T("开始删除玩家[%v]的roomType[%v] session[%v] ", s.GetUserId(), s.GetRoomType(), s)
		redisUtils.Del(getSessionKey(s.GetUserId(), s.GetRoomType()))
	}
}

//检测是否能够进入游戏房间
/**
	1,如果是进入朋友桌
		检测是否正在朋友桌玩，如果gid一样，roomType一样，则返回对应的session
	2,如果是进入金币本
		检测是否正在朋友桌玩，如果在朋友桌完，那么不能进入金币场
		检测是否正在金币场玩，如果在金币场完，那么检测gid,roomLevel是否一致，如果一致，那么可以玩，否则不可以玩
 */
func CanEnter(userId uint32, gid int32, roomType int32, roomLevel int32) (error, *ddproto.GameSession) {
	var getError = func(gid int32, roomType int32) error {
		msg := fmt.Sprintf("已经在游戏%v中了，结束之后再来吧!", utils.GetGameName(gid, roomType))
		return Error.NewError(consts.ACK_RESULT_ERROR, msg)
	}

	//判断朋友桌是否能进去
	//1,如果没有session表示没有在游戏中，可以直接进入
	s := GetSessionAuto(userId)
	if s == nil {
		return nil, nil
	}

	/**
		如果请求的是朋友桌子
		1，在玩金币场，可以进入朋友桌
		2，在朋友桌中，并且gid一样，那么可以进入朋友桌
	 */
	if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
		//如果玩家正在朋友桌子中完，但是gid 不一样，那么返回进入失败
		if s.GetRoomType() == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) &&
			(s.GetGameId() != gid || s.GetRoomLevel() != roomLevel) {
			return getError(s.GetGameId(), s.GetRoomType()), nil
		}
	}

	/**
		如果请求的是金币场
		1，在玩金币场，并且游戏id，等级一样的表示可以进入
	 */
	if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
		//如果正在朋友桌子中玩，返回进入金币场失败
		if s.GetRoomType() == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
			return getError(s.GetGameId(), s.GetRoomType()), nil
		}

		//如果gid 不一样，或者level不一样，返回进入失败
		if s.GetGameId() != gid || s.GetRoomLevel() != roomLevel {
			return getError(s.GetGameId(), s.GetRoomType()), nil
		}
	}

	//其他条件可以进入房间
	return nil, s
}

//通过suerId roomType 删除玩家的session
func DelSessionByKey(userId uint32, roomType int32, gid int32, deskId int32) {
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

	if s.GetDeskId() != deskId {
		log.E("删除玩家[%v]roomType[%v],gid[%v],dekid[%v] 的时候出错 session[%v],",
			userId, roomType, gid, deskId, s)
		return
	}
	delSession(s)
}

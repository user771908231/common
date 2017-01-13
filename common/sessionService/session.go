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
const (
	GAMESESSION_STATUS_NOGAME int32 = 1; //没有在游戏中
	GAMESESSION_STATUS_GAMING int32 = 2; //没有在游戏中

	GAMEID_DEZHOU    int32 = 1                                         //德州
	GAMEID_MAJAING   int32 = int32(ddproto.CommonEnumGame_GID_MAHJONG) //麻将
	GAMEID_DOUZIZHU  int32 = int32(ddproto.CommonEnumGame_GID_DDZ)     //斗地主
	GAMEID_ZHAJINHUA int32 = int32(ddproto.CommonEnumGame_GID_ZJH)     //扎金花
)

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

//更新用户的session信息，具体更新什么信息待定
func UpdateSession(userId uint32, gameStatus int32, gameId int32, gameNumber int32, roomId int32, deskId int32, gameCustomStatus int32, isBreak bool, isLeave bool, roomType int32, roomPass string) (*ddproto.GameSession, error) {
	session := GetSession(userId, roomType)
	if session == nil {
		log.E("没有找到user[%v]的session,需要重新申请一个并保存...", userId)
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

package sessionService

import (
	"strings"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"casino_common/common/log"
	"casino_common/proto"
	"casino_common/proto/funcsInit"
)
// session相关的...
const (
	GAMESESSION_STATUS_NOGAME int32 = 1; //没有在游戏中
	GAMESESSION_STATUS_GAMING int32 = 2; //没有在游戏中

	GAMEID_DEZHOU = 1; //德州
	GAMEID_MAJAING = 2; //麻将
	GAMEID_DOUZIZHU = 3; //斗地主
	GAMEID_ZHAJINHUA = 4; //这金花
)

var MJSESSION_KEY_PRE = "redis_game_session"

func getSessionKey(userId uint32) string {
	idstr, _ := numUtils.Uint2String(userId)
	ret := strings.Join([]string{MJSESSION_KEY_PRE, idstr}, "_")
	return ret
}

func GetSession(userId uint32) *casinoCommonProto.GameSession {
	s := redisUtils.GetObj(getSessionKey(userId), &casinoCommonProto.GameSession{})
	if s != nil {
		return s.(*casinoCommonProto.GameSession)
	} else {
		return nil
	}
}

//更新用户的session信息，具体更新什么信息待定
func UpdateSession(userId uint32, gameStatus int32, gameId int32, gameNumber int32, roomId int32, deskId int32, gameCustomStatus int32, isBreak bool, isLeave bool) (*casinoCommonProto.GameSession, error) {
	session := GetSession(userId)
	if session == nil {
		log.T("没有找到user[%v]的session,需要重新申请一个并保存...", userId)
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

	//保存session
	log.T("保存到redis的session【%v】", session)
	redisUtils.SetObj(getSessionKey(userId), session)
	return session, nil
}

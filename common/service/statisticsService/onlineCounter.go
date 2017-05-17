package statisticsService

import (
	"casino_common/common/consts"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"fmt"
)

//在线玩家的统计
func IncrOnline(userId uint32, gameId int32, ) {
	redisUtils.SADD(consts.RKEY_ALL_ONLINE_USERS, fmt.Sprintf("%v", userId))
	redisUtils.SADD(redisUtils.KINT32(consts.RKEY_PRE_ONLINE_SINGE_GAME, gameId), fmt.Sprintf("%v", userId))
}

func DecrOnline(userId uint32, gameId int32, ) {
	redisUtils.SREM(consts.RKEY_ALL_ONLINE_USERS, fmt.Sprintf("%v", userId))
	redisUtils.SREM(redisUtils.KINT32(consts.RKEY_PRE_ONLINE_SINGE_GAME, gameId), fmt.Sprintf("%v", userId))
}

func OnlineCountAll() int64 {
	return redisUtils.SCARD(consts.RKEY_ALL_ONLINE_USERS)
}

func OnlineCountByGame(gameId int32) int64 {
	return redisUtils.SCARD(redisUtils.KINT32(consts.RKEY_ALL_ONLINE_USERS, gameId))
}

func OnlineCountUsers() []uint32 {
	ret := make([]uint32, 0)
	ret2 := redisUtils.SMEMBERS(consts.RKEY_ALL_ONLINE_USERS)
	for _, r := range ret2 {
		ret = append(ret, numUtils.String2Uint32(r))
	}
	return ret
}

func OnlineCountUsersByGame(gameId int32) []uint32 {
	ret := make([]uint32, 0)
	ret2 := redisUtils.SMEMBERS(redisUtils.KINT32(consts.RKEY_ALL_ONLINE_USERS, gameId))
	for _, r := range ret2 {
		ret = append(ret, numUtils.String2Uint32(r))
	}
	return ret
}

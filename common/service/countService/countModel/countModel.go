package countModel

import (
	"casino_common/proto/ddproto"
	"casino_common/common/service/countService/countType"
	"casino_common/common/service/countService/gameCounter"
)

//获取用户在某游戏内消费的房费
func GetUserDayCoinFeeByGameId(userId uint32, gameId ddproto.CommonEnumGame) *gameCounter.Counter {
	var count_type countType.CountType
	switch gameId {
	case ddproto.CommonEnumGame_GID_MAHJONG:
		count_type = countType.MJ_COIN_FEE_COUNT
	case ddproto.CommonEnumGame_GID_DDZ:
		count_type = countType.DDZ_COIN_FEE_COUNT
	case ddproto.CommonEnumGame_GID_ZJH:
		count_type = countType.ZJH_COIN_FEE_COUNT
	default:
		return nil
	}

	return gameCounter.GetUserCounterByType(userId, count_type)
}

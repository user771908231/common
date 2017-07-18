package countService

import (
	"casino_common/common/model/hall"
	"time"
)

//朋友桌房卡游戏内真实消耗统计
func AddFriendRoomCardTrueConsume(owner_id uint32, num int64, gid int32) {
	hall.T_statistics_roomcard{
				Time:          time.Now(),
				UserId:        owner_id,
				Gid:           gid,
				Memo:          "真实房卡消耗",
				RoomCardCount: num,
	}.Insert()
}

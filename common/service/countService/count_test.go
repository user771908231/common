package countService

import (
	"testing"
	"casino_common/common/service/countService/countType"
	ldb "casino_common/common/db"
	"casino_common/common/service/countService/gameCounter"
	"casino_common/proto/ddproto"
	"time"
	"casino_common/common/service/taskService/task"
)

func init() {
	ldb.InitMongoDb("192.168.199.200", 27017, "test", "id",[]string{})
}

//
func TestGetUserCounterByType(t *testing.T) {
	gameCounter.Add(1, countType.ZJH_GAME_COUNT, 1)
	t.Log(gameCounter.GetAllCounter(1))
	t.Log(gameCounter.GetUserCounterByType(1, countType.ZJH_GAME_COUNT))
	t.Log(gameCounter.GetDayCounter(1))
}

//触发统计
func TestT_game_log_DoCountAndTask(t *testing.T) {
	task.InitTask()
	row := T_game_log{
		UserId: 11,
		GameId:ddproto.CommonEnumGame_GID_ZJH,
		GameNumber: 323541,
		RoomType: ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN,
		RoomLevel:1,
		Bill: 640,
		IsWine: true,
		StartTime:time.Now().Unix() - int64(10*60*time.Second),
		EndTime: time.Now().Unix(),
	}
	t.Log(gameCounter.GetDayCounter(11))
	row.DoCountAndTask()
	t.Log(gameCounter.GetDayCounter(11))
}

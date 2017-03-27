package countService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/service/taskService"
	"casino_common/common/service/countService/countType"
	"casino_common/common/service/pushService"
	"github.com/golang/protobuf/proto"
	"casino_common/common/service/countService/gameCounter"
	"fmt"
)

//游戏记录表
type T_game_log struct {
	UserId uint32  //用户名
	GameId ddproto.CommonEnumGame   //游戏
	GameNumber int32  //游戏编号
	RoomType ddproto.COMMON_ENUM_ROOMTYPE //房间类型
	RoomLevel int32  //房间等级
	Bill float64    //账单--输赢多少钱
	IsWine bool  //是否赢得比赛
	CoinFee int32   //消耗的门票
	Data interface{}  //附加数据
	StartTime int64    //开始时间
	EndTime int64    //结束时间
}

//插入到数据库
func (t *T_game_log)Insert() error {
	data := pushService.AssembleData(ddproto.HallEnumProtoId_HALL_PID_PUSH_GAME_COUNT, &ddproto.CommonSrvMsgGameCount{
		UserId: proto.Uint32(t.UserId),
		GameId: t.GameId.Enum(),
		GameNumber: proto.Int32(t.GameNumber),
		RoomType: t.RoomType.Enum(),
		RoomLevel: proto.Int32(t.RoomLevel),
		Bill: proto.Float64(t.Bill),
		IsWine: proto.Bool(t.IsWine),
		StartTime: proto.Int64(t.StartTime),
		EndTime: proto.Int64(t.EndTime),
		CoinFee: proto.Int32(t.CoinFee),
	})
	//推送到大厅服务器处理
	return pushService.Push(data)
}

//触发统计
func (t *T_game_log)DoCountAndTask() {
	//所有比赛局数统计
	gameCounter.Add(t.UserId, countType.ALL_GAME_COUNT, 1)
	taskService.OnTask(countType.ALL_GAME_COUNT, t.UserId)

	//所有赢的比赛统计
	if t.IsWine == true {
		gameCounter.Add(t.UserId, countType.ALL_GAME_WIN, 1)
		taskService.OnTask(countType.ALL_GAME_WIN, t.UserId)
	}

	//统计房费
	gameCounter.Add(t.UserId, countType.ALL_COIN_FEE_COUNT, int(t.CoinFee))
	switch t.GameId {
	case ddproto.CommonEnumGame_GID_DDZ:
		gameCounter.Add(t.UserId, countType.DDZ_COIN_FEE_COUNT, int(t.CoinFee))
	case ddproto.CommonEnumGame_GID_MAHJONG:
		gameCounter.Add(t.UserId, countType.MJ_COIN_FEE_COUNT, int(t.CoinFee))
	case ddproto.CommonEnumGame_GID_ZJH:
		gameCounter.Add(t.UserId, countType.ZJH_COIN_FEE_COUNT, int(t.CoinFee))
	}

	//不同level的房间比赛统计
	gameIdMap := map[ddproto.CommonEnumGame]string {
		ddproto.CommonEnumGame_GID_SRC: "all",
		ddproto.CommonEnumGame_GID_DDZ: "ddz",
		ddproto.CommonEnumGame_GID_HALL: "hall",
		ddproto.CommonEnumGame_GID_MAHJONG: "mj",
		ddproto.CommonEnumGame_GID_ZJH: "zjh",
		ddproto.CommonEnumGame_GID_PDK: "pdk",
	}
	isWinStr := "win"
	if !t.IsWine {
		isWinStr = "lose"
	}

	roomTypeStr := "friend"
	switch t.RoomType {
	case ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN:
		roomTypeStr = fmt.Sprintf("lv%d", t.RoomLevel)
	case ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND:
		roomTypeStr = "friend"
	}

	//游戏输赢统计
	count_type := countType.CountType(fmt.Sprintf("%s_%s_count", gameIdMap[t.GameId], isWinStr))
	gameCounter.Add(t.UserId, countType.CountType(count_type), 1)
	taskService.OnTask(count_type, t.UserId)

	//朋友桌和金币场房间内游戏局数统计
	countKeyStr := fmt.Sprintf("%s_%s_%s_count", gameIdMap[t.GameId], isWinStr, roomTypeStr)
	count_type = countType.CountType(countKeyStr)
	gameCounter.Add(t.UserId, count_type, 1)
	taskService.OnTask(count_type, t.UserId)


}

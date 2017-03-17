package countService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/service/taskService"
	"casino_common/common/service/countService/countType"
	"casino_common/common/service/pushService"
	"github.com/golang/protobuf/proto"
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
	})
	//推送到大厅服务器处理
	return pushService.Push(data)
}

//触发统计
func (t *T_game_log)DoCountAndTask() {
	//所有比赛局数
	Add(t.UserId, countType.ALL_GAME_COUNT, 1)
	taskService.OnTask(countType.ALL_GAME_COUNT, t.UserId)

	//赢比赛拆红包
	if t.IsWine == true {
		Add(t.UserId, countType.ALL_GAME_WIN, 1)
		taskService.OnTask(countType.ALL_GAME_WIN, t.UserId)
	}

}

package hall

import "time"

//统计房卡的消耗
type T_statistics_roomcard struct {
	Time          time.Time //使用的时间
	UserId        uint32    //使用的玩家
	Gid           int32     //游戏的Id
	Memo          string    //备注
	RoomCardCount int64     //消耗房卡的数量
}

//房卡的每日消耗记录
type T_statistics_roomcard_day_details struct {
	Time          string //时间，按天来统计
	Gid           int32  //游戏id
	RoomCardCount int64  //房卡的数量
}

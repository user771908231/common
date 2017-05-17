package hall

import "time"

//时间记录表

const (
	EVENT_DISSOLVE int32 = 1001 //解散房间
)

type T_statistics_event struct {
	EventId   int32     //事件id
	EventDes  string    //事件描述
	EventUser uint32    //操作时间的玩家
	Time      time.Time //操作事件的时间
}

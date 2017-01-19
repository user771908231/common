package countService

import (
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"casino_common/common/service/taskService"
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
	var err error = nil
	db.Query(func(d *mgo.Database) {
		err = d.C(tableName.DBT_T_GAME_LOG).Insert(t)
	})
	//触发统计任务
	t.doCount()
	t.doTask()
	return err
}

//触发统计
func (t *T_game_log)doCount() {
	counter := GetAllCounter(t.UserId)
	counter.All_Count += 1
	counter.Save()
}

//触发任务系统
func (t *T_game_log)doTask() {
	//总局数相关任务
	taskService.OnTask(taskService.TYPE_ALL_GAME_COUNT, t.UserId)
}

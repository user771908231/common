package countService

import (
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2/bson"
	"github.com/chanxuehong/time"
)

//游戏总计数表
type T_game_count struct {
	UserId uint32  //用户名
	All_Count int32  //所有局数

}

//游戏每日计数表
type T_game_day_count struct {
	Day string  //日期
	T_game_count //统计
}

//获得用户游戏计数器
func GetAllCounter(userId uint32) *T_game_count {
	newCounter := &T_game_count{
		UserId:userId,
		All_Count:0,
	}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_GAME_COUNT).Find(bson.M{
			"userid": userId,
		}).One(newCounter)
	})
	return newCounter
}

//保存
func (t *T_game_count)Save() {
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_GAME_COUNT).Upsert(bson.M{
			"userid": t.UserId,
		},t)
	})
}

//获得用户游戏每日计数器
func GetDayCounter(userId uint32) *T_game_day_count {
	dayStr := time.Now().Format("2006-01-02")
	newCounter := &T_game_day_count{
		Day: dayStr,
		T_game_count: T_game_count{
			UserId:userId,
		},
	}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_GAME_DAY_COUNT).Find(bson.M{
			"userid": userId,
			"day": dayStr,
		}).One(newCounter)
	})
	return newCounter
}

//保存
func (t *T_game_day_count)Save() {
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_GAME_DAY_COUNT).Upsert(bson.M{
			"userid": t.UserId,
			"day": time.Now().Format("2006-01-02"),
		},t)
	})
}

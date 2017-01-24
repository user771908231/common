package countService

import (
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2/bson"
	"github.com/chanxuehong/time"
	"casino_common/common/service/countService/countType"
)

//游戏总计数表
type T_game_count struct {
	UserId uint32  //用户名
	All_Game_Count int32  //所有局数
}



//游戏每日计数表
type T_game_day_count struct {
	Day string  //日期
	T_game_count //统计
}

//统计器
type Counter struct {
	CountType countType.CountType  //统计类型
	AllCount  int32                //所有数据
	DayCount  int32                //每日数据
}

//获得用户游戏计数器
func GetAllCounter(userId uint32) *T_game_count {
	newCounter := &T_game_count{
		UserId:userId,
	}
	var err error = nil
	db.Query(func(d *mgo.Database) {
		err = d.C(tableName.DBT_T_GAME_COUNT).Find(bson.M{
			"userid": userId,
		}).One(newCounter)
	})
	if err != nil {
		db.Query(func(d *mgo.Database) {
			d.C(tableName.DBT_T_GAME_COUNT).Insert(newCounter)
		})
	}
	return newCounter
}

//统计数自增
func Add(userId uint32,countType countType.CountType, num int) {
	GetAllCounter(userId)
	GetDayCounter(userId)
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_GAME_COUNT).Update(bson.M{
			"userid": userId,
		},bson.M{
			"$inc":bson.M{string(countType): num},
		})
		d.C(tableName.DBT_T_GAME_DAY_COUNT).Update(bson.M{
			string("t_game_count.userid"): userId,
			"day": time.Now().Format("2006-01-02"),
		},bson.M{
			"$inc":bson.M{string("t_game_count."+countType): num},
		})
	})
}

//根据统计类型返回
func GetUserCounterByType(userId uint32, countType countType.CountType) *Counter {
	counter := &Counter{
		CountType:countType,
		AllCount:0,
		DayCount:0,
	}
	var num map[string]int32
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_GAME_COUNT).Find(bson.M{
			"userid": userId,
		}).Select(bson.M{
			string(countType): 1,
		}).One(&num)
		if nums,ok := num[string(countType)];ok {
			counter.AllCount = nums
		}
		type t_day_counter struct {
			T_game_count map[string]int32
		}
		var t_count_day t_day_counter
		d.C(tableName.DBT_T_GAME_DAY_COUNT).Find(bson.M{
			string("t_game_count.userid"): userId,
			"day":time.Now().Format("2006-01-02"),
		}).Select(bson.M{
			string("t_game_count."+countType): 1,
		}).One(&t_count_day)
		if nums,ok := t_count_day.T_game_count[string(countType)];ok {
			counter.DayCount = nums
		}
	})
	return counter
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
	var err error = nil
	newCounter := &T_game_day_count{
		Day: dayStr,
		T_game_count: T_game_count{
			UserId:userId,
		},
	}
	db.Query(func(d *mgo.Database) {
		err = d.C(tableName.DBT_T_GAME_DAY_COUNT).Find(bson.M{
			"t_game_count.userid": userId,
			"day": dayStr,
		}).One(newCounter)
	})
	if err != nil {
		db.Query(func(d *mgo.Database) {
			d.C(tableName.DBT_T_GAME_DAY_COUNT).Insert(newCounter)
		})
	}
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

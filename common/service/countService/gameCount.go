package countService

import (
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2/bson"
)

//游戏统计表
type T_game_count struct {
	UserId uint32  //用户名
	All_Count int32  //所有局数
}

//获得用户游戏统计器
func GetUserCounter(userId uint32) *T_game_count {
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

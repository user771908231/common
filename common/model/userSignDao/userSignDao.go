package userSignDao

import (
	"casino_common/common/model"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
)

//获取签到天数对应的奖品
func FindSignRewardByDay(day int32) *model.T_sign_reward {
	row := &model.T_sign_reward{}
	db.Query("", func(d *mgo.Database) {
		err := d.C(tableName.DBT_T_SIGN_REWARD).Find(bson.M{
			"day": bson.M{"$lte": day},
		}).Sort("-day").One(row)
		if err != nil {
			row = nil
		}
	})
	return row
}


//获取签到天数对应的奖品
func FindSignRewardList() []*model.T_sign_reward {
	row := []*model.T_sign_reward{}
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_T_SIGN_REWARD).Find(bson.M{}).All(&row)
	})
	return row
}

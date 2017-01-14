package drawLotteryDao

import (
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/common/model"
	"casino_common/utils/db"
)

func FindDrawLotteryItemByKV(key string, v interface{}) *model.T_Draw_Lottery {
	ret := &model.T_Draw_Lottery{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_DRAW_LOTTERY).Find(bson.M{key: v}).Sort("Sort").One(ret)
	})
	return ret
}

func FindDrawLotteryItemsByKV(key string, v interface{}) []*model.T_Draw_Lottery {
	ret := []*model.T_Draw_Lottery{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_DRAW_LOTTERY).Find(bson.M{key: v}).Sort("Sort").All(&ret)
	})
	return ret
}


func FindDrawLotteryFreshVersion() float32 {
	ret := &model.T_Draw_Lottery{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_DRAW_LOTTERY).Find(nil).Sort("-Version").One(ret)
	})
	return ret.Version
}

func FindFreshDrawLotteryItems() []*model.T_Draw_Lottery {
	v := FindDrawLotteryFreshVersion()
	return FindDrawLotteryItemsByKV("Version", v)
}

//新增一条
func InserDrawLotteryItem(item *model.T_Draw_Lottery) {
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_DRAW_LOTTERY).Insert(item)
	})
}

//更新




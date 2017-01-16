package userStrongBoxDao

import (
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/common/model"
	"casino_common/utils/db"
)

func FindUserStrongBoxByKV(key string, v interface{}) *model.T_user_strongBox {
	tuser := &model.T_user_strongBox{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER_STRONGBOX).Find(bson.M{key: v}).One(tuser)
	})
	if tuser.UserId <= 0 {
		return nil
	}
	return tuser
}

func FindUserStrongBoxByUserId(userId uint32) *model.T_user_strongBox {
	return FindUserStrongBoxByKV("userid", userId)
}

//插入一条数据
func InsertUserStrongBox(t *model.T_user_strongBox) {
	db.InsertMgoData(tableName.DBT_T_USER_STRONGBOX, t)
}

func UpdateUserStrongBoxByModel(t *model.T_user_strongBox) {
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER_STRONGBOX).Update(bson.M{
			"userid": t.UserId,
		}, t)
	})
}

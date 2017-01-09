package userAttachDao

import (
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"casino_common/common/model"
	"casino_common/utils/db"
)

func FindUserAttachByKV(key string, v interface{}) *model.T_user_attach {
	tuser := &model.T_user_attach{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER_ATTACH).Find(bson.M{key: v}).One(tuser)
	})
	if *tuser.UserId <= uint32(0) {
		return nil
	}
	return tuser
}

//通过id 查找一个user
func FindUserAttachByUserId(userId uint32) *model.T_user_attach {
	return FindUserAttachByKV("userId", userId)
}

//
func SaveUser2Mgo(user model.T_user_attach) error {
	return db.InsertMgoData(tableName.DBT_T_USER, user)
}

func UpdateUser2Mgo(user model.T_user_attach) error {
	return db.UpdateMgoDataU32(tableName.DBT_T_USER, user)
}


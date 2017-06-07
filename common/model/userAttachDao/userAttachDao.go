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
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER_ATTACH).Find(bson.M{key: v}).One(tuser)
	})
	if tuser.UserId > 0 {
		//log.T("return tuser userid[%v]", tuser.UserId)
		return tuser
	}
	return nil
}

//通过userId 查找一个attach
func FindUserAttachByUserId(userId uint32) *model.T_user_attach {
	userAttach := FindUserAttachByKV("userid", userId)
	if userAttach != nil {
		return userAttach
	}
	userAttach = &model.T_user_attach{}
	userAttach.UserId = userId
	InsertUserAttachByModel(userAttach)
	return userAttach
}

//新增
func InsertUserAttachByModel(t *model.T_user_attach) {
	db.Query("", func(d *mgo.Database){
		d.C(tableName.DBT_T_USER_ATTACH).Insert(t)
	})
}

//更新
func UpdateUserAttachByModel(t *model.T_user_attach) {
	db.Query("", func(d *mgo.Database){
		d.C(tableName.DBT_T_USER_ATTACH).Update(bson.M{
			"userid": t.UserId,
		}, t)
	})
}


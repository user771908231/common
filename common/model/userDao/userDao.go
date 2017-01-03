package userDao

import (
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
)

func FindUserByKV(key string, v interface{}) *ddproto.User {
	tuser := new(ddproto.User)
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER).Find(bson.M{key: v}).One(tuser)
	})
	if tuser.GetId() > 0 {
		return tuser
	} else {
		return nil
	}
}

//找到对应的所有user
func FindUsersByKV(key string, v interface{}) []*ddproto.User {
	var tuser []*ddproto.User
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER).Find(bson.M{key: v}).All(tuser)
	})
	return tuser
}

//通过id 查找一个user
func FindUserById(userId uint32) *ddproto.User {
	return FindUserByKV("id", userId)
}

//通过opend查找一个user
func FindUserByOpenId(openId string) *ddproto.User {
	return FindUserByKV("openid", openId)
}

//通过opend查找一个user
func FindUserByUnionId(unionid string) *ddproto.User {
	return FindUserByKV("unionid", unionid)
}

//
func SaveUser2Mgo(user *ddproto.User) error {
	return db.InsertMgoData(tableName.DBT_T_USER, user)
}

func UpdateUser2Mgo(user *ddproto.User) error {
	return db.UpdateMgoData(tableName.DBT_T_USER, user)
}

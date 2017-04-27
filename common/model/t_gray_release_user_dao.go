package model

import (
	"casino_hall/data/model"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
	"casino_common/common/log"
)

func init() {
	//GetGrayReleaseUsers() //初始化所有的灰度发布的白名单的玩家
}

var GrayUsers []model.T_gray_release_user

//得到所有白名单的玩家
func GetGrayReleaseUsers() []model.T_gray_release_user {
	var users []model.T_gray_release_user
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_GRAY_RELEASE_USER).Find(bson.M{"status": 1}).All(&users)
	})
	GrayUsers = users
	log.T("找到所有白名单的玩家:%v", GrayUsers) //初始化所有白名单的玩家...
	return users
}

//判断是否在灰度中的玩家
func IsWhiteUser(userId uint32) bool {
	for _, u := range GrayUsers {
		if u.UserId == userId {
			return true
		}
	}
	return false
}

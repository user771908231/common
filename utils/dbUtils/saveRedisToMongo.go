package dbUtils

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2/bson"
)

//将redis中的用户数据落地到mongo,方便查询和排序
func SaveAllRedisUserToMongo(confirm bool) {
	if confirm == false {
		return
	}
	log.T("开始执行，将redis中所有的user数据保存到mongo.")
	users := []*ddproto.User{}
	err, _ := db.C(tableName.DBT_T_USER).Page(bson.M{}, &users, "-id", 1, 1)

	max_user := new(ddproto.User)
	if len(users) > 0 {
		max_user = users[0]
	} else {
		return
	}

	const min_id uint32 = 10164
	if err != nil || max_user.GetId() <= min_id {
		log.E("读取最大userid出错:err:v max_id:%v", err, max_user.GetId())
		return
	}

	//开始批量更新
	success_count, fail_count := 0, 0
	for i := min_id; i <= max_user.GetId(); i++ {
		user := userService.GetUserById(i)
		if user != nil {
			err := db.C(tableName.DBT_T_USER).Update(bson.M{"id": user.GetId()}, bson.M{
				"$set": bson.M{
					"coin":     user.GetCoin(),
					"roomcard": user.GetRoomCard(),
					"diamond":  user.GetDiamond(),
					"diamond2": user.GetDiamond2(),
				},
			})
			if err == nil {
				log.T("Save User Success:%v", user)
				success_count++
			} else {
				log.E("Save User Fail:%v Error:%v", user, err)
				fail_count++
			}
		}
	}
	log.T("执行完毕，将redis中所有的user数据保存到mongo,共同步成功%d个用户的数据，失败%d个。", success_count, fail_count)
}

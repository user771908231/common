package agentModel

import (
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"fmt"
	"github.com/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//
type UserInfo struct {
	Info oauth2.UserInfo //微信信息

}

//代理模型
func GetUserIdByUnionId(op_id string) uint32 {
	user := ddproto.User{}
	err := db.C(tableName.DBT_T_USER).Find(bson.M{"unionid": op_id}, &user)
	if err != nil {
		log.Println(err)
		return 0
	}
	return user.GetId()
}

//查找代理下面的玩家id
func GetAgentInvUsersId(agent_id uint32) []uint32 {
	resp := struct {
		Users []uint32
	}{
		Users: []uint32{},
	}

	db.C(tableName.DBT_T_USER).Pipe([]bson.M{
		bson.M{
			"$match": bson.M{"invcode": fmt.Sprintf("%d", agent_id)},
		},
		bson.M{
			"$group": bson.M{"_id": nil, "users": bson.M{"$addToSet": "$id"}},
		},
	}, &resp)

	return resp.Users
}

//获取代理的推荐用户
func GetAgentInvUsersPage(agent_id uint32, limit int, page int) (list []ddproto.User) {
	db.C(tableName.DBT_T_USER).Page(bson.M{
		"id": bson.M{"$in": GetAgentInvUsersId(agent_id)},
	}, &list, "", page, limit)
	return
}

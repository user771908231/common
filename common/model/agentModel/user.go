package agentModel

import (
	"github.com/chanxuehong/wechat.v2/mp/oauth2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"log"
)


//
type UserInfo struct {
	Info oauth2.UserInfo  //微信信息

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

package agentModel
import (
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
)

/**
用户充值统计
 */


//获取单个人的充值钻石总和
func GetInvUserPayRoomcardCount(userId uint32) int64 {
	resp := struct {
		Sum int64
	}{}
	query := []bson.M{
		bson.M{"$match": bson.M{
			"userid": userId,
			"status": ddproto.PayEnumTradeStatus_PAY_S_SUCC,
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$diamond"},
		}},
	}
	db.C(tableName.DBT_T_PAYBASEDETAILS).Pipe(query, &resp)

	return resp.Sum
}

//获取单个用户一共充值多少张房卡
func GetAgentChildsPayRoomcardCount(agent_id uint32, minTime int64, maxTime int64) int64 {
	//代理的用户
	invUsers := GetAgentInvUsersId(agent_id)

	resp := struct {
		Sum int64
	}{}
	query := []bson.M{
		bson.M{"$match": bson.M{
			"userid": bson.M{"$in": invUsers},
			"status": ddproto.PayEnumTradeStatus_PAY_S_SUCC,
			"time": bson.M{
				"$gte": minTime,
				"$lte": maxTime,
			},
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$diamond"},
		}},
	}
	db.C(tableName.DBT_T_PAYBASEDETAILS).Pipe(query, &resp)

	return resp.Sum
}

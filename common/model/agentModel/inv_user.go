package agentModel

import (
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"gopkg.in/mgo.v2/bson"
)

/**
用户充值统计
*/

//获取单个人的充值钻石总和
func GetInvUserPayDiamondCount(userId uint32) float64 {
	resp := struct {
		Sum float64
	}{}
	query := []bson.M{
		bson.M{"$match": bson.M{
			"userid": userId,
			"status": ddproto.PayEnumTradeStatus_PAY_S_SUCC,
			"time":   bson.M{"$gt": 0},
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$money"},
		}},
	}
	db.C(tableName.DBT_T_PAYBASEDETAILS).Pipe(query, &resp)

	return numUtils.Float64Format(resp.Sum)
}

//获取多个用户
func GetInvUserListPayDiamondCount(user_list []uint32, minTime int64, maxTime int64) float64 {
	resp := struct {
		Sum float64
	}{}
	query := []bson.M{
		bson.M{"$match": bson.M{
			"userid": bson.M{"$in": user_list},
			"status": ddproto.PayEnumTradeStatus_PAY_S_SUCC,
			"time": bson.M{
				"$gte": minTime,
				"$lte": maxTime,
			},
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$money"},
		}},
	}
	db.C(tableName.DBT_T_PAYBASEDETAILS).Pipe(query, &resp)

	return numUtils.Float64Format(resp.Sum)
}

//获取单个代理的玩家一共充值多少张房卡
func GetAgentChildsGamerPayDiamondCount(agent_id uint32, minTime int64, maxTime int64) float64 {
	//代理的用户
	invUsers := GetAgentInvUsersId(agent_id)

	return GetInvUserListPayDiamondCount(invUsers, minTime, maxTime)
}

//获取某个代理的子代理的玩家一共充值了多少钻石
func GetAgentChildAgentGamerPayDiamondCount(agent_id uint32, minTime int64, maxTime int64) float64 {
	//代理的用户
	invUsers := []uint32{}

	for _, agent := range GetAgentChildrens(agent_id) {
		invUsers = append(invUsers, GetAgentInvUsersId(agent.UserId)...)
	}

	return GetInvUserListPayDiamondCount(invUsers, minTime, maxTime)
}

//获取代理的一级、二级、三级..子代理充值钻石数
func GetAgentChildAgentTreeGamerPayDiamondCount(agent_id uint32, minTime int64, maxTime int64) float64 {
	//代理的用户
	invUsers := []uint32{}

	for _, agent := range GetAgentChildsTree(agent_id) {
		invUsers = append(invUsers, GetAgentInvUsersId(agent.UserId)...)
	}

	return GetInvUserListPayDiamondCount(invUsers, minTime, maxTime)
}

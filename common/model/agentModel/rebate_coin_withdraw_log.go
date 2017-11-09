package agentModel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"casino_common/utils/numUtils"
)

type RebateCoinWithdrawLog struct {
	Id bson.ObjectId
	AgentId uint32
	Amount float64  //金额
	Status int32  //1.审核中 2.已通过 3.已拒绝
	Time time.Time
}

//插入
func (t RebateCoinWithdrawLog) Insert() error {
	t.Id = bson.NewObjectId()
	t.Time = time.Now()
	return db.Log(tableName.DBT_AGENT_COIN_FEE_WITHDRAW_LOG).Insert(t)
}

//根据状态返回提现金额之和
func GetAgentCoinFeeWithdrawAmountSumByStatus(agent_id uint32, status int32) float64 {
	var all_rebate struct {
		Sum float64
	}
	//总佣金
	db.Log(tableName.DBT_AGENT_COIN_FEE_WITHDRAW_LOG).Pipe([]bson.M{
		bson.M{"$match":bson.M{
			"agentid": bson.M{"$eq": agent_id},
			"status": bson.M{"$eq": status},
		}},
		bson.M{"$group":bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$amount"},
		}},
	}, &all_rebate)
	return numUtils.Float64Format(all_rebate.Sum)
}

//返回已提现的金额
func GetAgentCoinFeeRebateIsWithdrawedSum(agent_id uint32) float64 {

	return 0
}

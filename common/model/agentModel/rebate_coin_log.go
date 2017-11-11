package agentModel

import (
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//金币场代理返利日志
type CoinFeeRebateLog struct {
	Id       bson.ObjectId `bson:"_id"`
	UserId   uint32        //玩家id
	UserName string        //玩家昵称
	GameId   int32         //所玩游戏
	CoinFee  int64         //房费
	Rebate   float64       //返利
	Time     time.Time     //时间
	AgentId  uint32        //领取的代理
	//FromAgent uint32    //来自多级返利的代理
	//Level    int        //级别
}

//插入一条领取记录
func (t CoinFeeRebateLog) Insert() error {
	t.Id = bson.NewObjectId()
	t.Time = time.Now()
	return db.Log(tableName.DBT_AGENT_COIN_FEE_REBATE_LOG).Insert(t)
}

//获取返利佣金总和
func GetAgentCoinFeeRebateAll(agent_id uint32) float64 {
	var all_rebate struct {
		Sum float64
	}
	//总佣金
	db.Log(tableName.DBT_AGENT_COIN_FEE_REBATE_LOG).Pipe([]bson.M{
		bson.M{"$match": bson.M{
			"agentid": bson.M{"$eq": agent_id},
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$rebate"},
		}},
	}, &all_rebate)
	return numUtils.Float64Format(all_rebate.Sum)
}

//获取今日佣金
func GetAgentCoinFeeRebateToday(agent_id uint32) float64 {
	var today_rebate struct {
		Sum float64
	}
	//今日佣金
	today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	db.Log(tableName.DBT_AGENT_COIN_FEE_REBATE_LOG).Pipe([]bson.M{
		bson.M{"$match": bson.M{
			"agentid": bson.M{"$eq": agent_id},
			"time": bson.M{
				"$gte": today,
				"$lt":  today.AddDate(0, 0, 1),
			},
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$rebate"},
		}},
	}, &today_rebate)

	return numUtils.Float64Format(today_rebate.Sum)
}

//获取昨日佣金
func GetAgentCoinFeeRebateYesterday(agent_id uint32) float64 {
	var today_rebate struct{
		Sum float64
	}
	//今日佣金
	today,_ := time.Parse("2006-01-02", time.Now().AddDate(0,0,-1).Format("2006-01-02"))
	db.Log(tableName.DBT_AGENT_COIN_FEE_REBATE_LOG).Pipe([]bson.M{
		bson.M{"$match":bson.M{
			"agentid": bson.M{"$eq": agent_id},
			"time": bson.M{
				"$gte": today,
				"$lt": today.AddDate(0,0,1),
			},
		}},
		bson.M{"$group":bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$rebate"},
		}},
	}, &today_rebate)

	return numUtils.Float64Format(today_rebate.Sum)
}

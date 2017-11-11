package agentModel

import (
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//销售记录表
type SalesLog struct {
	Id      bson.ObjectId `bson:"_id"` //销售编号
	AgentId uint32        //代理商id
	UserId  uint32        //用户id
	Type    GoodsType     //充值类型 金币 钻石 房卡 等
	Num     int64         //数量
	Money   float64       //金额
	Remark  string        //备注
	AddTime time.Time     //销售时间
}

//新增代理商销售记录
func AddNewSalesLog(agent_id uint32, user_id uint32, goods_type GoodsType, num int64, money float64, remark string) (err error) {
	err = db.C(tableName.DBT_AGENT_SALES_LOG).Insert(SalesLog{
		Id:      bson.NewObjectId(),
		AgentId: agent_id,
		UserId:  user_id,
		Type:    goods_type,
		Num:     num,
		Money:   money,
		Remark:  remark,
		AddTime: time.Now(),
	})

	return err
}

//获取某代理商总出售的房卡数
func GetAgentSalesCount(agent_id uint32) int64 {
	resp := struct {
		Sum int64
	}{}
	query := []bson.M{
		bson.M{"$match": bson.M{
			"agentid": agent_id,
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$num"},
		}},
	}
	db.C(tableName.DBT_AGENT_SALES_LOG).Pipe(query, &resp)
	return resp.Sum
}

//获取某代理商用户总购买的房卡数
func GetAgentUserSalesCount(agent_id uint32, user_id uint32) int64 {
	resp := struct {
		Sum int64
	}{}
	query := []bson.M{
		bson.M{"$match": bson.M{
			"userid":  user_id,
			"agentid": agent_id,
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$num"},
		}},
	}
	db.C(tableName.DBT_AGENT_SALES_LOG).Pipe(query, &resp)
	return resp.Sum
}

//获取某代理商用户本周购买的房卡数
func GetAgentThisWeekOneUserSalesCount(agent_id uint32, user_id uint32) int64 {
	now := time.Now()
	y, m, d := now.Date()
	week_one := time.Date(y, m, d-int(now.Weekday())+1, 0, 0, 0, 0, time.Local)
	resp := struct {
		Sum int64
	}{}
	query := []bson.M{
		bson.M{"$match": bson.M{
			"userid":  user_id,
			"agentid": agent_id,
			"addtime": bson.M{
				"$gt": week_one,
				"$lt": now,
			},
		}},
		bson.M{"$group": bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$num"},
		}},
	}
	db.C(tableName.DBT_AGENT_SALES_LOG).Pipe(query, &resp)
	return resp.Sum
}

package agentModel

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//充值记录表
type RechargeLog struct {
	Id       bson.ObjectId `bson:"_id"` //编号
	Detail   string        //订单详情
	GoodsId  int32         //商品Id
	GoodsNum int64         //商品数量
	AgentId  uint32        //代理商id
	Money    float64       //总价格
	IsPay    bool          //是否已支付
	AddTime  time.Time     //充值时间
	DoneTime time.Time     //付款时间
}

//表更新
func (r *RechargeLog) Save() error {
	err := db.C(tableName.DBT_AGENT_RECHARGE_LOG).Update(bson.M{
		"_id": r.Id,
	}, r)
	return err
}

//新增一个代理商充值订单
func AddNewRechargeLog(agent_id uint32, goods_id int32, goods_num int64) *RechargeLog {
	goods := GetGoodsInfoById(goods_id)
	if goods == nil {
		return nil
	}
	newOrder := &RechargeLog{
		Id:       bson.NewObjectId(),
		GoodsId:  goods_id,
		GoodsNum: goods_num,
		Detail:   fmt.Sprintf("代理充值-%s × %d", goods.Name, goods_num),
		AgentId:  agent_id,
		Money:    goods.Price * float64(goods_num),
		AddTime:  time.Now(),
		DoneTime: time.Time{},
		IsPay:    false,
	}
	err := db.C(tableName.DBT_AGENT_RECHARGE_LOG).Insert(newOrder)
	if err != nil {
		return nil
	}
	return newOrder
}

//获取一个充值订单信息
func GetRechargeOrderId(order_id string) *RechargeLog {
	order := new(RechargeLog)
	err := db.C(tableName.DBT_AGENT_RECHARGE_LOG).Find(bson.M{"_id": bson.ObjectIdHex(order_id)}, order)
	if err != nil {
		return nil
	}
	return order
}

//设置订单为已完成，并发货。
func SetRechargeDone(order_id string) error {
	order := GetRechargeOrderId(order_id)
	if order == nil {
		return errors.New(fmt.Sprintf("未找到订单%s！", order_id))
	}
	goods := GetGoodsInfoById(order.GoodsId)
	goods_num := goods.Num * order.GoodsNum
	var err error
	if goods.Type == RoomCard {
		_, err = userService.INCRUserRoomcard(order.AgentId, goods_num, int32(ddproto.CommonEnumGame_GID_SRC), "代理商充值")
	}
	if err == nil {
		err = db.C(tableName.DBT_AGENT_RECHARGE_LOG).Update(bson.M{"_id": bson.ObjectIdHex(order_id)}, bson.M{"$set": bson.M{
			"ispay":    true,
			"donetime": time.Now(),
		}})
	}
	return err
}

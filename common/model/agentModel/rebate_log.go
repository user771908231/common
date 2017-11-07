package agentModel

import (
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"errors"
	"github.com/golang/protobuf/proto"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//返利配置
var RebateMan int64
var RebateSong int64


func init() {
	//返利配置初始化
	RebateConfig = []RebateConfigItem{
		//卖满10张房卡则奖励1张房卡
		RebateConfigItem{
			Num: RebateMan,
			Reward: []*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:   ddproto.HallEnumTradeType_PROPS_FANGKA.Enum(),
					Amount: proto.Float64(float64(RebateSong)),
				},
			},
		},
	}
}

//返利配置
type RebateConfigItem struct {
	Num    int64 //数量
	Reward []*ddproto.HallBagItem
}

var RebateConfig []RebateConfigItem

//返利发放记录表
type RebateRecord struct {
	Id                   bson.ObjectId `bson:"_id"`
	AgentId              uint32        //代理商id
	GiveNum              int64         //返利的房卡数
	CurrRoomCardSalesNum int64         //当前出售房卡数
	IsCheck              bool          //是否已领取返利
	Time                 time.Time     //发放时间
}

//插入一张新表
func (r *RebateRecord) Insert() error {
	r.Id = bson.NewObjectId()
	r.Time = time.Now()
	return db.C(tableName.DBT_AGENT_REBATE_LOG).Insert(r)
}

//编辑一张表
func (r *RebateRecord)Save() error {
	return db.C(tableName.DBT_AGENT_REBATE_LOG).Update(bson.M{
		"_id": r.Id,
	}, r)
}

//获取上次发放返利时的出售房卡数
func GetLastRebateRoomCardSalesNum(agent_id uint32) int64 {
	list := []RebateRecord{}
	db.C(tableName.DBT_AGENT_REBATE_LOG).Page(bson.M{
		"agentid": agent_id,
	}, &list, "-time", 1, 1)
	if len(list) > 0 {
		return list[0].CurrRoomCardSalesNum
	}
	return 0
}

//获取代理已领取的房卡返利数
func GetAgentAllRebateRoomCardNum(agent_id uint32) int64 {
	resp := struct {
		Sum int64
	}{}
	query := []bson.M{
		bson.M{"$match":bson.M{
			"agentid": agent_id,
		}},
		bson.M{"$group":bson.M{
			"_id": nil,
			"sum": bson.M{"$sum": "$givenum"},
		}},
	}
	db.C(tableName.DBT_AGENT_REBATE_LOG).Pipe(query, &resp)
	return resp.Sum
}

//尝试领取返利并插入记录
func DoGetRebateRoomCard(agent_id uint32) (err error) {
	all_sales_num := GetAgentSalesCount(agent_id)
	last_rebate_card_num := GetLastRebateRoomCardSalesNum(agent_id)

	num := all_sales_num - last_rebate_card_num

	var reward_num int64 = 0
	var card_num int64 = 0
	for i, _ := range RebateConfig {
		item := RebateConfig[-i]
		if item.Num <= num && len(item.Reward) > 0 {
			reward_num = int64(item.Reward[0].GetAmount())
			card_num = item.Num
			break
		}
	}

	//领取奖励
	if card_num > 0 && reward_num > 0 {
		//发放奖励
		//userService.INCRUserRoomcard(agent_id, reward_num)
		//生成记录
		new_row := &RebateRecord{
			AgentId:              agent_id,
			GiveNum:              reward_num,
			CurrRoomCardSalesNum: last_rebate_card_num + card_num,
		}
		new_row.Insert()
		return nil
	}

	return errors.New("not rebate.")
}

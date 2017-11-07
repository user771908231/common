package agentModel

import (
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
)

//商品
type Goods struct {
	GoodsId   int32     //商品编号
	Name string    //商品名
	Description string //详细描述
	Type GoodsType //商品类型
	Num  int64     //商品数量
	Price float64    //价格 RMB
	GiveNum int64     //赠送的数量
}

//商品类型
type GoodsType string

const (
	Coin GoodsType = "金币"
	Diamond GoodsType = "钻石"
	RoomCard GoodsType = "房卡"
)

//获取一个商品信息
func GetGoodsInfoById(goods_id int32) *Goods {
	goods := new(Goods)
	err := db.C(tableName.DBT_AGENT_GOODS).Find(bson.M{
		"goodsid": goods_id,
	},goods)
	if err != nil {
		return nil
	}
	return goods
}

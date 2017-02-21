package model

import (
	"gopkg.in/mgo.v2/bson"
	"casino_common/proto/ddproto"
)

//商品分类枚举
type CateEnum int

const (
	CATE_COIN_ZONE     CateEnum = 1 //金币兑换
	CATE_DIAMOND_ZONE  CateEnum = 2 //钻石购买
	CATE_EXCHANGE_ZONE CateEnum = 3 //兑换专区
	CATE_PROPS_ZONE    CateEnum = 4 //道具专区
	CATE_VIP_ZONE      CateEnum = 5 //会员购买
)

//商品表结构
type T_Goods_Row struct {
	GoodsId   int32
	ObjId     bson.ObjectId `bson:"_id"`
	Name      string
	Category  CateEnum
	PriceType ddproto.HallEnumTradeType
	Price     float64
	GoodsType ddproto.HallEnumTradeType
	Amount    float64 //amount 代表数量，不是现金， 1转世兑换100金币,,amount == 100
	Discount  string
	Image     string
	IsShow    bool
	Sort      int32
}

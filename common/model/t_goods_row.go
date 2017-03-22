package model

import (
	"gopkg.in/mgo.v2/bson"
	"casino_common/proto/ddproto"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
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
	ObjId     bson.ObjectId `bson:"_id" binding:"Required"`
	GoodsId   int32 `binding:Required`
	Name      string `binding:Required`
	Category  CateEnum `binding:Required`
	PriceType ddproto.HallEnumTradeType `binding:Required`
	Price     float64 `binding:Required`
	GoodsType ddproto.HallEnumTradeType `binding:Required`
	Amount    float64 `binding:Required` //amount 代表数量，不是现金， 1转世兑换100金币,,amount == 100
	Discount  string `binding:Required`
	Image     string `binding:Required`
	IsShow    bool `binding:Required`
	Sort      int32 `binding:Required`
}

//更新
func (r *T_Goods_Row) Save() (error) {
	return db.C(tableName.DBT_T_GOODS_INFO).Update(bson.M{
		"_id": r.ObjId,
	}, r)
}

//删除
func (r *T_Goods_Row) Remove() error {
	return db.C(tableName.DBT_T_GOODS_INFO).Remove(bson.M{
		"_id": r.ObjId,
	})
}

//插入
func (r *T_Goods_Row) Insert() error {
	r.ObjId = bson.NewObjectId()
	return db.C(tableName.DBT_T_GOODS_INFO).Insert(r)
}

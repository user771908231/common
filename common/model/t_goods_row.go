package model

import (
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2/bson"
)

//商品分类枚举
type CateEnum int

const (
	CATE_COIN_ZONE       CateEnum = 1 //金币兑换
	CATE_DIAMOND_ZONE    CateEnum = 2 //钻石购买
	CATE_EXCHANGE_ZONE   CateEnum = 3 //兑换专区
	CATE_PROPS_ZONE      CateEnum = 4 //道具专区
	CATE_DIRECT_BUY_ZONE CateEnum = 5 //直接购买-用于兼容黄圣
)

//商品表结构
type T_Goods_Row struct {
	ObjId     bson.ObjectId             `bson:"_id" binding:"Required" title:"id"`
	GoodsId   int32                     `bson:"goodsid" binding:"Required" title:"商品Id"`
	Name      string                    `bson:"name" binding:"Required" title:"名称"`
	Category  CateEnum                  `bson:"category" binding:"Required" title:"分类Id"`
	PriceType ddproto.HallEnumTradeType `bson:"pricetype" binding:"Required" title:"价格类型"`
	Price     float64                   `bson:"price" binding:"Required" title:"价格"`
	GoodsType ddproto.HallEnumTradeType `bson:"goodstype" binding:"Required" title:"商品类型"`
	Amount    float64                   `bson:"amount" binding:"Required" title:"商品数量"` //amount 代表数量，不是现金， 1转世兑换100金币,,amount == 100
	Discount  string                    `bson:"discount" binding:"Required" title:"折扣"`
	Image     string                    `bson:"image" binding:"Required" title:"图片"`
	IsShow    bool                      `bson:"isshow" binding:"Required" title:"是否显示"`
	Sort      int32                     `bson:"sort" binding:"Required" title:"排序"`
}

//更新
func (r *T_Goods_Row) Save() error {
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

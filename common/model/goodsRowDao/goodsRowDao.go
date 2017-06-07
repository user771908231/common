package goodsRowDao

import (
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/proto/ddproto"
	"casino_common/common/model"
	"casino_common/utils/db"
)

//获取分类下的商品列表
func GetGoodsListByCategory(category model.CateEnum) []*model.T_Goods_Row {
	row := []*model.T_Goods_Row{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_GOODS_INFO).Find(bson.M{
			"category": category,
			"isshow":   true,
		}).Sort("sort").All(&row)
	})
	return row
}

//将数据库字段转换为协议
func GetGoodsListMsg(category model.CateEnum) []*ddproto.HallGoodsItemMsg {
	goods_list := GetGoodsListByCategory(category)
	msg_list := []*ddproto.HallGoodsItemMsg{}
	for _, goods := range goods_list {
		new_item := &ddproto.HallGoodsItemMsg{
			GoodsId:   &goods.GoodsId,
			Name:      &goods.Name,
			GoodsType: goods.GoodsType.Enum(),
			Amount:    &goods.Amount,
			PriceType: goods.PriceType.Enum(),
			Price:     &goods.Price,
			Discount:  &goods.Discount,
			Image:     &goods.Image,
		}
		msg_list = append(msg_list, new_item)
	}
	return msg_list
}

//获取单个商品的信息
func GetGoodsInfo(goods_id int32) *model.T_Goods_Row {
	row := &model.T_Goods_Row{}
	db.Query(func(d *mgo.Database) {
		err := d.C(tableName.DBT_T_GOODS_INFO).Find(bson.M{
			"goodsid": goods_id,
		}).One(row)
		if err != nil {
			row = nil
		}
	})
	return row
}

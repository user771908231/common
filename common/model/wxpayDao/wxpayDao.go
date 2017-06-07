package wxpayDao

import (
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"github.com/golang/protobuf/proto"
)

//保存订单信息
func UpsertDetail(d *ddproto.PayBaseDetails) error {
	if d.GetId() == 0 {
		//需要插入
		return insertDetail(d)
	} else {
		return updateDetail(d)
		//需要更新
	}
}

//插入充值订单
func insertDetail(d *ddproto.PayBaseDetails) error {
	id, _ := db.GetNextSeq(tableName.DBT_T_PAYBASEDETAILS)
	d.Id = proto.Int32(id)
	return db.InsertMgoData("", tableName.DBT_T_PAYBASEDETAILS, d)
}

//更新订单状态
func updateDetail(d *ddproto.PayBaseDetails) error {
	return db.UpdateMgoData("", tableName.DBT_T_PAYBASEDETAILS, d)
}

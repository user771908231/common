package service

import (
	"casino_common/proto/ddproto"
	"casino_common/common/log"
	"github.com/golang/protobuf/proto"
)

//充值套餐

func GetMealById(id int32) *ddproto.PayBaseMeal {
	//返回测试数据
	test := &ddproto.PayBaseMeal{
		Id:     proto.Int32(1),
		Money:  proto.Int64(1),
		Diamond:proto.Int64(5)}
	return test
}

func GetPayModel(id int32) *ddproto.PayBasePaymodel {
	//return getDefaultPayModel()
	return getDDZPayModel()
}

//返回一个默认的支付方式
func getDefaultPayModel() *ddproto.PayBasePaymodel {
	ret := &ddproto.PayBasePaymodel{
		Id:    proto.Int32(1),
		AppId: proto.String(WXConfig.AppId),
		MchId: proto.String(WXConfig.MchId),
		AppKey:proto.String(WXConfig.ApiKey)}
	return ret
}

func getDDZPayModel() *ddproto.PayBasePaymodel {
	ret := &ddproto.PayBasePaymodel{
		Id:    proto.Int32(1),
		AppId: proto.String("wx968f1e52358d6a29"),
		MchId: proto.String("1426502002"),
		AppKey:proto.String("SjDDze834Or9Lndw3FeNoj4IqPoda330")}
	return ret
}

//通过meal的信息更细user的信息
func UpdateUserByMeal(userId uint32, m *ddproto.PayBaseMeal) error {
	log.T("通过套餐[%v]更新玩家[%v]的账户信息", m, userId)
	return nil
}

package service

import (
	"casino_common/proto/ddproto"
	"casino_common/common/log"
	"github.com/golang/protobuf/proto"
	"casino_common/common/userService"
	"errors"
	"fmt"
	"casino_common/common/model/wxpayDao"
)

//充值套餐

func GetMealById(id int32) *ddproto.PayBaseProduct {
	//返回测试数据
	test := &ddproto.PayBaseProduct{
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
func UpdateUserByMeal(tradeNo string) error {
	//不是重复回调，开始做余额更新的处理
	detail := GetDetailsByTradeNo(tradeNo)
	if detail == nil {
		msg := fmt.Sprintf("没有在数据中找到订单号[%v]对应的套餐..", tradeNo)
		log.E(msg)
		return errors.New(msg)
	}

	//判断是否是重复回调
	if detail.GetStatus() == ddproto.PayEnumTradeStatus_PAY_S_SUCC {
		log.E("tradeNo[%v]重复回调", tradeNo)
		return nil
	}

	log.T("更新订单[%v]的回调信息，detail[%v]", tradeNo, detail)
	//找到套餐
	meal := GetMealById(detail.GetProductId())
	//根据套餐增加用户的余额
	userService.INCRUserDiamond(detail.GetUserId(), meal.GetDiamond())
	//更新订单状态
	UpdateDetailsStatus(tradeNo, ddproto.PayEnumTradeStatus_PAY_S_SUCC)
	//保存订单到数据库...
	go func() {
		wxpayDao.UpsertDetail(detail) //保存到数据库
		//DelDetails(tradeNo)           //保存到数据库之后删除//	app收到回复之后再删除
	}()
	return nil
}

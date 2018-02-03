package service

import (
	"casino_common/common/Error"
	"casino_common/common/log"
	"casino_common/common/model"
	"casino_common/common/model/goodsRowDao"
	"casino_common/common/model/wxpayDao"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

//充值套餐

func GetMealById(id int32) *model.T_Goods_Row {
	//返回测试数据
	return goodsRowDao.GetGoodsInfo(id)
}

func GetPayModel(id int32) *ddproto.PayBasePaymodel {
	//return getDefaultPayModel()
	return getDDZPayModel()
}

//返回一个默认的支付方式
func getDefaultPayModel() *ddproto.PayBasePaymodel {
	ret := &ddproto.PayBasePaymodel{
		Id:     proto.Int32(1),
		AppId:  proto.String(WXConfig.AppId),
		MchId:  proto.String(WXConfig.MchId),
		AppKey: proto.String(WXConfig.ApiKey)}
	return ret
}

func getDDZPayModel() *ddproto.PayBasePaymodel {
	ret := &ddproto.PayBasePaymodel{
		Id:     proto.Int32(1),
		AppId:  proto.String(WXConfig.AppId),
		MchId:  proto.String(WXConfig.MchId),
		AppKey: proto.String(WXConfig.ApiKey)}
	return ret
}

//通过meal的信息更细user的信息
func UpdateUserByMeal(tradeNo string, total_fee float64) error {
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
	if meal.PriceType == ddproto.HallEnumTradeType_TRADE_RMB {
		//如果pricetype为RMB，则为新版
		switch meal.GoodsType {
		case ddproto.HallEnumTradeType_TRADE_DIAMOND:
			userService.INCRUserDiamond(detail.GetUserId(), int64(meal.Amount), "商城微信支付充钻石")
		case ddproto.HallEnumTradeType_PROPS_FANGKA:
			userService.INCRUserRoomcard(detail.GetUserId(), int64(meal.Amount), int32(ddproto.CommonEnumGame_GID_HALL), "商城微信支付充房卡")
		}
	} else {
		//旧版，直接增加钻石
		//根据套餐增加用户的余额
		userService.INCRUserDiamond(detail.GetUserId(), int64(meal.Amount), "商城微信支付充钻石")
	}

	//更新订单状态
	UpdateDetailsStatus(tradeNo, ddproto.PayEnumTradeStatus_PAY_S_SUCC)
	//保存订单到数据库...
	log.T("微信支付成功，为用户%d充值%d钻石。", detail.GetUserId(), int64(meal.Amount))

	go func() {
		defer Error.ErrorRecovery("UpdateUserByMeal wxpayDao.UpsertDetail")
		detail.Status = ddproto.PayEnumTradeStatus_PAY_S_SUCC.Enum()
		detail.Money = proto.Float64(total_fee)
		wxpayDao.UpsertDetail(detail) //保存到数据库
		//DelDetails(tradeNo)           //保存到数据库之后删除//	app收到回复之后再删除
	}()
	return nil
}

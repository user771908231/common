package service

import (
	"casino_common/common/log"
	"casino_common/proto/ddproto"
)

//处理回信异步回调

//只需要订单号就可以了？
func DoAsynCb(tradeNo string) error {
	//删除同步机制
	defer clearWxpay(tradeNo)

	//微信返回的需要对比
	err := EndTradeSyncStatus(tradeNo)
	if err != nil {
		return err
	}

	//不是重复回调，开始做余额更新的处理
	detal := GetDetailsByTradeNo(tradeNo)

	//找到套餐
	meal := GetMealById(detal.GetMealId())

	//更新余额
	errUpdate := UpdateUserByMeal(detal.GetUserId(), meal)
	if errUpdate != nil {
		log.E(errUpdate.Error())
		return errUpdate
	}
	UpdateDetailsStatus(tradeNo, ddproto.PayEnumTradeStatus_PAY_S_SUCC)
	return nil
}

//完成之后的清理工作
func clearWxpay(tradeNo string) {
	//保存数据到数据库
	deal := GetDetailsByTradeNo(tradeNo)
	SaveDetails2Mgo(deal)
	//删除回调的同步机制
	//redisUtils.Del(GetTradeSyncKey(tradeNo))
	//删除redis中的数据
	//redisUtils.Del(tradeNo)
}

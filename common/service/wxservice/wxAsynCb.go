package service

import (
	"casino_common/common/log"
)

//处理回信异步回调

//只需要订单号就可以了？
func DoAsynCb(tradeNo string, total_fee float64) error {
	//微信返回的需要对比
	//todo LockPay总是返回"没有获得锁"错误，暂时注释掉
	//err := LockPay(tradeNo)
	//if err != nil {
	//	log.E("重复回调err:%v,暂时给微信回复已经成功...", err)
	//	return nil
	//}
	//defer UnLockPay(tradeNo)

	//更新余额
	errUpdate := UpdateUserByMeal(tradeNo, total_fee)
	if errUpdate != nil {
		log.E(errUpdate.Error())
		return errUpdate
	}

	return nil
}

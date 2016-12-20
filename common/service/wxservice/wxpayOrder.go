package service

import "time"
import (
	wx "github.com/chanxuehong/wechat.v2/mch/core"
	"casino_common/utils/numUtils"
	"strings"
	"casino_server/common/log"
	"casino_common/utils/rand"
)

//这里是否需要加锁
//var lock sync.Mutex = sync.Mutex{}

//生成支付订单号
//格式是: time+peymodel+userid+mealid+randstr，不合适的话需要和@kory商量
func GetWxpayTradeNo(payModelId int32, userId uint32, mealId int32, tnow time.Time) string {
	//支付方式
	payModelIdStr, _ := numUtils.Int2String(payModelId)

	//userId 字符串
	userIdStr, _ := numUtils.Uint2String(userId)

	//购买的套餐
	mealIdStr, _ := numUtils.Int2String(mealId)

	//时间字符串
	timeStr := wx.FormatTime(tnow)

	//随机数
	rand := rand.Rand(100, 999)
	randStr, _ := numUtils.Int2String(rand)

	result := strings.Join([]string{timeStr, payModelIdStr, userIdStr, mealIdStr, randStr}, "")
	log.T("玩家[%v]通过paymodel[%v]购买套餐[%v]时的订单号是:[%v]", userId, payModelId, mealId, result)
	return result
}

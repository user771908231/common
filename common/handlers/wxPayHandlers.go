package handlers

import (
	"casino_common/common/service/wxservice"
	"github.com/name5566/leaf/gate"
	"casino_common/proto/ddproto"
	"casino_common/utils/agentUtils"
	"casino_hall/wxpay"
)

//处理统一下单
func HandlerWxpayUnifiedOrder(args []interface{}) {
	m := args[0].(*ddproto.WxpayReqUnifiedorder)
	a := args[1].(gate.Agent)
	//通过meal知道 充值的数目等
	ack, _ := service.GetAppWxpayReqParams(m.GetPayModelId(), m.GetMeal(), m.GetHeader().GetUserId(), agentUtils.GetIP(a), wxpay.DEFAULT_DEVICEINFO)
	a.WriteMsg(ack)
}

//app 同步返回之后请求轮询
func HanlderAppSyncCb(args []interface{}) {
	m := args[0].(*ddproto.WxpayReqSyncChecker)
	a := args[1].(gate.Agent)
	service.InitChecker(m.GetTradeNo(), a)
}

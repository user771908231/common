package handlers

import (
	"casino_common/common/service/wxservice"
	"github.com/name5566/leaf/gate"
	"casino_common/proto/ddproto"
	"casino_common/utils/agentUtils"
	"casino_hall/wxpay"
	"casino_common/common/sys"
	"github.com/golang/protobuf/proto"
)

//处理统一下单
func HandlerWxpayUnifiedOrder(args []interface{}) {
	m := args[0].(*ddproto.WxpayReqUnifiedorder)
	a := args[1].(gate.Agent)
	//通过meal知道 充值的数目等
	ack, _ := service.GetAppWxpayReqParams(m.GetPayModelId(), m.GetMeal(), m.GetHeader().GetUserId(), agentUtils.GetIP(a), wxpay.DEFAULT_DEVICEINFO)
	if sys.GAMEENV.APPLEPAY_ONLY {
		//方便上架..设置使用苹果支付
		ack = new(ddproto.WxpayAckUnifiedorder)
		ack.Header = new(ddproto.ProtoHeader)
		ack.Header.Code = proto.Int32(int32(ddproto.PayReturn_ERR_ONLY_APPLE))
		a.WriteMsg(ack)
	} else {
		//生产环境使用微信支付...
		a.WriteMsg(ack)
	}
}

//app 同步返回之后请求轮询
func HanlderAppSyncCb(args []interface{}) {
	m := args[0].(*ddproto.WxpayReqSyncChecker)
	a := args[1].(gate.Agent)
	service.InitChecker(m.GetTradeNo(), a)
}

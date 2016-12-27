package handlers

import (
	"github.com/name5566/leaf/gate"
	"casino_common/proto/ddproto"
	"casino_common/common/userService"
	"casino_server/common/log"
	"github.com/golang/protobuf/proto"
)

//处理统一下单
func HandlerApplePayRechargeCb(args []interface{}) {
	m := args[0].(*ddproto.ApplepayReqRechargecb)
	a := args[1].(gate.Agent)
	//通过meal知道 充值的数目等
	userId := m.GetHeader().GetUserId()
	coin, err := userService.INCRUserCOIN(userId, m.GetCoin())
	if err != nil {
		//请求失败
		log.E("增加用户[%v]的coni[%v]余额失败...", userId, m.GetCoin())
		return
	} else {
		//增加用户金币成功之后，返回用户的金币
		ack := new(ddproto.ApplepayAcksRechargecb)
		ack.Coin = proto.Int64(coin)
		ack.RechargeCoin = proto.Int64(m.GetCoin())
		a.WriteMsg(ack)
	}

}

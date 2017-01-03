package handlers

import (
	"github.com/name5566/leaf/gate"
	"casino_common/proto/ddproto"
	"casino_common/common/userService"
	"github.com/golang/protobuf/proto"
	"casino_common/common/log"
	"casino_common/common/service/wxservice"
)

//处理统一下单
func HandlerApplePayRechargeCb(args []interface{}) {
	m := args[0].(*ddproto.ApplepayReqRechargecb)
	a := args[1].(gate.Agent)
	//通过meal知道 充值的数目等
	userId := m.GetHeader().GetUserId()

	//得到套餐
	p := service.GetMealById(m.GetProductId())

	//增加钻石
	d, err := userService.INCRUserDiamond(userId, p.GetDiamond())
	if err != nil {
		//请求失败
		log.E("增加用户[%v]的coni[%v]余额失败...", userId, p.GetDiamond())
		return
	} else {
		//增加用户金币成功之后，返回用户的金币
		ack := new(ddproto.ApplepayAcksRechargecb)
		ack.Diamond = proto.Int64(d)
		ack.RechargeDiamond = proto.Int64(p.GetDiamond())
		a.WriteMsg(ack)
	}

}

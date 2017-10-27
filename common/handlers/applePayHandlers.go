package handlers

import (
	"casino_common/common/log"
	"casino_common/common/service/wxservice"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
)

//处理统一下单
func HandlerApplePayRechargeCb(args []interface{}) {
	m := args[0].(*ddproto.ApplepayReqRechargecb)
	a := args[1].(gate.Agent)
	//通过meal知道 充值的数目等
	userId := m.GetHeader().GetUserId()

	//得到套餐
	p := service.GetMealById(m.GetProductId())

	if p == nil {
		ack := new(ddproto.ApplepayAcksRechargecb)
		*ack.Header.Code = -1
		*ack.Header.Error = "未找到该商品，购买失败！"
		a.WriteMsg(ack)
		return
	}

	//增加钻石
	d, err := userService.INCRUserDiamond(userId, int64(p.Amount), "商城苹果支付-购买钻石")
	if err != nil {
		//请求失败
		log.E("苹果支付失败，为用户%d充值%d个钻石失败，Err:%v", userId, p.Amount, err)
		return
	} else {
		//增加用户金币成功之后，返回用户的金币
		ack := new(ddproto.ApplepayAcksRechargecb)
		ack.Diamond = proto.Int64(d)
		ack.RechargeDiamond = proto.Int64(int64(p.Amount))
		a.WriteMsg(ack)
	}

	log.T("苹果支付成功，为用户%d充值%d个钻石。", userId, p.Amount)
}

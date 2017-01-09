package handlers

import (
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"casino_common/common/service/allowanceService"
)

//在线奖励的处理
func HandlerAllowance(args []interface{}) {
	m := args[0].(*ddproto.AwardReqOnline)
	a := args[1].(gate.Agent)
	//处理奖励
	allowanceService.DoAllowance(m.GetHeader().GetUserId(), a)
}

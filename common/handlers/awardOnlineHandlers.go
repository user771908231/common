package handlers

import (
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"casino_common/common/service/awardService"
)

//在线奖励的处理
func HandlerAwardOnline(args []interface{}) {
	m := args[0].(*ddproto.AwardReqOnline)
	a := args[1].(gate.Agent)
	//处理奖励
	awardService.DoAwardOnline(m.GetHeader().GetUserId(), a)
}

//在线奖励的信息
func HandlerAwardOnlineInfo(args []interface{}) {
	m := args[0].(*ddproto.AwardReqOnlineInfo)
	a := args[1].(gate.Agent)
	//处理奖励
	awardService.GetOnlineInfo(m.GetHeader().GetUserId(), a)
}

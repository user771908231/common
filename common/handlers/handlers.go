package handlers

import (
	"casino_common/proto/funcsInit"
	"github.com/name5566/leaf/gate"
	"casino_common/common/consts"
)

//处理心跳协议
func HandlerHeartBeat(args []interface{}) {
	a := args[1].(gate.Agent)
	//回复信息
	ack := commonNewPorot.NewHeartbeat()
	*ack.Header.Code = consts.ACK_RESULT_SUCC
	a.WriteMsg(ack)
}
package handlers

import (
	"casino_common/proto/funcsInit"
	"github.com/name5566/leaf/gate"
	"casino_common/common/consts"
	"casino_server/common/log"
	"casino_common/common/sys"
	"casino_common/proto/ddproto"
)

//处理心跳协议
func HandlerHeartBeat(args []interface{}) {
	m := args[0].(*ddproto.Heartbeat)
	a := args[1].(gate.Agent)
	//回复信息
	ack := commonNewPorot.NewHeartbeat()
	*ack.Header.Code = consts.ACK_RESULT_SUCC
	if sys.GAMEENV.DEVMODE {
		log.T("回复[%v]心跳", m.GetHeader().GetUserId())
	}
	a.WriteMsg(ack)
}

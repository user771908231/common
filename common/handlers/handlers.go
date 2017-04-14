package handlers

import (
	"github.com/name5566/leaf/gate"
	"casino_common/proto/ddproto"
)

//处理心跳协议
func HandlerHeartBeat(args []interface{}) {
	m := args[0].(*ddproto.Heartbeat)
	a := args[1].(gate.Agent)
	a.WriteMsg(m)
}

package handlers

import (
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"casino_common/common/noticeServer"
)

func HandlerGetCommonAckNotice(args []interface{}) {
	m := args[0].(*ddproto.CommonReqNotice)
	a := args[1].(gate.Agent)
	ack := noticeServer.GetCommonAckNotice(m.GetNoticeType())
	a.WriteMsg(ack)
}


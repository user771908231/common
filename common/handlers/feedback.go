package handlers

import (
	"casino_common/proto/ddproto"
	"casino_server/common/log"
)

func HandlerFeedback(args []interface{}) {
	m := args[0].(*ddproto.CommonReqFeedback)
	log.T("来自玩家的反馈:%v", m)
}

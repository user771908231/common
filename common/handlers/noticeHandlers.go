package handlers

import (
	"casino_common/common/service/noticeServer"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
)

func HandlerGetCommonAckNotice(args []interface{}) {
	m := args[0].(*ddproto.CommonReqNotice)
	a := args[1].(gate.Agent)
	userId := m.GetHeader().GetUserId()
	channelId := m.GetChannelId()
	if channelId == "" {
		//这里判断空的目的是因为：如果客户端没有传递channelId 那么需要从玩家信息获取
		if user := userService.GetUserById(userId); user != nil {
			channelId = user.GetRegChannel()
		}
	}
	ack := noticeServer.GetCommonAckNotice(m.GetNoticeType(), channelId)
	a.WriteMsg(ack)
}

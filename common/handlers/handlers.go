package handlers

import (
	"casino_common/proto/funcsInit"
	"github.com/name5566/leaf/gate"
	"casino_common/common/consts"
	"casino_common/proto/ddproto"
	"casino_common/common/log"
	"casino_common/common/userService"
)

//处理心跳协议
func HandlerHeartBeat(args []interface{}) {
	a := args[1].(gate.Agent)
	//回复信息
	ack := commonNewPorot.NewHeartbeat()
	*ack.Header.Code = consts.ACK_RESULT_SUCC
	a.WriteMsg(ack)
}

func HandlerGame_Login(args []interface{}) {
	m := args[0].(*ddproto.CommonReqGameLogin)
	a := args[1].(gate.Agent)

	log.T("请求handlerGame_Login  m[%v]", m)
	weixin := m.GetWxInfo()

	//不是初次登录
	if weixin == nil {
		//判断uerId
		userId := m.GetHeader().GetUserId()
		user := userService.GetUserById(userId)
		if user == nil {
			//登陆失败
			log.E("没有找到玩家[%v]的信息，login失败", userId)
			ack := commonNewPorot.NewCommonAckGameLogin()
			*ack.Header.Code = consts.ACK_RESULT_ERROR
			a.WriteMsg(ack)
		} else {
			//返回登陆成功的结果
			ack := commonNewPorot.NewCommonAckGameLogin()

			*ack.Header.Code = consts.ACK_RESULT_SUCC
			*ack.UserId = user.GetId()
			*ack.NickName = user.GetNickName()
			*ack.Chip = user.GetDiamond()
			a.WriteMsg(ack)
		}
		return
	}

	//1,首先通过weixinInfo 在数据库中查找 用户是否存在，如果用户存在，则表示，登陆成功
	user := userService.GetUserByOpenId(weixin.GetOpenId())
	if user == nil {
		//表示数据库中不存在次用户，新增加一个人后返回
		if weixin.GetOpenId() == "" || weixin.GetHeadUrl() == "" || weixin.GetNickName() == "" {
			ack := commonNewPorot.NewCommonAckGameLogin()

			*ack.Header.Code = consts.ACK_RESULT_ERROR
			a.WriteMsg(ack)
			return
		}
		//如果数据库中不存在用户，那么重新生成一个user
		user, _ = userService.NewUserAndSave(weixin.GetOpenId(), weixin.GetNickName(), weixin.GetHeadUrl(), weixin.GetSex(), weixin.GetCity())
		if user == nil {
			ack := commonNewPorot.NewCommonAckGameLogin()

			*ack.Header.Code = consts.ACK_RESULT_ERROR
			a.WriteMsg(ack)
			return
		}
	}

	//返回登陆成功的结果
	ack := commonNewPorot.NewCommonAckGameLogin()
	*ack.Header.Code = consts.ACK_RESULT_SUCC
	*ack.UserId = user.GetId()
	*ack.NickName = user.GetNickName()
	*ack.Chip = user.GetDiamond()

	a.WriteMsg(ack)
	return

}
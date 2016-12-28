package handlers

import (
	"casino_common/proto/funcsInit"
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"casino_common/common/log"
	"casino_common/common/userService"
	"casino_common/common/consts"
	"errors"
)

func HandlerGame_Login(args []interface{}) {
	m := args[0].(*ddproto.CommonReqGameLogin)
	a := args[1].(gate.Agent)
	log.T("请求handlerGame_Login  m[%v]", m)

	//调用登录方法
	user, err := DoLogin(m.GetWxInfo(), m.GetUserId())
	if err != nil || user == nil {
		log.E("登录出现出错..")
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
		//*ack.IsSignToday = signService.IsUserSignedToday(user)
		a.WriteMsg(ack)
	}
	return

}

//做登录的操作...
func DoLogin(weixin *ddproto.WeixinInfo, userId uint32) (*ddproto.User, error) {
	var user *ddproto.User
	//不是初次登录,需要用userId 来登录
	if weixin == nil {
		//判断uerId
		user := userService.GetUserById(userId)
		return user, nil
	} else {
		//微信不等于空的情况,表示是新用户
		//1,首先通过weixinInfo 在数据库中查找 用户是否存在，如果用户存在，则表示，登陆成功
		user = userService.GetUserByOpenId(weixin.GetUnionId())
		if user == nil {
			//表示数据库中不存在次用户，新增加一个人后返回
			if weixin.GetOpenId() == "" || weixin.GetHeadUrl() == "" || weixin.GetNickName() == "" {
				return nil, errors.New("新增用户失败...")
			}
			//如果数据库中不存在用户，那么重新生成一个user
			return userService.NewUserAndSave(weixin.GetUnionId(), weixin.GetOpenId(), weixin.GetNickName(), weixin.GetHeadUrl(), weixin.GetSex(), weixin.GetCity())
		}
	}
	//最终表示登录失败...
	return nil, nil
}

package loginService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/userService"
	"errors"
	"casino_common/common/service/signService"
	"casino_common/common/log"
)


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

func DoLoginSuccess(userId uint32) error {
	error := signService.DoSignLottery(userId)
	if error != nil {
		log.T("签到失败, 错误信息[%v]", error)
		return error
	}
	return nil
}

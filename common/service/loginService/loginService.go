package loginService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/userService"
	"casino_common/common/service/signService"
	"casino_common/common/log"
	"github.com/golang/protobuf/proto"
	"casino_common/utils/numUtils"
	"casino_common/utils/rand"
	"strings"
	"casino_common/common/Error"
	"casino_common/common/consts"
	"casino_common/common/model/userDao"
	"regexp"
)

//做登录的操作...
func DoLogin(weixin *ddproto.WeixinInfo, userId uint32) (*ddproto.User, error) {
	var user *ddproto.User
	//不是初次登录,需要用userId 来登录
	if userId > 0 {
		log.T("玩家使用userid【%v】登录..", userId)
		user := userService.GetUserById(userId)
		return user, nil
	} else if weixin != nil {
		log.T("微信注册【%v】登录..", userId)
		user = userService.GetUserByUnionId(weixin.GetUnionId())
		if user == nil {
			return nil, Error.NewError(consts.ACK_RESULT_ERROR, "登录失败")
		}
	}

	//最终表示登录失败...
	return nil, nil
}


//做登录的操作...
func DoLoginViaInput(userName, password string, userId uint32) (*ddproto.User, error) {
	var user *ddproto.User
	////不是初次登录,需要用userId 来登录
	//if userId > 0 {
	//	log.T("玩家使用userid【%v】登录..", userId)
	//	user := userService.GetUserById(userId)
	//	return user, nil
	//}
	log.T("用户名【%v】密码【%v】登录..", userName, password)
	user = userService.GetUserByUserName(userName)
	if user == nil {
		log.W("用户名[%v] 数据库中找不到该用户, 登录失败", userName)
		return nil, Error.NewError(consts.ACK_RESULT_ERROR, "登录失败")
	}

	if user.GetPwd() != password {
		log.W("用户名[%v] 数据库中密码[%v]与请求密码不符, 登录失败", userName, user.GetPwd(), password)
		return nil, Error.NewError(consts.ACK_RESULT_ERROR, "登录失败, 密码错误")
	}
	log.T("用户名【%v】密码【%v】登录成功", userName, password)
	//最终表示登录成功...
	return user, nil
}

func DoLoginSuccess(userId uint32) error {
	error := signService.DoSignLottery(userId)
	if error != nil {
		log.T("签到失败, 错误信息[%v]", error)
		return error
	}
	return nil
}

//游客注册
func TouristReg(channel string, regIp string) *ddproto.CommonAckReg {
	//设置游客昵称昵称
	nick, _ := numUtils.Int2String(rand.Rand(10000, 100000))
	nickName := strings.Join([]string{"游客", nick}, "")

	//游戏需要创建一个nickName    "游客+[5位随机数]"
	user, err := userService.NewUserAndSave("", "", nickName, "", 1, "", channel, regIp)
	if err != nil || user == nil {
		log.E("注册用户的时候失败...")
		return nil
	}

	log.T("游客登录的注册user: %v,nick:%v", user, user.GetNickName())

	//返回结果
	ack := new(ddproto.CommonAckReg)
	ack.Header = &ddproto.ProtoHeader{
		Code:  proto.Int32(consts.ACK_RESULT_SUCC),
		Error: proto.String("注册成功"),
	}
	ack.UserId = proto.Uint32(user.GetId())
	return ack
}

//微信注册
func WxReg(weixin *ddproto.WeixinInfo, channelId string, regIp string) *ddproto.CommonAckReg {
	//检测参数
	if weixin.GetOpenId() == "" || weixin.GetUnionId() == "" || weixin.GetNickName() == "" {
		log.E("玩家注册的时候失败，因为微信的信息[%v]不够...", weixin)
		return nil
	}

	//需要判断是否是转账号

	//微信注册的时候需要先判断是否已经注册过了，如果注册过了直接返回userId ,否则注册
	user := userService.GetUserByUnionId(weixin.GetUnionId())
	if user == nil {
		user, _ = userService.NewUserAndSave(weixin.GetUnionId(), weixin.GetOpenId(), weixin.GetNickName(), weixin.GetHeadUrl(), weixin.GetSex(), weixin.GetCity(), channelId, regIp)
	} else {
		log.W("玩家[%v],unionId[%v]已经注册成功了，不需要重复注册", user.GetId(), user.GetUnionId())
	}

	if user == nil {
		log.E("玩家注册的时候失败，注册的时候失败", weixin)
		return nil
	}

	//返回结果
	ack := new(ddproto.CommonAckReg)
	ack.Header = &ddproto.ProtoHeader{
		Code:  proto.Int32(consts.ACK_RESULT_SUCC),
		Error: proto.String("注册成功"),
	}
	ack.UserId = proto.Uint32(user.GetId())
	return ack

}

//微信注册
func TransWxReg(weixin *ddproto.WeixinInfo, userId uint32) *ddproto.CommonAckReg {
	//检测参数
	if weixin.GetOpenId() == "" || weixin.GetUnionId() == "" || weixin.GetNickName() == "" {
		log.E("玩家注册的时候失败，因为微信的信息[%v]不够...", weixin)
		return nil
	}

	//需要判断是否是转账号
	//微信注册的时候需要先判断是否已经注册过了，如果注册过了直接返回userId ,否则注册
	user := userService.GetUserById(userId)
	if user != nil {
		user.NickName = proto.String(weixin.GetNickName())
		user.HeadUrl = proto.String(weixin.GetHeadUrl())
		user.OpenId = proto.String(weixin.GetOpenId())
		user.UnionId = proto.String(weixin.GetUnionId())
		user.City = proto.String(weixin.GetCity())
		user.Sex = proto.Int32(weixin.GetSex())
		userService.UpdateUser2Mgo(user) //微信注册
	} else {
		log.E("玩家游客[%v]转微信的时候出现错误...没有在mongo 找到。")
	}

	//返回结果
	ack := new(ddproto.CommonAckReg)
	ack.Header = &ddproto.ProtoHeader{
		Code:  proto.Int32(consts.ACK_RESULT_SUCC),
		Error: proto.String("注册成功"),
	}
	ack.UserId = proto.Uint32(user.GetId())
	return ack
}



//用户名密码注册
func InputsReg(channel, regIp, userName, pwd string) *ddproto.CommonAckReg {
	log.T("用户名密码的方式请求注册注册信息:渠道[%v] regIp[%v] username[%v] pwd[%v]", channel, regIp, userName, pwd)

	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	if !rgx.MatchString(userName) {
		log.W("注册用户的时候失败, 用户名[%v]不是手机号...", userName)
		return nil
	}

	oldUser := userService.GetUserByUserName(userName)
	if oldUser != nil {
		log.W("注册用户的时候失败, 用户名[%v]已存在...", userName)
		return nil
	}

	//设置游客昵称昵称
	nick, _ := numUtils.Int2String(rand.Rand(10000, 100000))
	nickName := strings.Join([]string{"游客", nick}, "")

	//游戏需要创建一个nickName    "游客+[5位随机数]"
	user, err := userService.NewUserAndSave("", "", nickName, "", 1, "", channel, regIp)
	if err != nil || user == nil {
		log.E("注册用户的时候失败...")
		return nil
	}
	user.Pwd = proto.String(pwd)
	user.PhoneNumber = proto.String(userName) //用户名即电话号码

	//2保存数据到数据库
	err = userDao.UpdateUser2Mgo(user)
	if err != nil {
		log.E("保存用户到mongo的时候失败 error【%v】", err.Error())
		return nil
	}

	//3,保存到redis
	userService.SaveUser2Redis(user)

	log.T("游客登录的注册user: %v,nick:%v", user, user.GetNickName())

	//返回结果
	ack := new(ddproto.CommonAckReg)
	ack.Header = &ddproto.ProtoHeader{
		Code:  proto.Int32(consts.ACK_RESULT_SUCC),
		Error: proto.String("注册成功"),
	}
	ack.UserId = proto.Uint32(user.GetId())
	return ack
}
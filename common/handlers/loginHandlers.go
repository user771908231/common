package handlers

import (
	"casino_common/proto/funcsInit"
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"casino_common/common/log"
	"casino_common/common/consts"
	"casino_common/common/service/loginService"
	"casino_common/common/sessionService"
	"github.com/golang/protobuf/proto"
)

//注册的接口
func HandlerReg(args []interface{}) {
	m := args[0].(*ddproto.CommonReqReg)
	a := args[1].(gate.Agent)

	userId := m.GetHeader().GetUserId()
	regType := m.GetRegType()
	channel := m.GetChannelId() //渠道id
	var ack *ddproto.CommonAckReg
	if regType == int32(ddproto.CommonEnumReg_RET_TYPE_TOURIST) {
		//游客注册
		ack = loginService.TouristReg(channel)
	} else if regType == int32(ddproto.CommonEnumReg_RET_TYPE_WEIXIN) {
		//微信注册
		if userId > 0 {
			//转账号
			ack = loginService.TransWxReg(m.GetWxInfo(), userId)
		} else {
			ack = loginService.WxReg(m.GetWxInfo(), channel)
		}
	} else {

	}
	if ack == nil {
		ack = new(ddproto.CommonAckReg)
		ack.Header = &ddproto.ProtoHeader{
			Code:  proto.Int32(consts.ACK_RESULT_ERROR),
			Error: proto.String("注册的时候失败"),
		}
	}
	a.WriteMsg(ack)
}

func HandlerGame_Login(args []interface{}) {
	m := args[0].(*ddproto.CommonReqGameLogin)
	a := args[1].(gate.Agent)
	log.T("请求handlerGame_Login[%v]  登录IP[%v]", m, a.RemoteAddr())

	//调用登录方法
	user, err := loginService.DoLogin(m.GetWxInfo(), m.GetUserId())
	if err != nil || user == nil {
		log.E("登录出现出错..")
		ack := commonNewPorot.NewCommonAckGameLogin()
		*ack.Header.Code = consts.ACK_RESULT_ERROR
		a.WriteMsg(ack)
	} else {
		//返回登陆成功的结果
		ack := commonNewPorot.NewCommonAckGameLogin()
		*ack.Header.Code = consts.ACK_RESULT_SUCC //ack成功
		*ack.UserId = user.GetId()                //玩家的id
		*ack.NickName = user.GetNickName()        //玩家的昵称
		*ack.Chip = user.GetDiamond()             //玩家的钻石信息
		*ack.Coin = user.GetCoin()                //玩家的金币信息
		ack.RoomCard = user.RoomCard              //房卡的信息
		ack.NewUserAward = user.NewUserAward      //是否有新手奖励：如果用户领取之后，需要更新redis和mgo
		log.T("玩家%v登录之后返回ack:%v", user.GetId(), ack)
		a.WriteMsg(ack)
		loginService.DoLoginSuccess(user.GetId())
	}
	return

}

//获取游戏状态
func HandlerCommonReqGameState(args []interface{}) {
	m := args[0].(*ddproto.CommonReqGameState)
	a := args[1].(gate.Agent)

	ack := new(ddproto.CommonAckGameState)
	userId := m.GetHeader().GetUserId() //userId
	//reqRoomType := m.GetRoomType() 暂时没有使用

	//1,第一步获取金币场的session
	session := sessionService.GetSessionAuto(userId)
	if session != nil {
		//组装消息
		ack.GameId = proto.Int32(session.GetGameId())
		ack.GameStatus = proto.Int32(session.GetGameStatus())
		ack.RoomPassword = proto.String(session.GetRoomPassword())
		ack.RoomType = proto.Int32(session.GetRoomType())
		ack.RoomLevel = proto.Int32(session.GetRoomId())
	} else {
		ack.GameId = proto.Int32(0)
		ack.GameStatus = proto.Int32(int32(ddproto.COMMON_ENUM_GAMESTATUS_NOGAME))
		ack.RoomPassword = proto.String("")
		ack.RoomType = proto.Int32(0)
		ack.RoomLevel = proto.Int32(0)
	}

	//返回消息
	log.T("回复游戏状态 %v", ack)
	a.WriteMsg(ack)
}

package game

import (
	"github.com/name5566/leaf/gate"
	"github.com/golang/protobuf/proto"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
)

type GameUserApi interface {
	WriteMsg(p proto.Message)
	GetUserId() uint32
}

//通用的游戏玩家
type GameUser struct {
	Agent gate.Agent
	*ddproto.CommonSrvGameUser
}

func NewGameUser() *GameUser {
	ret := new(GameUser)
	ret.CommonSrvGameUser = commonNewPorot.NewCommonSrvGameUser()
	return ret
}

func (u *GameUser) WriteMsg(p proto.Message) {
	//判断用户是否为空
	if u == nil {
		log.T("给玩家发送proto[%v]失败,应为玩家为nil", p)
		return
	}

	//判断状态是否正确
	if u.GetIsBreak() || u.GetIsLeave() {
		return
	}
	//发送信息
	agent := u.Agent
	if agent != nil {
		agent.WriteMsg(p)
	}
}

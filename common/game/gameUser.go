package game

import (
	"github.com/name5566/leaf/gate"
	"github.com/golang/protobuf/proto"
	"casino_common/common/log"
)

//通用的游戏玩家
type GameUser struct {
	agent gate.Agent
}

func (u *GameUser) WriteMsg(p proto.Message) {
	//
	if u == nil {
		log.T("给玩家发送proto[%v]失败,应为玩家为nil", p)
		return
	}

	//发送信息
	agent := u.agent
	if agent != nil {
		agent.WriteMsg(p)
	}
}

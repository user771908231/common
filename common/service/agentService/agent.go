package agentService

import (
	"casino_common/common/log"
	"casino_common/common/userService"
	"casino_common/utils/agentUtils"
	"errors"
	"github.com/name5566/leaf/gate"
)

func init() {
	AgentMap = map[uint32]gate.Agent{}
}

var AgentMap map[uint32]gate.Agent

func GetAgent(userId uint32) (gate.Agent, error) {
	a, ok := AgentMap[userId]
	if ok {
		return a, nil
	}
	return nil, errors.New("not found agent.")
}

func SetAgent(userId uint32, a gate.Agent) {
	ex_a, ok := AgentMap[userId]
	if ok {
		if ex_a == a {
			return
		} else {
			log.T("Bind Agent %d <-> %p  old_agent:%p", userId, a, ex_a)
		}
	} else {
		log.T("Bind Agent %d <-> %p  old_agent:nil", userId, a)
	}
	AgentMap[userId] = a
}

func DelAgent(userId uint32) {
	_, ok := AgentMap[userId]
	if ok {
		delete(AgentMap, userId)
	}
}

//发消息
func WriteMsg(uid uint32, msg interface{}) error {
	if agent, ok := AgentMap[uid]; ok {
		agent.WriteMsg(msg)
		return nil
	} else {
		return errors.New("user agent not found.")
	}
}

//获取一个玩家的ip
func GetUserIp(userId uint32) string {
	if a, err := GetAgent(userId); err == nil && a != nil {
		return agentUtils.GetIP(a)
	}

	return userService.GetUserById(userId).GetLastIp()
}

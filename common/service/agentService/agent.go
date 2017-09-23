package agentService

import (
	"github.com/name5566/leaf/gate"
	"errors"
	//"casino_common/common/log"
)

func init() {
	AgentMap = map[uint32]gate.Agent{}
}

var AgentMap map[uint32]gate.Agent

func GetAgent(userId uint32) (gate.Agent, error) {
	a,ok := AgentMap[userId]
	if ok {
		return a, nil
	}
	return nil, errors.New("not found agent.")
}

func SetAgent(userId uint32, a gate.Agent) {
	AgentMap[userId] = a
	//log.T("set %v", AgentMap)
}

func DelAgent(userId uint32) {
	_,ok := AgentMap[userId]
	if ok {
		delete(AgentMap, userId)
	}
}

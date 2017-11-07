package countService

import (
	"casino_common/common/userService"
	"casino_common/common/model/agentModel"
	"casino_common/utils/numUtils"
	"errors"
)

//代理返利百分比
var AgentRebate int32 = 30

//房费代理结算
func DoCoinFeeAgentBill(user_id uint32, coin_fee int64, game_id int32) (error) {
	user_info := userService.GetUserById(user_id)
	agent_id := numUtils.String2Uint32(user_info.GetInvCode())

	agent_info := agentModel.GetAgentInfoById(agent_id)
	if agent_id<10000 || agent_info == nil {
		return errors.New("agent not found.")
	}


	return nil
}

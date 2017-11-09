package countService

import (
	"casino_common/common/userService"
	"casino_common/common/model/agentModel"
	"casino_common/utils/numUtils"
	"casino_common/common/Error"
)

//代理返利百分比
const AgentRebate float64 = 30

//房费代理结算
func DoCoinFeeAgentBill(user_id uint32, coin_fee int64, game_id int32) {
	go func(user_id uint32, coin_fee int64, game_id int32){
		defer Error.ErrorRecovery("金币场房费代理结算")
		user_info := userService.GetUserById(user_id)
		agent_id := numUtils.String2Uint32(user_info.GetInvCode())

		agent_info := agentModel.GetAgentInfoById(agent_id)
		if agent_id<10000 || agent_info == nil {
			return
		}

		//计算提成（30%）
		rebate_coin := float64(coin_fee) * (AgentRebate/100)
		//转换成RMB(1万抵1元)
		rebate_coin = (rebate_coin / 10000) * 1

		//更新提成
		if err := agent_info.Inc("coinfeerebate", rebate_coin); err != nil {
			return
		}

		//插入记录
		agentModel.CoinFeeRebateLog{
			UserId: user_id,
			UserName: user_info.GetNickName(),
			GameId: game_id,
			CoinFee: coin_fee,
			Rebate: rebate_coin,
			AgentId: agent_id,
		}.Insert()

		return
	}(user_id, coin_fee, game_id)
}


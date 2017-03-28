package task

import (
	"casino_common/common/service/taskService/taskType"
	"casino_common/common/service/countService/gameCounter"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"casino_common/common/service/countService/countModel"
)


//公共统计类型的验证函数
func ValidateCountFun(user_task *taskType.UserTask) {
	counter := gameCounter.GetUserCounterByType(user_task.UserId, user_task.TaskType)
	user_task.SumNo = counter.DayCount
	if user_task.SumNo >= user_task.TaskSum {
		user_task.SumNo = user_task.TaskSum
		user_task.IsDone = true
	}
	user_task.SetUserState(user_task.UserId, user_task.TaskState)
}

//红包任务验证函数
func ValidateBonusFun(user_task *taskType.UserTask)  {
	counter := gameCounter.GetUserCounterByType(user_task.UserId, user_task.TaskType)
	user_task.SumNo = counter.DayCount - user_task.LastCheckSumNo
	if user_task.SumNo >= user_task.TaskSum {
		user_task.SumNo = user_task.TaskSum
		user_task.IsDone = true
	}
	user_task.SetUserState(user_task.UserId, user_task.TaskState)
}

//红包任务奖励
func GetRewardBonusFun(user_task *taskType.UserTask) []*ddproto.HallBagItem {
	tody_coin_fee := countModel.GetUserDayCoinFeeByGameId(user_task.UserId, user_task.GameId).DayCount

	//赢5局比赛所消耗的房费总和
	var last_coin_fee int32 = 0
	if m_last_coin_fee,ok := user_task.Data["last_coin_fee"];ok {
		last_coin_fee = int32(m_last_coin_fee)
	}
	use_coin_fee := float64(tody_coin_fee - last_coin_fee)

	var reward float64 = 0
	conf := GetBonusTaskComputeConfigByGameId(user_task.GameId)
	if conf != nil {
		reward = (use_coin_fee/10000)*conf.CoinPercent*conf.BonusPercent
		if reward < conf.Min {
			reward = conf.Min
		}else if reward > conf.Max {
			reward = conf.Max
		}
	}

	return []*ddproto.HallBagItem{
		&ddproto.HallBagItem{
			Type: ddproto.HallEnumTradeType_TRADE_BONUS.Enum(),
			Amount: proto.Float64(reward),
		},
	}
}

//领取红包后,重置门票统计
func ResetCoinFeeCount(user_task *taskType.UserTask) {
	coinfee_counter := countModel.GetUserDayCoinFeeByGameId(user_task.UserId, user_task.GameId)
	user_task.Data["last_coin_fee"] = float64(coinfee_counter.DayCount)
	user_task.SetUserState(user_task.UserId, user_task.TaskState)
}

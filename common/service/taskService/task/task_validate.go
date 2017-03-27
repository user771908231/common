package task

import (
	"casino_common/common/service/taskService/taskType"
	"casino_common/common/service/countService/gameCounter"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
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
	return []*ddproto.HallBagItem{
		&ddproto.HallBagItem{
			Type: ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
			Amount: proto.Float64(13),
		},
	}
}

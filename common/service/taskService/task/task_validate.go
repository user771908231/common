package task

import (
	"casino_common/common/service/countService"
	"casino_common/common/service/taskService"
)

//所有任务相关的验证函数

//所有游戏局数相关
func ValidateAllGameCount(user_task *taskService.UserTask) {
	counter := countService.GetAllCounter(user_task.UserId)
	user_task.SumNo = counter.All_Game_Count
	user_task.SetUserState(user_task.UserId, user_task.TaskState)
	if user_task.SumNo >= user_task.TaskSum {
		user_task.SumNo = user_task.TaskSum
	}
}

//临时： 统计类型的验证函数
func ValidateCount(user_task *taskService.UserTask) {
	counter := countService.GetUserCounterByType(user_task.UserId, user_task.TaskType)
	user_task.SumNo = counter.AllCount
	if user_task.SumNo >= user_task.TaskSum {
		user_task.SumNo = user_task.TaskSum
	}
}

package task

import (
	"casino_common/common/service/countService"
	"casino_common/common/service/taskService"
)

//所有任务相关的验证函数

//所有游戏局数相关
func ValidateAllGameCount(user_task *taskService.UserTask) bool {
	counter := countService.GetUserCounter(user_task.UserId)
	user_task.SumNo = counter.All_Count
	user_task.SetUserState(user_task.UserId, user_task.TaskState)
	if user_task.SumNo >= user_task.TaskSum {
		return true
	}
	return false
}

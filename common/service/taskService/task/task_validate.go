package task

import (
	"casino_common/common/service/countService"
	"casino_common/common/service/taskService"
)


//公共统计类型的验证函数
func ValidateCount(user_task *taskService.UserTask) {
	counter := countService.GetUserCounterByType(user_task.UserId, user_task.TaskType)
	user_task.SumNo = counter.AllCount
	if user_task.SumNo >= user_task.TaskSum {
		user_task.SumNo = user_task.TaskSum
		user_task.IsDone = true
	}
	user_task.SetUserState(user_task.UserId, user_task.TaskState)
}

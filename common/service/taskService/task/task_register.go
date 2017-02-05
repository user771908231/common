package task

import "casino_common/common/service/taskService"

//注册任务列表
func InitTask() {
	//完成所有游戏总局数相关任务
	taskService.RegistTask(taskService.Task{
		TaskInfo: taskService.GetTaskInfo(1),
		Validate: ValidateCount,
	})
	taskService.RegistTask(taskService.Task{
		TaskInfo: taskService.GetTaskInfo(2),
		Validate: ValidateCount,
	})
	taskService.RegistTask(taskService.Task{
		TaskInfo: taskService.GetTaskInfo(3),
		Validate: ValidateCount,
	})
	taskService.RegistTask(taskService.Task{
		TaskInfo: taskService.GetTaskInfo(4),
		Validate: ValidateCount,
	})
	taskService.RegistTask(taskService.Task{
		TaskInfo: taskService.GetTaskInfo(5),
		Validate: ValidateCount,
	})
	taskService.RegistTask(taskService.Task{
		TaskInfo: taskService.GetTaskInfo(6),
		Validate: ValidateCount,
	})
}

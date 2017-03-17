package task

import "casino_common/common/service/taskService"

//注册任务列表
func InitTask() {
	//先清空配置
	taskService.TaskList = []*taskService.Task{}
	//注册所有的计数类型任务
	all_task := taskService.GetTaskInfoList()
	for _,task := range all_task{
		switch task.TaskId {
		case 101:
			//大厅每日分享任务
			taskService.RegistTask(taskService.Task{
				TaskInfo: task,
				Validate:func(task *taskService.UserTask) {
					task.IsDone = true
					task.SetUserState(task.UserId, task.TaskState)
				},
			})

		case 102:
			//游戏内每日分享任务
			taskService.RegistTask(taskService.Task{
				TaskInfo: task,
				Validate:func(task *taskService.UserTask) {
					task.IsDone = true
					task.SetUserState(task.UserId, task.TaskState)
				},
			})

		default:
			//普通计数类任务
			taskService.RegistTask(taskService.Task{
				TaskInfo: task,
				Validate:ValidateCount,
			})

		}
	}

}

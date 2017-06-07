package taskService

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/service/countService/countType"
	"casino_common/common/service/countService/gameCounter"
	"casino_common/common/service/taskService/taskType"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//任务列表
var TaskList []*taskType.Task

/**
	获取展示给用户的任务列表
	后面两个过滤参数为空则不过滤
 */
func GetUserTaskShowList(userId uint32, fill_cateid int32, fill_type countType.CountType, fill_game ddproto.CommonEnumGame) []*taskType.UserTask {
	list := []*taskType.UserTask{}
	task_map := make(map[countType.CountType][]*taskType.UserTask)
	//按类型将任务保存到map方便处理
	for _, task := range GetUserTaskList(userId) {
		//过滤任务分类
		if fill_cateid != 0 && fill_cateid != task.CateId {
			continue
		}
		//过滤任务类型
		if fill_type != "" && fill_type != task.TaskType {
			continue
		}
		//过滤所属游戏
		if fill_game != ddproto.CommonEnumGame_GID_SRC && fill_game != task.GameId && task.GameId != ddproto.CommonEnumGame_GID_SRC {
			continue
		}
		if _, ok := task_map[task.TaskType]; ok {
			task_map[task.TaskType] = append(task_map[task.TaskType], task)
		} else {
			task_map[task.TaskType] = []*taskType.UserTask{task}
		}
	}
	//选出同类型的前3个有效任务
	for _, type_tasks := range task_map {
		var i int32 = 0
		for _, task := range type_tasks {
			if task.IsCheck == false {
				list = append(list, task)
				i++
			}
			if i >= 3 {
				break
			}
		}
	}
	//排序
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i].Sort < list[j].Sort {
				t := list[j]
				list[j] = list[i]
				list[i] = t
			}
		}
	}
	return list
}

//通过任务来类型筛选任务
func GetTaskListByType(userId uint32, fill_type countType.CountType) []*taskType.UserTask {
	list := []*taskType.UserTask{}
	for _, task := range GetUserTaskList(userId) {
		//过滤任务类型
		if fill_type != "" && fill_type != task.TaskType {
			continue
		}
		list = append(list, task)
	}
	return list
}

//注册任务
func RegistTask(task taskType.Task) {
	//避免从数据库中查出数据出错导致异常。
	if task.TaskInfo == nil {
		return
	}
	TaskList = append(TaskList, &task)
}

//从数据库获取任务信息
func GetTaskInfo(task_id int32) *taskType.TaskInfo {
	task_info := new(taskType.TaskInfo)
	var err error = nil
	db.Query("", func(d *mgo.Database) {
		err = d.C(tableName.DBT_T_TASK_INFO).Find(bson.M{
			"taskid": task_id,
		}).One(task_info)
	})
	if err != nil {
		return nil
	}
	return task_info
}

//从数据库获取所有的任务信息
func GetTaskInfoList() []*taskType.TaskInfo {
	task_info := []*taskType.TaskInfo{}
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_T_TASK_INFO).Find(bson.M{}).All(&task_info)
	})
	return task_info
}

//触发任务
func OnTask(taskType countType.CountType, userId uint32) {
	//触发该类型的任务
	for _, task := range GetTaskListByType(userId, taskType) {
		if task.SumNo != task.TaskSum && task.IsDone {
			task.IsDone = false
			task.SetUserState(userId, task.TaskState)
		}
		if !task.IsDone && task.ValidateFun != nil {
			task.ValidateFun(task)
			if task.IsDone == true {
				//推送任务完成广播
			}
		}
	}
}

//获取用户单个任务及状态
func GetUserTask(userId uint32, taskId int32) *taskType.UserTask {
	for _, task := range TaskList {
		if task.TaskId == taskId {
			user_task := &taskType.UserTask{}
			user_task.Task = task
			user_task.TaskState = task.GetState(userId)
			user_task.UserId = userId
			return user_task
		}
	}
	return nil
}

//用户任务列表
func GetUserTaskList(userId uint32) []*taskType.UserTask {
	list := []*taskType.UserTask{}
	for _, task := range TaskList {
		user_task := &taskType.UserTask{}
		user_task.Task = task
		user_task.TaskState = task.GetState(userId)
		user_task.UserId = userId
		list = append(list, user_task)
	}
	return list
}

//获取活动列表
func GetActiveList() []*ddproto.HallItemEvent {
	list := []*ddproto.HallItemEvent{}
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_T_ACTIVE_LIST).Find(bson.M{}).All(&list)
	})
	return list
}

//领取奖励
func CheckReward(userId uint32, rewards []*ddproto.HallBagItem) (err error, name string) {
	i := 0
	for _, reward := range rewards {
		if i > 0 {
			name += "、"
		}
		switch reward.GetType() {
		case ddproto.HallEnumTradeType_TRADE_COIN:
			//领取金币
			_, err = userService.INCRUserCOIN(userId, int64(reward.GetAmount()))
			name += fmt.Sprintf("%.0f金币", reward.GetAmount())
		case ddproto.HallEnumTradeType_TRADE_DIAMOND:
			//领取钻石
			_, err = userService.INCRUserDiamond(userId, int64(reward.GetAmount()))
			name += fmt.Sprintf("%.0f钻石", reward.GetAmount())
		case ddproto.HallEnumTradeType_PROPS_FANGKA:
			//领取房卡
			_, err = userService.INCRUserRoomcard(userId, int64(reward.GetAmount()), 0, "领取奖励")
			name += fmt.Sprintf("%.0f房卡", reward.GetAmount())
		case ddproto.HallEnumTradeType_TRADE_BONUS:
			//领取红包
			_, err = userService.INCRUserBonus(userId, reward.GetAmount())
			name += fmt.Sprintf("%.2f红包", reward.GetAmount())
		case ddproto.HallEnumTradeType_TRADE_TICKET:
			//领取奖券
			_, err = userService.INCRUserTicket(userId, int32(reward.GetAmount()))
			name += fmt.Sprintf("%.0f奖券", reward.GetAmount())
		default:
			switch {
			case reward.GetType() > 200 && reward.GetType() < 300:
				//领取道具
			case reward.GetType() > 300 && reward.GetType() < 400:
				//领取实物
			}
		}
		if err != nil {
			return
		}
		i++
	}
	return
}

//领取任务奖励
func CheckTaskReward(userId uint32, taskId int32) (error, string, []*ddproto.HallBagItem) {
	user_task := GetUserTask(userId, taskId)
	if user_task == nil {
		return errors.New("系统未找到该条任务，领取奖励失败！"), "", []*ddproto.HallBagItem{}
	}
	//刷新任务状态
	if user_task.ValidateFun != nil {
		user_task.ValidateFun(user_task)
	}
	//判断领取次数
	if user_task.RepeatSum > 0 && user_task.RepeatNo >= user_task.RepeatSum {
		return errors.New("任务奖励领取次数已达上限，无法继续领取！"), "", []*ddproto.HallBagItem{}
	}
	//是否已完成
	if !user_task.IsDone {
		return errors.New("任务未完成，领取奖励失败！"), "", []*ddproto.HallBagItem{}
	}
	//是否已领取
	if user_task.IsCheck {
		return errors.New("领取任务奖励失败！"), "", []*ddproto.HallBagItem{}
	}

	name := ""
	reward := []*ddproto.HallBagItem{}
	if user_task.GetRewardFun == nil {
		reward = user_task.Reward
	} else {
		reward = user_task.GetRewardFun(user_task)
	}
	_, name = CheckReward(userId, reward)

	user_task.IsCheck = true

	//处理可重复领取的任务
	if user_task.RepeatSum == 0 {
		user_task.RepeatSum = 1
	}
	if user_task.RepeatNo < user_task.RepeatSum {
		user_task.RepeatNo += 1
	}
	if user_task.RepeatNo < user_task.RepeatSum {
		user_task.IsDone = false
		user_task.IsCheck = false
		user_task.SumNo = 0
		user_task.StartTime = time.Now().Unix()
	}

	user_task.LastCheckSumNo = gameCounter.GetUserCounterByType(user_task.UserId, user_task.TaskType).DayCount
	//保存任务状态
	user_task.SetUserState(userId, user_task.TaskState)
	//刷新任务状态
	if user_task.ValidateFun != nil {
		user_task.ValidateFun(user_task)
	}

	//触发成功领取任务奖励后的回调函数
	if user_task.AfterRewardFun != nil {
		user_task.AfterRewardFun(user_task)
	}
	return nil, name, reward
}

/**
	获取距离完成条件最近的一个任务状态
	参数：game : ("ddz"|"zjh"|"mj")
	返回值：numer: 剩余局数 taskId: 任务id
	如果任务不存在则返回 （-1，-1）
 */
func GetUserNearBonusTask(userId uint32, game ddproto.CommonEnumGame) *taskType.UserTask {
	taskList := GetUserTaskShowList(userId, 2, "", game)
	num_map := make(map[*taskType.UserTask]int32)
	for _, task := range taskList {
		if !task.IsCheck {
			num_map[task] = task.TaskSum - task.SumNo
		}
	}
	var result *taskType.UserTask = nil
	//选出最接近完成的一个任务
	for task, sum := range num_map {
		if result == nil {
			result = task
			continue
		}
		if num_map[result] > sum {
			result = task
		}
	}
	return result
}

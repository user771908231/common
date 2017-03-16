package taskService

import (
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"casino_common/common/userService"
	"casino_hall/service/pack"
	"log"
	"casino_common/common/service/countService/countType"
	"errors"
	"time"
)

//任务列表
var TaskList []*Task

//任务详情
type TaskInfo struct {
	CateId      int32                  //任务分类 1.普通任务 2.红包任务
	TaskId      int32                  //任务id
	TaskType    countType.CountType    //统计类型
	GameType    string                 //所属的游戏  zjh ddz hall
	Title       string                 //任务名
	Sort        int32                  //排序：数字越小排在越前面
	TaskSum     int32                  //任务需要完成的次数
	RepeatSum   int32                  //可重复领取的次数
	Description string                 //任务详情
	Reward      []*ddproto.HallBagItem //任务的奖励
}

//任务结构
type Task struct {
	*TaskInfo
	Data     interface{}          //自定义数据
	Validate func(*UserTask) //单次任务完成的验证函数
}

/**
	获取展示给用户的任务列表
	后面两个过滤参数为空则不过滤
 */
func GetUserTaskShowList(userId uint32, fill_cateid int32, fill_type countType.CountType, fill_game string) []*UserTask {
	list := []*UserTask{}
	task_map := make(map[countType.CountType][]*UserTask)
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
		if fill_game != "" && fill_game != task.GameType && task.GameType != "all" {
			continue
		}
		if _, ok := task_map[task.TaskType]; ok {
			task_map[task.TaskType] = append(task_map[task.TaskType], task)
		} else {
			task_map[task.TaskType] = []*UserTask{task}
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

//获得用户状态
func (task *Task) GetState(userId uint32) *TaskState {
	userState := new(TaskState)
	db.Query(func(d *mgo.Database) {
		err := d.C(tableName.DBT_T_USER_TASK).Find(bson.M{
			"userid": userId,
			"taskid": task.TaskId,
		}).One(userState)
		if err != nil {
			userState.IsDone = false
			userState.IsCheck = false
			userState.SumNo = 0
			userState.RepeatNo = 0
		}
	})
	return userState
}

//设置状态
func (task *Task) SetUserState(userId uint32, state *TaskState) {
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER_TASK).Upsert(
			bson.M{
				"userid": userId,
				"taskid": task.TaskId,
			},
			bson.M{
				"userid":  userId,
				"taskid":  task.TaskId,
				"sumno":   state.SumNo,
				"isdone":  state.IsDone,
				"ischeck": state.IsCheck,
				"repeatno": state.RepeatNo,
			},
		)
	})
}

//用户单个任务状态
type TaskState struct {
	StartTime int64  //开始任务的时间
	IsDone  bool  //是否已完成任务条件
	IsCheck bool  //是否已领取任务奖励
	SumNo   int32 //当前完成任务的次数
	RepeatNo int32  //已领取的次数
}

//注册任务
func RegistTask(task Task) {
	//避免从数据库中查出数据出错导致异常。
	if task.TaskInfo == nil {
		return
	}
	TaskList = append(TaskList, &task)
}

//从数据库获取任务信息
func GetTaskInfo(task_id int32) *TaskInfo {
	task_info := new(TaskInfo)
	var err error = nil
	db.Query(func(d *mgo.Database) {
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
func GetTaskInfoList() []*TaskInfo {
	task_info := []*TaskInfo{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_TASK_INFO).Find(bson.M{}).All(&task_info)
	})
	return task_info
}

//触发任务
func OnTask(taskType countType.CountType, userId uint32) {
	//普通任务
	for _, task := range GetUserTaskShowList(userId, 1, taskType, "") {
		if task.SumNo != task.TaskSum && task.IsDone {
			task.IsDone = false
			task.SetUserState(userId, task.TaskState)
		}
		if !task.IsDone && task.Validate != nil{
			task.Validate(task)
			if task.IsDone == true {
				//推送任务完成广播
			}
		}
	}
	//红包任务
	for _, task := range GetUserTaskShowList(userId, 2, taskType, "") {
		log.Println(task.UserId, *task.Task, *task.TaskState)
		if task.SumNo != task.TaskSum && task.IsDone {
			task.IsDone = false
			task.SetUserState(userId, task.TaskState)
		}
		if !task.IsDone && task.Validate != nil{
			task.Validate(task)
			if task.IsDone == true {
				//推送任务完成广播
			}
		}
	}
}

//用户任务状态
type UserTask struct {
	UserId uint32 //用户id
	*Task         //任务
	*TaskState    //任务状态
}

//获取用户单个任务及状态
func GetUserTask(userId uint32, taskId int32) *UserTask {
	for _, task := range TaskList {
		if task.TaskId == taskId {
			user_task := &UserTask{}
			user_task.Task = task
			user_task.TaskState = task.GetState(userId)
			user_task.UserId = userId
			return user_task
		}
	}
	return nil
}

//用户任务列表
func GetUserTaskList(userId uint32) []*UserTask {
	list := []*UserTask{}
	for _, task := range TaskList {
		user_task := &UserTask{}
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
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_ACTIVE_LIST).Find(bson.M{}).All(&list)
	})
	return list
}

//领取奖励
func CheckReward(userId uint32, rewards []*ddproto.HallBagItem) {
	for _, reward := range rewards {
		switch reward.GetType() {
		case ddproto.HallEnumTradeType_TRADE_COIN:
			//领取金币
			log.Println("coin")
			userService.INCRUserCOIN(userId, int64(reward.GetAmount()))
		case ddproto.HallEnumTradeType_TRADE_DIAMOND:
			//领取钻石
			userService.INCRUserDiamond(userId, int64(reward.GetAmount()))
		case ddproto.HallEnumTradeType_PROPS_FANGKA:
			//领取房卡
			userService.INCRUserRoomcard(userId, int64(reward.GetAmount()))
		case ddproto.HallEnumTradeType_TRADE_BONUS:
			//领取红包
			userService.INCRUserBonus(userId, reward.GetAmount())
		case ddproto.HallEnumTradeType_TRADE_TICKET:
			//领取奖券
			userService.INCRUserTicket(userId, int32(reward.GetAmount()))
		default:
			switch {
			case reward.GetType() > 200 && reward.GetType() < 300:
				//领取道具
				pack.DoUserPropsAdd(userId, reward.GetType(), int32(reward.GetAmount()))
			case reward.GetType() > 300 && reward.GetType() < 400:
				//领取实物
			}
		}
	}
}

//领取任务奖励
func CheckTaskReward(userId uint32, taskId int32) error {
	user_task := GetUserTask(userId, taskId)
	if user_task == nil {
		return errors.New("系统未找到该条任务，领取奖励失败！")
	}
	if user_task.Validate != nil {
		user_task.Validate(user_task)
	}
	if !user_task.IsDone {
		return errors.New("任务未完成，领取奖励失败！")
	}

	if user_task.IsCheck {
		if user_task.RepeatSum > 1 {
			return errors.New("任务奖励领取次数已达上限，无法继续领取！")
		}else {
			return errors.New("领取任务奖励失败！")
		}
	}

	CheckReward(userId, user_task.Reward)
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

	//保存任务状态
	user_task.SetUserState(userId, user_task.TaskState)
	return nil
}

/**
	获取距离完成条件最近的一个任务状态
	参数：game : ("ddz"|"zjh"|"mj")
	返回值：numer: 剩余局数 taskId: 任务id
	如果任务不存在则返回 （-1，-1）
 */
func GetUserNearBonusTask(userId uint32, game ddproto.CommonEnumGame) *UserTask {
	game_str := "all"
	switch game {
	case ddproto.CommonEnumGame_GID_DDZ:
		game_str = "ddz"
	case ddproto.CommonEnumGame_GID_ZJH:
		game_str = "zjh"
	case ddproto.CommonEnumGame_GID_MAHJONG:
		game_str = "mj"
	}
	taskList := GetUserTaskShowList(userId, 2, "", game_str)
	num_map := make(map[*UserTask]int32)
	for _,task := range taskList {
		if !task.IsCheck {
			num_map[task] = task.TaskSum-task.SumNo
		}
	}
	var result *UserTask = nil
	//选出最接近完成的一个任务
	for task,sum := range num_map{
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

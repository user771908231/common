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
)

//任务类型
type TaskType string

//警告：任务类型常量一经定义后就不应该修改或删除，以防止把数据库关系搞乱
const (
	TYPE_COIN_COUNT     TaskType = "coin_count"     //跟金币数量相关
	TYPE_WIN_COUNT      TaskType = "win_count"      //跟赢比赛的总局数相关
	TYPE_ZJH_GAME_COUNT TaskType = "zjh_game_count" //跟扎金花赢的比赛总局数相关
	TYPE_ALL_GAME_COUNT TaskType = "all_game_count" //所有游戏总局数
)

//任务列表
var TaskList []*Task

//任务详情
type TaskInfo struct {
	TaskId      int32                  //任务id
	TaskType    TaskType               //任务类型
	GameType    string                 //所属的游戏  zjh ddz hall
	Title       string                 //任务名
	Sort        int32                  //排序：数字越小排在越前面
	TaskSum     int32                  //任务需要完成的次数
	Description string                 //任务详情
	Reward      []*ddproto.HallBagItem //任务的奖励
}

//任务结构
type Task struct {
	*TaskInfo
	Data     interface{}          //自定义数据
	Validate func(*UserTask) bool //单次任务完成的验证函数
}

//获取展示给用户的任务列表  后面两个过滤参数为空则不过滤
func GetUserTaskShowList(userId uint32, filt_type TaskType, filt_game string) []*UserTask {
	list := []*UserTask{}
	task_map := make(map[TaskType][]*UserTask)
	//按类型将任务保存到map方便处理
	for _, task := range GetUserTaskList(userId) {
		//过滤任务类型
		if filt_type != "" && filt_type != task.TaskType {
			continue
		}
		//过滤所属游戏
		if filt_game != "" && filt_game != task.GameType {
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

//触发任务
func OnTask(taskType TaskType, userId uint32) {
	for _, task := range GetUserTaskShowList(userId, taskType, "") {
		if task.SumNo != task.TaskSum && task.IsDone {
			task.IsDone = false
			task.SetUserState(userId, task.TaskState)
		}
		if !task.IsDone {
			if task.Validate(GetUserTask(userId, task.TaskId)) {
				task.IsDone = true
				task.SetUserState(userId, task.TaskState)
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
			pack.DoUserPropsAdd(userId, ddproto.HallEnumTradeType_TRADE_BONUS, int32(reward.GetAmount()))
		case ddproto.HallEnumTradeType_TRADE_TICKET:
			//领取奖券
			pack.DoUserPropsAdd(userId, ddproto.HallEnumTradeType_TRADE_TICKET, int32(reward.GetAmount()))
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

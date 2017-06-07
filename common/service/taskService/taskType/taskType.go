package taskType

import (
	"casino_common/common/service/countService/countType"
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
)

//任务详情
type TaskInfo struct {
	Id bson.ObjectId `bson:"_id" binding:"Required"`
	CateId      int32   `binding:"Required"`               //任务分类 1.普通任务 2.红包任务
	TaskId      int32    `binding:"Required"`              //任务id
	TaskType    countType.CountType  `binding:"Required"`  //统计类型
	GameId      ddproto.CommonEnumGame                     //所属的游戏  zjh ddz hall
	Title       string        `binding:"Required"`         //任务名
	Sort        int32                                      //排序：数字越小排在越前面
	TaskSum     int32       `binding:"Required"`           //任务需要完成的次数
	RepeatSum   int32      `binding:"Required"`            //可重复领取的次数
	Description string     `binding:"Required"`            //任务详情
	Reward      []*ddproto.HallBagItem //任务的奖励
}

//任务结构
type Task struct {
	*TaskInfo
	ValidateFun  func(*UserTask)                       //单次任务完成的验证函数
	GetRewardFun func(*UserTask)[]*ddproto.HallBagItem //动态获取任务奖励
	AfterRewardFun func(*UserTask)                     //领取任务奖励后
}

//获得用户状态
func (task *Task) GetState(userId uint32) *TaskState {
	userState := new(TaskState)
	db.Query("", func(d *mgo.Database) {
		err := d.C(tableName.DBT_T_USER_TASK).Find(bson.M{
			"userid": userId,
			"taskid": task.TaskId,
		}).One(userState)
		if err != nil {
			userState.IsDone = false
			userState.IsCheck = false
			userState.SumNo = 0
			userState.RepeatNo = 0
			userState.LastCheckSumNo = 0
			userState.Data = map[string]float64{}
		}
	})
	return userState
}

//设置状态
func (task *Task) SetUserState(userId uint32, state *TaskState) {
	db.Query("", func(d *mgo.Database) {
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
				"lastchecksumno": state.LastCheckSumNo,
				"data": state.Data,
			},
		)
	})
}

//用户任务状态
type UserTask struct {
	UserId uint32 //用户id
	*Task         //任务
	*TaskState    //任务状态
}

//该组合任务是否已全部完成并已领取奖励
func (user_task *UserTask) IsComplete() bool {
	if user_task.RepeatSum > 0 {
		if user_task.RepeatNo >= user_task.RepeatSum {
			return true
		}else {
			if user_task.IsDone && user_task.IsCheck {
				return true
			}else {
				return false
			}
		}
	}else {
		if user_task.IsDone && user_task.IsCheck {
			return true
		}else {
			return false
		}
	}
}

//用户单个任务状态
type TaskState struct {
	StartTime int64  //开始任务的时间
	IsDone  bool  //是否已完成任务条件
	IsCheck bool  //是否已领取任务奖励
	SumNo   int32 //当前完成任务的次数
	LastCheckSumNo    int32  //上次领取奖励时的计数值
	RepeatNo int32  //已领取的次数
	Data map[string]float64  //任务的私有状态数据
}

package taskService

import (
	ldb "casino_common/common/db"
	"testing"
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"github.com/golang/protobuf/proto"
	"casino_common/common/service/countService/countType"
	"casino_common/common/service/taskService/taskType"
)

func init() {
	ldb.InitMongoDb("192.168.199.200", 27017, "test", "id",[]string{})
	RegistTask(taskType.Task{
		TaskInfo: GetTaskInfo(1),
	})
	RegistTask(taskType.Task{
		TaskInfo: GetTaskInfo(2),
	})
	RegistTask(taskType.Task{
		TaskInfo: GetTaskInfo(3),
	})
	RegistTask(taskType.Task{
		TaskInfo: GetTaskInfo(4),
	})
	RegistTask(taskType.Task{
		TaskInfo: GetTaskInfo(5),
	})
	RegistTask(taskType.Task{
		TaskInfo: GetTaskInfo(6),
	})
	RegistTask(taskType.Task{
		TaskInfo: GetTaskInfo(7),
	})
	RegistTask(taskType.Task{
		TaskInfo: GetTaskInfo(8),
	})

}

//测试任务
func TestTask(t *testing.T) {
	//OnTask(countType.ALL_GAME_COUNT,1)
	//t.Log(GetUserTaskList(1))
	t.Log(GetUserTaskShowList(1,2,"", ddproto.CommonEnumGame_GID_ZJH))
}

//测试任务列表
func TestHandlerTaskListReq(t *testing.T) {
	items := []*ddproto.HallItemTask{}
	list := GetUserTaskShowList(1, 1,"", ddproto.CommonEnumGame_GID_SRC)
	for _,task := range list{
		new_task := ddproto.HallItemTask{
			TaskId:&task.TaskId,
			Title:&task.Title,
			Description:&task.Description,
			Addion:task.Reward,
			TaskSum:&task.TaskSum,
			SumNo:&task.SumNo,
			IsDone:&task.IsDone,
			IsCheck:&task.IsCheck,
		}
		items = append(items, &new_task)
	}
	msg := ddproto.HallAckTask{
		TaskList:items,
	}
	t.Log(&msg)
}

//测试：获取用户的任务状态列表
func TestGetUserTaskShowList(t *testing.T) {
	list := GetUserTaskShowList(1, 1, countType.ALL_GAME_COUNT, ddproto.CommonEnumGame_GID_ZJH)
	t.Log(list)
	for _,task := range list{
		t.Log(task.UserId, task.TaskInfo, task.TaskState)
	}
}

//测试：插入用户列表
func TestInsertTaskInfo(t *testing.T) {
	list := []taskType.TaskInfo{
		taskType.TaskInfo{
			TaskId:      1,
			TaskType:    countType.ALL_GAME_COUNT,
			GameId:    ddproto.CommonEnumGame_GID_SRC,
			TaskSum:     5,
			Title:       "完成5局比赛",
			Description: "奖励1000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Float64(1000),
				},
			},
		},taskType.TaskInfo{
			TaskId:      2,
			TaskType:    countType.ALL_GAME_COUNT,
			GameId:    ddproto.CommonEnumGame_GID_SRC,
			TaskSum:     10,
			Title:       "完成10局比赛",
			Description: "奖励1000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Float64(1000),
				},
			},
		},taskType.TaskInfo{
			TaskId:      3,
			TaskType:    countType.ALL_GAME_COUNT,
			TaskSum:     15,
			GameId:    ddproto.CommonEnumGame_GID_SRC,
			Title:       "完成15局比赛",
			Description: "奖励1000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Float64(1000),
				},
			},
		},taskType.TaskInfo{
			TaskId:      4,
			TaskType:    countType.ALL_GAME_COUNT,
			TaskSum:     20,
			GameId:    ddproto.CommonEnumGame_GID_SRC,
			Title:       "完成20局比赛",
			Description: "奖励2000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Float64(2000),
				},
			},
		},taskType.TaskInfo{
			TaskId:      5,
			TaskType:    countType.ALL_GAME_COUNT,
			TaskSum:     25,
			GameId:    ddproto.CommonEnumGame_GID_SRC,
			Title:       "完成25局比赛",
			Description: "奖励2000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Float64(2000),
				},
			},
		},taskType.TaskInfo{
			TaskId:      6,
			TaskType:    countType.ALL_GAME_COUNT,
			TaskSum:     50,
			GameId:    ddproto.CommonEnumGame_GID_SRC,
			Title:       "完成50局比赛",
			Description: "奖励3000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Float64(3000),
				},
			},
		},
	}
	for _,task := range list {
		var err error = nil
		db.Query("", func(d *mgo.Database) {
			err = d.C(tableName.DBT_T_TASK_INFO).Insert(task)
		})
		t.Log(err)
	}
}

//测试：检查还差多少局领取红包 和 领取红包
func TestCheckBonus(t *testing.T) {
	var user_id uint32 = 13525
	game_id := ddproto.CommonEnumGame_GID_ZJH
	//还差多少局领取红包
	res := GetUserNearBonusTask(user_id, game_id)
	if res != nil {
		num := res.TaskSum - res.SumNo
		t.Log("还差",num,"局比赛才能领取红包！")
	}
	t.Log(res.UserId, *res.TaskInfo, *res.TaskState)
	////领取红包
	//err := CheckTaskReward(user_id, res.TaskId)
	//t.Log(err)
	//if err == nil {
	//	t.Log(*res.Reward[0], "领取红包奖励成功！")
	//}
}

func TestTaxkSum(t *testing.T) {
	list_task := GetUserTaskShowList(13525, 1,  "", ddproto.CommonEnumGame_GID_ZJH)
	list_bonus := GetUserTaskShowList(13525, 2,  "", ddproto.CommonEnumGame_GID_ZJH)

	var i,j int32
	for _, task := range list_task {
		if task.CateId == 1 {
			if task.IsDone == true && task.IsCheck == false {
				i++
			}
		}
	}

	for _, task := range list_bonus{
		if task.CateId == 2 {
			if task.IsDone == true && task.IsCheck == false {
				j++
			}
		}
	}

	t.Log(i,j)
}

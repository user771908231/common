package taskService

import (
	ldb "casino_common/common/db"
	"testing"
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"github.com/golang/protobuf/proto"
)

func init() {
	ldb.InitMongoDb("192.168.199.155", 27017, "test", "id",[]string{})
}

//测试任务
func TestTask(t *testing.T) {
	OnTask(TYPE_COIN_COUNT,1)
	t.Log(GetUserTaskShowList(1,"",""))
}

//测试任务列表
func TestHandlerTaskListReq(t *testing.T) {
	items := []*ddproto.HallItemTask{}
	list := GetUserTaskShowList(1,"","")
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
	list := GetUserTaskShowList(1, TYPE_COIN_COUNT, "zjh")
	t.Log(list)
	for _,task := range list{
		t.Log(task.UserId, task.TaskInfo, task.TaskState)
	}
}

//测试：插入用户列表
func TestInsertTaskInfo(t *testing.T) {
	list := []TaskInfo{
		TaskInfo{
			TaskId:      1,
			TaskType:    TYPE_ALL_GAME_COUNT,
			GameType:    "all",
			TaskSum:     5,
			Title:       "完成5局比赛",
			Description: "奖励1000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Int32(1000),
				},
			},
		},TaskInfo{
			TaskId:      2,
			TaskType:    TYPE_ALL_GAME_COUNT,
			GameType:    "all",
			TaskSum:     10,
			Title:       "完成10局比赛",
			Description: "奖励1000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Int32(1000),
				},
			},
		},TaskInfo{
			TaskId:      3,
			TaskType:    TYPE_ALL_GAME_COUNT,
			TaskSum:     15,
			GameType:    "all",
			Title:       "完成15局比赛",
			Description: "奖励1000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Int32(1000),
				},
			},
		},TaskInfo{
			TaskId:      4,
			TaskType:    TYPE_ALL_GAME_COUNT,
			TaskSum:     20,
			GameType:    "all",
			Title:       "完成20局比赛",
			Description: "奖励2000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Int32(2000),
				},
			},
		},TaskInfo{
			TaskId:      5,
			TaskType:    TYPE_ALL_GAME_COUNT,
			TaskSum:     25,
			GameType:    "all",
			Title:       "完成25局比赛",
			Description: "奖励2000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Int32(2000),
				},
			},
		},TaskInfo{
			TaskId:      6,
			TaskType:    TYPE_ALL_GAME_COUNT,
			TaskSum:     50,
			GameType:    "all",
			Title:       "完成50局比赛",
			Description: "奖励3000金币。",
			Reward:[]*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type:ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount:proto.Int32(3000),
				},
			},
		},
	}
	for _,task := range list {
		var err error = nil
		db.Query(func(d *mgo.Database) {
			err = d.C(tableName.DBT_T_TASK_INFO).Insert(task)
		})
		t.Log(err)
	}
}
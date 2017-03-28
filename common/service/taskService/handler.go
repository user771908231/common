package taskService

import (
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"github.com/golang/protobuf/proto"
	"casino_common/proto/funcsInit"
)

//活动列表请求
func HandlerReqEvent(req *ddproto.HallReqEvent, agent gate.Agent) {
	res := &ddproto.HallAckEvent{
		EventList: GetActiveList(),
	}
	agent.WriteMsg(res)
}

//任务列表ACK
func HandlerTaskListReq(req *ddproto.HallReqTask, agent gate.Agent) {
	list := GetUserTaskShowList(req.GetHeader().GetUserId(), 1,  "", ddproto.CommonEnumGame_GID_SRC)
	items := []*ddproto.HallItemTask{}
	for _, task := range list {
		new_task := ddproto.HallItemTask{
			TaskId:     &task.TaskId,
			Title:      &task.Title,
			Description:&task.Description,
			Addion:     task.Reward,
			TaskSum:    &task.TaskSum,
			SumNo:      &task.SumNo,
			IsDone:     &task.IsDone,
			IsCheck:    &task.IsCheck,
		}
		items = append(items, &new_task)
	}
	msg := ddproto.HallAckTask{
		TaskList:items,
	}
	agent.WriteMsg(&msg)
}

//领取任务奖励
func HandlerCheckTaskReq(req *ddproto.HallReqCheckTask, agent gate.Agent) {
	user_id := req.GetHeader().GetUserId()
	task_id := req.GetTaskId()

	msg := &ddproto.HallAckCheckTask{
		Header:commonNewPorot.NewHeader(),
	}
	defer agent.WriteMsg(msg)

	err,name := CheckTaskReward(user_id, task_id)
	if err != nil {
		msg.Header.Code = proto.Int32(-1)
		msg.Header.Error = proto.String(err.Error())
		return
	}
	msg.Header.Code = proto.Int32(1)
	msg.Header.Error = proto.String("恭喜你,成功领取 "+ name +" 奖励！")
}

//未领取的任务数
func HandlerTaskSumReq(req *ddproto.HallReqTaskSum, agent gate.Agent) {
	fill_game := ddproto.CommonEnumGame_GID_SRC
	var task_id int32 = 0

	switch req.GetTaskType() {
	case ddproto.HallEnumTaskType_TYPE_DDZ:
		fill_game = ddproto.CommonEnumGame_GID_DDZ
		task_id = 202
	case ddproto.HallEnumTaskType_TYPE_MJ:
		fill_game = ddproto.CommonEnumGame_GID_MAHJONG
		task_id = 201
	case ddproto.HallEnumTaskType_TYPE_ZJH:
		fill_game = ddproto.CommonEnumGame_GID_ZJH
		task_id = 203
	default:
		fill_game = ddproto.CommonEnumGame_GID_SRC
	}
	list_task := GetUserTaskShowList(req.Header.GetUserId(), 1,  "", fill_game)

	var i,j int32
	for _, task := range list_task {
		if task.CateId == 1 {
			if task.IsDone == true && task.IsCheck == false {
				i++
			}
		}
	}

	bonus_task := GetUserTask(req.Header.GetUserId(), task_id)
	msg := ddproto.HallAckTaskSum{
		TaskSum: &i,
		BonusSum: &j,
	}
	msg.BonusNext = proto.Int32(-1)
	if bonus_task != nil {
		*msg.BonusNext = bonus_task.TaskSum - bonus_task.SumNo
	}
	if *msg.BonusNext == 0 {
		*msg.BonusSum = 1
	}
	agent.WriteMsg(&msg)
}

//领取红包
func HandlerCheckBonusReq(req *ddproto.HallReqCheckBonus, agent gate.Agent) {
	var task_id int32 = 0
	switch req.GetTaskType() {
	case ddproto.HallEnumTaskType_TYPE_DDZ:
		task_id = 202
	case ddproto.HallEnumTaskType_TYPE_MJ:
		task_id = 201
	case ddproto.HallEnumTaskType_TYPE_ZJH:
		task_id = 203
	}
	user_id := req.GetHeader().GetUserId()
	msg := ddproto.HallAckCheckBonus{
		Header: commonNewPorot.NewHeader(),
		GiveBonus: proto.Float64(0),
	}

	if task_id == 0 {
		*msg.Header.Code = -1
		*msg.Header.Error = "TaskType参数错误！"
		agent.WriteMsg(&msg)
		return
	}

	err, name := CheckTaskReward(user_id, task_id)

	if err != nil {
		*msg.Header.Code = -2
		*msg.Header.Error = err.Error()
	}else {
		*msg.Header.Code = 1
		*msg.Header.Error = "恭喜你，成功领取" + name + "."
	}
	agent.WriteMsg(&msg)
}

package taskService

import (
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"github.com/golang/protobuf/proto"
	"casino_common/proto/funcsInit"
	"fmt"
	"casino_common/common/userService"
	"casino_common/common/service/taskService/taskType"
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
	var user_task *taskType.UserTask = nil
	switch req.GetTaskType() {
	case ddproto.HallEnumTaskType_TYPE_DDZ:
		fill_game = ddproto.CommonEnumGame_GID_DDZ
		user_task = GetUserNearBonusTask(req.Header.GetUserId(), fill_game)
	case ddproto.HallEnumTaskType_TYPE_MJ:
		fill_game = ddproto.CommonEnumGame_GID_MAHJONG
		user_task = GetUserNearBonusTask(req.Header.GetUserId(), fill_game)
	case ddproto.HallEnumTaskType_TYPE_ZJH:
		fill_game = ddproto.CommonEnumGame_GID_ZJH
		user_task = GetUserNearBonusTask(req.Header.GetUserId(), fill_game)
	default:
		fill_game = ddproto.CommonEnumGame_GID_SRC
		user_task = GetUserNearBonusTask(req.Header.GetUserId(), fill_game)
	}
	list_task := GetUserTaskShowList(req.Header.GetUserId(), 1,  "", fill_game)
	list_bonus := GetUserTaskShowList(req.Header.GetUserId(), 2,  "", fill_game)

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
	msg := ddproto.HallAckTaskSum{
		TaskSum: &i,
		BonusSum: &j,
	}
	msg.BonusNext = proto.Int32(-1)
	if user_task != nil {
		*msg.BonusNext = user_task.TaskSum - user_task.SumNo
	}
	agent.WriteMsg(&msg)
}

//领取红包
func HandlerCheckBonusReq(req *ddproto.HallReqCheckBonus, agent gate.Agent) {
	fill_game := ddproto.CommonEnumGame_GID_SRC
	switch req.GetTaskType() {
	case ddproto.HallEnumTaskType_TYPE_DDZ:
		fill_game = ddproto.CommonEnumGame_GID_DDZ
	case ddproto.HallEnumTaskType_TYPE_MJ:
		fill_game = ddproto.CommonEnumGame_GID_MAHJONG
	case ddproto.HallEnumTaskType_TYPE_ZJH:
		fill_game = ddproto.CommonEnumGame_GID_MAHJONG
	default:
		fill_game = ddproto.CommonEnumGame_GID_SRC
	}
	user_id := req.GetHeader().GetUserId()
	list := GetUserTaskShowList(user_id, 2,  "", fill_game)
	msg := &ddproto.HallAckCheckBonus{
		Header: commonNewPorot.NewHeader(),
	}
	defer agent.WriteMsg(msg)
	msg.GiveBonus = proto.Float64(0)
	if len(list) >= 1 {
		if list[0].IsDone == true && list[0].IsCheck == false {
			var bonus_count float64 = 0
			if len(list[0].Reward) >= 1 {
				bonus_count = list[0].Reward[0].GetAmount()
			}
			*msg.Header.Code = 1
			msg.GiveBonus = &bonus_count
			list[0].IsCheck = true
			list[0].SetUserState(list[0].UserId, list[0].TaskState)
			//发放红包
			userService.INCRUserBonus(user_id, bonus_count)
			*msg.Header.Error = fmt.Sprintf("成功领取%d个红包！", bonus_count)
			return
		}else {
			*msg.Header.Code = -2
			*msg.Header.Error = fmt.Sprintf("再赢%d局比赛才能领红包哦！", list[0].TaskSum - list[0].SumNo)
			return
		}
	}else {
		*msg.Header.Code = -1
		*msg.Header.Error = "您今天的红包被领完了~"
		return
	}

}

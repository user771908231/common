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
	list := GetUserTaskShowList(req.GetHeader().GetUserId(), 1,  "", "")
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
	task := GetUserTask(user_id, task_id)
	msg := &ddproto.HallAckCheckTask{
		Header:commonNewPorot.NewHeader(),
	}
	defer agent.WriteMsg(msg)
	if task == nil {
		msg.Header.Code = proto.Int32(-1)
		msg.Header.Error = proto.String("该任务不存在！")
		return
	}
	if !task.IsDone {
		msg.Header.Code = proto.Int32(-2)
		msg.Header.Error = proto.String("该任务未完成，领取奖励失败！")
		return
	}
	if task.IsCheck {
		msg.Header.Code = proto.Int32(-3)
		msg.Header.Error = proto.String("该任务已领取，重复领取奖励失败！")
		return
	}
	CheckReward(user_id, task.Reward)
	msg.Header.Code = proto.Int32(1)
	msg.Header.Error = proto.String("领取奖励成功！")
	//更新领取状态
	task.TaskState.IsCheck = true
	task.SetUserState(user_id, task.TaskState)
}

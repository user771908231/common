package game

import (
	"github.com/golang/protobuf/proto"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
)

//通用的游戏desk
type GameDesk struct {
	*ddproto.CommonSrvGameDesk
	AllUsers func() []GameUserApi
}

func NewGameDesk() *GameDesk {
	ret := new(GameDesk)
	ret.CommonSrvGameDesk = commonNewPorot.NewCommonSrvGameDesk()
	return ret
}

//发送信息
func (d *GameDesk) SendMessage(msg *ddproto.CommonReqMessage) {
	//获取到请求的信息
	result := commonNewPorot.NewCommonBcMessage()
	*result.UserId = msg.GetHeader().GetUserId()
	*result.Id = msg.GetId()
	*result.Msg = msg.GetMsg()
	*result.MsgType = msg.GetMsgType()
	//发送消息
	d.BroadCastProtoExclusive(result, msg.GetUserId())
}

//广播
func (d *GameDesk) BroadCastProto(msg proto.Message) {
	log.T("开始发送广播[%v]", msg)
	d.EverUserDo(func(u GameUserApi) error {
		u.WriteMsg(msg)
		return nil
	})
}

//广播，排除某人
func (d *GameDesk) BroadCastProtoExclusive(msg proto.Message, userId uint32) {
	d.EverUserDo(func(u GameUserApi) error {
		if u.GetUserId() != userId {
			log.T("开始给user[%v]发送信息", u.GetUserId())
			u.WriteMsg(msg)
		}
		//不用返回错误信息
		return nil
	})
}

//桌子的每个玩家做什么事情
func (d *GameDesk) EverUserDo(do func(u GameUserApi) error) {
	if d == nil {
		log.E("desk 为空不能调用EverUserDo...")
		return
	}
	for _, u := range d.AllUsers() {
		err := do(u)
		if err != nil {
			log.E("玩家在执行任务的时候出错err[%v]", err)

		}
	}
}

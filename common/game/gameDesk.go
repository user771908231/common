package game

import (
	"github.com/golang/protobuf/proto"
	"casino_server/common/log"
)

//通用的游戏desk
type GameDesk struct {
	GetUsers func() []GameUserApi
}

//广播
func (d *GameDesk) BroadCastProto(msg proto.Message) {
	d.EverUserDo(func(u GameUserApi) {
		u.WriteMsg(msg)
	})
}

//广播，排除某人
func (d *GameDesk) BroadCastProtoExclusive(msg proto.Message, userId uint32) {
	d.EverUserDo(func(u GameUserApi) {
		if u.GetUserId() != userId {
			u.WriteMsg(msg)
		}
	})
}
//桌子的每个玩家做什么事情
func (d *GameDesk) EverUserDo(do func(u GameUserApi) error) {
	for _, u := range d.GetUsers() {
		err := do(u)
		if err != nil {
			log.E("玩家在执行任务的时候出错err[%v]", err)

		}
	}
}

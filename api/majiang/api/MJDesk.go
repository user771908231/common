package api

import (
	"casino_common/api/majiang"
	"casino_common/common/Error"
	"casino_common/common/consts"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/sessionService"
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
	"casino_common/utils/db"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/module"
	"sync"
)

//麻将桌子的定义
type MJDesk interface {
	EnterUser(...interface{}) error                 //玩家进入desk，不定参数
	BeginStart() error                              //开始打牌
	ActOut(...interface{}) error                    //出牌的user和牌型
	ActPeng(...interface{}) error                   //碰
	ActGuo(...interface{}) error                    //过
	ActGang(...interface{}) error                   //杠
	ActBu(...interface{}) error                     //补
	ActChi(...interface{}) error                    //吃
	ActHu(...interface{}) error                     //胡
	ActBaoTing(...interface{}) error                //报听
	ActPiao(...interface{}) error                   //飘
	ActShangGa(...interface{}) error                //上嘎
	ActFly(...interface{}) error                    //飞
	ActTi(...interface{}) error                     //提
	ActPreOpeningPlayOption(...interface{}) error   //开局前的玩法选择
	ActLeave(...interface{}) error                  //离开房间
	ActReady(userId uint32) error                   //准备
	Lottery() error                                 //一局结束
	Dissolve(...interface{}) error                  //解散
	ApplyDissolve(...interface{}) error             //申请解散
	ApplyDissolveBack(...interface{}) error         //申请解散回复
	SendMessage(interface{}) error                  //聊天
	Break(...interface{}) error                     //断线的处理
	GetDeskId() int32                               //得到desk id
	GetRoom() MJRoom                                //得到一个room
	GetPassword() string                            //得到房间号
	GetCfg() interface{}                            //的牌桌子的配置信息
	GetUsers() []MJUser                             //得到所有的玩家
	GetParser() MJParser                            //得到解析器
	BroadCastProto(message proto.Message) error     //发送广播
	DlogDes() string                                //打印日志用到的tag
	GetUserById(userId uint32) MJUser               //得到一个User
	GetDeskSkeleton() interface{}                   //得到骨架desk
	GetAllMingPai(userId uint32) []*majiang.MJPAI   //得到对玩家来说的明牌
	GetCfgStr() string                              //得到配置信息的文字描述
	RmUser(userId uint32)                           //从桌子里删除一个玩家
	KickOutUser(userId, kickOutUserId uint32) error //踢出玩家
	EnterAgentMode(userId uint32) error             //请求进入托管模式
	QuitAgentMode(userId uint32) error              //请求退出托管模式
	GetNextPai(MJUser) *majiang.MJPAI               //获取下一张牌
	WipeDesk() error                                //备份桌子信息 数据恢复
}

type MJDeskCore struct {
	Room     MJRoom
	S        *module.Skeleton
	DeskId   int32
	Password string //房间号
	Parser   MJParser
	*MJDeskCfg
	Users   []MJUser
	TopDesk MJDesk //上层desk
	sync.Mutex
}

func NewMJDeskCore(s *module.Skeleton) *MJDeskCore {
	desk := &MJDeskCore{
		S: s,
	}
	//main key
	desk.DeskId, _ = db.GetNextSeq(tableName.DBT_MJ_DESK)
	return desk
}

//开始打牌
func (d *MJDeskCore) BeginStart() error {
	log.W("MJDeskCore BeginStart Nothing to Do")
	return nil
}

func (d *MJDeskCore) GetNextPai(MJUser) *majiang.MJPAI {
	log.W("MJDeskCore GetNextPai Nothing to Do")
	return nil
}

func (d *MJDeskCore) WipeDesk() error {
	log.W("MJDeskCore WipeDesk Nothing to Do")
	return nil
}

//离开房间
func (d *MJDeskCore) ActLeave(...interface{}) error {
	log.W("MJDeskCore ActLeave Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActChi(...interface{}) error {
	log.W("MJDeskCore ActChi Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActBaoTing(...interface{}) error {
	log.W("MJDeskCore ActBaoTing Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActBu(...interface{}) error {
	log.W("MJDeskCore ActBu Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActOut(...interface{}) error {
	log.W("MJDeskCore ActOut Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActShangGa(...interface{}) error {
	log.W("MJDeskCore ActShangGa Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActPeng(...interface{}) error {
	log.W("MJDeskCore ActPeng Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActReady(userId uint32) error {
	log.W("MJDeskCore ActReady Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActGang(...interface{}) error {
	log.W("MJDeskCore ActGang Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActHu(...interface{}) error {
	log.W("MJDeskCore ActHu Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActGuo(...interface{}) error {
	log.W("MJDeskCore ActGuo Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActPiao(...interface{}) error {
	log.W("MJDeskCore ActPiao Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActFly(...interface{}) error {
	log.W("MJDeskCore ActFly Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActTi(...interface{}) error {
	log.W("MJDeskCore ActTi Nothing to Do")
	return nil
}

func (d *MJDeskCore) ActPreOpeningPlayOption(...interface{}) error {
	log.W("MJDeskCore ActPreOpeningPlayOption Nothing to Do")
	return nil
}

func (d *MJDeskCore) Lottery() error {
	log.W("MJDeskCore Lottery Nothing to Do")
	return nil
}

//断线
func (d *MJDeskCore) Break(...interface{}) error {
	log.W("MJDeskCore Break Nothing to Do")
	return nil
}

func (d *MJDeskCore) LeafS() *module.Skeleton {
	return d.S
}

func (d *MJDeskCore) GetRoom() MJRoom {
	return d.Room
}

func (d *MJDeskCore) SetRoom(r MJRoom) {
	d.Room = r
}

//胡牌的 解析器
func (d *MJDeskCore) GetParser() MJParser {
	return d.Parser
}

func (d *MJDeskCore) SetParser(p MJParser) {
	d.Parser = p
}

func (d *MJDeskCore) GetDeskId() int32 {
	return d.DeskId
}

func (d *MJDeskCore) SetGid(gid int32) {
	d.Gid = gid
}

func (d *MJDeskCore) GetPassword() string {
	return d.Password
}
func (d *MJDeskCore) GetUserById(userId uint32) MJUser {
	for _, u := range d.Users {
		if u != nil && u.GetUserId() == userId {
			return u
		}
	}
	return nil
}

func (d *MJDeskCore) GetIndexByUserId(userId uint32) int {
	for i, u := range d.GetUsers() {
		if u != nil && u.GetUserId() == userId {
			return i
		}
	}
	return -1
}

func (d *MJDeskCore) GetUsers() []MJUser {
	return d.Users
}

func (d *MJDeskCore) DlogDes() string {
	return fmt.Sprintf("[desk%v-r%v] ", d.GetDeskId(), d.CurrRound)
}

func (d *MJDeskCore) BroadCastProto(p proto.Message) error {
	for _, u := range d.Users {
		if u != nil {
			u.WriteMsg2(p)
		}
	}
	return nil
}

//发送广播
func (d *MJDeskCore) BroadCastProtoExclusive(p proto.Message, userId uint32) {
	for _, u := range d.Users {
		if u != nil && u.GetUserId() != userId {
			u.WriteMsg2(p)
		}
	}
}

var ERR_ADDUSERBEAN error = Error.NewError(consts.ACK_RESULT_ERROR, "进入失败，房间已满")
var ERR_ADDUSERBEAN2 error = Error.NewError(consts.ACK_RESULT_ERROR, "进入失败")

func (d *MJDeskCore) AddUserBean(user MJUser) error {

	//判断
	if user == nil {
		return ERR_ADDUSERBEAN2
	}
	//根据房间类型判断人数是否已满
	for i := 0; i < len(d.Users); i++ {
		if d.Users[i] == nil {
			d.Users[i] = user
			return nil
		}
	}
	return ERR_ADDUSERBEAN
}

func (d *MJDeskCore) GetUserIds() (ids []uint32) {
	for _, u := range d.Users {
		if u != nil {
			ids = append(ids, u.GetUserId())
		}
	}
	return
}

func (d *MJDeskCore) SendMessage(m interface{}) error {
	msg := m.(*ddproto.CommonReqMessage)
	result := commonNewPorot.NewCommonBcMessage()
	*result.UserId = msg.GetHeader().GetUserId()
	*result.Id = msg.GetId()
	*result.Msg = msg.GetMsg()
	*result.MsgType = msg.GetMsgType()
	*result.ToUserId = msg.GetToUserId()
	//发送消息
	d.BroadCastProto(result)
	return nil
}

func (d *MJDeskCore) ApplyDissolve(...interface{}) error {
	log.W("MJDeskCore ApplyDissolve Nothing to Do")
	return nil
}

func (d *MJDeskCore) ApplyDissolveBack(...interface{}) error {
	log.W("MJDeskCore ApplyDissolveBack Nothing to Do")
	return nil
}

func (d *MJDeskCore) Dissolve(...interface{}) error {
	log.W("MJDeskCore Dissolve Nothing to Do")
	return nil
}

func (d *MJDeskCore) EnterUser(...interface{}) error {
	log.W("MJDeskCore EnterUser Nothing to Do")
	return nil
}

func (d *MJDeskCore) GetCfg() interface{} {
	log.W("MJDeskCore GetCfg Nothing to Do")
	return nil
}

func (d *MJDeskCore) GetCfgStr() string {
	log.W("MJDeskCore GetCfgStr Nothing to Do")
	return ""
}

func (d *MJDeskCore) GetDeskSkeleton() interface{} {
	log.W("MJDeskCore GetDeskSkeleton Nothing to Do")
	return nil
}

func (d *MJDeskCore) GetAllMingPai(userId uint32) []*majiang.MJPAI {
	log.W("MJDeskCore GetAllMingPai Nothing to Do")
	return nil
}

func (d *MJDeskCore) RmUser(userId uint32) {
	for i, u := range d.Users {
		if u != nil && u.GetUserId() == userId {
			d.Users[i] = nil
			sessionService.DelSessionByKey(
				u.GetUserId(),
				int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND),
				d.Gid,
				d.GetDeskId())
		}
	}
}

var ERR_KICKOUT_ERROR error = Error.NewError(consts.ACK_RESULT_ERROR, "踢出玩家失败")

func (d *MJDeskCore) KickOutUser(userId, kickOutUserId uint32) error {
	log.T("%v玩家[%v]请求踢出玩家[%v]开始处理", d.DlogDes(), userId, kickOutUserId)
	if d.Owner != userId {
		log.W("%v玩家[%v]请求踢出玩家[%v]失败 玩家[%v]不是房主[%v]", d.DlogDes(), userId, kickOutUserId, userId, d.Owner)
		return ERR_KICKOUT_ERROR
	}
	if d.CurrRound != 0 {
		log.W("%v玩家[%v]请求踢出玩家[%v]失败 游戏进行中 局数[%v]", d.DlogDes(), userId, kickOutUserId, d.CurrRound)
		return ERR_KICKOUT_ERROR //踢出房间失败 游戏进行中
	}

	kickOutUser := d.GetUserById(kickOutUserId)
	if kickOutUser == nil {
		log.W("%v玩家[%v]请求踢出玩家[%v]失败 找不到被踢出玩家", d.DlogDes(), userId, kickOutUserId)
		return ERR_KICKOUT_ERROR
	}

	d.RmUser(kickOutUserId)

	bc := commonNewPorot.NewBCKickOut()
	*bc.Type = ddproto.COMMON_ENUM_KICKOUT_K_OWNER
	*bc.UserId = kickOutUserId

	kickOutUser.WriteMsg2(bc)

	d.BroadCastProtoExclusive(bc, kickOutUserId)

	log.T("%v玩家[%v]请求踢出玩家[%v]处理完毕", d.DlogDes(), userId, kickOutUserId)
	return nil
}

func (d *MJDeskCore) EnterAgentMode(userId uint32) error {
	log.W("MJDeskCore EnterAgentMode Nothing to Do")
	return nil
}

func (d *MJDeskCore) QuitAgentMode(userId uint32) error {
	log.W("MJDeskCore QuitAgentMode Nothing to Do")
	return nil
}

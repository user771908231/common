package api

import (
	"casino_common/common/Error"
	"casino_common/common/consts"
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
	"casino_common/utils/db"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/module"
	"sync"
	"casino_common/api/majiang"
)

//麻将桌子的定义
type MJDesk interface {
	EnterUser(...interface{}) error               //玩家进入desk，不定参数
	ActOut(...interface{}) error                  //出牌的user和牌型
	ActPeng(...interface{}) error                 //碰
	ActGuo(...interface{}) error                  //过
	ActGang(...interface{}) error                 //杠
	ActBu(...interface{}) error                   //补
	ActChi(...interface{}) error                  //吃
	ActHu(...interface{}) error                   //胡
	ActBaoTing(...interface{}) error              //报听
	ActPiao(...interface{}) error                 //飘
	ActFly(...interface{}) error                  //飞
	ActTi(...interface{}) error                   //提
	ActLeave(...interface{}) error                //离开房间
	ActReady(userId uint32) error                 //准备
	Dissolve(...interface{}) error                //解散
	ApplyDissolve(...interface{}) error           //申请解散
	ApplyDissolveBack(...interface{}) error       //申请解散回复
	SendMessage(interface{}) error                //聊天
	Break(...interface{}) error                   //断线的处理
	GetDeskId() int32                             //得到desk id
	GetRoom() MJRoom                              //得到一个room
	GetPassword() string                          //得到房间号
	GetCfg() interface{}                          //的牌桌子的配置信息
	GetUsers() []MJUser                           //得到所有的玩家
	GetParser() MJParser                          //得到解析器
	BroadCastProto(message proto.Message) error   //发送广播
	DlogDes() string                              //打印日志用到的tag
	GetUserById(userId uint32) MJUser             //得到一个User
	GetDeskSkeleton() interface{}                 //得到骨架desk
	GetAllMingPai(userId uint32) []*majiang.MJPAI //得到对玩家来说的明牌
	GetCfgStr() string                            //得到配置信息的文字描述
}

type MJDeskCore struct {
	Room     MJRoom
	S        *module.Skeleton
	DeskId   int32
	Password string //房间号
	Parser   MJParser
	Users    []MJUser
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

//离开房间
func (d *MJDeskCore) ActLeave(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActChi(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActBaoTing(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActBu(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActOut(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActPeng(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActReady(userId uint32) error {
	return nil
}

func (d *MJDeskCore) ActGang(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActHu(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActGuo(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActPiao(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActFly(...interface{}) error {
	return nil
}

func (d *MJDeskCore) ActTi(...interface{}) error {
	return nil
}

//断线
func (d *MJDeskCore) Break(...interface{}) error {
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

func (d *MJDeskCore) BroadCastProto(p proto.Message) error {
	for _, u := range d.Users {
		if u != nil {
			u.WriteMsg(p)
		}
	}
	return nil
}

//发送广播
func (d *MJDeskCore) BroadCastProtoExclusive(p proto.Message, userId uint32) {
	for _, u := range d.Users {
		if u.GetUserId() != userId {
			u.WriteMsg(p)
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
	return nil
}

func (d *MJDeskCore) ApplyDissolveBack(...interface{}) error {
	return nil
}

func (d *MJDeskCore) Dissolve(...interface{}) error {
	return nil
}

func (d *MJDeskCore) EnterUser(...interface{}) error {
	return nil
}

func (d *MJDeskCore) GetCfg() interface{} {
	return nil
}

func (d *MJDeskCore) GetCfgStr() string {
	return ""
}

func (d *MJDeskCore) GetDeskSkeleton() interface{} {
	return nil
}

func (d *MJDeskCore) GetAllMingPai(userId uint32) []*majiang.MJPAI {
	return nil
}

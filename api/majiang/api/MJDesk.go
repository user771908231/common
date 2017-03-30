package api

import (
	"github.com/name5566/leaf/module"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"sync"
	"casino_common/common/Error"
	"casino_common/common/consts"
	"github.com/golang/protobuf/proto"
)

//麻将桌子的定义
type MJDesk interface {
	EnterUser(...interface{}) error            //玩家进入desk，不定参数
	ActOut(userId uint32, p interface{}) error //出牌的user和牌型
	ActPeng(...interface{}) error              //碰
	ActGuo(...interface{}) error               //过
	ActGang(...interface{}) error              //杠
	ActBu(...interface{}) error                //补
	ActHu(...interface{}) error                //胡
	ActReady(userId uint32) error              //准备
	GetDeskId() int32                          //得到desk id
	GetPassword() string                       //得到房间号
	GetCfg() interface{}                       //的牌桌子的配置信息
	GetParser() MJParser                       //得到解析器
	DlogDes() string                           //打印日志用到的tag
}

type MJDeskCore struct {
	s        *module.Skeleton
	deskId   int32
	password string //房间号
	parser   MJParser
	users    []MJUser
	sync.Mutex
}

func NewMJDeskCore(s *module.Skeleton) *MJDeskCore {
	desk := &MJDeskCore{
		s: s,
	}
	//main key
	desk.deskId, _ = db.GetNextSeq(tableName.DBT_MJ_DESK)
	return desk
}

//胡牌的 解析器
func (d *MJDeskCore) GetParser() MJParser {
	return d.parser
}

func (d *MJDeskCore) SetParser(p MJParser) {
	d.parser = p
}

func (d *MJDeskCore) GetDeskId() int32 {
	return d.deskId
}

func (d *MJDeskCore) GetPassword() string {
	return d.password
}
func (d *MJDeskCore) GetUserById(userId uint32) MJUser {
	for _, u := range d.users {
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
	return d.users
}

func (d *MJDeskCore) BroadCastProto(p proto.Message) {
	for _, u := range d.users {
		u.WriteMsg(p)
	}
}

//发送广播
func (d *MJDeskCore) BroadCastProtoExclusive(p proto.Message, userId uint32) {
	for _, u := range d.users {
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
	for i := 0; i < len(d.users); i++ {
		if d.users[i] == nil {
			d.users[i] = user
			return nil
		}
	}
	return ERR_ADDUSERBEAN
}

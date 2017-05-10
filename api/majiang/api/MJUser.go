package api

import (
	"github.com/golang/protobuf/proto"
	"casino_common/common/userService"
	"github.com/name5566/leaf/gate"
	"casino_common/utils/agentUtils"
	"sync/atomic"
)

type MJUser interface {
	//所有基本信息的优化
	GetUserId() uint32
	GetCoin() int64
	GetSex() int32
	GetHeadUrl() string
	GetOpenId() string
	GetNickName() string
	GetUserSkeleton() interface{}
	GetGameData() interface{}
	GetDesk() MJDesk //得到桌子

	//基本功能
	SendOverTurn(p proto.Message) error         //发送overTurn
	WriteMsg2(p proto.Message) error             //发送信息
	DoReady() error                             //准备
	DoOut(interface{}) error                    //玩家出牌
	DoPeng(...interface{}) (interface{}, error) //碰牌
	DoChi(...interface{}) (interface{}, error)  //吃牌
	DoBu(...interface{}) (interface{}, error)   //补牌
	DoHu(interface{}) (interface{}, error)      //胡牌
	DoGang(...interface{}) (interface{}, error) //杠牌

	//业务
	AddBillBean(interface{}) error //增加账单
}

//核心User
type MJUserCore struct {
	UserId  uint32
	uRedis  userService.U_REDIS //这里可以使用redis的方法
	Desk    MJDesk              //关联的desk
	coin    int64               //金币
	Break   bool                //是否断线
	Ready   bool                //是否已经准备
	Banker  bool                //是否是庄
	IsOwner bool                //是否是房主
	CanOut  bool                //是否可以打牌
	S       int32               //状态
	Leave   bool                //是否离开
	gate.Agent                  //agent
}

//New一个CoreUser
func NewMJUserCore(userId uint32, a gate.Agent) *MJUserCore {
	return &MJUserCore{
		UserId: userId,
		uRedis: userService.U_REDIS(userId),
		Agent:  a,
	}
}

func (u *MJUserCore) DoChi(...interface{}) (interface{}, error) {
	return nil, nil
}

func (u *MJUserCore) GetUserId() uint32 {
	return u.UserId
}

func (u *MJUserCore) GetSex() int32 {
	return u.uRedis.GetSex()
}

func (u *MJUserCore) GetHeadUrl() string {
	return u.uRedis.GetHeadUrl()
}

func (u *MJUserCore) GetOpenId() string {
	return u.uRedis.GetOpenId()
}

func (u *MJUserCore) GetNickName() string {
	return u.uRedis.GetNickName()
}

func (u *MJUserCore) GetIp() string {
	return agentUtils.GetIP(u.Agent)
}

func (u *MJUserCore) GetCoin() int64 {
	return u.coin
}

func (u *MJUserCore) AddCoin(c int64) int64 {
	return atomic.AddInt64(&u.coin, c)
}

func (u *MJUserCore) GetBreak() bool {
	return u.Break
}

func (u *MJUserCore) GetLeave() bool {
	return u.Leave
}

func (u *MJUserCore) WriteMsg2(p proto.Message) error {
	return nil
}

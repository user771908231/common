package phzData

import (
	"casino_common/common/userService"
	"casino_common/utils/agentUtils"
	"github.com/name5566/leaf/gate"
	"sync/atomic"
)

type User struct {
	gate.Agent                     //leaf agent
	URedis     userService.U_REDIS //redis中的微信相关的方法
	UserId     uint32              //玩家ID
	Score      int64               //玩家游戏中的分数
	GameData   *UserGameData
	GameStatus *UserGameStatus
}

func (u *User) GetUserId() uint32 {
	return u.UserId
}

func (u *User) GetNikcName() string {
	return u.URedis.GetNickName()
}

func (u *User) GetSex() int32 {
	return u.URedis.GetSex()
}

func (u *User) GetHeadUrl() string {
	return u.URedis.GetHeadUrl()
}

func (u *User) GetOpenId() string {
	return u.URedis.GetOpenId()
}

func (u *User) GetIp() string {
	return agentUtils.GetIP(u.Agent)
}

func (u *User) GetScore() int64 {
	return u.Score
}

func (u *User) AddScore(s int64) int64 {
	return atomic.AddInt64(&u.Score, s)
}

func (u *User) GetGameStatus() int32 {
	return u.GameStatus.S
}

func (u *User) SetGameStatus(s int32) {
	u.GameStatus.S = s
}

type UserGameData struct {
	Bills      *Bill       //账单
	HandPokers []*PHZPoker //牌
}

type UserGameStatus struct {
	IsBreak       bool //掉线
	IsLeave       bool //离开
	IsReady       bool
	IsBanker      bool
	IsOwner       bool
	S             int32 //玩家的状态
	ApplyDissolve int32
}

type BillBean struct {
	Round  int32  //第几局
	UserId uint32 //谁的账单
	Banker uint32 //庄家
	Winner uint32 //赢家
	Score  int64  //输赢分数
}

type Bill struct {
	Score     int64       //所有单局账单输赢分数的总和
	BillBeans []*BillBean //存放每一局的账单
}

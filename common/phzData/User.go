package phzData

import (
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/agentUtils"
	"github.com/name5566/leaf/gate"
	ltimer "github.com/name5566/leaf/timer"
	"sync/atomic"
)

type User struct {
	gate.Agent                        //leaf agent
	URedis        userService.U_REDIS //redis中的微信相关的方法
	UserId        uint32              //玩家ID
	GameData      *UserGameData
	GameStatus    *UserGameStatus
	ApplyDisTimer *ltimer.Timer
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
	return u.GameData.Score
}

func (u *User) AddScore(s int64) int64 {
	return atomic.AddInt64(&u.GameData.Score, s)
}

func (u *User) GetGameStatus() int32 {
	return u.GameStatus.S
}

func (u *User) SetGameStatus(s int32) {
	u.GameStatus.S = s
}

type UserGameData struct {
	HandPokers    []*PHZPoker           //手牌
	PengPai       []*PengPai            //碰的牌
	ChiPai        []*ChiPai             //吃的牌
	TiPai         []*TiPai              //提的牌
	HuInfo        *HuInfo               //胡牌信息
	ChouPais      []*PHZPoker           //忍碰、忍吃后的牌为臭牌
	CheckCase     []*UserCheckCase      //玩家的checkcase， 每个玩家同一张牌可以有多个checkcase
	PaoPais       []*PaoPai             //跑牌
	WeiPais       []*WeiPai             //偎的牌
	PrePengResult *PengPai              //预处理请求碰牌的结果
	PreChiResult  *ddproto.PhzAckChiPai //预处理请求吃牌的结果
	Score         int64                 //总得分
	HuXi          int32                 //总的胡息数
	Bill          *Bill                 //账单
	Statistic     *UserStatistic        //统计信息
}

func (d *UserGameData) GetPengPai() []*PengPai {
	return d.PengPai
}

func (d *UserGameData) GetChiPai() []*ChiPai {
	return d.ChiPai
}

func (d *UserGameData) GetHandPokers() []*PHZPoker {
	return d.HandPokers
}

type UserGameStatus struct {
	IsBreak       bool //掉线
	IsLeave       bool //离开
	IsReady       bool
	IsBanker      bool
	IsOwner       bool
	IsPaoPai      bool  //用于纪录玩家是否跑过牌的信号量
	IsTiPai       bool  //用于纪录玩家是否已经提过牌
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

type UserStatisticBean struct {
	Round    int32 //第几局
	CDianPao int32 //点炮的次数 总次数
	CZimo    int32 //自摸的次数 总次数
	CBeiZimo int32 //被自摸的次数
	CHu      int32 //胡牌的次数

	TScore    int64 //输赢的金币数
	CAnGang   int32 //暗杠的次数
	CDianGang int32 //点杠的次数
	CMingGang int32 //点杠的次数
}

type UserStatistic struct {
	TcountDianPao  int32 //点炮的次数 总次数
	TcountHu       int32 //胡的次数 总次数
	TcountZimo     int32 //自摸的次数 总次数
	TcountBeiZimo  int32 //被自摸的次数
	TCoin          int64 //输赢的金币数
	TcountAnGang   int32 //暗杠的次数
	TcountDianGang int32 //点杠的次数
	TcountMingGang int32 //明杠的次数
	TcountZhaMa    int32 //扎码中的码数
	//每一局的次数
	Beans map[int32]*UserStatisticBean //通过map来存储
}

type UserCheckCase struct {
	CheckID     int32       //id checkcase的唯一标识
	UserIdOut   uint32      //出牌的人
	CheckUserId uint32      //正在check的玩家
	CheckStatus int32       //check的状态：checking和checked
	CheckPai    *PHZPoker   //check的牌
	CanPeng     bool        //是否可以碰牌
	PengPais    []*PHZPoker //可以碰的牌
	CanChi      bool        //是否可以吃牌
	CanHu       bool        //是否可以胡牌
	HuXi        int32       //胡息数
}

func NewCheckCase() *UserCheckCase {
	ret := &UserCheckCase{}
	atomic.AddInt32(&ret.CheckID, 1)
	return ret
}

package phzData

import (
	"casino_common/common/userService"
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

func (u *User) GetGameStatus() int32 {
	return u.GameStatus.S
}

func (u *User) SetGameStatus(s int32) {
	u.GameStatus.S = s
}

type UserGameData struct {
	HandPokers   []*PHZPoker    //手牌
	OutCards     []*PHZPoker    //玩家打出去的牌
	PengPai      []*PengPai     //碰的牌
	ChiPai       []*ChiPai      //吃的牌
	TiPai        []*TiPai       //提的牌
	HuInfo       *HuInfo        //胡牌信息
	PengChouPais []*PHZPoker    //忍碰、忍吃后的牌为臭牌
	ChiChouPais  []*PHZPoker    //吃的臭牌
	CheckCase    *UserCheckCase //玩家的checkcase， 每个玩家同一张牌可以有多个checkcase
	PaoPais      []*PaoPai      //跑牌
	WeiPais      []*WeiPai      //偎的牌
	Score        int64          //总的得分
	PaoScore     int64          //放炮分,根据这个分数判断是否可以结算
	RoundHuXi    int32          //单局胡息数，用于计算是否能胡牌
	GameHuXi     int32          //总的胡息，够200或者400时结束游戏
	Bill         *Bill          //账单
	Statistic    *UserStatistic //统计信息
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
	IsBreak        bool //掉线
	IsLeave        bool //离开
	IsReady        bool
	IsBanker       bool
	IsOwner        bool
	IsTiPai        bool  //用于纪录玩家是否已经提过牌
	IsReconnect    bool  //是否是断线重连
	IsShouldChuPai bool  //过胡后是否需要继续出牌
	S              int32 //玩家的状态
	ApplyDissolve  int32
}

type BillBean struct {
	Round         int32      //第几局
	UserId        uint32     //谁的账单
	Banker        uint32     //庄家
	IsHuangZhuang bool       //是否是荒庄
	Winner        uint32     //赢家
	Score         int64      //输赢分数
	PaoScore      int64      //放炮罚的分数，根据这个分数判断是否牌局结束
	HuXi          int32      //单局的胡息数
	HuInfo        *HuInfo    //胡牌信息
	ChiPais       []*ChiPai  //吃了的牌
	PengPais      []*PengPai //碰了的牌
	WeiPais       []*WeiPai  //偎了的牌
	PaoPais       []*PaoPai  //跑了的牌
	TiPais        []*TiPai   //提了的牌
}

type Bill struct {
	Score     int64       //所有单局账单输赢分数的总和
	HuXi      int32       //当前玩家总胡息
	BillBeans []*BillBean //存放每一局的账单
}

type UserStatistic struct {
	TcountHu      int32 //胡的次数 总次数
	TcountZiMo    int32 //自摸的次数 总次数
	TcountDianPao int32 //点炮的次数 总次数
	TcountTiPai   int32 //提牌次数
	TcountPaoPai  int32 //跑牌次数
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
	ChiPais     []*ChiPai   //可以吃的牌
	CanHu       bool        //是否可以胡牌
	PengHuXi    int32       //胡息数
	ChiHuXi     int32       //吃牌的胡息
	HuHuXi      int32       //胡的胡息
}

func NewCheckCase() *UserCheckCase {
	ret := &UserCheckCase{
		CanHu:   false,
		CanChi:  false,
		CanPeng: false,
		ChiPais: nil,
	}
	atomic.AddInt32(&ret.CheckID, 1)
	return ret
}

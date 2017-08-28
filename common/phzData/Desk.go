package phzData

import (
	"github.com/name5566/leaf/module"
	"sync"
	"time"
)

//desk
type Desk struct {
	S          *module.Skeleton //leaf骨架
	DeskId     int32            //唯一ID
	Cfg        *DeskCfg         //配置器
	GameData   *DeskGameData    //数据
	GameStatus *DeskGameStatus  //状态
	Parser     PHZParser        //解析器
	sync.Mutex
}

//deskCfg
type DeskCfg struct {
	Owner            uint32    //房主
	Banker           uint32    //庄家
	NextBanker       uint32    //下一个庄家
	CreateFee        int64     //房费
	GameNumber       int32     //游戏编号
	ActiveUser       uint32    //活动玩家
	BoardsCout       int32     //总场次数
	UserCountLimit   int32     //玩家的人数
	CurrRound        int32     //当前的场次数
	Password         string    //房间号
	BeginTime        time.Time //开始时间
	RoomType         int32     //房间类型
	TotalPokerNum    int32     //扑克牌
	HandPokerInitNum int32     //初始手牌的数量
	ApplyDis         bool      //申请解散状态
	HuXi             int32     //胡息数
	PaiCursor        int32
	CardsNum         int32 //手牌数量
	IsDaiKai         bool  //是否代开
}

type DeskGameData struct {
	AllPokers    []*PHZPoker //所有的牌
	RemainPokers []*PHZPoker //桌面上剩余的牌
	LiangZhang   *PHZPoker   //每轮的亮张:摸的牌
}

type DeskGameStatus struct {
	status int32 //Desk状态
}

func (d *Desk) GetStatus() int32 {
	return d.GameStatus.status
}

func (d *Desk) SetStatus(s int32) {
	d.GameStatus.status = s
}

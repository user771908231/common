package phzData

import (
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"github.com/name5566/leaf/module"
	"sync"
)

//desk
type Desk struct {
	S            *module.Skeleton               //leaf骨架
	DeskId       int32                          //唯一ID
	Cfg          *DeskCfg                       //配置器
	GameData     *DeskGameData                  //数据
	GameStatus   *DeskGameStatus                //状态
	HuUser       uint32                         //纪录每局可以胡牌的玩家
	PlayBackData []*ddproto.PhzPlaybackSnapshot //数据快照
	NeedCheckPao bool                           //是否需要checkpao， 用于过胡后
	sync.Mutex
}

//deskCfg
type DeskCfg struct {
	Owner            uint32 //房主
	Banker           uint32 //庄家
	NextBanker       uint32 //下一个庄家
	CreateFee        int64  //房费
	GameNumber       int32  //游戏编号
	ActiveUser       uint32 //活动玩家
	BoardsCout       int32  //总场次数
	UserCountLimit   int32  //玩家的人数
	CurrRound        int32  //当前的场次数
	Password         string //房间号
	GameStartTime    int64  //游戏开始时间
	GameEndTime      int64  //游戏结束时间
	BeginTime        int64  //单局开始时间
	EndTime          int64  //单局结束时间
	RoomType         int32  //房间类型
	TotalPokerNum    int32  //扑克牌
	HandPokerInitNum int32  //初始手牌的数量
	ApplyDis         bool   //申请解散状态
	HuXi             int32  //胡息数
	PaiCursor        int32
	CardsNum         int32                                  //手牌数量
	IsDaiKai         bool                                   //是否代开
	ApplyDisUser     uint32                                 //申请解散的玩家
	DiceValue        int32                                  //每局随机的概率值
	RoomCardBillType ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE //扣房卡模式：房主扣卡AA扣卡大赢家扣卡等
	ChanelId         int32                                  //建房的渠道

	IsNeedTianDiHu  bool //是否需要天地胡
	IsNeedHongHeiHu bool //是否需要红黑胡
}

type DeskGameData struct {
	AllPokers      []*PHZPoker //所有的牌
	RemainPokers   []*PHZPoker //桌面上剩余的牌
	LiangZhang     *OutCard    //每轮的亮张:摸的牌
	FirstTimeMoPai bool        //是否是每局的第一次摸牌
	NoWinnerTimes  int32       //荒庄次数
	FirstOutCard   *PHZPoker   //庄家打出的第一张牌
}

type OutCard struct {
	Pai      *PHZPoker //桌面上的亮张
	UserId   uint32    //这张牌是谁摸的，或者谁打出来的
	IsOutPai bool      //是否是玩家打出来的牌
}

type DeskGameStatus struct {
	Status int32 //Desk状态
}

func NewDesk(cfg *DeskCfg, s *module.Skeleton) *Desk {
	desk := &Desk{
		S:          s,
		Cfg:        cfg,
		GameData:   &DeskGameData{},
		GameStatus: &DeskGameStatus{},
	}
	desk.DeskId, _ = db.GetNextSeq(tableName.DBT_MJ_DESK)
	return desk
}

func (d *Desk) GetStatus() int32 {
	return d.GameStatus.Status
}

func (d *Desk) SetStatus(s int32) {
	d.GameStatus.Status = s
}

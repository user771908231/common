package sichuan

import "casino_common/api/majiang"

//四川麻将的游戏数据
type BillBean struct {
	UserId, OutUserId uint32
	Type              int32
	Des               string
	Amount            int64
	Pai               *majiang.MJPAI
}

type Bill struct {
	BillBeans []*BillBean
}

type GuoHuInfo struct {
	SendUserId uint32
	Fan        int32
	Pai        *majiang.MJPAI
}

//胡牌返回的信息
type HuInfo struct {
	WinUser    uint32 //
	LoseUser   uint32
	IsZimo     bool
	IsMengQing bool
	GetUserId  uint32
	SendUserId uint32
	HuDesc     string
	Fan        int32
	Score      int64
	Pai        *majiang.MJPAI
	HuType     int32
	PaiType    []int32 //牌的类型
}

//玩家的用户数据
type SCMJUserData struct {
	HandPais      []*majiang.MJPAI   //所有的手牌
	MoPai         *majiang.MJPAI     //摸的牌
	PengPais      []*majiang.PengPai //碰的牌
	GangPais      []*majiang.GangPai //杠的牌
	HuInfo        []*HuInfo          //胡牌信息
	ChiPais       []*majiang.ChiPai  //吃的牌
	OutPais       []*majiang.MJPAI   //打出去的牌
	GuoHuInfo     *GuoHuInfo         //过胡的信息
	PreMoGangInfo *majiang.GangPai
}

func (d *SCMJUserData) GetPreMoGangInfo() *majiang.GangPai {
	return d.PreMoGangInfo
}

func (d *SCMJUserData) GetGuoHuInfo() *GuoHuInfo {
	return d.GuoHuInfo
}

func (d *SCMJUserData) GetPengPais() []*majiang.PengPai {
	return d.PengPais
}

func (d *SCMJUserData) GetGangPais() []*majiang.GangPai {
	return d.GangPais
}

func (d *SCMJUserData) GetChiPais() []*majiang.ChiPai {
	return d.ChiPais
}

func (d *SCMJUserData) GetHandPais() []*majiang.MJPAI {
	return d.HandPais
}

func (d *SCMJUserData) GetOutPais() []*majiang.MJPAI {
	return d.OutPais
}

func (d *SCMJUserData) GetHuPais() []*HuInfo {
	return d.HuInfo
}

func (d *SCMJUserData) GetMoPai() *majiang.MJPAI {
	return d.MoPai
}

////四川麻将的状态信息
type SCMJUserStatus struct {
	IsBreak   bool //掉线
	IsLeave   bool //离开
	IsReady   bool
	IsBanker  bool //
	IsQianNiu bool //是否是牵牛的玩家
	IsOwner   bool
	S         int32 //玩家的状态
	CanOut    bool  //是否能打牌
}

type UserStatisticBean struct {
	Round     int32 //第几局
	CDianPao  int32 //点炮的次数 总次数
	CZimo     int32 //自摸的次数 总次数
	CBeiZimo  int32 //被自摸的次数
	CHu       int32 //胡牌的次数
	TCoin     int64 //输赢的金币数
	CAnGang   int32 //暗杠的次数
	CDianGang int32 //点杠的次数
	CMingGang int32 //点杠的次数
}

//玩家统计
type UserStatistic struct {
	TcountDianPao  int32 //点炮的次数 总次数
	TcountHu       int32 //胡的次数 总次数
	TcountZimo     int32 //自摸的次数 总次数
	TcountBeiZimo  int32 //被自摸的次数
	TCoin          int64 //输赢的金币数
	TcountAnGang   int32 //暗杠的次数
	TcountDianGang int32 //点杠的次数
	TcountMingGang int32 //明杠的次数
	//每一局的次数
	Beans map[int32]*UserStatisticBean //通过map来存储
}

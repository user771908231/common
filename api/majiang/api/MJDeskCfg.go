package api

import "casino_common/common/userService"

type MJDeskCfg struct {
	Gid             int32               //游戏类型
	Owner           uint32              //房主
	owner           userService.U_REDIS //房主
	Banker          uint32
	NextBanker      uint32 //下一个庄
	ActiveUser      uint32
	GameNumber      int32  //游戏编号
	BeginTime       string //开始时间
	BoardsCout      int32  //总场次数
	UserCountLimit  int32  //玩家的人数
	CurrRound       int32  //当前的场次数 每次初始化一局就加一
	CreateFee       int64  //消耗的房费
	MJPaiCursor     int32
	PassWord        string
	TotalMjPaiCount int32 //麻将牌的总张数
	BaseValue       int64
	MjRoomType      int32 //麻将玩法类型
	CardsNum        int32 //牌的张数

	CapMax int64 //番数封顶

	IsDaiKai bool //是否是代开房间

	//
	ActUser  uint32
	ActType  int32
	ApplyDis bool //申请状态
}

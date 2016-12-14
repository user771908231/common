// Code generated by protoc-gen-go.
// source: ddz_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// 打出去的牌
type DdzSrvOutPokerPais struct {
	KeyValue         *int32               `protobuf:"varint,1,opt,name=keyValue" json:"keyValue,omitempty"`
	PokerPais        []*CommonSrvPokerPai `protobuf:"bytes,2,rep,name=pokerPais" json:"pokerPais,omitempty"`
	Type             *int32               `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	IsBomb           *bool                `protobuf:"varint,4,opt,name=isBomb" json:"isBomb,omitempty"`
	CountDuizi       *int32               `protobuf:"varint,5,opt,name=countDuizi" json:"countDuizi,omitempty"`
	CountSanzhang    *int32               `protobuf:"varint,6,opt,name=countSanzhang" json:"countSanzhang,omitempty"`
	CountSizhang     *int32               `protobuf:"varint,7,opt,name=countSizhang" json:"countSizhang,omitempty"`
	CountYizhang     *int32               `protobuf:"varint,8,opt,name=countYizhang" json:"countYizhang,omitempty"`
	UserId           *uint32              `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzSrvOutPokerPais) Reset()                    { *m = DdzSrvOutPokerPais{} }
func (m *DdzSrvOutPokerPais) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvOutPokerPais) ProtoMessage()               {}
func (*DdzSrvOutPokerPais) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *DdzSrvOutPokerPais) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetPokerPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.PokerPais
	}
	return nil
}

func (m *DdzSrvOutPokerPais) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetIsBomb() bool {
	if m != nil && m.IsBomb != nil {
		return *m.IsBomb
	}
	return false
}

func (m *DdzSrvOutPokerPais) GetCountDuizi() int32 {
	if m != nil && m.CountDuizi != nil {
		return *m.CountDuizi
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountSanzhang() int32 {
	if m != nil && m.CountSanzhang != nil {
		return *m.CountSanzhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountSizhang() int32 {
	if m != nil && m.CountSizhang != nil {
		return *m.CountSizhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetCountYizhang() int32 {
	if m != nil && m.CountYizhang != nil {
		return *m.CountYizhang
	}
	return 0
}

func (m *DdzSrvOutPokerPais) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 对 desk的统计
type DdzSrvDeskTongJi struct {
	Bombs            []*DdzSrvOutPokerPais `protobuf:"bytes,1,rep,name=bombs" json:"bombs,omitempty"`
	CountQiangDiZhu  *int32                `protobuf:"varint,2,opt,name=countQiangDiZhu" json:"countQiangDiZhu,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvDeskTongJi) Reset()                    { *m = DdzSrvDeskTongJi{} }
func (m *DdzSrvDeskTongJi) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDeskTongJi) ProtoMessage()               {}
func (*DdzSrvDeskTongJi) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *DdzSrvDeskTongJi) GetBombs() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.Bombs
	}
	return nil
}

func (m *DdzSrvDeskTongJi) GetCountQiangDiZhu() int32 {
	if m != nil && m.CountQiangDiZhu != nil {
		return *m.CountQiangDiZhu
	}
	return 0
}

// desk
type DdzSrvDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	Key              *string              `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	UserCountLimit   *int32               `protobuf:"varint,3,opt,name=userCountLimit" json:"userCountLimit,omitempty"`
	AllPokerPai      []*CommonSrvPokerPai `protobuf:"bytes,4,rep,name=allPokerPai" json:"allPokerPai,omitempty"`
	DiPokerPai       []*CommonSrvPokerPai `protobuf:"bytes,5,rep,name=diPokerPai" json:"diPokerPai,omitempty"`
	OutPai           *DdzSrvOutPokerPais  `protobuf:"bytes,6,opt,name=outPai" json:"outPai,omitempty"`
	Owner            *uint32              `protobuf:"varint,7,opt,name=owner" json:"owner,omitempty"`
	DizhuPaiUser     *uint32              `protobuf:"varint,8,opt,name=dizhuPaiUser" json:"dizhuPaiUser,omitempty"`
	DiZhuUserId      *uint32              `protobuf:"varint,9,opt,name=diZhuUserId" json:"diZhuUserId,omitempty"`
	ActiveUserId     *uint32              `protobuf:"varint,10,opt,name=activeUserId" json:"activeUserId,omitempty"`
	Tongji           *DdzSrvDeskTongJi    `protobuf:"bytes,11,opt,name=tongji" json:"tongji,omitempty"`
	BaseValue        *int64               `protobuf:"varint,12,opt,name=baseValue" json:"baseValue,omitempty"`
	QingDizhuValue   *int64               `protobuf:"varint,13,opt,name=qingDizhuValue" json:"qingDizhuValue,omitempty"`
	WinValue         *int64               `protobuf:"varint,14,opt,name=winValue" json:"winValue,omitempty"`
	DdzType          *int32               `protobuf:"varint,15,opt,name=ddzType" json:"ddzType,omitempty"`
	RoomType         *int32               `protobuf:"varint,16,opt,name=RoomType" json:"RoomType,omitempty"`
	BoardsCount      *int32               `protobuf:"varint,17,opt,name=BoardsCount" json:"BoardsCount,omitempty"`
	CapMax           *int64               `protobuf:"varint,18,opt,name=CapMax" json:"CapMax,omitempty"`
	IsJiaoFen        *bool                `protobuf:"varint,19,opt,name=IsJiaoFen" json:"IsJiaoFen,omitempty"`
	RoomId           *int32               `protobuf:"varint,20,opt,name=roomId" json:"roomId,omitempty"`
	FootRate         *int32               `protobuf:"varint,21,opt,name=footRate" json:"footRate,omitempty"`
	PlayRate         *int32               `protobuf:"varint,22,opt,name=playRate" json:"playRate,omitempty"`
	GameStatus       *int32               `protobuf:"varint,23,opt,name=GameStatus" json:"GameStatus,omitempty"`
	InitRoomCoin     *int64               `protobuf:"varint,24,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32               `protobuf:"varint,25,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32               `protobuf:"varint,26,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	PlayerNum        *int32               `protobuf:"varint,27,opt,name=playerNum" json:"playerNum,omitempty"`
	TimeOutSecond    *int32               `protobuf:"varint,28,opt,name=timeOutSecond" json:"timeOutSecond,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzSrvDesk) Reset()                    { *m = DdzSrvDesk{} }
func (m *DdzSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvDesk) ProtoMessage()               {}
func (*DdzSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *DdzSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzSrvDesk) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *DdzSrvDesk) GetUserCountLimit() int32 {
	if m != nil && m.UserCountLimit != nil {
		return *m.UserCountLimit
	}
	return 0
}

func (m *DdzSrvDesk) GetAllPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.AllPokerPai
	}
	return nil
}

func (m *DdzSrvDesk) GetDiPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.DiPokerPai
	}
	return nil
}

func (m *DdzSrvDesk) GetOutPai() *DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *DdzSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *DdzSrvDesk) GetDizhuPaiUser() uint32 {
	if m != nil && m.DizhuPaiUser != nil {
		return *m.DizhuPaiUser
	}
	return 0
}

func (m *DdzSrvDesk) GetDiZhuUserId() uint32 {
	if m != nil && m.DiZhuUserId != nil {
		return *m.DiZhuUserId
	}
	return 0
}

func (m *DdzSrvDesk) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *DdzSrvDesk) GetTongji() *DdzSrvDeskTongJi {
	if m != nil {
		return m.Tongji
	}
	return nil
}

func (m *DdzSrvDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *DdzSrvDesk) GetQingDizhuValue() int64 {
	if m != nil && m.QingDizhuValue != nil {
		return *m.QingDizhuValue
	}
	return 0
}

func (m *DdzSrvDesk) GetWinValue() int64 {
	if m != nil && m.WinValue != nil {
		return *m.WinValue
	}
	return 0
}

func (m *DdzSrvDesk) GetDdzType() int32 {
	if m != nil && m.DdzType != nil {
		return *m.DdzType
	}
	return 0
}

func (m *DdzSrvDesk) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *DdzSrvDesk) GetBoardsCount() int32 {
	if m != nil && m.BoardsCount != nil {
		return *m.BoardsCount
	}
	return 0
}

func (m *DdzSrvDesk) GetCapMax() int64 {
	if m != nil && m.CapMax != nil {
		return *m.CapMax
	}
	return 0
}

func (m *DdzSrvDesk) GetIsJiaoFen() bool {
	if m != nil && m.IsJiaoFen != nil {
		return *m.IsJiaoFen
	}
	return false
}

func (m *DdzSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *DdzSrvDesk) GetFootRate() int32 {
	if m != nil && m.FootRate != nil {
		return *m.FootRate
	}
	return 0
}

func (m *DdzSrvDesk) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvDesk) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *DdzSrvDesk) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *DdzSrvDesk) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *DdzSrvDesk) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *DdzSrvDesk) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *DdzSrvDesk) GetTimeOutSecond() int32 {
	if m != nil && m.TimeOutSecond != nil {
		return *m.TimeOutSecond
	}
	return 0
}

// 游戏数据
type DdzSrvGameData struct {
	HandPokers       []*CommonSrvPokerPai  `protobuf:"bytes,1,rep,name=handPokers" json:"handPokers,omitempty"`
	OutPaiList       []*DdzSrvOutPokerPais `protobuf:"bytes,2,rep,name=outPaiList" json:"outPaiList,omitempty"`
	OutPai           *DdzSrvOutPokerPais   `protobuf:"bytes,3,opt,name=outPai" json:"outPai,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvGameData) Reset()                    { *m = DdzSrvGameData{} }
func (m *DdzSrvGameData) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvGameData) ProtoMessage()               {}
func (*DdzSrvGameData) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *DdzSrvGameData) GetHandPokers() []*CommonSrvPokerPai {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *DdzSrvGameData) GetOutPaiList() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPaiList
	}
	return nil
}

func (m *DdzSrvGameData) GetOutPai() *DdzSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

type DdzSrvBillBean struct {
	// 斗地主的账单
	Coin             *int64  `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	LoseUser         *uint32 `protobuf:"varint,2,opt,name=loseUser" json:"loseUser,omitempty"`
	WinUser          *uint32 `protobuf:"varint,3,opt,name=winUser" json:"winUser,omitempty"`
	Desc             *string `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DdzSrvBillBean) Reset()                    { *m = DdzSrvBillBean{} }
func (m *DdzSrvBillBean) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBillBean) ProtoMessage()               {}
func (*DdzSrvBillBean) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *DdzSrvBillBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvBillBean) GetLoseUser() uint32 {
	if m != nil && m.LoseUser != nil {
		return *m.LoseUser
	}
	return 0
}

func (m *DdzSrvBillBean) GetWinUser() uint32 {
	if m != nil && m.WinUser != nil {
		return *m.WinUser
	}
	return 0
}

func (m *DdzSrvBillBean) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

type DdzSrvBill struct {
	// 斗地主的账单
	WinCoin          *int64            `protobuf:"varint,1,opt,name=winCoin" json:"winCoin,omitempty"`
	BillBean         []*DdzSrvBillBean `protobuf:"bytes,2,rep,name=billBean" json:"billBean,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *DdzSrvBill) Reset()                    { *m = DdzSrvBill{} }
func (m *DdzSrvBill) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBill) ProtoMessage()               {}
func (*DdzSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *DdzSrvBill) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzSrvBill) GetBillBean() []*DdzSrvBillBean {
	if m != nil {
		return m.BillBean
	}
	return nil
}

type DdzSrvUserStatisticsRound struct {
	// 玩家每一局的统计信息
	Round            *int32                `protobuf:"varint,1,opt,name=round" json:"round,omitempty"`
	WinCoin          *int64                `protobuf:"varint,2,opt,name=winCoin" json:"winCoin,omitempty"`
	CountBomb        *int32                `protobuf:"varint,3,opt,name=countBomb" json:"countBomb,omitempty"`
	BomBean          []*DdzSrvOutPokerPais `protobuf:"bytes,4,rep,name=bomBean" json:"bomBean,omitempty"`
	PlayRate         *int32                `protobuf:"varint,5,opt,name=playRate" json:"playRate,omitempty"`
	IsSpring         *bool                 `protobuf:"varint,6,opt,name=isSpring" json:"isSpring,omitempty"`
	IsDizhu          *bool                 `protobuf:"varint,7,opt,name=isDizhu" json:"isDizhu,omitempty"`
	IsWin            *bool                 `protobuf:"varint,8,opt,name=isWin" json:"isWin,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUserStatisticsRound) Reset()                    { *m = DdzSrvUserStatisticsRound{} }
func (m *DdzSrvUserStatisticsRound) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatisticsRound) ProtoMessage()               {}
func (*DdzSrvUserStatisticsRound) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *DdzSrvUserStatisticsRound) GetRound() int32 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetBomBean() []*DdzSrvOutPokerPais {
	if m != nil {
		return m.BomBean
	}
	return nil
}

func (m *DdzSrvUserStatisticsRound) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUserStatisticsRound) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

func (m *DdzSrvUserStatisticsRound) GetIsDizhu() bool {
	if m != nil && m.IsDizhu != nil {
		return *m.IsDizhu
	}
	return false
}

func (m *DdzSrvUserStatisticsRound) GetIsWin() bool {
	if m != nil && m.IsWin != nil {
		return *m.IsWin
	}
	return false
}

type DdzSrvUserStatistics struct {
	// 玩家统计信息
	RoundBean        []*DdzSrvUserStatisticsRound `protobuf:"bytes,1,rep,name=roundBean" json:"roundBean,omitempty"`
	CountWin         *int32                       `protobuf:"varint,2,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32                       `protobuf:"varint,3,opt,name=countLose" json:"countLose,omitempty"`
	CountSpring      *int32                       `protobuf:"varint,4,opt,name=countSpring" json:"countSpring,omitempty"`
	CountDizhu       *int32                       `protobuf:"varint,5,opt,name=countDizhu" json:"countDizhu,omitempty"`
	CountBomb        *int32                       `protobuf:"varint,6,opt,name=countBomb" json:"countBomb,omitempty"`
	MaxWinCoin       *int64                       `protobuf:"varint,7,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *DdzSrvUserStatistics) Reset()                    { *m = DdzSrvUserStatistics{} }
func (m *DdzSrvUserStatistics) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUserStatistics) ProtoMessage()               {}
func (*DdzSrvUserStatistics) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *DdzSrvUserStatistics) GetRoundBean() []*DdzSrvUserStatisticsRound {
	if m != nil {
		return m.RoundBean
	}
	return nil
}

func (m *DdzSrvUserStatistics) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountSpring() int32 {
	if m != nil && m.CountSpring != nil {
		return *m.CountSpring
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountDizhu() int32 {
	if m != nil && m.CountDizhu != nil {
		return *m.CountDizhu
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzSrvUserStatistics) GetMaxWinCoin() int64 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

// user
type DdzSrvUser struct {
	UserId           *uint32               `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	GameData         *DdzSrvGameData       `protobuf:"bytes,2,opt,name=gameData" json:"gameData,omitempty"`
	Status           *int32                `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	IsBreak          *bool                 `protobuf:"varint,4,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool                 `protobuf:"varint,5,opt,name=isLeave" json:"isLeave,omitempty"`
	QiangDiZhuStatus *int32                `protobuf:"varint,7,opt,name=qiangDiZhuStatus" json:"qiangDiZhuStatus,omitempty"`
	JiabeiStatus     *int32                `protobuf:"varint,8,opt,name=jiabeiStatus" json:"jiabeiStatus,omitempty"`
	Bill             *DdzSrvBill           `protobuf:"bytes,9,opt,name=bill" json:"bill,omitempty"`
	DeskId           *int32                `protobuf:"varint,10,opt,name=deskId" json:"deskId,omitempty"`
	RoomId           *int32                `protobuf:"varint,11,opt,name=roomId" json:"roomId,omitempty"`
	Coin             *int64                `protobuf:"varint,12,opt,name=coin" json:"coin,omitempty"`
	Statistics       *DdzSrvUserStatistics `protobuf:"bytes,13,opt,name=statistics" json:"statistics,omitempty"`
	LaDaoStatus      *int32                `protobuf:"varint,14,opt,name=laDaoStatus" json:"laDaoStatus,omitempty"`
	PlayRate         *int32                `protobuf:"varint,15,opt,name=playRate" json:"playRate,omitempty"`
	JiaoScore        *int32                `protobuf:"varint,16,opt,name=jiaoScore" json:"jiaoScore,omitempty"`
	TimeOutCount     *int32                `protobuf:"varint,17,opt,name=timeOutCount" json:"timeOutCount,omitempty"`
	IsAgent          *bool                 `protobuf:"varint,18,opt,name=isAgent" json:"isAgent,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzSrvUser) Reset()                    { *m = DdzSrvUser{} }
func (m *DdzSrvUser) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvUser) ProtoMessage()               {}
func (*DdzSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *DdzSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzSrvUser) GetGameData() *DdzSrvGameData {
	if m != nil {
		return m.GameData
	}
	return nil
}

func (m *DdzSrvUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *DdzSrvUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *DdzSrvUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *DdzSrvUser) GetQiangDiZhuStatus() int32 {
	if m != nil && m.QiangDiZhuStatus != nil {
		return *m.QiangDiZhuStatus
	}
	return 0
}

func (m *DdzSrvUser) GetJiabeiStatus() int32 {
	if m != nil && m.JiabeiStatus != nil {
		return *m.JiabeiStatus
	}
	return 0
}

func (m *DdzSrvUser) GetBill() *DdzSrvBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *DdzSrvUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzSrvUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *DdzSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzSrvUser) GetStatistics() *DdzSrvUserStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *DdzSrvUser) GetLaDaoStatus() int32 {
	if m != nil && m.LaDaoStatus != nil {
		return *m.LaDaoStatus
	}
	return 0
}

func (m *DdzSrvUser) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *DdzSrvUser) GetJiaoScore() int32 {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return 0
}

func (m *DdzSrvUser) GetTimeOutCount() int32 {
	if m != nil && m.TimeOutCount != nil {
		return *m.TimeOutCount
	}
	return 0
}

func (m *DdzSrvUser) GetIsAgent() bool {
	if m != nil && m.IsAgent != nil {
		return *m.IsAgent
	}
	return false
}

// room
type DdzSrvRoom struct {
	RoomId           *int32 `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DdzSrvRoom) Reset()                    { *m = DdzSrvRoom{} }
func (m *DdzSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvRoom) ProtoMessage()               {}
func (*DdzSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *DdzSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

// 备份专用...
type DdzSrvBak struct {
	Desk             *DdzSrvDesk   `protobuf:"bytes,1,opt,name=desk" json:"desk,omitempty"`
	Users            []*DdzSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DdzSrvBak) Reset()                    { *m = DdzSrvBak{} }
func (m *DdzSrvBak) String() string            { return proto.CompactTextString(m) }
func (*DdzSrvBak) ProtoMessage()               {}
func (*DdzSrvBak) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *DdzSrvBak) GetDesk() *DdzSrvDesk {
	if m != nil {
		return m.Desk
	}
	return nil
}

func (m *DdzSrvBak) GetUsers() []*DdzSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*DdzSrvOutPokerPais)(nil), "ddproto.ddz_srv_outPokerPais")
	proto.RegisterType((*DdzSrvDeskTongJi)(nil), "ddproto.ddz_srv_deskTongJi")
	proto.RegisterType((*DdzSrvDesk)(nil), "ddproto.ddz_srv_desk")
	proto.RegisterType((*DdzSrvGameData)(nil), "ddproto.ddz_srv_gameData")
	proto.RegisterType((*DdzSrvBillBean)(nil), "ddproto.ddz_srv_billBean")
	proto.RegisterType((*DdzSrvBill)(nil), "ddproto.ddz_srv_bill")
	proto.RegisterType((*DdzSrvUserStatisticsRound)(nil), "ddproto.ddz_srv_userStatisticsRound")
	proto.RegisterType((*DdzSrvUserStatistics)(nil), "ddproto.ddz_srv_userStatistics")
	proto.RegisterType((*DdzSrvUser)(nil), "ddproto.ddz_srv_user")
	proto.RegisterType((*DdzSrvRoom)(nil), "ddproto.ddz_srv_room")
	proto.RegisterType((*DdzSrvBak)(nil), "ddproto.ddz_srv_bak")
}

var fileDescriptor8 = []byte{
	// 1218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xeb, 0x6e, 0x23, 0x35,
	0x14, 0x26, 0xb7, 0x26, 0x71, 0x9a, 0x6d, 0x31, 0x6d, 0x99, 0x5e, 0x58, 0xa2, 0xd1, 0x6a, 0x15,
	0x40, 0xea, 0x8f, 0xae, 0x56, 0x48, 0x88, 0x8b, 0x68, 0x2b, 0x50, 0x57, 0x05, 0x8a, 0xbb, 0x4b,
	0x05, 0x7f, 0x22, 0x27, 0x63, 0x52, 0x37, 0x99, 0x71, 0x76, 0x2e, 0xe9, 0x36, 0xcf, 0xc0, 0x6f,
	0x9e, 0x07, 0x09, 0x89, 0xe7, 0xe0, 0x51, 0xd0, 0x39, 0xb6, 0x67, 0x3c, 0x69, 0xa4, 0x96, 0x5f,
	0x33, 0xe7, 0x3b, 0x9f, 0xed, 0x73, 0xb7, 0xc9, 0x66, 0x10, 0x2c, 0x06, 0x89, 0x88, 0xe7, 0x22,
	0x3e, 0x9c, 0xc5, 0x2a, 0x55, 0xb4, 0x19, 0x04, 0xf8, 0xb3, 0xb7, 0x3b, 0x52, 0x61, 0xa8, 0x22,
	0xa3, 0x1d, 0xcc, 0xd4, 0xc4, 0x72, 0xfc, 0xbf, 0xaa, 0x64, 0x0b, 0x17, 0xc6, 0xf3, 0x81, 0xca,
	0xd2, 0x0b, 0x50, 0x5d, 0x70, 0x99, 0xd0, 0x3d, 0xd2, 0x9a, 0x88, 0xbb, 0x5f, 0xf8, 0x34, 0x13,
	0x5e, 0xa5, 0x57, 0xe9, 0x37, 0x58, 0x2e, 0xd3, 0x2f, 0x48, 0x7b, 0x66, 0x89, 0x5e, 0xb5, 0x57,
	0xeb, 0x77, 0x8e, 0x0e, 0x0e, 0xcd, 0x61, 0x87, 0xf6, 0xac, 0x78, 0x3e, 0xb0, 0x24, 0x56, 0xd0,
	0x29, 0x25, 0xf5, 0xf4, 0x6e, 0x26, 0xbc, 0x1a, 0xee, 0x89, 0xff, 0x74, 0x87, 0xac, 0xc9, 0xe4,
	0x58, 0x85, 0x43, 0xaf, 0xde, 0xab, 0xf4, 0x5b, 0xcc, 0x48, 0xf4, 0x29, 0x21, 0x23, 0x95, 0x45,
	0xe9, 0x69, 0x26, 0x17, 0xd2, 0x6b, 0xe0, 0x0a, 0x07, 0xa1, 0xcf, 0x48, 0x17, 0xa5, 0x4b, 0x1e,
	0x2d, 0xae, 0x79, 0x34, 0xf6, 0xd6, 0x90, 0x52, 0x06, 0xa9, 0x4f, 0xd6, 0x35, 0x20, 0x35, 0xa9,
	0x89, 0xa4, 0x12, 0x96, 0x73, 0x7e, 0x35, 0x9c, 0x96, 0xc3, 0x31, 0x18, 0x58, 0x99, 0x25, 0x22,
	0x3e, 0x0b, 0xbc, 0x76, 0xaf, 0xd2, 0xef, 0x32, 0x23, 0xf9, 0x09, 0xa1, 0x36, 0x82, 0x81, 0x48,
	0x26, 0xaf, 0x55, 0x34, 0x7e, 0x25, 0xe9, 0x0b, 0xd2, 0x18, 0xaa, 0x70, 0x98, 0x78, 0x15, 0x8c,
	0xcf, 0x47, 0x79, 0x7c, 0x56, 0x45, 0x9b, 0x69, 0x2e, 0xed, 0x93, 0x0d, 0x3c, 0xf2, 0x67, 0xc9,
	0xa3, 0xf1, 0xa9, 0xfc, 0xed, 0x3a, 0xf3, 0xaa, 0x68, 0xc9, 0x32, 0xec, 0xff, 0xdb, 0x24, 0xeb,
	0xee, 0xa9, 0x60, 0x1d, 0x7c, 0xcf, 0x02, 0x93, 0x2d, 0x23, 0xd1, 0x4d, 0x52, 0x9b, 0x88, 0x3b,
	0xdc, 0xa6, 0xcd, 0xe0, 0x97, 0x3e, 0x27, 0x4f, 0xc0, 0xf2, 0x13, 0xd8, 0xf1, 0x5c, 0x86, 0x32,
	0x35, 0xb9, 0x58, 0x42, 0xe9, 0xd7, 0xa4, 0xc3, 0xa7, 0x53, 0x6b, 0xa3, 0x57, 0x7f, 0x44, 0x9e,
	0xdd, 0x05, 0xf4, 0x4b, 0x42, 0x02, 0x99, 0x2f, 0x6f, 0x3c, 0x62, 0xb9, 0xc3, 0xa7, 0x2f, 0xc9,
	0x1a, 0x44, 0x88, 0x4b, 0x4c, 0xea, 0x83, 0x01, 0x34, 0x64, 0xba, 0x45, 0x1a, 0xea, 0x36, 0x12,
	0x31, 0x66, 0xb9, 0xcb, 0xb4, 0x00, 0xe9, 0x0d, 0xe4, 0xe2, 0x3a, 0xbb, 0xe0, 0xf2, 0x4d, 0x22,
	0x62, 0x4c, 0x6f, 0x97, 0x95, 0x30, 0xda, 0x23, 0x9d, 0x00, 0x42, 0xfb, 0xc6, 0xcd, 0xb1, 0x0b,
	0xc1, 0x2e, 0x7c, 0x94, 0xca, 0xb9, 0x30, 0x14, 0xa2, 0x77, 0x71, 0x31, 0xfa, 0x82, 0xac, 0xa5,
	0x2a, 0x1a, 0xdf, 0x48, 0xaf, 0x83, 0x66, 0xef, 0xdf, 0x33, 0xbb, 0xa8, 0x11, 0x66, 0xa8, 0xf4,
	0x80, 0xb4, 0x87, 0x3c, 0x11, 0xba, 0xd9, 0xd6, 0x7b, 0x95, 0x7e, 0x8d, 0x15, 0x00, 0xe4, 0xeb,
	0xad, 0x84, 0xbc, 0x2f, 0xae, 0x33, 0x4d, 0xe9, 0x22, 0x65, 0x09, 0x85, 0x8e, 0xbd, 0x95, 0x91,
	0x66, 0x3c, 0x41, 0x46, 0x2e, 0x53, 0x8f, 0x34, 0x83, 0x60, 0xf1, 0x1a, 0x1a, 0x6f, 0x03, 0x93,
	0x6d, 0x45, 0x58, 0xc5, 0x94, 0x0a, 0x51, 0xb5, 0xa9, 0xfb, 0xdc, 0xca, 0x10, 0x92, 0x63, 0xc5,
	0xe3, 0x20, 0xc1, 0xaa, 0xf0, 0xde, 0x47, 0xb5, 0x0b, 0x41, 0xd5, 0x9d, 0xf0, 0xd9, 0x0f, 0xfc,
	0x9d, 0x47, 0xf1, 0x44, 0x23, 0x81, 0x47, 0x67, 0xc9, 0x2b, 0xc9, 0xd5, 0x77, 0x22, 0xf2, 0x3e,
	0xc0, 0xa6, 0x2e, 0x00, 0x58, 0x15, 0x2b, 0x15, 0x9e, 0x05, 0xde, 0x96, 0xae, 0x55, 0x2d, 0x81,
	0x2d, 0xbf, 0x2b, 0x95, 0x32, 0x9e, 0x0a, 0x6f, 0x5b, 0xdb, 0x62, 0x65, 0xd0, 0xcd, 0xa6, 0xfc,
	0x0e, 0x75, 0x3b, 0x5a, 0x67, 0x65, 0x98, 0x13, 0xdf, 0xf3, 0x50, 0x5c, 0xa6, 0x3c, 0xcd, 0x12,
	0xef, 0x43, 0x3d, 0x27, 0x0a, 0x04, 0x12, 0x27, 0x23, 0x99, 0x82, 0x5f, 0x27, 0x4a, 0x46, 0x9e,
	0x87, 0xb6, 0x96, 0x30, 0x9c, 0x25, 0x59, 0x1c, 0x5f, 0x4c, 0xf9, 0x9d, 0xf6, 0x76, 0xd7, 0xcc,
	0x12, 0x17, 0x84, 0x5c, 0xa4, 0x2a, 0xe5, 0xd3, 0x82, 0xb6, 0xa7, 0x7b, 0xa7, 0x8c, 0x82, 0xff,
	0x60, 0x9d, 0x88, 0x7f, 0xcc, 0x42, 0x6f, 0x1f, 0x29, 0x05, 0x00, 0x67, 0xa5, 0x32, 0x14, 0x3f,
	0x65, 0xe9, 0xa5, 0x18, 0xa9, 0x28, 0xf0, 0x0e, 0xf4, 0x59, 0x25, 0xd0, 0xff, 0xa7, 0x62, 0x66,
	0x7a, 0x3c, 0x1f, 0x8c, 0x79, 0x28, 0x4e, 0x79, 0xca, 0xa1, 0xa9, 0xae, 0x79, 0x14, 0x60, 0xe1,
	0xdb, 0xd9, 0xf2, 0x40, 0x53, 0x15, 0x7c, 0xfa, 0x15, 0x21, 0xba, 0x4f, 0xce, 0x65, 0x92, 0x9a,
	0xc9, 0xfd, 0x40, 0x63, 0x39, 0x0b, 0x9c, 0x9e, 0xac, 0xfd, 0x8f, 0x9e, 0xf4, 0x67, 0x85, 0x1f,
	0x43, 0x39, 0x9d, 0x1e, 0x0b, 0x1e, 0xc1, 0x35, 0x30, 0x82, 0x54, 0x54, 0x30, 0x15, 0xf8, 0x0f,
	0x29, 0x9e, 0xaa, 0x04, 0x3b, 0x09, 0xe7, 0x55, 0x97, 0xe5, 0x32, 0x14, 0xf0, 0xad, 0x8c, 0x50,
	0x55, 0x43, 0x95, 0x15, 0x61, 0xa7, 0x40, 0x24, 0x23, 0xbc, 0x3a, 0xda, 0x0c, 0xff, 0xfd, 0x41,
	0x31, 0x1c, 0xe1, 0x44, 0xb3, 0xfa, 0xa4, 0x38, 0xd0, 0x8a, 0xf4, 0x25, 0x69, 0x59, 0x9b, 0x4c,
	0x3c, 0x76, 0xef, 0x39, 0x65, 0x09, 0x2c, 0xa7, 0xfa, 0x7f, 0x54, 0xc9, 0xbe, 0x55, 0xc3, 0xd8,
	0x84, 0x42, 0x93, 0x49, 0x2a, 0x47, 0x09, 0x53, 0x59, 0x14, 0xc0, 0x18, 0x8a, 0xe1, 0xc7, 0x0c,
	0x63, 0x2d, 0xb8, 0x66, 0x54, 0xcb, 0x66, 0x1c, 0x90, 0x36, 0x4e, 0x78, 0xbc, 0x04, 0xf5, 0x38,
	0x2e, 0x00, 0xfa, 0x39, 0x69, 0x0e, 0x55, 0x88, 0x36, 0xd6, 0x1f, 0x93, 0x33, 0xcb, 0x2e, 0x35,
	0x4d, 0x63, 0xa9, 0x69, 0xf6, 0x48, 0x4b, 0x26, 0x97, 0xb3, 0x58, 0x9a, 0x7b, 0xb3, 0xc5, 0x72,
	0x19, 0x0c, 0x95, 0x09, 0x8e, 0x16, 0x9c, 0xa3, 0x2d, 0x66, 0x45, 0x70, 0x4c, 0x26, 0x57, 0x32,
	0xc2, 0x11, 0xda, 0x62, 0x5a, 0xf0, 0xff, 0xac, 0x92, 0x9d, 0xd5, 0xe1, 0xa0, 0xc7, 0xa4, 0x8d,
	0xce, 0xa3, 0xf5, 0xba, 0x5e, 0x9f, 0xdd, 0xb3, 0x7e, 0x45, 0x08, 0x59, 0xb1, 0x0c, 0x4c, 0xc5,
	0x60, 0x5c, 0x99, 0xc0, 0x35, 0x58, 0x2e, 0xe7, 0x91, 0x3b, 0x57, 0x89, 0x28, 0x45, 0x0e, 0x00,
	0x98, 0x60, 0xfa, 0x9e, 0xd7, 0x7e, 0xd6, 0xf5, 0x04, 0x73, 0xa0, 0xe2, 0x8d, 0x81, 0xde, 0x96,
	0xde, 0x18, 0xe8, 0x70, 0x29, 0x33, 0x6b, 0xcb, 0x99, 0x79, 0x4a, 0x48, 0xc8, 0xdf, 0x5d, 0x99,
	0xa4, 0x36, 0x31, 0xa9, 0x0e, 0xe2, 0xff, 0x5d, 0x2f, 0x2a, 0x11, 0x9c, 0x74, 0x1e, 0x11, 0x15,
	0xf7, 0x11, 0x01, 0x75, 0x68, 0x7b, 0x1c, 0x5d, 0x5c, 0x55, 0x87, 0x96, 0xc0, 0x72, 0x2a, 0x6c,
	0x97, 0xe8, 0xa9, 0xa7, 0x5d, 0x37, 0x92, 0x4e, 0xe0, 0x71, 0x2c, 0xf8, 0xc4, 0x3c, 0xa9, 0xac,
	0xa8, 0x35, 0xe7, 0x82, 0xcf, 0x75, 0x45, 0xa0, 0x06, 0x45, 0xfa, 0x29, 0xd9, 0x7c, 0x9b, 0x3f,
	0x30, 0xcc, 0x2c, 0xd5, 0x6f, 0xa5, 0x7b, 0x38, 0x4c, 0xd4, 0x1b, 0xc9, 0x87, 0x42, 0x1a, 0x9e,
	0x79, 0x2f, 0xb9, 0x18, 0xfd, 0x84, 0xd4, 0xa1, 0x5f, 0xf0, 0x26, 0xed, 0x1c, 0x6d, 0xaf, 0x6c,
	0x2b, 0x86, 0x14, 0xe7, 0xf1, 0x42, 0x4a, 0x8f, 0x97, 0xe2, 0xa2, 0xe8, 0x94, 0x2e, 0x0a, 0x3b,
	0x3d, 0xd6, 0x9d, 0xe9, 0xf1, 0x0d, 0x21, 0x49, 0x5e, 0x42, 0x78, 0x45, 0x76, 0x8e, 0x3e, 0x7e,
	0xa8, 0xd2, 0x9c, 0x25, 0x50, 0x2b, 0x53, 0x7e, 0xca, 0x95, 0x71, 0xe9, 0x89, 0xae, 0x15, 0x07,
	0x2a, 0xb5, 0xd3, 0xc6, 0x52, 0x3b, 0x1d, 0x90, 0xf6, 0x8d, 0xe4, 0xea, 0x72, 0xa4, 0x62, 0x7b,
	0x91, 0x16, 0x00, 0xc4, 0xcb, 0x0c, 0x77, 0xf7, 0x2a, 0x2d, 0x61, 0x3a, 0x33, 0xdf, 0x8e, 0x45,
	0x94, 0xe2, 0x65, 0x8a, 0x99, 0x41, 0xd1, 0x7f, 0x5e, 0x14, 0x11, 0x04, 0xc0, 0x09, 0x4b, 0xc5,
	0x0d, 0x8b, 0x2f, 0x48, 0x27, 0x0f, 0x2e, 0x9f, 0x40, 0x02, 0x20, 0x8e, 0x48, 0x5a, 0x95, 0x00,
	0x50, 0xe2, 0xc0, 0x9c, 0xd0, 0xcf, 0x48, 0x03, 0x22, 0x63, 0x5f, 0xf3, 0xdb, 0x2b, 0xe3, 0xc6,
	0x34, 0xe7, 0xe2, 0xbd, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x99, 0x44, 0xed, 0x6e, 0x6c, 0x0c,
	0x00, 0x00,
}

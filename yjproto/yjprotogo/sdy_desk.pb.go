// Code generated by protoc-gen-go.
// source: sdy_desk.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of cm_offline from common_client.proto

// Ignoring public import of cm_hearbeat from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of sdy_base_userPaiIds from sdy_base.proto

// Ignoring public import of sdy_base_roomTypeInfo from sdy_base.proto

// Ignoring public import of sdy_base_timerInfo from sdy_base.proto

// Ignoring public import of sdy_base_playerInfo from sdy_base.proto

// Ignoring public import of sdy_base_commonRateInfo from sdy_base.proto

// Ignoring public import of sdy_base_deskInfo from sdy_base.proto

// Ignoring public import of sdy_enum_protoId from sdy_base.proto

// Ignoring public import of sdy_enum_errorCode from sdy_base.proto

// Ignoring public import of sdy_enum_actType from sdy_base.proto

// Ignoring public import of sdy_enum_deskStatus from sdy_base.proto

// Ignoring public import of sdy_enum_userStatus from sdy_base.proto

// Ignoring public import of sdy_enum_enterType from sdy_base.proto

// Ignoring public import of sdy_enum_Option from sdy_base.proto

// Ignoring public import of sdy_enum_coinRoomLevel from sdy_base.proto

// Ignoring public import of sdy_enum_flowers from sdy_base.proto

// 准备游戏
type SdyReqReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqReady) Reset()                    { *m = SdyReqReady{} }
func (m *SdyReqReady) String() string            { return proto.CompactTextString(m) }
func (*SdyReqReady) ProtoMessage()               {}
func (*SdyReqReady) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *SdyReqReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 准备游戏的结果
type SdyAckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckReady) Reset()                    { *m = SdyAckReady{} }
func (m *SdyAckReady) String() string            { return proto.CompactTextString(m) }
func (*SdyAckReady) ProtoMessage()               {}
func (*SdyAckReady) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *SdyAckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *SdyAckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 赢牌信息：谁赢了多少
type SdyBaseWinCoinInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	TotalScore       *int32       `protobuf:"varint,2,opt,name=totalScore" json:"totalScore,omitempty"`
	RoundScore       *int32       `protobuf:"varint,3,opt,name=roundScore" json:"roundScore,omitempty"`
	WinNum           *int32       `protobuf:"varint,4,opt,name=winNum" json:"winNum,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBaseWinCoinInfo) Reset()                    { *m = SdyBaseWinCoinInfo{} }
func (m *SdyBaseWinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBaseWinCoinInfo) ProtoMessage()               {}
func (*SdyBaseWinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *SdyBaseWinCoinInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBaseWinCoinInfo) GetTotalScore() int32 {
	if m != nil && m.TotalScore != nil {
		return *m.TotalScore
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetRoundScore() int32 {
	if m != nil && m.RoundScore != nil {
		return *m.RoundScore
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetWinNum() int32 {
	if m != nil && m.WinNum != nil {
		return *m.WinNum
	}
	return 0
}

// 本局结果(广播)
type SdyBcCurrentResult struct {
	WinCoinInfo      []*SdyBaseWinCoinInfo `protobuf:"bytes,1,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	CurrentRound     *int32                `protobuf:"varint,2,opt,name=currentRound" json:"currentRound,omitempty"`
	BoardsCout       *int32                `protobuf:"varint,3,opt,name=boardsCout" json:"boardsCout,omitempty"`
	FootPokers       []int32               `protobuf:"varint,4,rep,name=footPokers" json:"footPokers,omitempty"`
	HuanDiPokers     []int32               `protobuf:"varint,5,rep,name=huanDiPokers" json:"huanDiPokers,omitempty"`
	BaseValue        *int32                `protobuf:"varint,6,opt,name=baseValue" json:"baseValue,omitempty"`
	IsWangKou        *bool                 `protobuf:"varint,7,opt,name=isWangKou" json:"isWangKou,omitempty"`
	IsPoPai          *bool                 `protobuf:"varint,8,opt,name=isPoPai" json:"isPoPai,omitempty"`
	IsKouDi          *bool                 `protobuf:"varint,9,opt,name=isKouDi" json:"isKouDi,omitempty"`
	IsGuang          *bool                 `protobuf:"varint,10,opt,name=isGuang" json:"isGuang,omitempty"`
	IsShangChe       *bool                 `protobuf:"varint,11,opt,name=isShangChe" json:"isShangChe,omitempty"`
	IsChengPai       *bool                 `protobuf:"varint,12,opt,name=isChengPai" json:"isChengPai,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SdyBcCurrentResult) Reset()                    { *m = SdyBcCurrentResult{} }
func (m *SdyBcCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*SdyBcCurrentResult) ProtoMessage()               {}
func (*SdyBcCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *SdyBcCurrentResult) GetWinCoinInfo() []*SdyBaseWinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

func (m *SdyBcCurrentResult) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *SdyBcCurrentResult) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *SdyBcCurrentResult) GetFootPokers() []int32 {
	if m != nil {
		return m.FootPokers
	}
	return nil
}

func (m *SdyBcCurrentResult) GetHuanDiPokers() []int32 {
	if m != nil {
		return m.HuanDiPokers
	}
	return nil
}

func (m *SdyBcCurrentResult) GetBaseValue() int32 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *SdyBcCurrentResult) GetIsWangKou() bool {
	if m != nil && m.IsWangKou != nil {
		return *m.IsWangKou
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsPoPai() bool {
	if m != nil && m.IsPoPai != nil {
		return *m.IsPoPai
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsKouDi() bool {
	if m != nil && m.IsKouDi != nil {
		return *m.IsKouDi
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsGuang() bool {
	if m != nil && m.IsGuang != nil {
		return *m.IsGuang
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsShangChe() bool {
	if m != nil && m.IsShangChe != nil {
		return *m.IsShangChe
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsChengPai() bool {
	if m != nil && m.IsChengPai != nil {
		return *m.IsChengPai
	}
	return false
}

type SdyBaseBill struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WXInfo           *WeixinInfo  `protobuf:"bytes,2,opt,name=wXInfo" json:"wXInfo,omitempty"`
	IsBigWin         *bool        `protobuf:"varint,3,opt,name=isBigWin" json:"isBigWin,omitempty"`
	IsOwner          *bool        `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int32       `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	ChengPai         *int32       `protobuf:"varint,6,opt,name=chengPai" json:"chengPai,omitempty"`
	PoPai            *int32       `protobuf:"varint,7,opt,name=poPai" json:"poPai,omitempty"`
	KouDi            *int32       `protobuf:"varint,8,opt,name=kouDi" json:"kouDi,omitempty"`
	GuangPai         *int32       `protobuf:"varint,9,opt,name=guangPai" json:"guangPai,omitempty"`
	ShangChe         *int32       `protobuf:"varint,10,opt,name=shangChe" json:"shangChe,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBaseBill) Reset()                    { *m = SdyBaseBill{} }
func (m *SdyBaseBill) String() string            { return proto.CompactTextString(m) }
func (*SdyBaseBill) ProtoMessage()               {}
func (*SdyBaseBill) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *SdyBaseBill) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBaseBill) GetWXInfo() *WeixinInfo {
	if m != nil {
		return m.WXInfo
	}
	return nil
}

func (m *SdyBaseBill) GetIsBigWin() bool {
	if m != nil && m.IsBigWin != nil {
		return *m.IsBigWin
	}
	return false
}

func (m *SdyBaseBill) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *SdyBaseBill) GetWinCoin() int32 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *SdyBaseBill) GetChengPai() int32 {
	if m != nil && m.ChengPai != nil {
		return *m.ChengPai
	}
	return 0
}

func (m *SdyBaseBill) GetPoPai() int32 {
	if m != nil && m.PoPai != nil {
		return *m.PoPai
	}
	return 0
}

func (m *SdyBaseBill) GetKouDi() int32 {
	if m != nil && m.KouDi != nil {
		return *m.KouDi
	}
	return 0
}

func (m *SdyBaseBill) GetGuangPai() int32 {
	if m != nil && m.GuangPai != nil {
		return *m.GuangPai
	}
	return 0
}

func (m *SdyBaseBill) GetShangChe() int32 {
	if m != nil && m.ShangChe != nil {
		return *m.ShangChe
	}
	return 0
}

// 全局结果
type SdyBcEndLotteryInfo struct {
	UserBills        []*SdyBaseBill `protobuf:"bytes,1,rep,name=userBills" json:"userBills,omitempty"`
	CurrentRound     *int32         `protobuf:"varint,2,opt,name=currentRound" json:"currentRound,omitempty"`
	BoardsCout       *int32         `protobuf:"varint,3,opt,name=boardsCout" json:"boardsCout,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SdyBcEndLotteryInfo) Reset()                    { *m = SdyBcEndLotteryInfo{} }
func (m *SdyBcEndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBcEndLotteryInfo) ProtoMessage()               {}
func (*SdyBcEndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *SdyBcEndLotteryInfo) GetUserBills() []*SdyBaseBill {
	if m != nil {
		return m.UserBills
	}
	return nil
}

func (m *SdyBcEndLotteryInfo) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *SdyBcEndLotteryInfo) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

// 出的牌,出牌的玩家id
type SdyReUserOutPai struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	PokerId          *int32  `protobuf:"varint,2,opt,name=pokerId" json:"pokerId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReUserOutPai) Reset()                    { *m = SdyReUserOutPai{} }
func (m *SdyReUserOutPai) String() string            { return proto.CompactTextString(m) }
func (*SdyReUserOutPai) ProtoMessage()               {}
func (*SdyReUserOutPai) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *SdyReUserOutPai) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReUserOutPai) GetPokerId() int32 {
	if m != nil && m.PokerId != nil {
		return *m.PokerId
	}
	return 0
}

// 准备阶段需要发的消息
type SdyReReady struct {
	HandPokersId     []int32 `protobuf:"varint,1,rep,name=handPokersId" json:"handPokersId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReReady) Reset()                    { *m = SdyReReady{} }
func (m *SdyReReady) String() string            { return proto.CompactTextString(m) }
func (*SdyReReady) ProtoMessage()               {}
func (*SdyReReady) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *SdyReReady) GetHandPokersId() []int32 {
	if m != nil {
		return m.HandPokersId
	}
	return nil
}

type SdyReHuanDi struct {
	DeskFootPokersId []int32 `protobuf:"varint,1,rep,name=deskFootPokersId" json:"deskFootPokersId,omitempty"`
	HuanDiPokersId   []int32 `protobuf:"varint,2,rep,name=huanDiPokersId" json:"huanDiPokersId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReHuanDi) Reset()                    { *m = SdyReHuanDi{} }
func (m *SdyReHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyReHuanDi) ProtoMessage()               {}
func (*SdyReHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *SdyReHuanDi) GetDeskFootPokersId() []int32 {
	if m != nil {
		return m.DeskFootPokersId
	}
	return nil
}

func (m *SdyReHuanDi) GetHuanDiPokersId() []int32 {
	if m != nil {
		return m.HuanDiPokersId
	}
	return nil
}

type SdyRePlay struct {
	OutPai           []*SdyReUserOutPai `protobuf:"bytes,1,rep,name=outPai" json:"outPai,omitempty"`
	ScorePai         []int32            `protobuf:"varint,2,rep,name=scorePai" json:"scorePai,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SdyRePlay) Reset()                    { *m = SdyRePlay{} }
func (m *SdyRePlay) String() string            { return proto.CompactTextString(m) }
func (*SdyRePlay) ProtoMessage()               {}
func (*SdyRePlay) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *SdyRePlay) GetOutPai() []*SdyReUserOutPai {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *SdyRePlay) GetScorePai() []int32 {
	if m != nil {
		return m.ScorePai
	}
	return nil
}

type SdyReJiaoFen struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	JiaoScore        *int32  `protobuf:"varint,2,opt,name=jiaoScore" json:"jiaoScore,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReJiaoFen) Reset()                    { *m = SdyReJiaoFen{} }
func (m *SdyReJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyReJiaoFen) ProtoMessage()               {}
func (*SdyReJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

func (m *SdyReJiaoFen) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReJiaoFen) GetJiaoScore() int32 {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return 0
}

type SdyReLenHandPokers struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	PokersLength     *int32  `protobuf:"varint,2,opt,name=pokersLength" json:"pokersLength,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SdyReLenHandPokers) Reset()                    { *m = SdyReLenHandPokers{} }
func (m *SdyReLenHandPokers) String() string            { return proto.CompactTextString(m) }
func (*SdyReLenHandPokers) ProtoMessage()               {}
func (*SdyReLenHandPokers) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{11} }

func (m *SdyReLenHandPokers) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReLenHandPokers) GetPokersLength() int32 {
	if m != nil && m.PokersLength != nil {
		return *m.PokersLength
	}
	return 0
}

// 用于断线重连回来后发送的消息
type SdyBcReconnectInfo struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       []*SdyBasePlayerInfo  `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	DeskInfo         *SdyBaseDeskInfo      `protobuf:"bytes,3,opt,name=deskInfo" json:"deskInfo,omitempty"`
	ReconnectUser    *uint32               `protobuf:"varint,4,opt,name=reconnectUser" json:"reconnectUser,omitempty"`
	ZhuFlower        *int32                `protobuf:"varint,5,opt,name=zhuFlower" json:"zhuFlower,omitempty"`
	ActiveUser       *uint32               `protobuf:"varint,6,opt,name=activeUser" json:"activeUser,omitempty"`
	DeskStatus       *int32                `protobuf:"varint,7,opt,name=deskStatus" json:"deskStatus,omitempty"`
	BankerFen        *int32                `protobuf:"varint,8,opt,name=bankerFen" json:"bankerFen,omitempty"`
	Ready            *SdyReReady           `protobuf:"bytes,9,opt,name=ready" json:"ready,omitempty"`
	HuanDi           *SdyReHuanDi          `protobuf:"bytes,10,opt,name=huanDi" json:"huanDi,omitempty"`
	OutPai           *SdyRePlay            `protobuf:"bytes,11,opt,name=outPai" json:"outPai,omitempty"`
	CurrentResult    *SdyBcCurrentResult   `protobuf:"bytes,12,opt,name=currentResult" json:"currentResult,omitempty"`
	JiaoFen          []*SdyReJiaoFen       `protobuf:"bytes,13,rep,name=jiaoFen" json:"jiaoFen,omitempty"`
	PaiLen           []*SdyReLenHandPokers `protobuf:"bytes,14,rep,name=paiLen" json:"paiLen,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SdyBcReconnectInfo) Reset()                    { *m = SdyBcReconnectInfo{} }
func (m *SdyBcReconnectInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBcReconnectInfo) ProtoMessage()               {}
func (*SdyBcReconnectInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{12} }

func (m *SdyBcReconnectInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetPlayerInfo() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetDeskInfo() *SdyBaseDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetReconnectUser() uint32 {
	if m != nil && m.ReconnectUser != nil {
		return *m.ReconnectUser
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetZhuFlower() int32 {
	if m != nil && m.ZhuFlower != nil {
		return *m.ZhuFlower
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetActiveUser() uint32 {
	if m != nil && m.ActiveUser != nil {
		return *m.ActiveUser
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetDeskStatus() int32 {
	if m != nil && m.DeskStatus != nil {
		return *m.DeskStatus
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetBankerFen() int32 {
	if m != nil && m.BankerFen != nil {
		return *m.BankerFen
	}
	return 0
}

func (m *SdyBcReconnectInfo) GetReady() *SdyReReady {
	if m != nil {
		return m.Ready
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetHuanDi() *SdyReHuanDi {
	if m != nil {
		return m.HuanDi
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetOutPai() *SdyRePlay {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetCurrentResult() *SdyBcCurrentResult {
	if m != nil {
		return m.CurrentResult
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetJiaoFen() []*SdyReJiaoFen {
	if m != nil {
		return m.JiaoFen
	}
	return nil
}

func (m *SdyBcReconnectInfo) GetPaiLen() []*SdyReLenHandPokers {
	if m != nil {
		return m.PaiLen
	}
	return nil
}

// 用于玩家断线重连回来后广播给其他玩家，通知该玩家是否在线
type SdyBcIsOnLine struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsOnline         *bool        `protobuf:"varint,3,opt,name=isOnline" json:"isOnline,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcIsOnLine) Reset()                    { *m = SdyBcIsOnLine{} }
func (m *SdyBcIsOnLine) String() string            { return proto.CompactTextString(m) }
func (*SdyBcIsOnLine) ProtoMessage()               {}
func (*SdyBcIsOnLine) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{13} }

func (m *SdyBcIsOnLine) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcIsOnLine) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcIsOnLine) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func init() {
	proto.RegisterType((*SdyReqReady)(nil), "yjprotogo.sdy_req_ready")
	proto.RegisterType((*SdyAckReady)(nil), "yjprotogo.sdy_ack_ready")
	proto.RegisterType((*SdyBaseWinCoinInfo)(nil), "yjprotogo.sdy_base_winCoinInfo")
	proto.RegisterType((*SdyBcCurrentResult)(nil), "yjprotogo.sdy_bc_currentResult")
	proto.RegisterType((*SdyBaseBill)(nil), "yjprotogo.sdy_base_bill")
	proto.RegisterType((*SdyBcEndLotteryInfo)(nil), "yjprotogo.sdy_bc_endLotteryInfo")
	proto.RegisterType((*SdyReUserOutPai)(nil), "yjprotogo.sdy_re_userOutPai")
	proto.RegisterType((*SdyReReady)(nil), "yjprotogo.sdy_re_ready")
	proto.RegisterType((*SdyReHuanDi)(nil), "yjprotogo.sdy_re_huanDi")
	proto.RegisterType((*SdyRePlay)(nil), "yjprotogo.sdy_re_play")
	proto.RegisterType((*SdyReJiaoFen)(nil), "yjprotogo.sdy_re_jiaoFen")
	proto.RegisterType((*SdyReLenHandPokers)(nil), "yjprotogo.sdy_re_lenHandPokers")
	proto.RegisterType((*SdyBcReconnectInfo)(nil), "yjprotogo.sdy_bc_reconnectInfo")
	proto.RegisterType((*SdyBcIsOnLine)(nil), "yjprotogo.sdy_bc_isOnLine")
}

var fileDescriptor11 = []byte{
	// 977 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x6f, 0x1b, 0xc5,
	0x13, 0xff, 0x3a, 0xae, 0x1d, 0x7b, 0x1c, 0xe7, 0x5b, 0x8e, 0xb6, 0x1c, 0x51, 0x54, 0xac, 0x13,
	0x42, 0x11, 0x52, 0x23, 0x14, 0x10, 0xf0, 0x84, 0x44, 0xd2, 0x9a, 0x44, 0x8d, 0x88, 0xb5, 0x11,
	0x84, 0xb7, 0xd3, 0xe5, 0xbc, 0x3d, 0x6f, 0x7d, 0xd9, 0x35, 0xbb, 0x77, 0x18, 0xf3, 0x6f, 0xf0,
	0xce, 0x33, 0x12, 0xff, 0x03, 0x7f, 0x1b, 0x9a, 0xd9, 0xdd, 0xbb, 0x73, 0xec, 0x3e, 0xb4, 0xea,
	0x93, 0x6f, 0x3e, 0x9f, 0xd9, 0x99, 0xd9, 0xf9, 0xb5, 0x86, 0x7d, 0x33, 0x5d, 0xc5, 0x53, 0x6e,
	0xe6, 0xc7, 0x0b, 0xad, 0x0a, 0x15, 0xf4, 0x57, 0xaf, 0xe9, 0x23, 0x53, 0x07, 0x1f, 0xa6, 0xea,
	0xee, 0x4e, 0xc9, 0x38, 0xcd, 0x05, 0x97, 0x85, 0xe5, 0x0f, 0x48, 0xff, 0x36, 0x31, 0xdc, 0xca,
	0xd1, 0x0d, 0x0c, 0x11, 0xd1, 0xfc, 0xd7, 0x58, 0xf3, 0x64, 0xba, 0x0a, 0x8e, 0xa1, 0x3b, 0xe3,
	0xc9, 0x94, 0xeb, 0xb0, 0x35, 0x6a, 0x1d, 0x0d, 0x4e, 0x9e, 0x1c, 0x57, 0x16, 0x8f, 0x27, 0xf8,
	0x7b, 0x4e, 0x2c, 0x73, 0x5a, 0xc1, 0x13, 0xe8, 0x96, 0x86, 0xeb, 0x8b, 0x69, 0xb8, 0x33, 0x6a,
	0x1d, 0x0d, 0x99, 0x93, 0x22, 0x61, 0x0d, 0x27, 0xe9, 0xfc, 0x1d, 0x0d, 0x3f, 0x84, 0xf6, 0x9d,
	0xc9, 0xc8, 0x6a, 0x9f, 0xe1, 0x67, 0xc3, 0x55, 0x7b, 0xcd, 0xd5, 0x5f, 0x2d, 0x78, 0xe4, 0xaf,
	0x15, 0x2f, 0x85, 0x3c, 0x53, 0x42, 0x5e, 0xc8, 0x57, 0xea, 0xad, 0x5d, 0x3e, 0x05, 0x28, 0x54,
	0x91, 0xe4, 0xd7, 0xa9, 0xd2, 0x9c, 0x3c, 0x77, 0x58, 0x03, 0x41, 0x5e, 0xab, 0x52, 0x4e, 0x2d,
	0xdf, 0xb6, 0x7c, 0x8d, 0x60, 0x80, 0x4b, 0x21, 0x7f, 0x2c, 0xef, 0xc2, 0x07, 0xc4, 0x39, 0x29,
	0xfa, 0xa7, 0xed, 0x02, 0x4c, 0xe3, 0xb4, 0xd4, 0x9a, 0xcb, 0x82, 0x71, 0x53, 0xe6, 0x45, 0xf0,
	0x3d, 0x0c, 0x1a, 0xf1, 0x86, 0xad, 0x51, 0xfb, 0x68, 0x70, 0xf2, 0x49, 0x23, 0xca, 0x6d, 0xd7,
	0x62, 0xcd, 0x33, 0x41, 0x04, 0x7b, 0xde, 0x26, 0x06, 0xe2, 0xa2, 0x5e, 0xc3, 0x30, 0xee, 0x5b,
	0x95, 0xe8, 0xa9, 0x39, 0x53, 0x65, 0xe1, 0xe3, 0xae, 0x11, 0xe4, 0x5f, 0x29, 0x55, 0x4c, 0xd4,
	0x9c, 0x6b, 0x13, 0x3e, 0x18, 0xb5, 0x91, 0xaf, 0x11, 0xf4, 0x31, 0x2b, 0x13, 0xf9, 0x5c, 0x38,
	0x8d, 0x0e, 0x69, 0xac, 0x61, 0xc1, 0x21, 0xf4, 0x31, 0xd0, 0x9f, 0x93, 0xbc, 0xe4, 0x61, 0x97,
	0x5c, 0xd4, 0x00, 0xb2, 0xc2, 0xdc, 0x24, 0x32, 0x7b, 0xa9, 0xca, 0x70, 0x77, 0xd4, 0x3a, 0xea,
	0xb1, 0x1a, 0x08, 0x42, 0xd8, 0x15, 0x66, 0xa2, 0x26, 0x89, 0x08, 0x7b, 0xc4, 0x79, 0xd1, 0x32,
	0x2f, 0x55, 0xf9, 0x5c, 0x84, 0x7d, 0xcf, 0x90, 0x68, 0x99, 0x1f, 0xca, 0x44, 0x66, 0x21, 0x78,
	0x86, 0x44, 0xbc, 0x8d, 0x30, 0xd7, 0xb3, 0x44, 0x66, 0x67, 0x33, 0x1e, 0x0e, 0x88, 0x6c, 0x20,
	0x96, 0x3f, 0x9b, 0x71, 0x99, 0xa1, 0xc3, 0x3d, 0xcf, 0x7b, 0x24, 0xfa, 0x77, 0xc7, 0xb6, 0x2e,
	0xe5, 0xfd, 0x56, 0xe4, 0xf9, 0x5b, 0xf7, 0xd1, 0x33, 0xe8, 0x2e, 0x7f, 0xa1, 0x8a, 0xee, 0x90,
	0xfe, 0xe3, 0x86, 0xfe, 0x0d, 0x17, 0xbf, 0xbb, 0x3a, 0x3a, 0xa5, 0xe0, 0x00, 0x7a, 0xc2, 0x9c,
	0x8a, 0xec, 0x46, 0x48, 0x2a, 0x4e, 0x8f, 0x55, 0xb2, 0xbd, 0xe6, 0xd5, 0x52, 0x72, 0x4d, 0x3d,
	0x45, 0xd7, 0x24, 0x11, 0x19, 0xd7, 0x07, 0x61, 0x87, 0xd2, 0xed, 0x45, 0xb4, 0x97, 0xfa, 0xeb,
	0xd9, 0x4a, 0x54, 0x72, 0xf0, 0x08, 0x3a, 0x0b, 0x4a, 0xf4, 0x2e, 0x11, 0x56, 0x40, 0x74, 0x4e,
	0x49, 0xee, 0x59, 0x94, 0x04, 0xb4, 0x93, 0x61, 0x46, 0x51, 0xbd, 0x6f, 0xed, 0x78, 0x19, 0x39,
	0xe3, 0x53, 0x0c, 0x96, 0xf3, 0x72, 0xf4, 0x67, 0x0b, 0x1e, 0xbb, 0x76, 0xe7, 0x72, 0x7a, 0xa9,
	0x8a, 0x82, 0xeb, 0x15, 0xdd, 0xf4, 0x6b, 0xe8, 0xe3, 0xcc, 0x9e, 0x8a, 0x3c, 0x37, 0xae, 0xdb,
	0xc3, 0x6d, 0xdd, 0x8e, 0x59, 0x67, 0xb5, 0xea, 0xfb, 0x68, 0xf2, 0xe8, 0x05, 0x7c, 0x60, 0x37,
	0x5d, 0x8c, 0x76, 0xaf, 0xca, 0x02, 0xaf, 0x51, 0xaf, 0x94, 0x56, 0x73, 0xa5, 0x60, 0x72, 0x17,
	0xd8, 0xd7, 0x17, 0xde, 0x97, 0x17, 0xa3, 0x13, 0xd8, 0x73, 0x66, 0xec, 0x5a, 0xc3, 0xd9, 0x48,
	0xe4, 0xd4, 0x4e, 0x01, 0xd9, 0xb1, 0xb3, 0xd1, 0xc0, 0xa2, 0xd4, 0x2f, 0xd9, 0xd8, 0x8e, 0x4c,
	0xf0, 0x39, 0x3c, 0xc4, 0x9d, 0x3d, 0xae, 0x46, 0xac, 0x3a, 0xb8, 0x81, 0x07, 0x9f, 0xc1, 0x7e,
	0x73, 0xd0, 0x28, 0x22, 0xd4, 0xbc, 0x87, 0x46, 0x31, 0x0c, 0x9c, 0x93, 0x45, 0x9e, 0xac, 0x82,
	0xaf, 0xa0, 0xab, 0xe8, 0x8e, 0x2e, 0xcf, 0x87, 0xf7, 0xf2, 0xbc, 0x96, 0x07, 0xe6, 0x74, 0xa9,
	0xac, 0xb8, 0xca, 0xf0, 0x9c, 0x75, 0x53, 0xc9, 0xd1, 0xd8, 0x3e, 0x36, 0x9a, 0xc7, 0xaf, 0x45,
	0xa2, 0xc6, 0x5c, 0xbe, 0x31, 0x7b, 0x87, 0xd0, 0x47, 0x95, 0xe6, 0x1a, 0xad, 0x81, 0x88, 0xd9,
	0x65, 0xa8, 0x79, 0x9c, 0x73, 0x79, 0x5e, 0xe5, 0xe9, 0x8d, 0xd6, 0x22, 0xd8, 0xa3, 0xe4, 0x9b,
	0x4b, 0x2e, 0xb3, 0x62, 0xe6, 0x8b, 0xdf, 0xc4, 0xa2, 0xbf, 0x3b, 0xd5, 0x86, 0xd5, 0x3c, 0x55,
	0x52, 0xf2, 0xb4, 0x78, 0xa7, 0x27, 0xe0, 0x3b, 0x00, 0x4c, 0x1f, 0xd7, 0x6e, 0x7c, 0x31, 0x75,
	0x4f, 0xb7, 0xb5, 0x68, 0xad, 0xc5, 0x1a, 0x27, 0x82, 0x6f, 0xa1, 0x87, 0x15, 0xa4, 0xd3, 0x6d,
	0xf2, 0x78, 0xb8, 0xed, 0xb4, 0xd7, 0x61, 0x95, 0x76, 0xf0, 0x29, 0x0c, 0xab, 0xd0, 0x7f, 0x32,
	0x6e, 0xde, 0x87, 0x6c, 0x1d, 0xc4, 0xd4, 0xfe, 0x31, 0x2b, 0xc7, 0xb9, 0x5a, 0x72, 0xed, 0xe6,
	0xbe, 0x06, 0x70, 0x06, 0x92, 0xb4, 0x10, 0xbf, 0x71, 0x32, 0xd0, 0x25, 0x03, 0x0d, 0x04, 0x79,
	0xf4, 0x77, 0x5d, 0x24, 0x45, 0x69, 0xdc, 0x0a, 0x68, 0x20, 0x76, 0x89, 0xcb, 0x39, 0xd7, 0x63,
	0x2e, 0xdd, 0x2e, 0xa8, 0x81, 0xe0, 0x19, 0x74, 0xa8, 0xe7, 0x69, 0x19, 0x0c, 0x4e, 0x3e, 0xda,
	0xec, 0x28, 0xa2, 0x99, 0xd5, 0x0a, 0xbe, 0x80, 0xae, 0x6d, 0x51, 0x5a, 0x10, 0x9b, 0x93, 0x5e,
	0x8d, 0x03, 0x73, 0x7a, 0x58, 0x2c, 0xd7, 0xb3, 0x83, 0x8d, 0x62, 0x35, 0x7a, 0xbb, 0xea, 0xd6,
	0x17, 0x30, 0x5c, 0x7b, 0x4f, 0x69, 0x99, 0x6f, 0x79, 0x40, 0xef, 0x3d, 0xbb, 0x6c, 0xfd, 0x54,
	0xf0, 0x25, 0xec, 0xba, 0x8e, 0x0e, 0x87, 0x54, 0xf0, 0x8f, 0x37, 0xfd, 0x3a, 0x05, 0xe6, 0x35,
	0x83, 0x6f, 0xa0, 0xbb, 0x48, 0xc4, 0x25, 0x97, 0xe1, 0xfe, 0xd6, 0x57, 0xfb, 0x7e, 0x7b, 0x33,
	0xa7, 0x1e, 0x95, 0xf0, 0x7f, 0x17, 0x94, 0x30, 0x57, 0xf2, 0x52, 0x48, 0xfe, 0xbe, 0xfe, 0x73,
	0xd9, 0x87, 0xe4, 0x4a, 0xe6, 0x42, 0xf2, 0xfa, 0x21, 0xb1, 0xf2, 0xe9, 0xce, 0x79, 0x7b, 0xf2,
	0xbf, 0x49, 0xeb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xd6, 0x52, 0xa1, 0x31, 0x0a, 0x00,
	0x00,
}

// Code generated by protoc-gen-go.
// source: sdy_play.proto
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

type SdyEnumJdScore int32

const (
	SdyEnumJdScore_SDY_SIXTY       SdyEnumJdScore = 1
	SdyEnumJdScore_SDY_SIXTYFIVE   SdyEnumJdScore = 2
	SdyEnumJdScore_SDY_SEVENTY     SdyEnumJdScore = 3
	SdyEnumJdScore_SDY_SEVENTYFIVE SdyEnumJdScore = 4
	SdyEnumJdScore_SDY_NONE        SdyEnumJdScore = 5
)

var SdyEnumJdScore_name = map[int32]string{
	1: "SDY_SIXTY",
	2: "SDY_SIXTYFIVE",
	3: "SDY_SEVENTY",
	4: "SDY_SEVENTYFIVE",
	5: "SDY_NONE",
}
var SdyEnumJdScore_value = map[string]int32{
	"SDY_SIXTY":       1,
	"SDY_SIXTYFIVE":   2,
	"SDY_SEVENTY":     3,
	"SDY_SEVENTYFIVE": 4,
	"SDY_NONE":        5,
}

func (x SdyEnumJdScore) Enum() *SdyEnumJdScore {
	p := new(SdyEnumJdScore)
	*p = x
	return p
}
func (x SdyEnumJdScore) String() string {
	return proto.EnumName(SdyEnumJdScore_name, int32(x))
}
func (x *SdyEnumJdScore) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SdyEnumJdScore_value, data, "SdyEnumJdScore")
	if err != nil {
		return err
	}
	*x = SdyEnumJdScore(value)
	return nil
}
func (SdyEnumJdScore) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

// 开局（接收服务端消息）
type SdyBcOpening struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Ticket           *int64               `protobuf:"varint,3,opt,name=ticket" json:"ticket,omitempty"`
	UserCoin         *int64               `protobuf:"varint,4,opt,name=userCoin" json:"userCoin,omitempty"`
	UserRoomCard     *int64               `protobuf:"varint,5,opt,name=userRoomCard" json:"userRoomCard,omitempty"`
	PlayerInfos      []*SdyBasePlayerInfo `protobuf:"bytes,6,rep,name=playerInfos" json:"playerInfos,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcOpening) Reset()                    { *m = SdyBcOpening{} }
func (m *SdyBcOpening) String() string            { return proto.CompactTextString(m) }
func (*SdyBcOpening) ProtoMessage()               {}
func (*SdyBcOpening) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *SdyBcOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcOpening) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcOpening) GetTicket() int64 {
	if m != nil && m.Ticket != nil {
		return *m.Ticket
	}
	return 0
}

func (m *SdyBcOpening) GetUserCoin() int64 {
	if m != nil && m.UserCoin != nil {
		return *m.UserCoin
	}
	return 0
}

func (m *SdyBcOpening) GetUserRoomCard() int64 {
	if m != nil && m.UserRoomCard != nil {
		return *m.UserRoomCard
	}
	return 0
}

func (m *SdyBcOpening) GetPlayerInfos() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

// 发牌
type SdyBcDealCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerPokers     []int32      `protobuf:"varint,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcDealCards) Reset()                    { *m = SdyBcDealCards{} }
func (m *SdyBcDealCards) String() string            { return proto.CompactTextString(m) }
func (*SdyBcDealCards) ProtoMessage()               {}
func (*SdyBcDealCards) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *SdyBcDealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcDealCards) GetPlayerPokers() []int32 {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

// 叫分
type SdyReqJiaoFen struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32         `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Jiao             *bool           `protobuf:"varint,3,opt,name=jiao" json:"jiao,omitempty"`
	Score            *SdyEnumJdScore `protobuf:"varint,4,opt,name=score,enum=yjprotogo.SdyEnumJdScore" json:"score,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SdyReqJiaoFen) Reset()                    { *m = SdyReqJiaoFen{} }
func (m *SdyReqJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyReqJiaoFen) ProtoMessage()               {}
func (*SdyReqJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *SdyReqJiaoFen) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqJiaoFen) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReqJiaoFen) GetJiao() bool {
	if m != nil && m.Jiao != nil {
		return *m.Jiao
	}
	return false
}

func (m *SdyReqJiaoFen) GetScore() SdyEnumJdScore {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return SdyEnumJdScore_SDY_SIXTY
}

// 叫分回复
type SdyAckJiaoFen struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32         `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Jiao             *bool           `protobuf:"varint,3,opt,name=jiao" json:"jiao,omitempty"`
	Score            *SdyEnumJdScore `protobuf:"varint,4,opt,name=score,enum=yjprotogo.SdyEnumJdScore" json:"score,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SdyAckJiaoFen) Reset()                    { *m = SdyAckJiaoFen{} }
func (m *SdyAckJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyAckJiaoFen) ProtoMessage()               {}
func (*SdyAckJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *SdyAckJiaoFen) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckJiaoFen) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyAckJiaoFen) GetJiao() bool {
	if m != nil && m.Jiao != nil {
		return *m.Jiao
	}
	return false
}

func (m *SdyAckJiaoFen) GetScore() SdyEnumJdScore {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return SdyEnumJdScore_SDY_SIXTY
}

// 叫地主结束，开始游戏 (广播)
type SdyBcStartPlay struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	FootPokers       []int32      `protobuf:"varint,2,rep,name=footPokers" json:"footPokers,omitempty"`
	Banker           *uint32      `protobuf:"varint,3,opt,name=banker" json:"banker,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcStartPlay) Reset()                    { *m = SdyBcStartPlay{} }
func (m *SdyBcStartPlay) String() string            { return proto.CompactTextString(m) }
func (*SdyBcStartPlay) ProtoMessage()               {}
func (*SdyBcStartPlay) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *SdyBcStartPlay) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcStartPlay) GetFootPokers() []int32 {
	if m != nil {
		return m.FootPokers
	}
	return nil
}

func (m *SdyBcStartPlay) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

// 出牌
type SdyReqOutCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	OutCards         []int32      `protobuf:"varint,2,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqOutCards) Reset()                    { *m = SdyReqOutCards{} }
func (m *SdyReqOutCards) String() string            { return proto.CompactTextString(m) }
func (*SdyReqOutCards) ProtoMessage()               {}
func (*SdyReqOutCards) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *SdyReqOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqOutCards) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

type SdyAckOutCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	RemainPokers     *int32       `protobuf:"varint,3,opt,name=remainPokers" json:"remainPokers,omitempty"`
	OutCards         []int32      `protobuf:"varint,4,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckOutCards) Reset()                    { *m = SdyAckOutCards{} }
func (m *SdyAckOutCards) String() string            { return proto.CompactTextString(m) }
func (*SdyAckOutCards) ProtoMessage()               {}
func (*SdyAckOutCards) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *SdyAckOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckOutCards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyAckOutCards) GetRemainPokers() int32 {
	if m != nil && m.RemainPokers != nil {
		return *m.RemainPokers
	}
	return 0
}

func (m *SdyAckOutCards) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

// 轮到谁操作
type SdyBcOverTurn struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Time             *int32               `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	ActType          *SdyEnumActType      `protobuf:"varint,4,opt,name=actType,enum=yjprotogo.SdyEnumActType" json:"actType,omitempty"`
	CanOutCards      *bool                `protobuf:"varint,5,opt,name=canOutCards" json:"canOutCards,omitempty"`
	PlayerInfos      []*SdyBasePlayerInfo `protobuf:"bytes,6,rep,name=playerInfos" json:"playerInfos,omitempty"`
	JiaoScore        *SdyEnumJdScore      `protobuf:"varint,7,opt,name=jiaoScore,enum=yjprotogo.SdyEnumJdScore" json:"jiaoScore,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcOverTurn) Reset()                    { *m = SdyBcOverTurn{} }
func (m *SdyBcOverTurn) String() string            { return proto.CompactTextString(m) }
func (*SdyBcOverTurn) ProtoMessage()               {}
func (*SdyBcOverTurn) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *SdyBcOverTurn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcOverTurn) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcOverTurn) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *SdyBcOverTurn) GetActType() SdyEnumActType {
	if m != nil && m.ActType != nil {
		return *m.ActType
	}
	return SdyEnumActType_SDY_T_NORMAL_ACT
}

func (m *SdyBcOverTurn) GetCanOutCards() bool {
	if m != nil && m.CanOutCards != nil {
		return *m.CanOutCards
	}
	return false
}

func (m *SdyBcOverTurn) GetPlayerInfos() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

func (m *SdyBcOverTurn) GetJiaoScore() SdyEnumJdScore {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return SdyEnumJdScore_SDY_SIXTY
}

// 游戏信息(广播)
type SdyBcGameInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	SenderUserId     *uint32              `protobuf:"varint,2,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *bool                `protobuf:"varint,3,opt,name=isReconnect" json:"isReconnect,omitempty"`
	PlayerInfo       []*SdyBasePlayerInfo `protobuf:"bytes,4,rep,name=playerInfo" json:"playerInfo,omitempty"`
	DeskInfo         *SdyBaseDeskInfo     `protobuf:"bytes,5,opt,name=deskInfo" json:"deskInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcGameInfo) Reset()                    { *m = SdyBcGameInfo{} }
func (m *SdyBcGameInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBcGameInfo) ProtoMessage()               {}
func (*SdyBcGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *SdyBcGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *SdyBcGameInfo) GetIsReconnect() bool {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return false
}

func (m *SdyBcGameInfo) GetPlayerInfo() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *SdyBcGameInfo) GetDeskInfo() *SdyBaseDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

type SdyBcHuandi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcHuandi) Reset()                    { *m = SdyBcHuandi{} }
func (m *SdyBcHuandi) String() string            { return proto.CompactTextString(m) }
func (*SdyBcHuandi) ProtoMessage()               {}
func (*SdyBcHuandi) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *SdyBcHuandi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcHuandi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type SdyHuandi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	FootPokers       []int32      `protobuf:"varint,3,rep,name=footPokers" json:"footPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyHuandi) Reset()                    { *m = SdyHuandi{} }
func (m *SdyHuandi) String() string            { return proto.CompactTextString(m) }
func (*SdyHuandi) ProtoMessage()               {}
func (*SdyHuandi) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{10} }

func (m *SdyHuandi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyHuandi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyHuandi) GetFootPokers() []int32 {
	if m != nil {
		return m.FootPokers
	}
	return nil
}

type SdyReqHuanDi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	OutCards         []int32      `protobuf:"varint,3,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqHuanDi) Reset()                    { *m = SdyReqHuanDi{} }
func (m *SdyReqHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyReqHuanDi) ProtoMessage()               {}
func (*SdyReqHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{11} }

func (m *SdyReqHuanDi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqHuanDi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReqHuanDi) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

type SdyAckHuandi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsHuanDiOver     *bool        `protobuf:"varint,3,opt,name=isHuanDiOver" json:"isHuanDiOver,omitempty"`
	OutCards         []int32      `protobuf:"varint,4,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckHuandi) Reset()                    { *m = SdyAckHuandi{} }
func (m *SdyAckHuandi) String() string            { return proto.CompactTextString(m) }
func (*SdyAckHuandi) ProtoMessage()               {}
func (*SdyAckHuandi) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{12} }

func (m *SdyAckHuandi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckHuandi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyAckHuandi) GetIsHuanDiOver() bool {
	if m != nil && m.IsHuanDiOver != nil {
		return *m.IsHuanDiOver
	}
	return false
}

func (m *SdyAckHuandi) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

func init() {
	proto.RegisterType((*SdyBcOpening)(nil), "yjprotogo.sdy_bc_opening")
	proto.RegisterType((*SdyBcDealCards)(nil), "yjprotogo.sdy_bc_dealCards")
	proto.RegisterType((*SdyReqJiaoFen)(nil), "yjprotogo.sdy_req_jiaoFen")
	proto.RegisterType((*SdyAckJiaoFen)(nil), "yjprotogo.sdy_ack_jiaoFen")
	proto.RegisterType((*SdyBcStartPlay)(nil), "yjprotogo.sdy_bc_startPlay")
	proto.RegisterType((*SdyReqOutCards)(nil), "yjprotogo.sdy_req_outCards")
	proto.RegisterType((*SdyAckOutCards)(nil), "yjprotogo.sdy_ack_outCards")
	proto.RegisterType((*SdyBcOverTurn)(nil), "yjprotogo.sdy_bc_overTurn")
	proto.RegisterType((*SdyBcGameInfo)(nil), "yjprotogo.sdy_bc_gameInfo")
	proto.RegisterType((*SdyBcHuandi)(nil), "yjprotogo.sdy_bc_huandi")
	proto.RegisterType((*SdyHuandi)(nil), "yjprotogo.sdy_huandi")
	proto.RegisterType((*SdyReqHuanDi)(nil), "yjprotogo.sdy_req_huanDi")
	proto.RegisterType((*SdyAckHuandi)(nil), "yjprotogo.sdy_ack_huandi")
	proto.RegisterEnum("yjprotogo.SdyEnumJdScore", SdyEnumJdScore_name, SdyEnumJdScore_value)
}

var fileDescriptor7 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4f, 0x4f, 0xdb, 0x3e,
	0x18, 0xfe, 0xa5, 0x69, 0xa1, 0xbc, 0x6d, 0xa1, 0x3f, 0x23, 0xa1, 0xa8, 0x9b, 0x50, 0xe5, 0x53,
	0xb5, 0x43, 0xa5, 0x21, 0x4d, 0xda, 0x2e, 0xd3, 0x34, 0x28, 0xa2, 0x17, 0xa8, 0x4c, 0xc7, 0xc6,
	0x65, 0x91, 0x49, 0x0c, 0x84, 0x36, 0x76, 0xe7, 0xb8, 0x48, 0xdd, 0x6d, 0xdf, 0x62, 0xda, 0x69,
	0xdf, 0x65, 0xdf, 0x68, 0x9f, 0x60, 0xb2, 0xe3, 0xa6, 0x49, 0xb5, 0x81, 0x60, 0x3d, 0xec, 0x14,
	0xbf, 0x8f, 0x1f, 0xbf, 0x7f, 0x1e, 0xbf, 0x7e, 0x03, 0x9b, 0x49, 0x38, 0xf3, 0x27, 0x63, 0x3a,
	0xeb, 0x4e, 0xa4, 0x50, 0x02, 0x6d, 0xcc, 0x6e, 0xcc, 0xe2, 0x4a, 0xb4, 0xb6, 0x03, 0x11, 0xc7,
	0x82, 0xfb, 0xc1, 0x38, 0x62, 0x5c, 0xa5, 0xfb, 0x2d, 0xc3, 0xbf, 0xa0, 0x09, 0x4b, 0x6d, 0xfc,
	0xd3, 0x49, 0x5d, 0x5c, 0x04, 0xbe, 0x98, 0x30, 0x1e, 0xf1, 0x2b, 0xd4, 0x85, 0xb5, 0x6b, 0x46,
	0x43, 0x26, 0x3d, 0xa7, 0xed, 0x74, 0x6a, 0x7b, 0x3b, 0xdd, 0xcc, 0x67, 0x77, 0xa0, 0xbf, 0x47,
	0x66, 0x97, 0x58, 0x16, 0xda, 0x81, 0xb5, 0x69, 0xc2, 0x64, 0x3f, 0xf4, 0x4a, 0x6d, 0xa7, 0xd3,
	0x20, 0xd6, 0xd2, 0xb8, 0x8a, 0x82, 0x11, 0x53, 0x9e, 0xdb, 0x76, 0x3a, 0x2e, 0xb1, 0x16, 0x6a,
	0x41, 0x55, 0x33, 0xf6, 0x45, 0xc4, 0xbd, 0xb2, 0xd9, 0xc9, 0x6c, 0x84, 0xa1, 0xae, 0xd7, 0x44,
	0x88, 0x78, 0x9f, 0xca, 0xd0, 0xab, 0x98, 0xfd, 0x02, 0x86, 0xde, 0x40, 0x4d, 0x17, 0xcc, 0x64,
	0x9f, 0x5f, 0x8a, 0xc4, 0x5b, 0x6b, 0xbb, 0x9d, 0xda, 0xde, 0x6e, 0x2e, 0xc9, 0x79, 0x89, 0xfe,
	0x82, 0x46, 0xf2, 0x47, 0xf0, 0x25, 0x34, 0x6d, 0xcd, 0x21, 0xa3, 0x63, 0xed, 0x34, 0x79, 0x70,
	0xd5, 0x18, 0xea, 0xa9, 0xcb, 0x81, 0x18, 0x31, 0x99, 0x78, 0xa5, 0xb6, 0xdb, 0xa9, 0x90, 0x02,
	0x86, 0xbf, 0x3b, 0xb0, 0xa5, 0x03, 0x49, 0xf6, 0xc9, 0xbf, 0x89, 0xa8, 0x38, 0x64, 0x7c, 0x65,
	0xea, 0x22, 0x28, 0x6b, 0x97, 0x46, 0xdb, 0x2a, 0x31, 0x6b, 0xf4, 0x1c, 0x2a, 0x49, 0x20, 0x24,
	0x33, 0xb2, 0x6e, 0xee, 0x3d, 0x59, 0xd2, 0x84, 0xf1, 0x69, 0xec, 0xdf, 0x84, 0xa7, 0x9a, 0x42,
	0x52, 0x66, 0x96, 0x22, 0x0d, 0x46, 0xff, 0x6a, 0x8a, 0x9f, 0xb3, 0xdb, 0x4a, 0x14, 0x95, 0x6a,
	0x30, 0xa6, 0xb3, 0x07, 0xa7, 0xb8, 0x0b, 0x70, 0x29, 0x84, 0x2a, 0xdc, 0x55, 0x0e, 0xd1, 0x25,
	0x5c, 0x50, 0x3e, 0x62, 0xd2, 0x24, 0xdb, 0x20, 0xd6, 0xc2, 0x1f, 0xd3, 0xd8, 0xfa, 0x02, 0xc5,
	0x54, 0x3d, 0xae, 0x53, 0x5a, 0x50, 0x9d, 0x9f, 0xb5, 0x91, 0x33, 0x1b, 0x7f, 0x73, 0xd2, 0x00,
	0x5a, 0xfe, 0x47, 0x07, 0xf8, 0x93, 0xfe, 0x18, 0xea, 0x92, 0xc5, 0x34, 0xe2, 0xb6, 0x6c, 0x5d,
	0x5a, 0x85, 0x14, 0xb0, 0x42, 0x72, 0xe5, 0xa5, 0xe4, 0x7e, 0x94, 0xd2, 0xde, 0xd0, 0xb3, 0xe1,
	0x96, 0xc9, 0xe1, 0x54, 0xae, 0xb4, 0x37, 0x54, 0x14, 0x33, 0x9b, 0x93, 0x59, 0xa3, 0x17, 0xb0,
	0x4e, 0x03, 0x35, 0x9c, 0x4d, 0xee, 0xec, 0x0e, 0x4b, 0x21, 0x73, 0x2e, 0x6a, 0x43, 0x2d, 0xa0,
	0xfc, 0x64, 0x5e, 0x45, 0xc5, 0x74, 0x5b, 0x1e, 0xfa, 0xfb, 0x89, 0x81, 0x5e, 0xc1, 0x86, 0x6e,
	0x5f, 0xd3, 0x97, 0xde, 0xfa, 0xfd, 0xad, 0xbb, 0x60, 0xe3, 0x2f, 0x0b, 0x15, 0xaf, 0x68, 0xcc,
	0xb4, 0xbf, 0xc7, 0x0c, 0x9b, 0x84, 0xf1, 0x90, 0xc9, 0x77, 0x79, 0x2d, 0x0b, 0x98, 0x96, 0x21,
	0x4a, 0x08, 0x0b, 0x04, 0xe7, 0x2c, 0x50, 0xf6, 0xd1, 0xe5, 0x21, 0xf4, 0x1a, 0x60, 0x51, 0x93,
	0xb9, 0xed, 0xfb, 0x55, 0xc8, 0x9d, 0x40, 0x2f, 0xa1, 0x1a, 0xb2, 0x64, 0x64, 0x4e, 0x57, 0x4c,
	0xde, 0x4f, 0x7f, 0x77, 0x7a, 0xce, 0x21, 0x19, 0x1b, 0xbf, 0x87, 0x86, 0x95, 0xe0, 0x7a, 0x4a,
	0x79, 0x18, 0xad, 0xaa, 0x8d, 0xb0, 0x02, 0xd0, 0x8e, 0x57, 0xeb, 0x75, 0x69, 0x5a, 0xb8, 0xcb,
	0xd3, 0x02, 0xab, 0xf4, 0x9f, 0xa9, 0xa7, 0x82, 0x8e, 0x7c, 0xb0, 0xba, 0xc8, 0xf9, 0xe7, 0xe8,
	0x2e, 0x3d, 0xc7, 0xaf, 0xf6, 0x57, 0xad, 0x67, 0xc5, 0x8a, 0x0b, 0xc6, 0x50, 0x8f, 0x92, 0x23,
	0x53, 0xca, 0xc9, 0xad, 0x1d, 0x82, 0x55, 0x52, 0xc0, 0xee, 0x9a, 0x14, 0xcf, 0xae, 0xd3, 0x29,
	0x96, 0x7f, 0x02, 0xa8, 0x01, 0x1b, 0xa7, 0x07, 0xe7, 0xfe, 0x69, 0xff, 0xc3, 0xf0, 0xbc, 0xe9,
	0xa0, 0xff, 0xa1, 0x91, 0x99, 0x87, 0xfd, 0xb3, 0x5e, 0xb3, 0x84, 0xb6, 0xa0, 0x66, 0xa0, 0xde,
	0x59, 0xef, 0x78, 0x78, 0xde, 0x74, 0xd1, 0x36, 0x6c, 0xe5, 0x00, 0xc3, 0x2a, 0xa3, 0x3a, 0x54,
	0x35, 0x78, 0x7c, 0x72, 0xdc, 0x6b, 0x56, 0xde, 0x96, 0x8e, 0xdc, 0xc1, 0x7f, 0x03, 0xe7, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x59, 0x5f, 0x04, 0xf8, 0x08, 0x00, 0x00,
}

// Code generated by protoc-gen-go.
// source: pez_play.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of pez_base_PaiInfo from pez_base.proto

// Ignoring public import of pez_base_PlayOptions from pez_base.proto

// Ignoring public import of pez_base_PlayConf from pez_base.proto

// Ignoring public import of pez_base_RoomTypeInfo from pez_base.proto

// Ignoring public import of pez_base_PaiValue from pez_base.proto

// Ignoring public import of pez_base_PlayerCard from pez_base.proto

// Ignoring public import of pez_base_PlayerInfo from pez_base.proto

// Ignoring public import of pez_base_DeskGameInfo from pez_base.proto

// Ignoring public import of pez_enum_protoId from pez_base.proto

// Ignoring public import of pez_enum_ErrorCode from pez_base.proto

// Ignoring public import of pez_enum_Option from pez_base.proto

// Ignoring public import of pez_RoomType from pez_base.proto

// Ignoring public import of pez_enum_mjColor from pez_base.proto

// Ignoring public import of pez_enum_PaiType from pez_base.proto

// Ignoring public import of pez_enum_UserGameStatus from pez_base.proto

// Ignoring public import of pez_enum_DeskGameStatus from pez_base.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// 链接类型
type PEZ_RECONNECT_TYPE int32

const (
	PEZ_RECONNECT_TYPE_PEZ_NORMAL    PEZ_RECONNECT_TYPE = 1
	PEZ_RECONNECT_TYPE_PEZ_RECONNECT PEZ_RECONNECT_TYPE = 2
)

var PEZ_RECONNECT_TYPE_name = map[int32]string{
	1: "PEZ_NORMAL",
	2: "PEZ_RECONNECT",
}
var PEZ_RECONNECT_TYPE_value = map[string]int32{
	"PEZ_NORMAL":    1,
	"PEZ_RECONNECT": 2,
}

func (x PEZ_RECONNECT_TYPE) Enum() *PEZ_RECONNECT_TYPE {
	p := new(PEZ_RECONNECT_TYPE)
	*p = x
	return p
}
func (x PEZ_RECONNECT_TYPE) String() string {
	return proto.EnumName(PEZ_RECONNECT_TYPE_name, int32(x))
}
func (x *PEZ_RECONNECT_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PEZ_RECONNECT_TYPE_value, data, "PEZ_RECONNECT_TYPE")
	if err != nil {
		return err
	}
	*x = PEZ_RECONNECT_TYPE(value)
	return nil
}
func (PEZ_RECONNECT_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor38, []int{0} }

// 积分
type PezUserCoinBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64  `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezUserCoinBean) Reset()                    { *m = PezUserCoinBean{} }
func (m *PezUserCoinBean) String() string            { return proto.CompactTextString(m) }
func (*PezUserCoinBean) ProtoMessage()               {}
func (*PezUserCoinBean) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{0} }

func (m *PezUserCoinBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezUserCoinBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// 开局（接收服务端消息）
type PezGame_Opening struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CurrPlayCount    *int32             `protobuf:"varint,2,opt,name=CurrPlayCount" json:"CurrPlayCount,omitempty"`
	Dice1            *int32             `protobuf:"varint,3,opt,name=dice1" json:"dice1,omitempty"`
	Dice2            *int32             `protobuf:"varint,4,opt,name=dice2" json:"dice2,omitempty"`
	UserCoinBeans    []*PezUserCoinBean `protobuf:"bytes,5,rep,name=userCoinBeans" json:"userCoinBeans,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *PezGame_Opening) Reset()                    { *m = PezGame_Opening{} }
func (m *PezGame_Opening) String() string            { return proto.CompactTextString(m) }
func (*PezGame_Opening) ProtoMessage()               {}
func (*PezGame_Opening) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{1} }

func (m *PezGame_Opening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezGame_Opening) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *PezGame_Opening) GetDice1() int32 {
	if m != nil && m.Dice1 != nil {
		return *m.Dice1
	}
	return 0
}

func (m *PezGame_Opening) GetDice2() int32 {
	if m != nil && m.Dice2 != nil {
		return *m.Dice2
	}
	return 0
}

func (m *PezGame_Opening) GetUserCoinBeans() []*PezUserCoinBean {
	if m != nil {
		return m.UserCoinBeans
	}
	return nil
}

// 发牌
type PezGame_DealCards struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerCard       []*PezBase_PlayerCard `protobuf:"bytes,2,rep,name=playerCard" json:"playerCard,omitempty"`
	PaiCount         *int32                `protobuf:"varint,3,opt,name=paiCount" json:"paiCount,omitempty"`
	DealerUserId     *uint32               `protobuf:"varint,4,opt,name=dealerUserId" json:"dealerUserId,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PezGame_DealCards) Reset()                    { *m = PezGame_DealCards{} }
func (m *PezGame_DealCards) String() string            { return proto.CompactTextString(m) }
func (*PezGame_DealCards) ProtoMessage()               {}
func (*PezGame_DealCards) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{2} }

func (m *PezGame_DealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezGame_DealCards) GetPlayerCard() []*PezBase_PlayerCard {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *PezGame_DealCards) GetPaiCount() int32 {
	if m != nil && m.PaiCount != nil {
		return *m.PaiCount
	}
	return 0
}

func (m *PezGame_DealCards) GetDealerUserId() uint32 {
	if m != nil && m.DealerUserId != nil {
		return *m.DealerUserId
	}
	return 0
}

// 押注
type PezGame_Bet struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	BetNum           *int32  `protobuf:"varint,2,opt,name=betNum" json:"betNum,omitempty"`
	Time             *int32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezGame_Bet) Reset()                    { *m = PezGame_Bet{} }
func (m *PezGame_Bet) String() string            { return proto.CompactTextString(m) }
func (*PezGame_Bet) ProtoMessage()               {}
func (*PezGame_Bet) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{3} }

func (m *PezGame_Bet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezGame_Bet) GetBetNum() int32 {
	if m != nil && m.BetNum != nil {
		return *m.BetNum
	}
	return 0
}

func (m *PezGame_Bet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type PezGame_AckBet struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	BetCount         *int32  `protobuf:"varint,2,opt,name=betCount" json:"betCount,omitempty"`
	Time             *int32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezGame_AckBet) Reset()                    { *m = PezGame_AckBet{} }
func (m *PezGame_AckBet) String() string            { return proto.CompactTextString(m) }
func (*PezGame_AckBet) ProtoMessage()               {}
func (*PezGame_AckBet) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{4} }

func (m *PezGame_AckBet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezGame_AckBet) GetBetCount() int32 {
	if m != nil && m.BetCount != nil {
		return *m.BetCount
	}
	return 0
}

func (m *PezGame_AckBet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

// 开牌
type PezGame_OpenCard struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CardId           *int32       `protobuf:"varint,2,opt,name=cardId" json:"cardId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PezGame_OpenCard) Reset()                    { *m = PezGame_OpenCard{} }
func (m *PezGame_OpenCard) String() string            { return proto.CompactTextString(m) }
func (*PezGame_OpenCard) ProtoMessage()               {}
func (*PezGame_OpenCard) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{5} }

func (m *PezGame_OpenCard) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezGame_OpenCard) GetCardId() int32 {
	if m != nil && m.CardId != nil {
		return *m.CardId
	}
	return 0
}

type PezGame_AckOpenCard struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Result           *int32             `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	PaiType          *int32             `protobuf:"varint,3,opt,name=paiType" json:"paiType,omitempty"`
	UserId           *uint32            `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	Card             *PezBase_PaiInfo   `protobuf:"bytes,5,opt,name=card" json:"card,omitempty"`
	Card2            *PezBase_PaiInfo   `protobuf:"bytes,6,opt,name=card2" json:"card2,omitempty"`
	UserCoinBeans    []*PezUserCoinBean `protobuf:"bytes,7,rep,name=userCoinBeans" json:"userCoinBeans,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *PezGame_AckOpenCard) Reset()                    { *m = PezGame_AckOpenCard{} }
func (m *PezGame_AckOpenCard) String() string            { return proto.CompactTextString(m) }
func (*PezGame_AckOpenCard) ProtoMessage()               {}
func (*PezGame_AckOpenCard) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{6} }

func (m *PezGame_AckOpenCard) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezGame_AckOpenCard) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *PezGame_AckOpenCard) GetPaiType() int32 {
	if m != nil && m.PaiType != nil {
		return *m.PaiType
	}
	return 0
}

func (m *PezGame_AckOpenCard) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezGame_AckOpenCard) GetCard() *PezBase_PaiInfo {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *PezGame_AckOpenCard) GetCard2() *PezBase_PaiInfo {
	if m != nil {
		return m.Card2
	}
	return nil
}

func (m *PezGame_AckOpenCard) GetUserCoinBeans() []*PezUserCoinBean {
	if m != nil {
		return m.UserCoinBeans
	}
	return nil
}

// 发送游戏信息(广播)
type PezGame_SendGameInfo struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 1. 首先是牌桌的玩家数据（玩家数据包括其id昵称筹码头像等基本信息，其手牌数据，以及自己打开牌的数据，还有状态是否已经押注了，玩家在整局的总输赢）
	PlayerInfo []*PezBase_PlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	// 2. 桌面信息（包括：游戏是否结束，当前哪个玩家未押注，倒计时剩余时间）
	DeskGameInfo *PezBase_DeskGameInfo `protobuf:"bytes,3,opt,name=deskGameInfo" json:"deskGameInfo,omitempty"`
	//
	SenderUserId     *uint32             `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *PEZ_RECONNECT_TYPE `protobuf:"varint,5,opt,name=isReconnect,enum=ddproto.PEZ_RECONNECT_TYPE" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *PezGame_SendGameInfo) Reset()                    { *m = PezGame_SendGameInfo{} }
func (m *PezGame_SendGameInfo) String() string            { return proto.CompactTextString(m) }
func (*PezGame_SendGameInfo) ProtoMessage()               {}
func (*PezGame_SendGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{7} }

func (m *PezGame_SendGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PezGame_SendGameInfo) GetPlayerInfo() []*PezBase_PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *PezGame_SendGameInfo) GetDeskGameInfo() *PezBase_DeskGameInfo {
	if m != nil {
		return m.DeskGameInfo
	}
	return nil
}

func (m *PezGame_SendGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *PezGame_SendGameInfo) GetIsReconnect() PEZ_RECONNECT_TYPE {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return PEZ_RECONNECT_TYPE_PEZ_NORMAL
}

func init() {
	proto.RegisterType((*PezUserCoinBean)(nil), "ddproto.pez_user_coin_bean")
	proto.RegisterType((*PezGame_Opening)(nil), "ddproto.pez_game_Opening")
	proto.RegisterType((*PezGame_DealCards)(nil), "ddproto.pez_game_DealCards")
	proto.RegisterType((*PezGame_Bet)(nil), "ddproto.pez_game_Bet")
	proto.RegisterType((*PezGame_AckBet)(nil), "ddproto.pez_game_AckBet")
	proto.RegisterType((*PezGame_OpenCard)(nil), "ddproto.pez_game_OpenCard")
	proto.RegisterType((*PezGame_AckOpenCard)(nil), "ddproto.pez_game_AckOpenCard")
	proto.RegisterType((*PezGame_SendGameInfo)(nil), "ddproto.pez_game_SendGameInfo")
	proto.RegisterEnum("ddproto.PEZ_RECONNECT_TYPE", PEZ_RECONNECT_TYPE_name, PEZ_RECONNECT_TYPE_value)
}

var fileDescriptor38 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xce, 0x4f, 0xab, 0x49, 0x53, 0xda, 0xa5, 0x54, 0x26, 0x20, 0x14, 0x59, 0x1c, 0x22,
	0x04, 0x41, 0xf8, 0xc2, 0x05, 0x24, 0x12, 0x37, 0xa2, 0x91, 0x20, 0xb1, 0x96, 0x70, 0x28, 0x17,
	0x6b, 0x63, 0x0f, 0xc5, 0xaa, 0xbd, 0xb6, 0xfc, 0x73, 0x08, 0x8f, 0xc4, 0x8d, 0x37, 0xe0, 0x25,
	0x78, 0x1f, 0xb4, 0xeb, 0x1f, 0x6c, 0x42, 0xa9, 0xd2, 0x4b, 0xe2, 0x6f, 0xf6, 0x9b, 0xd9, 0xf9,
	0xbe, 0xd9, 0x81, 0xc3, 0x08, 0xbf, 0xd9, 0x91, 0xcf, 0x36, 0xe3, 0x28, 0x0e, 0xd3, 0x90, 0xec,
	0xb9, 0xae, 0xfc, 0x18, 0xdc, 0x73, 0xc2, 0x20, 0x08, 0xb9, 0xed, 0xf8, 0x1e, 0xf2, 0x34, 0x3f,
	0x1d, 0x48, 0xf6, 0x9a, 0x25, 0x58, 0xe0, 0xe3, 0x82, 0x84, 0x3c, 0x0b, 0xf2, 0x90, 0xfe, 0x16,
	0x88, 0x20, 0x65, 0x09, 0xc6, 0xb6, 0x13, 0x7a, 0xdc, 0x5e, 0x23, 0xe3, 0xe4, 0x14, 0xba, 0x22,
	0x32, 0x77, 0x35, 0x65, 0xa8, 0x8c, 0xfa, 0xb4, 0x40, 0x84, 0x40, 0x5b, 0x90, 0x34, 0x75, 0xa8,
	0x8c, 0x5a, 0x54, 0x7e, 0xeb, 0xbf, 0x14, 0x38, 0x12, 0x25, 0x2e, 0x59, 0x80, 0xf6, 0x32, 0x42,
	0xee, 0xf1, 0x4b, 0xf2, 0x0c, 0xba, 0x5f, 0x91, 0xb9, 0x18, 0xcb, 0x02, 0x3d, 0xe3, 0x64, 0x5c,
	0x34, 0x3a, 0xb6, 0xc4, 0xef, 0xb9, 0x3c, 0xa3, 0x05, 0x87, 0x3c, 0x81, 0xbe, 0x99, 0xc5, 0xb1,
	0xe5, 0xb3, 0x8d, 0x19, 0x66, 0x3c, 0x95, 0xf5, 0x3b, 0xb4, 0x19, 0x24, 0x27, 0xd0, 0x71, 0x3d,
	0x07, 0x5f, 0x6a, 0x2d, 0x79, 0x9a, 0x83, 0x32, 0x6a, 0x68, 0xed, 0x3f, 0x51, 0x83, 0x4c, 0xa0,
	0x2f, 0x5a, 0x36, 0x43, 0x8f, 0x4f, 0x91, 0xf1, 0x44, 0xeb, 0x0c, 0x5b, 0xa3, 0x9e, 0xf1, 0xb0,
	0x6a, 0x63, 0x5b, 0x34, 0x6d, 0x66, 0xe8, 0x3f, 0x95, 0xdc, 0x1a, 0xa9, 0xeb, 0x0c, 0x99, 0x6f,
	0xb2, 0xd8, 0x4d, 0x76, 0x54, 0xf6, 0x1a, 0x40, 0x4c, 0x0b, 0x63, 0x91, 0xac, 0xa9, 0xb2, 0x89,
	0x47, 0x8d, 0x26, 0xc4, 0x78, 0x6c, 0xab, 0xe2, 0xd0, 0x1a, 0x9f, 0x0c, 0x60, 0x3f, 0x62, 0x5e,
	0x6e, 0x49, 0x2e, 0xba, 0xc2, 0x44, 0x87, 0x03, 0x17, 0x99, 0x8f, 0xf1, 0xa7, 0x7c, 0x50, 0x6d,
	0x39, 0xa8, 0x46, 0x4c, 0xa7, 0x70, 0x50, 0x29, 0x98, 0x62, 0x7a, 0xed, 0x58, 0x4f, 0xa1, 0xbb,
	0xc6, 0x74, 0x91, 0x05, 0x85, 0xf1, 0x05, 0x12, 0xe3, 0x4e, 0xbd, 0x00, 0x8b, 0xbb, 0xe5, 0xb7,
	0x7e, 0x01, 0x77, 0xab, 0x9a, 0x13, 0xe7, 0xea, 0x7f, 0x65, 0x07, 0xb0, 0xbf, 0xc6, 0xb4, 0x3e,
	0xd1, 0x0a, 0x5f, 0x53, 0xfa, 0xb8, 0xf1, 0x90, 0xa4, 0x07, 0xbb, 0xf9, 0x7d, 0x0a, 0x5d, 0x87,
	0xc5, 0xee, 0xdc, 0x2d, 0x95, 0xe4, 0x48, 0xff, 0xa1, 0xc2, 0x49, 0xbd, 0xed, 0xdb, 0x97, 0x8f,
	0x31, 0xc9, 0xfc, 0x52, 0x4f, 0x81, 0x88, 0x06, 0x7b, 0x11, 0xf3, 0x56, 0x9b, 0xa8, 0x14, 0x54,
	0xc2, 0x9a, 0x37, 0xed, 0x86, 0x37, 0xcf, 0xa1, 0x2d, 0x5a, 0xd3, 0x3a, 0xf2, 0xd6, 0x07, 0xff,
	0x78, 0x12, 0xcc, 0x9b, 0xf3, 0x2f, 0x21, 0x95, 0x34, 0xf2, 0x02, 0x3a, 0xe2, 0xdf, 0xd0, 0xba,
	0x37, 0xf1, 0x73, 0xde, 0xf6, 0x02, 0xec, 0xed, 0xbc, 0x00, 0xdf, 0x55, 0xb8, 0x5f, 0x79, 0xf6,
	0x11, 0xb9, 0xfb, 0x8e, 0x05, 0x28, 0xee, 0xb8, 0xed, 0x0e, 0x88, 0xdc, 0x9b, 0x76, 0x40, 0x6a,
	0xa8, 0xf1, 0xc9, 0x54, 0xbc, 0xf3, 0xe4, 0xaa, 0xbc, 0x5b, 0xfa, 0xdb, 0x33, 0x1e, 0x6f, 0xe7,
	0x9f, 0xd5, 0x58, 0xb4, 0x91, 0x23, 0x76, 0x25, 0x41, 0xee, 0xfe, 0xbd, 0x2b, 0xf5, 0x18, 0x79,
	0x03, 0x3d, 0x2f, 0xa1, 0xe8, 0x84, 0x9c, 0xa3, 0x93, 0xca, 0xb9, 0x1c, 0xd6, 0xec, 0xb2, 0x66,
	0x9f, 0x6d, 0x3a, 0x33, 0x97, 0x8b, 0xc5, 0xcc, 0x5c, 0xd9, 0xab, 0x0b, 0x6b, 0x46, 0xeb, 0xfc,
	0xa7, 0xaf, 0x80, 0x6c, 0x53, 0xc8, 0x21, 0x80, 0x88, 0x2e, 0x96, 0xf4, 0xc3, 0xe4, 0xfd, 0x91,
	0x42, 0x8e, 0xa1, 0xdf, 0x60, 0x1d, 0xa9, 0x53, 0xf5, 0xbc, 0x65, 0xdd, 0xb1, 0x14, 0x4b, 0xfd,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xba, 0x3f, 0xfe, 0xdc, 0x05, 0x00, 0x00,
}

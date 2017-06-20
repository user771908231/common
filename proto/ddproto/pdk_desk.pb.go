// Code generated by protoc-gen-go.
// source: pdk_desk.proto
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

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 房主解散房间(未开局)
type PdkReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkReqDissolveDesk) Reset()                    { *m = PdkReqDissolveDesk{} }
func (m *PdkReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkReqDissolveDesk) ProtoMessage()               {}
func (*PdkReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{0} }

func (m *PdkReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 解散房间回复
type PdkAckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=passWord" json:"passWord,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkAckDissolveDesk) Reset()                    { *m = PdkAckDissolveDesk{} }
func (m *PdkAckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkAckDissolveDesk) ProtoMessage()               {}
func (*PdkAckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{1} }

func (m *PdkAckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkAckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkAckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PdkAckDissolveDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

// 准备游戏
type PdkReqReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsShowHandPokers *bool        `protobuf:"varint,3,opt,name=isShowHandPokers" json:"isShowHandPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkReqReady) Reset()                    { *m = PdkReqReady{} }
func (m *PdkReqReady) String() string            { return proto.CompactTextString(m) }
func (*PdkReqReady) ProtoMessage()               {}
func (*PdkReqReady) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{2} }

func (m *PdkReqReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkReqReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkReqReady) GetIsShowHandPokers() bool {
	if m != nil && m.IsShowHandPokers != nil {
		return *m.IsShowHandPokers
	}
	return false
}

// 准备游戏的结果
type PdkAckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkAckReady) Reset()                    { *m = PdkAckReady{} }
func (m *PdkAckReady) String() string            { return proto.CompactTextString(m) }
func (*PdkAckReady) ProtoMessage()               {}
func (*PdkAckReady) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{3} }

func (m *PdkAckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkAckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *PdkAckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 赢牌信息：谁赢了多少
type PdkBaseWinCoinInfo struct {
	NickName         *string            `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	BaseValue        *int32             `protobuf:"varint,3,opt,name=baseValue" json:"baseValue,omitempty"`
	WinCoin          *int64             `protobuf:"varint,4,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64             `protobuf:"varint,5,opt,name=coin" json:"coin,omitempty"`
	IsDiZhu          *bool              `protobuf:"varint,6,opt,name=isDiZhu" json:"isDiZhu,omitempty"`
	Rate             *int32             `protobuf:"varint,7,opt,name=rate" json:"rate,omitempty"`
	Description      *string            `protobuf:"bytes,8,opt,name=description" json:"description,omitempty"`
	Bomb             *int32             `protobuf:"varint,9,opt,name=bomb" json:"bomb,omitempty"`
	IsSpring         *bool              `protobuf:"varint,10,opt,name=isSpring" json:"isSpring,omitempty"`
	HandPokers       []*ClientBasePoker `protobuf:"bytes,11,rep,name=handPokers" json:"handPokers,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *PdkBaseWinCoinInfo) Reset()                    { *m = PdkBaseWinCoinInfo{} }
func (m *PdkBaseWinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkBaseWinCoinInfo) ProtoMessage()               {}
func (*PdkBaseWinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{4} }

func (m *PdkBaseWinCoinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PdkBaseWinCoinInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetBaseValue() int32 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetIsDiZhu() bool {
	if m != nil && m.IsDiZhu != nil {
		return *m.IsDiZhu
	}
	return false
}

func (m *PdkBaseWinCoinInfo) GetRate() int32 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *PdkBaseWinCoinInfo) GetBomb() int32 {
	if m != nil && m.Bomb != nil {
		return *m.Bomb
	}
	return 0
}

func (m *PdkBaseWinCoinInfo) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

func (m *PdkBaseWinCoinInfo) GetHandPokers() []*ClientBasePoker {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

// 本局结果(广播)
type PdkBaCurrentResult struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WinCoinInfo      []*PdkBaseWinCoinInfo `protobuf:"bytes,2,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	RemainPais       []int32               `protobuf:"varint,3,rep,name=RemainPais" json:"RemainPais,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PdkBaCurrentResult) Reset()                    { *m = PdkBaCurrentResult{} }
func (m *PdkBaCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*PdkBaCurrentResult) ProtoMessage()               {}
func (*PdkBaCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{5} }

func (m *PdkBaCurrentResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkBaCurrentResult) GetWinCoinInfo() []*PdkBaseWinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

func (m *PdkBaCurrentResult) GetRemainPais() []int32 {
	if m != nil {
		return m.RemainPais
	}
	return nil
}

type PdkBaseEndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	BigWin           *bool   `protobuf:"varint,3,opt,name=bigWin" json:"bigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	MaxWinCoin       *int32  `protobuf:"varint,6,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	CountBomb        *int32  `protobuf:"varint,7,opt,name=countBomb" json:"countBomb,omitempty"`
	CountWin         *int32  `protobuf:"varint,8,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32  `protobuf:"varint,9,opt,name=countLose" json:"countLose,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PdkBaseEndLotteryInfo) Reset()                    { *m = PdkBaseEndLotteryInfo{} }
func (m *PdkBaseEndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkBaseEndLotteryInfo) ProtoMessage()               {}
func (*PdkBaseEndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{6} }

func (m *PdkBaseEndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PdkBaseEndLotteryInfo) GetBigWin() bool {
	if m != nil && m.BigWin != nil {
		return *m.BigWin
	}
	return false
}

func (m *PdkBaseEndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PdkBaseEndLotteryInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetMaxWinCoin() int32 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *PdkBaseEndLotteryInfo) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

// 牌局结束(广播)
type PdkBcEndLottery struct {
	Header           *ProtoHeader             `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CoinInfo         []*PdkBaseEndLotteryInfo `protobuf:"bytes,2,rep,name=coinInfo" json:"coinInfo,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *PdkBcEndLottery) Reset()                    { *m = PdkBcEndLottery{} }
func (m *PdkBcEndLottery) String() string            { return proto.CompactTextString(m) }
func (*PdkBcEndLottery) ProtoMessage()               {}
func (*PdkBcEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor37, []int{7} }

func (m *PdkBcEndLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkBcEndLottery) GetCoinInfo() []*PdkBaseEndLotteryInfo {
	if m != nil {
		return m.CoinInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PdkReqDissolveDesk)(nil), "ddproto.pdk_req_dissolveDesk")
	proto.RegisterType((*PdkAckDissolveDesk)(nil), "ddproto.pdk_ack_dissolveDesk")
	proto.RegisterType((*PdkReqReady)(nil), "ddproto.pdk_req_ready")
	proto.RegisterType((*PdkAckReady)(nil), "ddproto.pdk_ack_ready")
	proto.RegisterType((*PdkBaseWinCoinInfo)(nil), "ddproto.pdk_base_winCoinInfo")
	proto.RegisterType((*PdkBaCurrentResult)(nil), "ddproto.pdk_ba_currentResult")
	proto.RegisterType((*PdkBaseEndLotteryInfo)(nil), "ddproto.pdk_base_endLotteryInfo")
	proto.RegisterType((*PdkBcEndLottery)(nil), "ddproto.pdk_bc_endLottery")
}

var fileDescriptor37 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x71, 0x9d, 0x0f, 0x7b, 0x4c, 0x4b, 0xeb, 0x56, 0x62, 0x55, 0x09, 0x29, 0xb2, 0x38,
	0xf4, 0x94, 0x43, 0x1f, 0xa1, 0xf4, 0x90, 0x4a, 0x11, 0x44, 0x41, 0x22, 0x12, 0x17, 0x6b, 0xb3,
	0xbb, 0x24, 0x2b, 0xc7, 0xbb, 0x66, 0xd7, 0x26, 0xf4, 0xce, 0x4b, 0xf1, 0x52, 0x3c, 0x03, 0xfb,
	0x11, 0xbb, 0x89, 0xca, 0x21, 0x20, 0x2e, 0x51, 0x66, 0xe6, 0xbf, 0xff, 0x99, 0xf9, 0x8d, 0xe1,
	0xac, 0xa2, 0x45, 0x4e, 0x99, 0x2e, 0xc6, 0x95, 0x92, 0xb5, 0x4c, 0x87, 0x94, 0xba, 0x3f, 0xd7,
	0x97, 0x44, 0x96, 0xa5, 0x14, 0x39, 0xd9, 0x70, 0x26, 0x6a, 0x5f, 0xcd, 0xa6, 0x70, 0x65, 0xf5,
	0x8a, 0x7d, 0xcd, 0x29, 0xd7, 0x5a, 0x6e, 0xbe, 0xb1, 0x7b, 0xf3, 0x36, 0x7d, 0x0b, 0x83, 0x35,
	0xc3, 0x94, 0x29, 0x14, 0x8c, 0x82, 0x9b, 0xe4, 0xf6, 0x6a, 0xbc, 0xb3, 0x19, 0xcf, 0xec, 0xef,
	0xc4, 0xd5, 0xd2, 0x33, 0x18, 0x34, 0x9a, 0xa9, 0x07, 0x8a, 0x4e, 0x8c, 0xea, 0x34, 0x13, 0xde,
	0x0d, 0x93, 0xe2, 0x3f, 0xb8, 0xd9, 0xd8, 0xee, 0x61, 0xe2, 0xd0, 0xc4, 0xfd, 0xf4, 0x1c, 0xa2,
	0x0a, 0x6b, 0xbd, 0x90, 0x8a, 0xa2, 0x9e, 0xc9, 0xc4, 0x59, 0x0e, 0xa7, 0xed, 0xf4, 0xca, 0x78,
	0x3c, 0xfe, 0x63, 0x23, 0x04, 0xe7, 0x5c, 0x7f, 0x5c, 0xcb, 0xed, 0x04, 0x0b, 0x3a, 0x93, 0x05,
	0x53, 0xda, 0xb5, 0x8c, 0xb2, 0xb9, 0x6f, 0x60, 0x17, 0xfa, 0x9b, 0x06, 0x09, 0x84, 0xa5, 0x5e,
	0x39, 0xf7, 0x78, 0xaf, 0x5b, 0xe8, 0x20, 0xfd, 0x0a, 0x3c, 0xa5, 0x25, 0xd6, 0x2c, 0xdf, 0x72,
	0xf1, 0x4e, 0x72, 0xf1, 0x20, 0xbe, 0x48, 0xbb, 0x9f, 0xe0, 0xa4, 0x78, 0x8f, 0x4b, 0xe6, 0xdc,
	0xe3, 0x67, 0x83, 0x5e, 0x40, 0x6c, 0x5f, 0x7d, 0xc2, 0x9b, 0x86, 0xed, 0xa0, 0xbc, 0x82, 0xe1,
	0xce, 0xc3, 0x31, 0x09, 0xd3, 0x97, 0xd0, 0x23, 0x36, 0xea, 0xbb, 0xc8, 0x94, 0xb9, 0xbe, 0xe7,
	0x9f, 0xd7, 0x0d, 0x1a, 0xd8, 0x8d, 0x6c, 0x59, 0xe1, 0x9a, 0xa1, 0xa1, 0x7b, 0x7d, 0x09, 0x89,
	0x41, 0x4c, 0x14, 0xaf, 0x6a, 0x2e, 0x05, 0x8a, 0x5c, 0x57, 0x23, 0x59, 0xca, 0x72, 0x89, 0xe2,
	0x96, 0xba, 0x81, 0x53, 0x29, 0x2e, 0x56, 0x08, 0x9c, 0xc5, 0x18, 0x60, 0xfd, 0x04, 0x2a, 0x19,
	0x85, 0x86, 0xc3, 0x75, 0xc7, 0xc1, 0x7f, 0x5e, 0x7e, 0xbb, 0xca, 0x4a, 0xb2, 0x1f, 0xdd, 0xc2,
	0x39, 0x69, 0x94, 0x32, 0xd5, 0x39, 0xd3, 0xcd, 0xa6, 0x3e, 0x12, 0xe6, 0x2d, 0x24, 0x7b, 0x94,
	0x0c, 0x09, 0xdb, 0xef, 0x4d, 0x27, 0xfd, 0x23, 0xca, 0x14, 0x60, 0xce, 0x4a, 0xcc, 0xc5, 0x0c,
	0x73, 0x7b, 0xcb, 0xf0, 0xa6, 0x9f, 0xfd, 0x0c, 0xe0, 0x75, 0x27, 0x66, 0x82, 0x4e, 0x65, 0x5d,
	0x33, 0xf5, 0xe8, 0xf4, 0x4f, 0xa0, 0x03, 0x07, 0x7a, 0xff, 0x14, 0xdd, 0x15, 0x97, 0x7c, 0xb5,
	0x30, 0x60, 0xdd, 0x97, 0xe1, 0xc1, 0x7e, 0xd8, 0x0a, 0x33, 0x7c, 0xaf, 0x4d, 0xb4, 0x87, 0xf0,
	0xe8, 0xcd, 0x0c, 0x25, 0xfe, 0xbe, 0xd8, 0xe5, 0x06, 0x0e, 0xa6, 0x39, 0x20, 0x91, 0x8d, 0xa8,
	0xef, 0x2c, 0xdf, 0x61, 0xcb, 0xd7, 0xa5, 0xac, 0x75, 0x74, 0x20, 0x9a, 0x4a, 0xcd, 0xfc, 0x11,
	0xb2, 0x12, 0x2e, 0xdc, 0xe8, 0x64, 0x6f, 0xf0, 0xa3, 0xf1, 0x45, 0xe4, 0x90, 0xdd, 0xe8, 0x39,
	0xbb, 0x43, 0x1c, 0x77, 0x27, 0x93, 0x70, 0xf6, 0xe2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa,
	0xb6, 0xbe, 0x90, 0x4a, 0x04, 0x00, 0x00,
}

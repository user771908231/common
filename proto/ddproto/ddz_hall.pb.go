// Code generated by protoc-gen-go.
// source: ddz_hall.proto
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

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of ddz_base_roomTypeInfo from ddz_base.proto

// Ignoring public import of ddz_base_playerInfo from ddz_base.proto

// Ignoring public import of ddz_base_deskInfo from ddz_base.proto

// Ignoring public import of ddz_enum_protoId from ddz_base.proto

// Ignoring public import of ddz_enum_errorCode from ddz_base.proto

// Ignoring public import of ddz_enum_paiType from ddz_base.proto

// Ignoring public import of ddz_enum_actType from ddz_base.proto

// Ignoring public import of ddz_enum_gameStatus from ddz_base.proto

// Ignoring public import of ddz_enum_playerStatus from ddz_base.proto

// Ignoring public import of ddz_enum_roomType from ddz_base.proto

// Ignoring public import of ddz_enum_deskGameStatus from ddz_base.proto

// 创建房间
type DdzReqCreateDesk struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *DdzBaseRoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzReqCreateDesk) Reset()                    { *m = DdzReqCreateDesk{} }
func (m *DdzReqCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzReqCreateDesk) ProtoMessage()               {}
func (*DdzReqCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *DdzReqCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqCreateDesk) GetRoomTypeInfo() *DdzBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 创建房间回复的信息
type DdzAckCreateDesk struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32               `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string              `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserBalance      *int64               `protobuf:"varint,4,opt,name=userBalance" json:"userBalance,omitempty"`
	CreateFee        *int64               `protobuf:"varint,5,opt,name=createFee" json:"createFee,omitempty"`
	RoomTypeInfo     *DdzBaseRoomTypeInfo `protobuf:"bytes,6,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzAckCreateDesk) Reset()                    { *m = DdzAckCreateDesk{} }
func (m *DdzAckCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzAckCreateDesk) ProtoMessage()               {}
func (*DdzAckCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *DdzAckCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckCreateDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzAckCreateDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *DdzAckCreateDesk) GetUserBalance() int64 {
	if m != nil && m.UserBalance != nil {
		return *m.UserBalance
	}
	return 0
}

func (m *DdzAckCreateDesk) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *DdzAckCreateDesk) GetRoomTypeInfo() *DdzBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 游戏战绩
type DdzReqGameRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32       `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqGameRecord) Reset()                    { *m = DdzReqGameRecord{} }
func (m *DdzReqGameRecord) String() string            { return proto.CompactTextString(m) }
func (*DdzReqGameRecord) ProtoMessage()               {}
func (*DdzReqGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *DdzReqGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *DdzReqGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

//
type DdzBaseUserRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=NickName" json:"NickName,omitempty"`
	WinAmount        *int64       `protobuf:"varint,4,opt,name=WinAmount" json:"WinAmount,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzBaseUserRecord) Reset()                    { *m = DdzBaseUserRecord{} }
func (m *DdzBaseUserRecord) String() string            { return proto.CompactTextString(m) }
func (*DdzBaseUserRecord) ProtoMessage()               {}
func (*DdzBaseUserRecord) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *DdzBaseUserRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBaseUserRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzBaseUserRecord) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *DdzBaseUserRecord) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

//
type DdzBaseGameRecord struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32               `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	DeskId           *int32               `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	BeginTime        *string              `protobuf:"bytes,4,opt,name=beginTime" json:"beginTime,omitempty"`
	Users            []*DdzBaseUserRecord `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzBaseGameRecord) Reset()                    { *m = DdzBaseGameRecord{} }
func (m *DdzBaseGameRecord) String() string            { return proto.CompactTextString(m) }
func (*DdzBaseGameRecord) ProtoMessage()               {}
func (*DdzBaseGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *DdzBaseGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBaseGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *DdzBaseGameRecord) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzBaseGameRecord) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *DdzBaseGameRecord) GetUsers() []*DdzBaseUserRecord {
	if m != nil {
		return m.Users
	}
	return nil
}

//
type DdzAckGameRecord struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Records          []*DdzBaseGameRecord `protobuf:"bytes,3,rep,name=records" json:"records,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DdzAckGameRecord) Reset()                    { *m = DdzAckGameRecord{} }
func (m *DdzAckGameRecord) String() string            { return proto.CompactTextString(m) }
func (*DdzAckGameRecord) ProtoMessage()               {}
func (*DdzAckGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *DdzAckGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckGameRecord) GetRecords() []*DdzBaseGameRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

// 客户端请求进入 room, 服务器返回DdzSendGameInfo
type DdzReqEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MatchId          *int32       `protobuf:"varint,2,opt,name=matchId" json:"matchId,omitempty"`
	TableId          *int32       `protobuf:"varint,3,opt,name=tableId" json:"tableId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=PassWord" json:"PassWord,omitempty"`
	UserId           *uint32      `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqEnterDesk) Reset()                    { *m = DdzReqEnterDesk{} }
func (m *DdzReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzReqEnterDesk) ProtoMessage()               {}
func (*DdzReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *DdzReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqEnterDesk) GetMatchId() int32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *DdzReqEnterDesk) GetTableId() int32 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *DdzReqEnterDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

func (m *DdzReqEnterDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type DdzAckEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzAckEnterDesk) Reset()                    { *m = DdzAckEnterDesk{} }
func (m *DdzAckEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzAckEnterDesk) ProtoMessage()               {}
func (*DdzAckEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *DdzAckEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*DdzReqCreateDesk)(nil), "ddproto.ddz_req_createDesk")
	proto.RegisterType((*DdzAckCreateDesk)(nil), "ddproto.ddz_ack_createDesk")
	proto.RegisterType((*DdzReqGameRecord)(nil), "ddproto.ddz_req_gameRecord")
	proto.RegisterType((*DdzBaseUserRecord)(nil), "ddproto.ddz_base_userRecord")
	proto.RegisterType((*DdzBaseGameRecord)(nil), "ddproto.ddz_base_gameRecord")
	proto.RegisterType((*DdzAckGameRecord)(nil), "ddproto.ddz_ack_gameRecord")
	proto.RegisterType((*DdzReqEnterDesk)(nil), "ddproto.ddz_req_enterDesk")
	proto.RegisterType((*DdzAckEnterDesk)(nil), "ddproto.ddz_ack_enterDesk")
}

var fileDescriptor7 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x25, 0x09, 0x69, 0xe9, 0x14, 0x56, 0xc2, 0x8b, 0x56, 0x51, 0x55, 0xa1, 0x28, 0xa7, 0x1e,
	0x50, 0x0f, 0x3d, 0x70, 0xdf, 0x0a, 0xa1, 0xed, 0x65, 0x55, 0x59, 0x8b, 0xf6, 0x58, 0xb9, 0xf1,
	0xb0, 0x0d, 0x49, 0xec, 0x62, 0x67, 0x85, 0xe0, 0x03, 0xf8, 0x00, 0x24, 0x3e, 0x82, 0x8f, 0xe0,
	0xc7, 0x38, 0x21, 0xc7, 0x71, 0x92, 0x6a, 0xc5, 0x81, 0x68, 0x2f, 0x51, 0xde, 0xcc, 0x78, 0xfc,
	0xde, 0x9b, 0x31, 0x9c, 0x71, 0xfe, 0x6d, 0x77, 0x60, 0x45, 0xb1, 0x3c, 0x2a, 0x59, 0x49, 0x32,
	0xe6, 0xbc, 0xfe, 0x99, 0x9d, 0xa7, 0xb2, 0x2c, 0xa5, 0xd8, 0xa5, 0x45, 0x86, 0xa2, 0xb2, 0xd9,
	0x59, 0x5d, 0xbd, 0x67, 0x1a, 0x2d, 0x4e, 0xbe, 0x7b, 0x40, 0x4c, 0x48, 0xe1, 0xe7, 0x5d, 0xaa,
	0x90, 0x55, 0xf8, 0x0e, 0x75, 0x4e, 0xde, 0xc0, 0xe8, 0x80, 0x8c, 0xa3, 0x8a, 0xbc, 0xd8, 0x5b,
	0x4c, 0x57, 0xaf, 0x96, 0x4d, 0xd7, 0xe5, 0xd6, 0x7c, 0xaf, 0xea, 0x1c, 0x6d, 0x6a, 0xc8, 0x1a,
	0x9e, 0x2b, 0x29, 0xcb, 0x9b, 0xaf, 0x47, 0xdc, 0x88, 0x8f, 0x32, 0xf2, 0xeb, 0x33, 0xaf, 0xdb,
	0x33, 0xee, 0xce, 0x5d, 0xbf, 0x8a, 0x9e, 0x9c, 0x49, 0xfe, 0x34, 0x44, 0x58, 0x9a, 0x0f, 0x27,
	0x72, 0x01, 0x23, 0x8e, 0x3a, 0xdf, 0xf0, 0x9a, 0x42, 0x48, 0x1b, 0x44, 0x66, 0xf0, 0xec, 0xc8,
	0xb4, 0xfe, 0x22, 0x15, 0x8f, 0x82, 0xd8, 0x5b, 0x4c, 0x68, 0x8b, 0x49, 0x0c, 0xd3, 0x7b, 0x8d,
	0x6a, 0xcd, 0x0a, 0x26, 0x52, 0x8c, 0x9e, 0xc6, 0xde, 0x22, 0xa0, 0xfd, 0x10, 0x99, 0xc3, 0xc4,
	0x32, 0x7a, 0x8f, 0x18, 0x85, 0x75, 0xbe, 0x0b, 0x3c, 0x10, 0x3f, 0x1a, 0x20, 0xfe, 0x53, 0x37,
	0x84, 0x3b, 0x56, 0x22, 0xc5, 0xd4, 0x30, 0xfb, 0x3f, 0xed, 0x67, 0xe0, 0x67, 0x4e, 0xb7, 0x9f,
	0x71, 0xe3, 0x85, 0x11, 0xb1, 0xb1, 0x8a, 0x5f, 0xd0, 0x06, 0x25, 0x3f, 0x3d, 0x38, 0x6f, 0x39,
	0x99, 0xd8, 0xa0, 0xdb, 0x2e, 0x60, 0xf4, 0xc1, 0x76, 0xf7, 0x6d, 0x77, 0x8b, 0x8c, 0xd3, 0xd7,
	0x59, 0x9a, 0x5f, 0xb3, 0x12, 0x9d, 0xd3, 0x0e, 0x1b, 0x1f, 0x6f, 0x33, 0x71, 0x59, 0xca, 0x7b,
	0x51, 0x35, 0x3e, 0x77, 0x81, 0xe4, 0x77, 0x9f, 0xd7, 0x63, 0xba, 0xd0, 0x6c, 0x44, 0x70, 0xb2,
	0x11, 0x73, 0x98, 0xec, 0xf1, 0x2e, 0x13, 0x37, 0x59, 0x69, 0x67, 0x3e, 0xa1, 0x5d, 0x80, 0xac,
	0x20, 0x34, 0xce, 0xe8, 0x28, 0x8c, 0x83, 0xc5, 0x74, 0x35, 0x7f, 0x38, 0xcc, 0xce, 0x38, 0x6a,
	0x4b, 0x93, 0x1f, 0xbd, 0x05, 0x1e, 0x4c, 0xbf, 0x1b, 0x9a, 0xdf, 0x1f, 0x1a, 0x79, 0x0b, 0x63,
	0x55, 0xf7, 0xd3, 0x51, 0xf0, 0x2f, 0x4a, 0xdd, 0xa5, 0xd4, 0x15, 0x27, 0xbf, 0x3c, 0x78, 0xe9,
	0x36, 0x0b, 0x45, 0x85, 0x6a, 0xc0, 0xa3, 0x8a, 0x60, 0x5c, 0xb2, 0x2a, 0x3d, 0xb4, 0xaf, 0xca,
	0x41, 0x93, 0xa9, 0xd8, 0xbe, 0xc0, 0xd6, 0x5d, 0x07, 0xcd, 0x1a, 0x6c, 0x99, 0xd6, 0xb7, 0xe6,
	0xc1, 0x59, 0x77, 0x5b, 0xdc, 0xd3, 0x18, 0x9e, 0x2c, 0xe6, 0xa5, 0xa5, 0x6a, 0xfc, 0x1b, 0x48,
	0x75, 0xed, 0x5f, 0x05, 0xdb, 0x27, 0x5b, 0xef, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x38,
	0x5a, 0x84, 0x14, 0x05, 0x00, 0x00,
}

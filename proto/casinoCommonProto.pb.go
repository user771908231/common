// Code generated by protoc-gen-go.
// source: casinoCommonProto.proto
// DO NOT EDIT!

/*
Package casinoCommonProto is a generated protocol buffer package.

It is generated from these files:
	casinoCommonProto.proto

It has these top-level messages:
	ProtoHeader
	Heartbeat
	User
	TNotice
	GameSession
	WeixinInfo
	Game_Login
*/
package casinoCommonProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ProtoHeader 需要在每个 Message 中作为第一个字段
type ProtoHeader struct {
	Version          *string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Code             *int32  `protobuf:"varint,3,opt,name=code" json:"code,omitempty"`
	Error            *string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoHeader) Reset()                    { *m = ProtoHeader{} }
func (m *ProtoHeader) String() string            { return proto.CompactTextString(m) }
func (*ProtoHeader) ProtoMessage()               {}
func (*ProtoHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoHeader) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ProtoHeader) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ProtoHeader) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ProtoHeader) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// 这里都是通用过的一些协议
type Heartbeat struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Heartbeat) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type User struct {
	Mid                *string `protobuf:"bytes,1,opt,name=mid" json:"mid,omitempty"`
	Pwd                *string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Coin               *int64  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                 *uint32 `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName           *string `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	LoginTurntable     *bool   `protobuf:"varint,6,opt,name=loginTurntable" json:"loginTurntable,omitempty"`
	LoginTurntableTime *string `protobuf:"bytes,7,opt,name=loginTurntableTime" json:"loginTurntableTime,omitempty"`
	Scores             *int32  `protobuf:"varint,8,opt,name=scores" json:"scores,omitempty"`
	LastSignTime       *string `protobuf:"bytes,9,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignCount          *int32  `protobuf:"varint,10,opt,name=signCount" json:"signCount,omitempty"`
	Diamond            *int64  `protobuf:"varint,11,opt,name=Diamond" json:"Diamond,omitempty"`
	OpenId             *string `protobuf:"bytes,12,opt,name=openId" json:"openId,omitempty"`
	HeadUrl            *string `protobuf:"bytes,13,opt,name=headUrl" json:"headUrl,omitempty"`
	City               *string `protobuf:"bytes,14,opt,name=city" json:"city,omitempty"`
	Sex                *int32  `protobuf:"varint,15,opt,name=sex" json:"sex,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetMid() string {
	if m != nil && m.Mid != nil {
		return *m.Mid
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *User) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *User) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *User) GetLoginTurntable() bool {
	if m != nil && m.LoginTurntable != nil {
		return *m.LoginTurntable
	}
	return false
}

func (m *User) GetLoginTurntableTime() string {
	if m != nil && m.LoginTurntableTime != nil {
		return *m.LoginTurntableTime
	}
	return ""
}

func (m *User) GetScores() int32 {
	if m != nil && m.Scores != nil {
		return *m.Scores
	}
	return 0
}

func (m *User) GetLastSignTime() string {
	if m != nil && m.LastSignTime != nil {
		return *m.LastSignTime
	}
	return ""
}

func (m *User) GetSignCount() int32 {
	if m != nil && m.SignCount != nil {
		return *m.SignCount
	}
	return 0
}

func (m *User) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *User) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *User) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

type TNotice struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NoticeType       *int32   `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string  `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string  `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string  `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Fileds           []string `protobuf:"bytes,6,rep,name=fileds" json:"fileds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TNotice) Reset()                    { *m = TNotice{} }
func (m *TNotice) String() string            { return proto.CompactTextString(m) }
func (*TNotice) ProtoMessage()               {}
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *TNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *TNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *TNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *TNotice) GetFileds() []string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

// 用户的seesion 主要是游戏的
type GameSession struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	RoomId           *int32  `protobuf:"varint,2,opt,name=RoomId" json:"RoomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,3,opt,name=DeskId" json:"DeskId,omitempty"`
	GameStatus       *int32  `protobuf:"varint,4,opt,name=GameStatus" json:"GameStatus,omitempty"`
	GameId           *int32  `protobuf:"varint,5,opt,name=GameId" json:"GameId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,6,opt,name=GameNumber" json:"GameNumber,omitempty"`
	GameCustomStatus *int32  `protobuf:"varint,7,opt,name=GameCustomStatus" json:"GameCustomStatus,omitempty"`
	IsBreak          *bool   `protobuf:"varint,8,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,9,opt,name=isLeave" json:"isLeave,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameSession) Reset()                    { *m = GameSession{} }
func (m *GameSession) String() string            { return proto.CompactTextString(m) }
func (*GameSession) ProtoMessage()               {}
func (*GameSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GameSession) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *GameSession) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *GameSession) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *GameSession) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *GameSession) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameSession) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *GameSession) GetGameCustomStatus() int32 {
	if m != nil && m.GameCustomStatus != nil {
		return *m.GameCustomStatus
	}
	return 0
}

func (m *GameSession) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *GameSession) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

// 微信信息
type WeixinInfo struct {
	OpenId           *string `protobuf:"bytes,1,opt,name=openId" json:"openId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	HeadUrl          *string `protobuf:"bytes,3,opt,name=headUrl" json:"headUrl,omitempty"`
	Sex              *int32  `protobuf:"varint,4,opt,name=sex" json:"sex,omitempty"`
	City             *string `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	UnionId          *string `protobuf:"bytes,6,opt,name=unionId" json:"unionId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WeixinInfo) Reset()                    { *m = WeixinInfo{} }
func (m *WeixinInfo) String() string            { return proto.CompactTextString(m) }
func (*WeixinInfo) ProtoMessage()               {}
func (*WeixinInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WeixinInfo) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *WeixinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *WeixinInfo) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *WeixinInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *WeixinInfo) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *WeixinInfo) GetUnionId() string {
	if m != nil && m.UnionId != nil {
		return *m.UnionId
	}
	return ""
}

// 登录的结构
type Game_Login struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	ProtoVersion     *int32       `protobuf:"varint,3,opt,name=protoVersion" json:"protoVersion,omitempty"`
	WxInfo           *WeixinInfo  `protobuf:"bytes,4,opt,name=wxInfo" json:"wxInfo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_Login) Reset()                    { *m = Game_Login{} }
func (m *Game_Login) String() string            { return proto.CompactTextString(m) }
func (*Game_Login) ProtoMessage()               {}
func (*Game_Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Game_Login) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_Login) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_Login) GetProtoVersion() int32 {
	if m != nil && m.ProtoVersion != nil {
		return *m.ProtoVersion
	}
	return 0
}

func (m *Game_Login) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoHeader)(nil), "casinoCommonProto.ProtoHeader")
	proto.RegisterType((*Heartbeat)(nil), "casinoCommonProto.Heartbeat")
	proto.RegisterType((*User)(nil), "casinoCommonProto.User")
	proto.RegisterType((*TNotice)(nil), "casinoCommonProto.TNotice")
	proto.RegisterType((*GameSession)(nil), "casinoCommonProto.GameSession")
	proto.RegisterType((*WeixinInfo)(nil), "casinoCommonProto.WeixinInfo")
	proto.RegisterType((*Game_Login)(nil), "casinoCommonProto.game_Login")
}

var fileDescriptor0 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0xd8, 0x89, 0xc7, 0x49, 0xda, 0x2e, 0x5f, 0x2b, 0x24, 0x10, 0xca, 0x89, 0x0b,
	0x39, 0xf4, 0xca, 0x8d, 0x54, 0x82, 0x8a, 0x12, 0x21, 0x9a, 0xc0, 0x11, 0x39, 0xf1, 0x36, 0xac,
	0x62, 0xef, 0x46, 0xbb, 0xeb, 0x36, 0xf9, 0x03, 0xdc, 0xf9, 0x29, 0xfc, 0x43, 0x66, 0xc7, 0x36,
	0x0d, 0x84, 0x0b, 0x97, 0x28, 0xf3, 0xfc, 0xe6, 0xeb, 0xcd, 0x5b, 0x78, 0xb2, 0xca, 0xac, 0x54,
	0x7a, 0xaa, 0xcb, 0x52, 0xab, 0x8f, 0x46, 0x3b, 0x3d, 0xd9, 0xfa, 0x5f, 0x76, 0x76, 0xf4, 0x61,
	0xfc, 0x1e, 0x52, 0xfa, 0xf3, 0x4e, 0x64, 0xb9, 0x30, 0xec, 0x04, 0x7a, 0xb7, 0xc2, 0x58, 0xa9,
	0x15, 0x0f, 0x5e, 0x04, 0x2f, 0x13, 0x36, 0x82, 0xb8, 0xb2, 0xc2, 0x5c, 0xe6, 0xbc, 0x83, 0xf1,
	0x90, 0x0d, 0xa0, 0xbb, 0xd2, 0xb9, 0xe0, 0x21, 0x46, 0x11, 0x1b, 0x42, 0x24, 0x8c, 0xd1, 0x86,
	0x77, 0x3d, 0x79, 0xfc, 0x1a, 0x12, 0xac, 0x63, 0xdc, 0x52, 0x64, 0x8e, 0x4d, 0x20, 0xfe, 0x46,
	0x45, 0xa9, 0x52, 0x7a, 0xfe, 0x7c, 0x72, 0x3c, 0xd6, 0x41, 0xeb, 0xf1, 0xf7, 0x0e, 0x74, 0x17,
	0xd8, 0x8a, 0xa5, 0x10, 0x96, 0x32, 0x6f, 0xfa, 0x63, 0xb0, 0xbd, 0xab, 0x9b, 0x27, 0x75, 0x73,
	0xa9, 0xa8, 0x79, 0xc8, 0x00, 0x3a, 0x48, 0xeb, 0xd2, 0x58, 0xa7, 0xd0, 0x57, 0x72, 0xb5, 0x99,
	0x65, 0xa5, 0xe0, 0x11, 0x71, 0x1f, 0xc3, 0xa8, 0xd0, 0x6b, 0xa9, 0xe6, 0x95, 0x51, 0x2e, 0x5b,
	0x16, 0x82, 0xc7, 0x88, 0xf7, 0xd9, 0x53, 0x60, 0x7f, 0xe2, 0x73, 0x89, 0x39, 0xbd, 0x76, 0x59,
	0xbb, 0xd2, 0x46, 0x58, 0xde, 0xa7, 0xf5, 0x1e, 0xc2, 0xa0, 0xc8, 0xac, 0xbb, 0x96, 0x6b, 0x45,
	0xac, 0x84, 0x58, 0x67, 0x90, 0x58, 0x44, 0xa6, 0xba, 0x52, 0x8e, 0x03, 0x11, 0x51, 0xb6, 0x0b,
	0x99, 0xe1, 0x5a, 0x39, 0x4f, 0x69, 0x36, 0xac, 0xa4, 0xb7, 0x42, 0xa1, 0x6c, 0x03, 0xca, 0x41,
	0x82, 0x17, 0x63, 0x61, 0x0a, 0x3e, 0xfc, 0xbd, 0x8a, 0x74, 0x7b, 0x3e, 0x6a, 0xb7, 0xb4, 0x62,
	0xc7, 0x4f, 0x7c, 0xb1, 0xf1, 0x1e, 0x7a, 0xf3, 0x99, 0x76, 0x72, 0x25, 0x9a, 0x15, 0x03, 0xea,
	0x81, 0x81, 0x22, 0x74, 0xbe, 0xdf, 0x0a, 0x12, 0x24, 0x62, 0x0f, 0x20, 0x6d, 0x30, 0xe9, 0x8a,
	0xfa, 0x28, 0x09, 0x7b, 0x04, 0xc3, 0x1a, 0x9c, 0x6a, 0xe5, 0x04, 0xce, 0x48, 0xc7, 0xb9, 0xcf,
	0xff, 0x20, 0x4a, 0xdd, 0x88, 0x84, 0x63, 0xde, 0xc8, 0x42, 0xe4, 0x16, 0xc5, 0x09, 0xf1, 0x80,
	0x3f, 0x03, 0x48, 0xdf, 0xa2, 0x86, 0xd7, 0xc2, 0x7a, 0x0f, 0xf8, 0xef, 0x8b, 0xfa, 0xfa, 0x01,
	0xc9, 0x8c, 0xf1, 0x27, 0xad, 0xcb, 0xc6, 0x0d, 0x91, 0x8f, 0x2f, 0x84, 0xdd, 0x60, 0x1c, 0xb6,
	0x33, 0x52, 0xba, 0xcb, 0x5c, 0x65, 0xa9, 0x2f, 0x71, 0x3c, 0x86, 0x9c, 0xe8, 0x90, 0x33, 0xab,
	0xca, 0x25, 0x7a, 0x23, 0x26, 0x8c, 0xc3, 0xa9, 0xc7, 0xa6, 0x95, 0x75, 0xba, 0x6c, 0xb2, 0x7b,
	0xad, 0xb2, 0xd2, 0xbe, 0x31, 0x22, 0xdb, 0xd0, 0x4d, 0xfa, 0x35, 0x70, 0x25, 0xb2, 0xdb, 0xfa,
	0x1c, 0xfd, 0xf1, 0x06, 0xe0, 0x8b, 0x90, 0x3b, 0xa9, 0x2e, 0xd5, 0x8d, 0x3e, 0x10, 0xbe, 0xf6,
	0xcf, 0xa1, 0x31, 0x3a, 0x7f, 0x9f, 0x22, 0x3c, 0x14, 0xbf, 0x9e, 0xb6, 0xbd, 0x4b, 0xd4, 0x72,
	0x2b, 0x85, 0x42, 0x60, 0xb9, 0x98, 0x1c, 0xfe, 0x23, 0x00, 0x58, 0x63, 0xad, 0xaf, 0x57, 0xde,
	0x43, 0xff, 0xeb, 0xf1, 0xa3, 0xd7, 0x84, 0x06, 0xa3, 0x97, 0xf9, 0xb9, 0x79, 0x73, 0xb5, 0x8a,
	0xaf, 0x20, 0xbe, 0xdb, 0xf9, 0x6d, 0x68, 0xa6, 0xf4, 0xfc, 0xd9, 0x3f, 0xaa, 0xde, 0xaf, 0xfc,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x93, 0x0b, 0x95, 0xef, 0x03, 0x00, 0x00,
}

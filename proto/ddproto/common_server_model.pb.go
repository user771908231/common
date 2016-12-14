// Code generated by protoc-gen-go.
// source: common_server_model.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	RoomCard           *int64  `protobuf:"varint,1,opt,name=RoomCard" json:"RoomCard,omitempty"`
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
	Diamond2           *int64  `protobuf:"varint,12,opt,name=Diamond2" json:"Diamond2,omitempty"`
	OpenId             *string `protobuf:"bytes,13,opt,name=openId" json:"openId,omitempty"`
	UnionId            *string `protobuf:"bytes,14,opt,name=UnionId" json:"UnionId,omitempty"`
	HeadUrl            *string `protobuf:"bytes,15,opt,name=headUrl" json:"headUrl,omitempty"`
	City               *string `protobuf:"bytes,16,opt,name=city" json:"city,omitempty"`
	Sex                *int32  `protobuf:"varint,17,opt,name=sex" json:"sex,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *User) GetRoomCard() int64 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
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

func (m *User) GetDiamond2() int64 {
	if m != nil && m.Diamond2 != nil {
		return *m.Diamond2
	}
	return 0
}

func (m *User) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *User) GetUnionId() string {
	if m != nil && m.UnionId != nil {
		return *m.UnionId
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
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

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

func init() {
	proto.RegisterType((*User)(nil), "ddproto.User")
	proto.RegisterType((*TNotice)(nil), "ddproto.TNotice")
}

var fileDescriptor2 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0xa6, 0xff, 0x32, 0xbb, 0x2d, 0x8b, 0x0f, 0x68, 0x40, 0x08, 0x45, 0x15, 0x42,
	0x39, 0xed, 0x81, 0x57, 0x28, 0x17, 0x0e, 0xec, 0xc1, 0xb4, 0xe7, 0x55, 0x88, 0x87, 0x62, 0x11,
	0x7b, 0x2a, 0xdb, 0x0b, 0xec, 0xa3, 0xf1, 0x12, 0x3c, 0x13, 0xf2, 0x34, 0xdd, 0xb6, 0x88, 0xdb,
	0xf7, 0xfb, 0xbe, 0x64, 0x6c, 0xcf, 0x07, 0x2f, 0x3b, 0x76, 0x8e, 0xfd, 0x7d, 0xa4, 0xf0, 0x83,
	0xc2, 0xbd, 0x63, 0x43, 0xfd, 0xed, 0x3e, 0x70, 0x62, 0x35, 0x33, 0x46, 0xc4, 0xea, 0x4f, 0x09,
	0xe3, 0x6d, 0xa4, 0xa0, 0x5e, 0xc1, 0x5c, 0x33, 0xbb, 0x75, 0x1b, 0x0c, 0x16, 0x75, 0xd1, 0x94,
	0xfa, 0x89, 0xd5, 0x0d, 0x94, 0xfb, 0x9f, 0x06, 0x47, 0x75, 0xd1, 0x54, 0x3a, 0x4b, 0xa5, 0x60,
	0xdc, 0xb1, 0xf5, 0x58, 0xca, 0x97, 0xa2, 0xd5, 0x12, 0x46, 0xd6, 0xe0, 0xb8, 0x2e, 0x9a, 0x85,
	0x1e, 0x59, 0x93, 0x27, 0x7a, 0xdb, 0x7d, 0xbf, 0x6b, 0x1d, 0xe1, 0x44, 0x7e, 0x7d, 0x62, 0xf5,
	0x0e, 0x96, 0x3d, 0xef, 0xac, 0xdf, 0x3c, 0x04, 0x9f, 0xda, 0x2f, 0x3d, 0xe1, 0xb4, 0x2e, 0x9a,
	0xb9, 0xfe, 0xc7, 0x55, 0xb7, 0xa0, 0x2e, 0x9d, 0x8d, 0x75, 0x84, 0x33, 0x99, 0xf6, 0x9f, 0x44,
	0xbd, 0x80, 0x69, 0xec, 0x38, 0x50, 0xc4, 0x79, 0x5d, 0x34, 0x13, 0x3d, 0x90, 0x5a, 0xc1, 0x75,
	0xdf, 0xc6, 0xf4, 0xd9, 0xee, 0xbc, 0x4c, 0xa8, 0x64, 0xc2, 0x85, 0xa7, 0x5e, 0x43, 0x15, 0xed,
	0xce, 0xaf, 0xf9, 0xc1, 0x27, 0x04, 0xf9, 0xfd, 0x64, 0x28, 0x84, 0xd9, 0x07, 0xdb, 0x3a, 0xf6,
	0x06, 0xaf, 0xe4, 0xd1, 0x47, 0xcc, 0xef, 0x1c, 0xe4, 0x7b, 0xbc, 0x3e, 0x6c, 0xee, 0xc8, 0xf9,
	0x3e, 0xbc, 0x27, 0xff, 0xd1, 0xe0, 0x42, 0x4e, 0x1c, 0x28, 0x4f, 0xdb, 0x7a, 0xcb, 0x39, 0x58,
	0x4a, 0x70, 0xc4, 0x9c, 0x7c, 0xa3, 0xd6, 0x6c, 0x43, 0x8f, 0xcf, 0x0e, 0xc9, 0x80, 0xb2, 0x73,
	0x9b, 0x1e, 0xf1, 0x46, 0x6c, 0xd1, 0xb9, 0x99, 0x48, 0xbf, 0xf0, 0xb9, 0xdc, 0x36, 0xcb, 0xd5,
	0xef, 0x02, 0x66, 0x9b, 0x3b, 0x4e, 0xb6, 0xa3, 0xa1, 0x91, 0x42, 0xc2, 0xdc, 0xc8, 0x1b, 0x00,
	0x2f, 0xc9, 0xe6, 0x71, 0x4f, 0x52, 0xe7, 0x44, 0x9f, 0x39, 0xaa, 0x86, 0xab, 0x81, 0x6c, 0xea,
	0x49, 0xca, 0xad, 0xf4, 0xb9, 0xa5, 0xde, 0xc2, 0xe2, 0x80, 0x6b, 0xf6, 0x89, 0x7c, 0x92, 0xba,
	0x2b, 0x7d, 0x69, 0x9e, 0xce, 0xf9, 0x44, 0x8e, 0x87, 0xee, 0xcf, 0x9c, 0xbc, 0x95, 0xaf, 0xb6,
	0x27, 0x13, 0x71, 0x5a, 0x97, 0x79, 0x2b, 0x07, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xc4,
	0x86, 0x15, 0xb1, 0x02, 0x00, 0x00,
}

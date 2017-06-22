// Code generated by protoc-gen-go.
// source: tdk_data.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TdkDeskUserData struct {
	Userid           *uint32 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Headurl          *string `protobuf:"bytes,3,opt,name=headurl" json:"headurl,omitempty"`
	Score            *int32  `protobuf:"varint,4,opt,name=score" json:"score,omitempty"`
	Already          *bool   `protobuf:"varint,5,opt,name=already" json:"already,omitempty"`
	Fold             *bool   `protobuf:"varint,6,opt,name=fold" json:"fold,omitempty"`
	Isowner          *bool   `protobuf:"varint,7,opt,name=isowner" json:"isowner,omitempty"`
	Pos              *int32  `protobuf:"varint,8,opt,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TdkDeskUserData) Reset()                    { *m = TdkDeskUserData{} }
func (m *TdkDeskUserData) String() string            { return proto.CompactTextString(m) }
func (*TdkDeskUserData) ProtoMessage()               {}
func (*TdkDeskUserData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *TdkDeskUserData) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *TdkDeskUserData) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TdkDeskUserData) GetHeadurl() string {
	if m != nil && m.Headurl != nil {
		return *m.Headurl
	}
	return ""
}

func (m *TdkDeskUserData) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *TdkDeskUserData) GetAlready() bool {
	if m != nil && m.Already != nil {
		return *m.Already
	}
	return false
}

func (m *TdkDeskUserData) GetFold() bool {
	if m != nil && m.Fold != nil {
		return *m.Fold
	}
	return false
}

func (m *TdkDeskUserData) GetIsowner() bool {
	if m != nil && m.Isowner != nil {
		return *m.Isowner
	}
	return false
}

func (m *TdkDeskUserData) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

type TdkUserPokerData struct {
	Userid           *uint32  `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Pokerlist        []uint32 `protobuf:"varint,2,rep,name=pokerlist" json:"pokerlist,omitempty"`
	Score            *uint32  `protobuf:"varint,3,opt,name=score" json:"score,omitempty"`
	Borrow           *bool    `protobuf:"varint,4,opt,name=borrow" json:"borrow,omitempty"`
	Pubscore         *uint32  `protobuf:"varint,5,opt,name=pubscore" json:"pubscore,omitempty"`
	Fold             *bool    `protobuf:"varint,6,opt,name=fold" json:"fold,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TdkUserPokerData) Reset()                    { *m = TdkUserPokerData{} }
func (m *TdkUserPokerData) String() string            { return proto.CompactTextString(m) }
func (*TdkUserPokerData) ProtoMessage()               {}
func (*TdkUserPokerData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *TdkUserPokerData) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *TdkUserPokerData) GetPokerlist() []uint32 {
	if m != nil {
		return m.Pokerlist
	}
	return nil
}

func (m *TdkUserPokerData) GetScore() uint32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *TdkUserPokerData) GetBorrow() bool {
	if m != nil && m.Borrow != nil {
		return *m.Borrow
	}
	return false
}

func (m *TdkUserPokerData) GetPubscore() uint32 {
	if m != nil && m.Pubscore != nil {
		return *m.Pubscore
	}
	return 0
}

func (m *TdkUserPokerData) GetFold() bool {
	if m != nil && m.Fold != nil {
		return *m.Fold
	}
	return false
}

type TdkDisUserData struct {
	Userid           *uint32 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Winnum           *uint32 `protobuf:"varint,2,opt,name=winnum" json:"winnum,omitempty"`
	Languonum        *uint32 `protobuf:"varint,3,opt,name=languonum" json:"languonum,omitempty"`
	Baozinum         *uint32 `protobuf:"varint,4,opt,name=baozinum" json:"baozinum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TdkDisUserData) Reset()                    { *m = TdkDisUserData{} }
func (m *TdkDisUserData) String() string            { return proto.CompactTextString(m) }
func (*TdkDisUserData) ProtoMessage()               {}
func (*TdkDisUserData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *TdkDisUserData) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *TdkDisUserData) GetWinnum() uint32 {
	if m != nil && m.Winnum != nil {
		return *m.Winnum
	}
	return 0
}

func (m *TdkDisUserData) GetLanguonum() uint32 {
	if m != nil && m.Languonum != nil {
		return *m.Languonum
	}
	return 0
}

func (m *TdkDisUserData) GetBaozinum() uint32 {
	if m != nil && m.Baozinum != nil {
		return *m.Baozinum
	}
	return 0
}

type TdkDeskConfig struct {
	Roundcount       *uint32 `protobuf:"varint,1,opt,name=roundcount" json:"roundcount,omitempty"`
	Playercount      *uint32 `protobuf:"varint,2,opt,name=playercount" json:"playercount,omitempty"`
	Allinrate        *uint32 `protobuf:"varint,3,opt,name=allinrate" json:"allinrate,omitempty"`
	Basecoin         *uint32 `protobuf:"varint,4,opt,name=basecoin" json:"basecoin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TdkDeskConfig) Reset()                    { *m = TdkDeskConfig{} }
func (m *TdkDeskConfig) String() string            { return proto.CompactTextString(m) }
func (*TdkDeskConfig) ProtoMessage()               {}
func (*TdkDeskConfig) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *TdkDeskConfig) GetRoundcount() uint32 {
	if m != nil && m.Roundcount != nil {
		return *m.Roundcount
	}
	return 0
}

func (m *TdkDeskConfig) GetPlayercount() uint32 {
	if m != nil && m.Playercount != nil {
		return *m.Playercount
	}
	return 0
}

func (m *TdkDeskConfig) GetAllinrate() uint32 {
	if m != nil && m.Allinrate != nil {
		return *m.Allinrate
	}
	return 0
}

func (m *TdkDeskConfig) GetBasecoin() uint32 {
	if m != nil && m.Basecoin != nil {
		return *m.Basecoin
	}
	return 0
}

type TdkZhanJiData struct {
	Userid           *uint32 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Winnum           *uint32 `protobuf:"varint,2,opt,name=winnum" json:"winnum,omitempty"`
	Languonum        *uint32 `protobuf:"varint,3,opt,name=languonum" json:"languonum,omitempty"`
	Baozinum         *uint32 `protobuf:"varint,4,opt,name=baozinum" json:"baozinum,omitempty"`
	Score            *int32  `protobuf:"varint,5,opt,name=score" json:"score,omitempty"`
	Isowner          *bool   `protobuf:"varint,6,opt,name=isowner" json:"isowner,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TdkZhanJiData) Reset()                    { *m = TdkZhanJiData{} }
func (m *TdkZhanJiData) String() string            { return proto.CompactTextString(m) }
func (*TdkZhanJiData) ProtoMessage()               {}
func (*TdkZhanJiData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *TdkZhanJiData) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *TdkZhanJiData) GetWinnum() uint32 {
	if m != nil && m.Winnum != nil {
		return *m.Winnum
	}
	return 0
}

func (m *TdkZhanJiData) GetLanguonum() uint32 {
	if m != nil && m.Languonum != nil {
		return *m.Languonum
	}
	return 0
}

func (m *TdkZhanJiData) GetBaozinum() uint32 {
	if m != nil && m.Baozinum != nil {
		return *m.Baozinum
	}
	return 0
}

func (m *TdkZhanJiData) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *TdkZhanJiData) GetIsowner() bool {
	if m != nil && m.Isowner != nil {
		return *m.Isowner
	}
	return false
}

func init() {
	proto.RegisterType((*TdkDeskUserData)(nil), "yjprotogo.TdkDeskUserData")
	proto.RegisterType((*TdkUserPokerData)(nil), "yjprotogo.TdkUserPokerData")
	proto.RegisterType((*TdkDisUserData)(nil), "yjprotogo.TdkDisUserData")
	proto.RegisterType((*TdkDeskConfig)(nil), "yjprotogo.TdkDeskConfig")
	proto.RegisterType((*TdkZhanJiData)(nil), "yjprotogo.TdkZhanJiData")
}

var fileDescriptor14 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x91, 0xb1, 0x6e, 0xc2, 0x30,
	0x10, 0x86, 0x45, 0x43, 0x20, 0x1c, 0x0d, 0x50, 0x77, 0xf1, 0x58, 0x65, 0xea, 0xd4, 0x97, 0x28,
	0x53, 0xa7, 0x0e, 0x65, 0xa9, 0x54, 0x55, 0x06, 0x1b, 0x70, 0x63, 0x7c, 0x91, 0xed, 0x08, 0xd1,
	0x67, 0xe8, 0x43, 0xf7, 0x6c, 0x08, 0x74, 0x61, 0xeb, 0x12, 0xf9, 0x7e, 0xdf, 0xfd, 0xf7, 0xe5,
	0x37, 0x4c, 0x82, 0xac, 0x3f, 0xa5, 0x08, 0xe2, 0xa9, 0x71, 0x18, 0x90, 0x8d, 0x0e, 0x5f, 0xe9,
	0xb0, 0xc1, 0xea, 0xa7, 0x07, 0xd3, 0x37, 0x59, 0xcf, 0x95, 0xaf, 0x17, 0x5e, 0xb9, 0x39, 0x35,
	0xb1, 0x09, 0x0c, 0x5a, 0x3a, 0x6b, 0xc9, 0x7b, 0x0f, 0xbd, 0xc7, 0x92, 0xdd, 0x42, 0xdf, 0x8a,
	0x9d, 0xe2, 0x37, 0x54, 0x8d, 0xd8, 0x14, 0x86, 0x5b, 0x25, 0x64, 0xeb, 0x0c, 0xcf, 0x92, 0x50,
	0x42, 0xee, 0x57, 0xe8, 0x14, 0xef, 0x53, 0x99, 0xc7, 0x7b, 0x61, 0x1c, 0x75, 0x1c, 0x78, 0x4e,
	0x42, 0x11, 0xc7, 0xd7, 0x68, 0x24, 0x1f, 0xa4, 0x8a, 0xae, 0xb5, 0xc7, 0xbd, 0x55, 0x8e, 0x0f,
	0x93, 0x30, 0x86, 0xac, 0x41, 0xcf, 0x8b, 0x38, 0x5c, 0x05, 0x98, 0x11, 0x4d, 0x24, 0x79, 0xc5,
	0xfa, 0x0a, 0xce, 0x1d, 0x8c, 0x9a, 0x78, 0x69, 0xb4, 0x0f, 0xc4, 0x94, 0x91, 0x74, 0x46, 0xc8,
	0x52, 0x07, 0x4d, 0x2c, 0xd1, 0x39, 0xdc, 0x27, 0xa4, 0x82, 0xcd, 0xa0, 0x68, 0xda, 0xe5, 0xb1,
	0x23, 0xef, 0x7e, 0xe9, 0xc2, 0x54, 0x2d, 0x60, 0x12, 0x33, 0xd0, 0xfe, 0x6a, 0x04, 0x54, 0xef,
	0xb5, 0xb5, 0xed, 0x2e, 0x85, 0x90, 0x18, 0x8c, 0xb0, 0x9b, 0x16, 0xa3, 0x74, 0x5c, 0x4a, 0x4b,
	0x96, 0x02, 0xbf, 0x75, 0x54, 0xe2, 0xda, 0xb2, 0xfa, 0x80, 0xf2, 0x14, 0xed, 0x33, 0xda, 0xb5,
	0xde, 0x30, 0x06, 0xe0, 0xb0, 0xb5, 0x72, 0x45, 0x9f, 0x70, 0x72, 0xbe, 0x87, 0x71, 0x63, 0xc4,
	0x41, 0xb9, 0xa3, 0x78, 0xb6, 0x17, 0xc6, 0x68, 0xeb, 0x44, 0x50, 0x7f, 0xed, 0xbd, 0x5a, 0xa1,
	0xb6, 0x27, 0xfb, 0x90, 0xec, 0xdf, 0xb7, 0xc2, 0xbe, 0xe8, 0x7f, 0x83, 0xbe, 0x44, 0x99, 0x77,
	0xaf, 0xd9, 0x3d, 0x57, 0xca, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0x01, 0xb9, 0x90, 0x8b, 0x4c,
	0x02, 0x00, 0x00,
}
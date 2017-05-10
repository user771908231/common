// Code generated by protoc-gen-go.
// source: niuniu_server.proto
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

// Ignoring public import of niuniu_client_poker from niuniu_base.proto

// Ignoring public import of niuniu_user_bill from niuniu_base.proto

// Ignoring public import of niuniu_desk_option from niuniu_base.proto

// Ignoring public import of niuniu_enum_protoid from niuniu_base.proto

// Ignoring public import of niuniu_enum_PokerType from niuniu_base.proto

// Ignoring public import of niuniu_enum_desk_state from niuniu_base.proto

// Ignoring public import of niuniu_enum_banker_rule from niuniu_base.proto

// 打出去的牌
type NiuniuSrvPoker struct {
	Pais             []*CommonSrvPokerPai  `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *NiuniuEnum_PokerType `protobuf:"varint,3,opt,name=type,enum=ddproto.NiuniuEnum_PokerType" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuSrvPoker) Reset()                    { *m = NiuniuSrvPoker{} }
func (m *NiuniuSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvPoker) ProtoMessage()               {}
func (*NiuniuSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{0} }

func (m *NiuniuSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *NiuniuSrvPoker) GetType() NiuniuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NiuniuEnum_PokerType_NO_NIU
}

// desk 的信息
type NiuniuSrvDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	DeskNumber       *string              `protobuf:"bytes,2,opt,name=deskNumber" json:"deskNumber,omitempty"`
	GameNumber       *int32               `protobuf:"varint,3,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32               `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	Status           *NiuniuEnumDeskState `protobuf:"varint,5,opt,name=status,enum=ddproto.NiuniuEnumDeskState" json:"status,omitempty"`
	LastWiner        *uint32              `protobuf:"varint,6,opt,name=lastWiner" json:"lastWiner,omitempty"`
	CircleNo         *int32               `protobuf:"varint,8,opt,name=circleNo" json:"circleNo,omitempty"`
	Owner            *uint32              `protobuf:"varint,10,opt,name=owner" json:"owner,omitempty"`
	CurrBanker       *uint32              `protobuf:"varint,11,opt,name=currBanker" json:"currBanker,omitempty"`
	DeskOption       *NiuniuDeskOption    `protobuf:"bytes,12,opt,name=deskOption" json:"deskOption,omitempty"`
	IsStart          *bool                `protobuf:"varint,13,opt,name=isStart" json:"isStart,omitempty"`
	IsOnDissolve     *bool                `protobuf:"varint,14,opt,name=isOnDissolve" json:"isOnDissolve,omitempty"`
	DissolveTime     *int64               `protobuf:"varint,15,opt,name=dissolveTime" json:"dissolveTime,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuSrvDesk) Reset()                    { *m = NiuniuSrvDesk{} }
func (m *NiuniuSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvDesk) ProtoMessage()               {}
func (*NiuniuSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{1} }

func (m *NiuniuSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDeskNumber() string {
	if m != nil && m.DeskNumber != nil {
		return *m.DeskNumber
	}
	return ""
}

func (m *NiuniuSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *NiuniuSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvDesk) GetStatus() NiuniuEnumDeskState {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER
}

func (m *NiuniuSrvDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *NiuniuSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *NiuniuSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *NiuniuSrvDesk) GetCurrBanker() uint32 {
	if m != nil && m.CurrBanker != nil {
		return *m.CurrBanker
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDeskOption() *NiuniuDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

func (m *NiuniuSrvDesk) GetIsStart() bool {
	if m != nil && m.IsStart != nil {
		return *m.IsStart
	}
	return false
}

func (m *NiuniuSrvDesk) GetIsOnDissolve() bool {
	if m != nil && m.IsOnDissolve != nil {
		return *m.IsOnDissolve
	}
	return false
}

func (m *NiuniuSrvDesk) GetDissolveTime() int64 {
	if m != nil && m.DissolveTime != nil {
		return *m.DissolveTime
	}
	return 0
}

// 用户属性
type NiuniuSrvUser struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Bill             *NiuniuUserBill `protobuf:"bytes,6,opt,name=bill" json:"bill,omitempty"`
	IsOnline         *bool           `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32          `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	Pokers           *NiuniuSrvPoker `protobuf:"bytes,13,opt,name=pokers" json:"pokers,omitempty"`
	BankerScore      *int32          `protobuf:"varint,14,opt,name=bankerScore" json:"bankerScore,omitempty"`
	DoubleScore      *int32          `protobuf:"varint,15,opt,name=doubleScore" json:"doubleScore,omitempty"`
	IsReady          *bool           `protobuf:"varint,16,opt,name=isReady" json:"isReady,omitempty"`
	LastScore        *int32          `protobuf:"varint,17,opt,name=lastScore" json:"lastScore,omitempty"`
	DissolveState    *int32          `protobuf:"varint,18,opt,name=dissolveState" json:"dissolveState,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NiuniuSrvUser) Reset()                    { *m = NiuniuSrvUser{} }
func (m *NiuniuSrvUser) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvUser) ProtoMessage()               {}
func (*NiuniuSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{2} }

func (m *NiuniuSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *NiuniuSrvUser) GetBill() *NiuniuUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *NiuniuSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *NiuniuSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *NiuniuSrvUser) GetPokers() *NiuniuSrvPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

func (m *NiuniuSrvUser) GetBankerScore() int32 {
	if m != nil && m.BankerScore != nil {
		return *m.BankerScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetDoubleScore() int32 {
	if m != nil && m.DoubleScore != nil {
		return *m.DoubleScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *NiuniuSrvUser) GetLastScore() int32 {
	if m != nil && m.LastScore != nil {
		return *m.LastScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetDissolveState() int32 {
	if m != nil && m.DissolveState != nil {
		return *m.DissolveState
	}
	return 0
}

// room 的信息
type NiuniuSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuniuSrvRoom) Reset()                    { *m = NiuniuSrvRoom{} }
func (m *NiuniuSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvRoom) ProtoMessage()               {}
func (*NiuniuSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{3} }

func (m *NiuniuSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func init() {
	proto.RegisterType((*NiuniuSrvPoker)(nil), "ddproto.niuniu_srv_poker")
	proto.RegisterType((*NiuniuSrvDesk)(nil), "ddproto.niuniu_srv_desk")
	proto.RegisterType((*NiuniuSrvUser)(nil), "ddproto.niuniu_srv_user")
	proto.RegisterType((*NiuniuSrvRoom)(nil), "ddproto.niuniu_srv_room")
}

func init() { proto.RegisterFile("niuniu_server.proto", fileDescriptor31) }

var fileDescriptor31 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xdd, 0x6e, 0xd3, 0x40,
	0x10, 0x85, 0x71, 0x12, 0xa7, 0xc9, 0xa4, 0xf9, 0xdb, 0x82, 0xb4, 0x2d, 0x08, 0xac, 0xdc, 0x60,
	0x10, 0x4a, 0xa5, 0x3c, 0x02, 0xe2, 0xa6, 0x12, 0x6a, 0x23, 0x12, 0x09, 0x71, 0x15, 0x39, 0xf6,
	0x08, 0xad, 0x62, 0xef, 0x5a, 0xbb, 0x76, 0x20, 0x6f, 0xc7, 0x7b, 0x71, 0x83, 0x66, 0xbc, 0xa6,
	0x3f, 0xea, 0xd5, 0x66, 0x4f, 0x3e, 0x9f, 0x99, 0xb3, 0x33, 0x70, 0xa1, 0x55, 0xad, 0x55, 0xbd,
	0x73, 0x68, 0x8f, 0x68, 0x97, 0xa5, 0x35, 0x95, 0x11, 0x67, 0x59, 0xc6, 0x3f, 0xae, 0x2e, 0x53,
	0x53, 0x14, 0x46, 0xfb, 0x7f, 0x77, 0xa5, 0x39, 0xb4, 0xcc, 0xd5, 0xdc, 0x7f, 0xb8, 0x4f, 0x1c,
	0x36, 0xd2, 0x22, 0x87, 0x59, 0xeb, 0x66, 0x8f, 0x0d, 0x2c, 0x3e, 0x42, 0xaf, 0x4c, 0x94, 0x93,
	0x9d, 0xa8, 0x1b, 0x8f, 0x56, 0x6f, 0x96, 0xde, 0x79, 0xd9, 0x1a, 0xb7, 0xe0, 0x3a, 0x51, 0xe2,
	0x13, 0xf4, 0xaa, 0x53, 0x89, 0xb2, 0x1b, 0x05, 0xf1, 0x64, 0xf5, 0xf6, 0x3f, 0xeb, 0x4d, 0x51,
	0xd7, 0xc5, 0x6e, 0x4d, 0xf0, 0xf6, 0x54, 0xe2, 0xe2, 0x4f, 0x07, 0xa6, 0x0f, 0xca, 0x65, 0xe8,
	0x0e, 0x62, 0x02, 0x7d, 0x3a, 0x6f, 0x32, 0x19, 0x44, 0x41, 0x1c, 0x0a, 0x01, 0x40, 0xf7, 0xdb,
	0xba, 0xd8, 0xa3, 0x95, 0x9d, 0x28, 0x88, 0x87, 0xa4, 0xfd, 0x4c, 0x0a, 0xf4, 0x5a, 0x97, 0xb9,
	0x09, 0xf4, 0xad, 0x31, 0xc5, 0x4d, 0x26, 0x7b, 0x7c, 0xbf, 0x86, 0xbe, 0xab, 0x92, 0xaa, 0x76,
	0x32, 0xe4, 0x5e, 0xde, 0x3d, 0xdb, 0x0b, 0x59, 0xef, 0x88, 0x43, 0x31, 0x87, 0x61, 0x9e, 0xb8,
	0xea, 0xbb, 0xd2, 0x68, 0x65, 0x3f, 0x0a, 0xe2, 0xb1, 0x98, 0xc1, 0x20, 0x55, 0x36, 0xcd, 0xf1,
	0xd6, 0xc8, 0x01, 0xbb, 0x8e, 0x21, 0x34, 0xbf, 0x08, 0x00, 0x06, 0x04, 0x40, 0x5a, 0x5b, 0xfb,
	0x39, 0xd1, 0x07, 0xb4, 0x72, 0xc4, 0xda, 0x75, 0xd3, 0xf0, 0x5d, 0x59, 0x29, 0xa3, 0xe5, 0x79,
	0x14, 0xc4, 0xa3, 0xd5, 0xeb, 0xa7, 0xc5, 0xb9, 0xae, 0x61, 0x44, 0x4c, 0xe1, 0x4c, 0xb9, 0x4d,
	0x95, 0xd8, 0x4a, 0x8e, 0xa3, 0x20, 0x1e, 0x88, 0x97, 0x70, 0xae, 0xdc, 0x9d, 0xfe, 0xa2, 0x9c,
	0x33, 0xf9, 0x11, 0xe5, 0xa4, 0x55, 0x33, 0xaf, 0x6c, 0x55, 0x81, 0x72, 0x1a, 0x05, 0x71, 0x77,
	0xf1, 0x37, 0x78, 0xf4, 0x84, 0xb5, 0x43, 0x4b, 0x4f, 0x41, 0xa7, 0x7f, 0xc2, 0xb1, 0x78, 0x0f,
	0xbd, 0xbd, 0xca, 0x73, 0x0e, 0x35, 0x5a, 0x5d, 0x3e, 0xed, 0x85, 0xd8, 0x1d, 0x01, 0x94, 0x97,
	0x0a, 0xe7, 0x4a, 0x23, 0x07, 0x1c, 0x50, 0x5e, 0xa5, 0x33, 0xfc, 0xcd, 0x39, 0x42, 0xf1, 0x01,
	0xfa, 0x3c, 0x6a, 0xc7, 0x9d, 0x3e, 0xe3, 0x75, 0xbf, 0x35, 0x17, 0x30, 0xda, 0xf3, 0xb3, 0x6c,
	0x52, 0x63, 0x9b, 0x0c, 0x21, 0x89, 0x99, 0xa9, 0xf7, 0x39, 0x36, 0xe2, 0x94, 0x45, 0xce, 0xff,
	0x0d, 0x93, 0xec, 0x24, 0x67, 0x5c, 0xd4, 0x4f, 0xa2, 0x61, 0xe6, 0xcc, 0xbc, 0x82, 0x71, 0x1b,
	0x7e, 0x43, 0xd3, 0x92, 0x82, 0xe4, 0xc5, 0x8f, 0x47, 0xe1, 0x69, 0xfe, 0x0f, 0xf6, 0xa0, 0xd9,
	0x9f, 0x19, 0x0c, 0xe8, 0x4e, 0xfb, 0xc6, 0xdb, 0x13, 0x92, 0x3d, 0x29, 0x5f, 0xf1, 0x88, 0xb9,
	0x5f, 0x1e, 0x2f, 0x6d, 0x55, 0x95, 0x23, 0xef, 0xcf, 0x70, 0xfd, 0x62, 0x1d, 0xfc, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xa4, 0x76, 0x7a, 0x07, 0x5a, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go.
// source: zjh_base.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EProtoId int32

const (
	// //////////////////////////////////
	EProtoId_ZJH_PID_HEARTBEAT           EProtoId = 0
	EProtoId_ZJH_PID_QUICKCOnn           EProtoId = 1
	EProtoId_ZJH_PID_QUICKCONNA_ACK      EProtoId = 2
	EProtoId_ZJH_PID_GAME_LOGIN          EProtoId = 3
	EProtoId_ZJH_PID_GAME_LOGIN_ACK      EProtoId = 4
	EProtoId_ZJH_PID_GET_ROOM_LIST       EProtoId = 5
	EProtoId_ZJH_PID_ENTER_ROOM_LIST_ACK EProtoId = 6
	EProtoId_ZJH_PID_AUTO_ENTER_DESK     EProtoId = 7
	EProtoId_ZJH_PID_AUTO_ENTER_DESK_ACK EProtoId = 8
	EProtoId_ZJH_PID_SEND_GAMEINFO       EProtoId = 9
	EProtoId_ZJH_PID_BC_NEWENTER         EProtoId = 10
	EProtoId_ZJH_PID_BC_OPENING          EProtoId = 11
	EProtoId_ZJH_PID_REQ_READY           EProtoId = 12
	EProtoId_ZJH_PID_BC_READY            EProtoId = 13
	EProtoId_ZJH_PID_REQ_LEAVEDESK       EProtoId = 14
	EProtoId_ZJH_PID_BC_LEAVEDESK        EProtoId = 15
	EProtoId_ZJH_PID_BC_NEXT             EProtoId = 16
	EProtoId_ZJH_PID_REQ_KANPAI          EProtoId = 17
	EProtoId_ZJH_PID_BC_KANPAI           EProtoId = 18
	EProtoId_ZJH_PID_REQ_QIPAI           EProtoId = 19
	EProtoId_ZJH_PID_BC_QIPAI            EProtoId = 20
	EProtoId_ZJH_PID_REQ_GENZHU          EProtoId = 21
	EProtoId_ZJH_PID_BC_GENZHU           EProtoId = 22
	EProtoId_ZJH_PID_REQ_FAQI_XUEPIN     EProtoId = 23
	EProtoId_ZJH_PID_BC_FAQI_XUEPIN      EProtoId = 24
	EProtoId_ZJH_PID_BC_XUEPIN_END       EProtoId = 25
	EProtoId_ZJH_PID_REQ_JIAZHU          EProtoId = 26
	EProtoId_ZJH_PID_BC_JIAZHU           EProtoId = 27
	EProtoId_ZJH_PID_REQ_BIPAI           EProtoId = 28
	EProtoId_ZJH_PID_BC_BIPAI            EProtoId = 29
	EProtoId_ZJH_PID_BC_GAME_END         EProtoId = 30
	EProtoId_ZJH_PID_REQ_SEND_MESSAGE    EProtoId = 31
	EProtoId_ZJH_PID_BC_MESSAGE          EProtoId = 32
	EProtoId_ZJH_PID_REQ_DASHANG         EProtoId = 33
	EProtoId_ZJH_PID_BC_DASHANG          EProtoId = 34
)

var EProtoId_name = map[int32]string{
	0:  "ZJH_PID_HEARTBEAT",
	1:  "ZJH_PID_QUICKCOnn",
	2:  "ZJH_PID_QUICKCONNA_ACK",
	3:  "ZJH_PID_GAME_LOGIN",
	4:  "ZJH_PID_GAME_LOGIN_ACK",
	5:  "ZJH_PID_GET_ROOM_LIST",
	6:  "ZJH_PID_ENTER_ROOM_LIST_ACK",
	7:  "ZJH_PID_AUTO_ENTER_DESK",
	8:  "ZJH_PID_AUTO_ENTER_DESK_ACK",
	9:  "ZJH_PID_SEND_GAMEINFO",
	10: "ZJH_PID_BC_NEWENTER",
	11: "ZJH_PID_BC_OPENING",
	12: "ZJH_PID_REQ_READY",
	13: "ZJH_PID_BC_READY",
	14: "ZJH_PID_REQ_LEAVEDESK",
	15: "ZJH_PID_BC_LEAVEDESK",
	16: "ZJH_PID_BC_NEXT",
	17: "ZJH_PID_REQ_KANPAI",
	18: "ZJH_PID_BC_KANPAI",
	19: "ZJH_PID_REQ_QIPAI",
	20: "ZJH_PID_BC_QIPAI",
	21: "ZJH_PID_REQ_GENZHU",
	22: "ZJH_PID_BC_GENZHU",
	23: "ZJH_PID_REQ_FAQI_XUEPIN",
	24: "ZJH_PID_BC_FAQI_XUEPIN",
	25: "ZJH_PID_BC_XUEPIN_END",
	26: "ZJH_PID_REQ_JIAZHU",
	27: "ZJH_PID_BC_JIAZHU",
	28: "ZJH_PID_REQ_BIPAI",
	29: "ZJH_PID_BC_BIPAI",
	30: "ZJH_PID_BC_GAME_END",
	31: "ZJH_PID_REQ_SEND_MESSAGE",
	32: "ZJH_PID_BC_MESSAGE",
	33: "ZJH_PID_REQ_DASHANG",
	34: "ZJH_PID_BC_DASHANG",
}
var EProtoId_value = map[string]int32{
	"ZJH_PID_HEARTBEAT":           0,
	"ZJH_PID_QUICKCOnn":           1,
	"ZJH_PID_QUICKCONNA_ACK":      2,
	"ZJH_PID_GAME_LOGIN":          3,
	"ZJH_PID_GAME_LOGIN_ACK":      4,
	"ZJH_PID_GET_ROOM_LIST":       5,
	"ZJH_PID_ENTER_ROOM_LIST_ACK": 6,
	"ZJH_PID_AUTO_ENTER_DESK":     7,
	"ZJH_PID_AUTO_ENTER_DESK_ACK": 8,
	"ZJH_PID_SEND_GAMEINFO":       9,
	"ZJH_PID_BC_NEWENTER":         10,
	"ZJH_PID_BC_OPENING":          11,
	"ZJH_PID_REQ_READY":           12,
	"ZJH_PID_BC_READY":            13,
	"ZJH_PID_REQ_LEAVEDESK":       14,
	"ZJH_PID_BC_LEAVEDESK":        15,
	"ZJH_PID_BC_NEXT":             16,
	"ZJH_PID_REQ_KANPAI":          17,
	"ZJH_PID_BC_KANPAI":           18,
	"ZJH_PID_REQ_QIPAI":           19,
	"ZJH_PID_BC_QIPAI":            20,
	"ZJH_PID_REQ_GENZHU":          21,
	"ZJH_PID_BC_GENZHU":           22,
	"ZJH_PID_REQ_FAQI_XUEPIN":     23,
	"ZJH_PID_BC_FAQI_XUEPIN":      24,
	"ZJH_PID_BC_XUEPIN_END":       25,
	"ZJH_PID_REQ_JIAZHU":          26,
	"ZJH_PID_BC_JIAZHU":           27,
	"ZJH_PID_REQ_BIPAI":           28,
	"ZJH_PID_BC_BIPAI":            29,
	"ZJH_PID_BC_GAME_END":         30,
	"ZJH_PID_REQ_SEND_MESSAGE":    31,
	"ZJH_PID_BC_MESSAGE":          32,
	"ZJH_PID_REQ_DASHANG":         33,
	"ZJH_PID_BC_DASHANG":          34,
}

func (x EProtoId) Enum() *EProtoId {
	p := new(EProtoId)
	*p = x
	return p
}
func (x EProtoId) String() string {
	return proto.EnumName(EProtoId_name, int32(x))
}
func (x *EProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EProtoId_value, data, "EProtoId")
	if err != nil {
		return err
	}
	*x = EProtoId(value)
	return nil
}
func (EProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

// 玩家当前状态
type ZjhEnumPlayerGameStatus int32

const (
	ZjhEnumPlayerGameStatus_zjh_SS_GAMING   ZjhEnumPlayerGameStatus = 1
	ZjhEnumPlayerGameStatus_zjh_SS_STANDING ZjhEnumPlayerGameStatus = 2
	ZjhEnumPlayerGameStatus_zjh_SS_SITED    ZjhEnumPlayerGameStatus = 3
	ZjhEnumPlayerGameStatus_zjh_SS_READY    ZjhEnumPlayerGameStatus = 4
	ZjhEnumPlayerGameStatus_zjh_SS_OFFLINE  ZjhEnumPlayerGameStatus = 5
)

var ZjhEnumPlayerGameStatus_name = map[int32]string{
	1: "zjh_SS_GAMING",
	2: "zjh_SS_STANDING",
	3: "zjh_SS_SITED",
	4: "zjh_SS_READY",
	5: "zjh_SS_OFFLINE",
}
var ZjhEnumPlayerGameStatus_value = map[string]int32{
	"zjh_SS_GAMING":   1,
	"zjh_SS_STANDING": 2,
	"zjh_SS_SITED":    3,
	"zjh_SS_READY":    4,
	"zjh_SS_OFFLINE":  5,
}

func (x ZjhEnumPlayerGameStatus) Enum() *ZjhEnumPlayerGameStatus {
	p := new(ZjhEnumPlayerGameStatus)
	*p = x
	return p
}
func (x ZjhEnumPlayerGameStatus) String() string {
	return proto.EnumName(ZjhEnumPlayerGameStatus_name, int32(x))
}
func (x *ZjhEnumPlayerGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumPlayerGameStatus_value, data, "ZjhEnumPlayerGameStatus")
	if err != nil {
		return err
	}
	*x = ZjhEnumPlayerGameStatus(value)
	return nil
}
func (ZjhEnumPlayerGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{1} }

// 桌面状态
type ZjhEnumDeskState int32

const (
	ZjhEnumDeskState_DESK_IS_GAMING ZjhEnumDeskState = 1
	ZjhEnumDeskState_DESK_IS_WAIT   ZjhEnumDeskState = 2
)

var ZjhEnumDeskState_name = map[int32]string{
	1: "DESK_IS_GAMING",
	2: "DESK_IS_WAIT",
}
var ZjhEnumDeskState_value = map[string]int32{
	"DESK_IS_GAMING": 1,
	"DESK_IS_WAIT":   2,
}

func (x ZjhEnumDeskState) Enum() *ZjhEnumDeskState {
	p := new(ZjhEnumDeskState)
	*p = x
	return p
}
func (x ZjhEnumDeskState) String() string {
	return proto.EnumName(ZjhEnumDeskState_name, int32(x))
}
func (x *ZjhEnumDeskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumDeskState_value, data, "ZjhEnumDeskState")
	if err != nil {
		return err
	}
	*x = ZjhEnumDeskState(value)
	return nil
}
func (ZjhEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{2} }

// 房间玩家状态
type ZjhEnumUserState int32

const (
	ZjhEnumUserState_USER_IS_GAMING ZjhEnumUserState = 1
	ZjhEnumUserState_USER_IS_STAND  ZjhEnumUserState = 2
	ZjhEnumUserState_USER_IS_SITED  ZjhEnumUserState = 3
)

var ZjhEnumUserState_name = map[int32]string{
	1: "USER_IS_GAMING",
	2: "USER_IS_STAND",
	3: "USER_IS_SITED",
}
var ZjhEnumUserState_value = map[string]int32{
	"USER_IS_GAMING": 1,
	"USER_IS_STAND":  2,
	"USER_IS_SITED":  3,
}

func (x ZjhEnumUserState) Enum() *ZjhEnumUserState {
	p := new(ZjhEnumUserState)
	*p = x
	return p
}
func (x ZjhEnumUserState) String() string {
	return proto.EnumName(ZjhEnumUserState_name, int32(x))
}
func (x *ZjhEnumUserState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumUserState_value, data, "ZjhEnumUserState")
	if err != nil {
		return err
	}
	*x = ZjhEnumUserState(value)
	return nil
}
func (ZjhEnumUserState) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{3} }

// 房间类型
type ZjhEnumRoomType int32

const (
	ZjhEnumRoomType_ROOM_TYPE_FRIEND   ZjhEnumRoomType = 1
	ZjhEnumRoomType_ROOM_TYPE_NORMAL   ZjhEnumRoomType = 2
	ZjhEnumRoomType_ROOM_TYPE_REDBLACK ZjhEnumRoomType = 3
)

var ZjhEnumRoomType_name = map[int32]string{
	1: "ROOM_TYPE_FRIEND",
	2: "ROOM_TYPE_NORMAL",
	3: "ROOM_TYPE_REDBLACK",
}
var ZjhEnumRoomType_value = map[string]int32{
	"ROOM_TYPE_FRIEND":   1,
	"ROOM_TYPE_NORMAL":   2,
	"ROOM_TYPE_REDBLACK": 3,
}

func (x ZjhEnumRoomType) Enum() *ZjhEnumRoomType {
	p := new(ZjhEnumRoomType)
	*p = x
	return p
}
func (x ZjhEnumRoomType) String() string {
	return proto.EnumName(ZjhEnumRoomType_name, int32(x))
}
func (x *ZjhEnumRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumRoomType_value, data, "ZjhEnumRoomType")
	if err != nil {
		return err
	}
	*x = ZjhEnumRoomType(value)
	return nil
}
func (ZjhEnumRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{4} }

func init() {
	proto.RegisterEnum("ddproto.EProtoId", EProtoId_name, EProtoId_value)
	proto.RegisterEnum("ddproto.ZjhEnumPlayerGameStatus", ZjhEnumPlayerGameStatus_name, ZjhEnumPlayerGameStatus_value)
	proto.RegisterEnum("ddproto.ZjhEnumDeskState", ZjhEnumDeskState_name, ZjhEnumDeskState_value)
	proto.RegisterEnum("ddproto.ZjhEnumUserState", ZjhEnumUserState_name, ZjhEnumUserState_value)
	proto.RegisterEnum("ddproto.ZjhEnumRoomType", ZjhEnumRoomType_name, ZjhEnumRoomType_value)
}

var fileDescriptor20 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x94, 0xcb, 0x73, 0xda, 0x3e,
	0x10, 0xc7, 0x7f, 0x21, 0xc9, 0x8f, 0x54, 0xcd, 0x43, 0x88, 0x77, 0x48, 0x9b, 0xb6, 0x47, 0x0e,
	0xfd, 0x03, 0x7a, 0x13, 0x78, 0x31, 0x0a, 0x46, 0x36, 0xb6, 0x28, 0x49, 0x2e, 0x1a, 0x3a, 0x78,
	0xa6, 0xd3, 0x96, 0xc7, 0xf0, 0x38, 0xa4, 0xbd, 0xf4, 0x4f, 0xaf, 0x24, 0x63, 0x50, 0x74, 0xf0,
	0x8c, 0xe7, 0xb3, 0xfb, 0xdd, 0xfd, 0xca, 0xda, 0x35, 0xba, 0xfe, 0xfd, 0xe3, 0xbb, 0xfc, 0x36,
	0xdd, 0xa4, 0x9f, 0x57, 0xeb, 0xe5, 0x76, 0x49, 0x8a, 0xb3, 0x99, 0x79, 0x69, 0xff, 0x2d, 0xa2,
	0x0b, 0x88, 0xf4, 0x2b, 0x9b, 0x91, 0x2a, 0x2a, 0x3d, 0x3f, 0xf4, 0x65, 0xc4, 0x3c, 0xd9, 0x07,
	0x1a, 0x8b, 0x0e, 0x50, 0x81, 0xff, 0xb3, 0xf1, 0x68, 0xcc, 0xba, 0x83, 0x6e, 0xb8, 0x58, 0xe0,
	0x13, 0x72, 0x8b, 0x6a, 0x0e, 0xe6, 0x9c, 0x4a, 0xda, 0x1d, 0xe0, 0x02, 0xa9, 0x21, 0x92, 0xc7,
	0x7c, 0x3a, 0x04, 0x19, 0x84, 0x3e, 0xe3, 0xf8, 0xd4, 0xd6, 0x1c, 0xb9, 0xd1, 0x9c, 0x91, 0x26,
	0xaa, 0x1e, 0x62, 0x20, 0x64, 0x1c, 0x86, 0x43, 0x19, 0xb0, 0x44, 0xe0, 0x73, 0x72, 0x8f, 0x5a,
	0x79, 0x08, 0xb8, 0x80, 0xf8, 0x18, 0x34, 0xda, 0xff, 0x49, 0x0b, 0xd5, 0xf3, 0x04, 0x3a, 0x16,
	0xe1, 0x3e, 0xcb, 0x83, 0x64, 0x80, 0x8b, 0xb6, 0xda, 0x09, 0x1a, 0xf5, 0x85, 0xdd, 0x39, 0x01,
	0x9e, 0x59, 0x63, 0xbc, 0x17, 0xe2, 0x37, 0xa4, 0x8e, 0xca, 0x79, 0xa8, 0xd3, 0x95, 0x1c, 0x26,
	0x46, 0x8c, 0x91, 0x7d, 0x42, 0x15, 0x08, 0x23, 0xe0, 0x8c, 0xfb, 0xf8, 0xad, 0xfd, 0xb1, 0x62,
	0x18, 0xa9, 0x87, 0x7a, 0x4f, 0xf8, 0x92, 0x54, 0x10, 0xb6, 0xd2, 0x33, 0x7a, 0x65, 0x37, 0xd6,
	0xc9, 0x01, 0xd0, 0xaf, 0x60, 0x4c, 0x5f, 0x93, 0x06, 0xaa, 0x58, 0x82, 0x63, 0xe4, 0x86, 0x94,
	0xd1, 0xcd, 0x2b, 0x4b, 0x8f, 0x02, 0x63, 0xdb, 0x8e, 0xae, 0x34, 0xa0, 0x3c, 0xa2, 0x0c, 0x97,
	0x6c, 0x3b, 0x2a, 0x79, 0x8f, 0x89, 0xeb, 0x72, 0xc4, 0x34, 0x2e, 0x3b, 0x2e, 0x33, 0x5a, 0x71,
	0x6b, 0xfb, 0xc0, 0x9f, 0xfb, 0x63, 0x5c, 0x75, 0x6a, 0xef, 0x71, 0xcd, 0xbe, 0x0b, 0x9d, 0xde,
	0xa3, 0x23, 0x26, 0x1f, 0xc7, 0x10, 0xa9, 0x01, 0xa8, 0xdb, 0x03, 0xa0, 0x34, 0x76, 0xac, 0x61,
	0x7f, 0x0d, 0x15, 0xcb, 0xb0, 0xba, 0x2c, 0x0f, 0x37, 0x5d, 0x0b, 0x0f, 0x8c, 0xea, 0x5e, 0xb7,
	0x8e, 0x85, 0x3d, 0x6e, 0xb9, 0xc7, 0xeb, 0x98, 0x83, 0xdc, 0x39, 0xc7, 0xcb, 0xe8, 0x3b, 0xe7,
	0x8a, 0xcd, 0x58, 0xea, 0xa6, 0xef, 0xc9, 0x1d, 0x6a, 0xd8, 0x55, 0xcc, 0x68, 0x0c, 0x21, 0x49,
	0xa8, 0x0f, 0xf8, 0xde, 0x19, 0x80, 0x9c, 0x7f, 0xb0, 0xcb, 0x69, 0x95, 0x47, 0x93, 0x3e, 0x55,
	0x93, 0xf1, 0xd1, 0x11, 0xe4, 0xfc, 0x53, 0xfb, 0x0f, 0x6a, 0xea, 0xed, 0x4c, 0x17, 0xbb, 0xb9,
	0x5c, 0xfd, 0x9a, 0xbe, 0xa4, 0x6b, 0x7f, 0x3a, 0x4f, 0x93, 0xed, 0x74, 0xbb, 0xdb, 0x90, 0x12,
	0xba, 0xd2, 0xc1, 0x24, 0xd1, 0xc6, 0xf4, 0x84, 0x9d, 0xe8, 0xfb, 0xdf, 0xa3, 0x44, 0x50, 0xee,
	0x69, 0x58, 0x20, 0x18, 0x5d, 0xe6, 0x90, 0x09, 0xf0, 0xd4, 0xaa, 0x1d, 0x49, 0x36, 0x6d, 0x67,
	0x84, 0x64, 0xbf, 0x01, 0x45, 0xc2, 0x5e, 0x2f, 0x60, 0x1c, 0xf0, 0x79, 0xfb, 0x0b, 0x22, 0x87,
	0xe6, 0xb3, 0x74, 0xf3, 0x53, 0xb7, 0x4d, 0x75, 0xa6, 0x59, 0x0f, 0x66, 0xb5, 0x55, 0xf5, 0x72,
	0x36, 0xa1, 0x4c, 0xe0, 0x42, 0x9b, 0x5b, 0xda, 0xdd, 0x26, 0x5d, 0x1f, 0xb4, 0xe3, 0x44, 0xed,
	0x97, 0xad, 0x55, 0xa7, 0xc8, 0x99, 0xf1, 0xac, 0x0c, 0xdb, 0x28, 0x73, 0xdc, 0x9e, 0xa0, 0xd2,
	0xa1, 0xde, 0x7a, 0xb9, 0x9c, 0x8b, 0x97, 0x55, 0xaa, 0xef, 0xcc, 0x2c, 0xbb, 0x78, 0x8a, 0x40,
	0xf6, 0x62, 0xa6, 0xaf, 0xe6, 0xe4, 0x35, 0xe5, 0x61, 0x3c, 0xa4, 0x41, 0xf6, 0xd7, 0x39, 0xd2,
	0x18, 0xbc, 0x4e, 0xa0, 0xf7, 0xfb, 0xf4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x8d, 0x2b,
	0x8f, 0xfe, 0x04, 0x00, 0x00,
}

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
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x94, 0xc9, 0x92, 0xd3, 0x30,
	0x10, 0x86, 0x99, 0xcc, 0x8a, 0x98, 0xa5, 0xa3, 0xd9, 0x17, 0x18, 0xce, 0x39, 0xf0, 0x00, 0xdc,
	0x64, 0xbb, 0xe3, 0x68, 0xe2, 0xc8, 0x8e, 0xa5, 0x90, 0x99, 0xb9, 0xa8, 0x42, 0xc5, 0x55, 0x14,
	0x90, 0xa5, 0xb2, 0x1c, 0x06, 0x1e, 0x92, 0x57, 0xa2, 0x24, 0xc7, 0x89, 0xa2, 0x9b, 0xeb, 0xfb,
	0xfb, 0x57, 0xff, 0x92, 0x5a, 0x26, 0xa7, 0x7f, 0x7e, 0xfe, 0xd0, 0xdf, 0x07, 0xf3, 0xe2, 0xcb,
	0x74, 0x36, 0x59, 0x4c, 0xe8, 0xe1, 0x70, 0x68, 0x3f, 0x1a, 0xff, 0x0e, 0xc8, 0x11, 0x66, 0xe6,
	0x93, 0x0f, 0xe9, 0x25, 0xa9, 0xbf, 0x3e, 0xb5, 0x74, 0xc6, 0x23, 0xdd, 0x42, 0x96, 0xab, 0x00,
	0x99, 0x82, 0x77, 0x2e, 0xee, 0xf6, 0x78, 0xd8, 0x0e, 0xd3, 0xf1, 0x18, 0x76, 0xe8, 0x1d, 0xb9,
	0xf2, 0xb0, 0x10, 0x4c, 0xb3, 0xb0, 0x0d, 0x35, 0x7a, 0x45, 0x68, 0xa5, 0xc5, 0xac, 0x83, 0x3a,
	0x49, 0x63, 0x2e, 0x60, 0xd7, 0xf5, 0x6c, 0xb8, 0xf5, 0xec, 0xd1, 0x5b, 0x72, 0xb9, 0xd6, 0x50,
	0xe9, 0x3c, 0x4d, 0x3b, 0x3a, 0xe1, 0x52, 0xc1, 0x3e, 0x7d, 0x24, 0xf7, 0x95, 0x84, 0x42, 0x61,
	0xbe, 0x11, 0xad, 0xf7, 0x80, 0xde, 0x93, 0xeb, 0xaa, 0x80, 0xf5, 0x54, 0xba, 0xaa, 0x8a, 0x50,
	0xb6, 0xe1, 0xd0, 0x75, 0x7b, 0xa2, 0x75, 0x1f, 0xb9, 0x9d, 0x25, 0x8a, 0x32, 0x1a, 0x17, 0xcd,
	0x14, 0xde, 0xd3, 0x6b, 0x72, 0x5e, 0x49, 0x41, 0xa8, 0x05, 0xf6, 0xad, 0x19, 0x88, 0xbb, 0xc3,
	0x20, 0xd4, 0x69, 0x86, 0x82, 0x8b, 0x18, 0x3e, 0xb8, 0x87, 0x95, 0x63, 0x57, 0xe7, 0xc8, 0xa2,
	0x17, 0x38, 0xa6, 0x17, 0x04, 0x9c, 0xf2, 0x92, 0x9e, 0xb8, 0x8d, 0x4d, 0x71, 0x82, 0xec, 0x1b,
	0xda, 0xd0, 0xa7, 0xf4, 0x86, 0x5c, 0x38, 0x86, 0x8d, 0x72, 0x46, 0xcf, 0xc9, 0xd9, 0x56, 0xa4,
	0x67, 0x05, 0xe0, 0xc6, 0x31, 0x2b, 0xb5, 0x99, 0xc8, 0x18, 0x87, 0xba, 0x1b, 0x27, 0x08, 0x2b,
	0x4c, 0xfd, 0x94, 0x5d, 0x6e, 0xf0, 0xb9, 0x97, 0xb2, 0xa4, 0x17, 0xfe, 0xda, 0x31, 0x8a, 0xd7,
	0x56, 0x0f, 0x2e, 0xbd, 0xb5, 0x57, 0xf8, 0xca, 0xbd, 0x0b, 0x53, 0xde, 0x64, 0x5d, 0xae, 0x9f,
	0x7b, 0x98, 0x71, 0x01, 0xd7, 0xee, 0x00, 0x04, 0xe1, 0x96, 0x76, 0xe3, 0x9e, 0x46, 0x10, 0xae,
	0xb0, 0x46, 0x11, 0xc1, 0xad, 0x1f, 0xe1, 0x89, 0x33, 0xd3, 0xeb, 0xce, 0x8b, 0xb0, 0xc2, 0xf7,
	0xfe, 0xf6, 0x02, 0xbb, 0x91, 0x07, 0x6f, 0x7b, 0x25, 0xfd, 0xe8, 0x5d, 0xb1, 0x1d, 0x4b, 0xd3,
	0xf4, 0x13, 0x7d, 0x20, 0x37, 0xee, 0x2a, 0x76, 0x34, 0x3a, 0x28, 0x25, 0x8b, 0x11, 0x1e, 0xbd,
	0x01, 0xa8, 0xf8, 0xe7, 0xc6, 0x5f, 0x72, 0x6b, 0x1e, 0x5b, 0x31, 0x5e, 0x8e, 0xf4, 0xf4, 0xf7,
	0xe0, 0xad, 0x98, 0xc5, 0x83, 0x51, 0x21, 0x17, 0x83, 0xc5, 0x72, 0x4e, 0xeb, 0xe4, 0xc4, 0x88,
	0x52, 0x9a, 0x3e, 0x66, 0x60, 0x76, 0xcc, 0x75, 0xae, 0x90, 0x54, 0x4c, 0x44, 0x06, 0xd6, 0x28,
	0x90, 0xe3, 0x0a, 0x72, 0x85, 0x11, 0xec, 0x3a, 0xa4, 0x1c, 0x9e, 0x3d, 0x4a, 0xcb, 0x57, 0x2d,
	0xa5, 0x4e, 0x9b, 0xcd, 0x84, 0x0b, 0x84, 0xfd, 0xc6, 0x57, 0x42, 0xd7, 0xcd, 0x87, 0xc5, 0xfc,
	0x97, 0x69, 0x5b, 0x98, 0x4a, 0x3b, 0xed, 0xdc, 0x69, 0x0b, 0xe4, 0xb8, 0x62, 0x7d, 0xc6, 0x15,
	0xd4, 0x1a, 0xc2, 0xf1, 0x2e, 0xe7, 0xc5, 0x6c, 0xed, 0xed, 0x49, 0xcc, 0xb7, 0xbc, 0x75, 0x72,
	0x52, 0x31, 0x9b, 0x19, 0x6a, 0x5b, 0xa8, 0x4c, 0xdc, 0xe8, 0x93, 0xfa, 0x7a, 0xbd, 0xd9, 0x64,
	0x32, 0x52, 0x6f, 0xd3, 0xc2, 0x5c, 0x81, 0x7d, 0xbb, 0xea, 0x25, 0x43, 0xdd, 0xcc, 0xb9, 0x39,
	0xe9, 0x9d, 0x6d, 0x2a, 0xd2, 0xbc, 0xc3, 0x92, 0xf2, 0x27, 0xb2, 0xa1, 0x39, 0x46, 0x41, 0x62,
	0x9e, 0xeb, 0xee, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x71, 0xa4, 0x77, 0xcd, 0x04, 0x00,
	0x00,
}

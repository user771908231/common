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
	EProtoId_ZJH_PID_REQ_LEAVEDESK       EProtoId = 12
	EProtoId_ZJH_PID_BC_LEAVEDESK        EProtoId = 13
	EProtoId_ZJH_PID_BC_NEXT             EProtoId = 14
	EProtoId_ZJH_PID_REQ_KANPAI          EProtoId = 15
	EProtoId_ZJH_PID_BC_KANPAI           EProtoId = 16
	EProtoId_ZJH_PID_REQ_QIPAI           EProtoId = 17
	EProtoId_ZJH_PID_BC_QIPAI            EProtoId = 18
	EProtoId_ZJH_PID_REQ_GENZHU          EProtoId = 19
	EProtoId_ZJH_PID_BC_GENZHU           EProtoId = 20
	EProtoId_ZJH_PID_REQ_FAQI_XUEPIN     EProtoId = 21
	EProtoId_ZJH_PID_BC_FAQI_XUEPIN      EProtoId = 22
	EProtoId_ZJH_PID_BC_XUEPIN_END       EProtoId = 23
	EProtoId_ZJH_PID_REQ_JIAZHU          EProtoId = 24
	EProtoId_ZJH_PID_BC_JIAZHU           EProtoId = 25
	EProtoId_ZJH_PID_REQ_BIPAI           EProtoId = 26
	EProtoId_ZJH_PID_BC_BIPAI            EProtoId = 27
	EProtoId_ZJH_PID_BC_GAME_END         EProtoId = 28
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
	12: "ZJH_PID_REQ_LEAVEDESK",
	13: "ZJH_PID_BC_LEAVEDESK",
	14: "ZJH_PID_BC_NEXT",
	15: "ZJH_PID_REQ_KANPAI",
	16: "ZJH_PID_BC_KANPAI",
	17: "ZJH_PID_REQ_QIPAI",
	18: "ZJH_PID_BC_QIPAI",
	19: "ZJH_PID_REQ_GENZHU",
	20: "ZJH_PID_BC_GENZHU",
	21: "ZJH_PID_REQ_FAQI_XUEPIN",
	22: "ZJH_PID_BC_FAQI_XUEPIN",
	23: "ZJH_PID_BC_XUEPIN_END",
	24: "ZJH_PID_REQ_JIAZHU",
	25: "ZJH_PID_BC_JIAZHU",
	26: "ZJH_PID_REQ_BIPAI",
	27: "ZJH_PID_BC_BIPAI",
	28: "ZJH_PID_BC_GAME_END",
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
	"ZJH_PID_REQ_LEAVEDESK":       12,
	"ZJH_PID_BC_LEAVEDESK":        13,
	"ZJH_PID_BC_NEXT":             14,
	"ZJH_PID_REQ_KANPAI":          15,
	"ZJH_PID_BC_KANPAI":           16,
	"ZJH_PID_REQ_QIPAI":           17,
	"ZJH_PID_BC_QIPAI":            18,
	"ZJH_PID_REQ_GENZHU":          19,
	"ZJH_PID_BC_GENZHU":           20,
	"ZJH_PID_REQ_FAQI_XUEPIN":     21,
	"ZJH_PID_BC_FAQI_XUEPIN":      22,
	"ZJH_PID_BC_XUEPIN_END":       23,
	"ZJH_PID_REQ_JIAZHU":          24,
	"ZJH_PID_BC_JIAZHU":           25,
	"ZJH_PID_REQ_BIPAI":           26,
	"ZJH_PID_BC_BIPAI":            27,
	"ZJH_PID_BC_GAME_END":         28,
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
func (EProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

// 玩家当前状态
type ZjhEnumPlayerGameStatus int32

const (
	ZjhEnumPlayerGameStatus_ZJH_TEMP ZjhEnumPlayerGameStatus = 0
)

var ZjhEnumPlayerGameStatus_name = map[int32]string{
	0: "ZJH_TEMP",
}
var ZjhEnumPlayerGameStatus_value = map[string]int32{
	"ZJH_TEMP": 0,
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
func (ZjhEnumPlayerGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

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
func (ZjhEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

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
func (ZjhEnumUserState) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

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
func (ZjhEnumRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func init() {
	proto.RegisterEnum("ddproto.EProtoId", EProtoId_name, EProtoId_value)
	proto.RegisterEnum("ddproto.ZjhEnumPlayerGameStatus", ZjhEnumPlayerGameStatus_name, ZjhEnumPlayerGameStatus_value)
	proto.RegisterEnum("ddproto.ZjhEnumDeskState", ZjhEnumDeskState_name, ZjhEnumDeskState_value)
	proto.RegisterEnum("ddproto.ZjhEnumUserState", ZjhEnumUserState_name, ZjhEnumUserState_value)
	proto.RegisterEnum("ddproto.ZjhEnumRoomType", ZjhEnumRoomType_name, ZjhEnumRoomType_value)
}

var fileDescriptor12 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x93, 0xcb, 0x72, 0xda, 0x3c,
	0x14, 0xc7, 0x03, 0x7c, 0xf9, 0x42, 0x4f, 0x09, 0x39, 0x88, 0x5b, 0x08, 0x9d, 0xe9, 0xba, 0x2c,
	0xfa, 0x00, 0xdd, 0xc9, 0xf6, 0x01, 0x14, 0x40, 0x36, 0xb6, 0x28, 0x69, 0x36, 0x1a, 0x3a, 0x78,
	0xa6, 0xd3, 0x96, 0xcb, 0x70, 0x59, 0xa4, 0xcf, 0xdc, 0x87, 0xe8, 0xc8, 0x60, 0x90, 0xbd, 0xd3,
	0xfc, 0xfe, 0xe7, 0x7f, 0x2e, 0x3a, 0x12, 0x54, 0xff, 0xfc, 0xfc, 0xa1, 0xbf, 0x2f, 0xf6, 0xf1,
	0xe7, 0xed, 0x6e, 0x73, 0xd8, 0xb0, 0xbb, 0xe5, 0x32, 0x39, 0xf4, 0xfe, 0xde, 0x42, 0x99, 0x02,
	0x73, 0x14, 0x4b, 0xd6, 0x84, 0xda, 0xeb, 0xf3, 0x50, 0x07, 0xc2, 0xd3, 0x43, 0xe2, 0xa1, 0x72,
	0x88, 0x2b, 0xbc, 0xb1, 0xf1, 0x74, 0x26, 0xdc, 0x91, 0xeb, 0xaf, 0xd7, 0x58, 0x60, 0x4f, 0xd0,
	0xca, 0x61, 0x29, 0xb9, 0xe6, 0xee, 0x08, 0x8b, 0xac, 0x05, 0x2c, 0xd5, 0x06, 0x7c, 0x42, 0x7a,
	0xec, 0x0f, 0x84, 0xc4, 0x92, 0xed, 0xb9, 0xf2, 0xc4, 0xf3, 0x1f, 0xeb, 0x40, 0xf3, 0xa2, 0x91,
	0xd2, 0xa1, 0xef, 0x4f, 0xf4, 0x58, 0x44, 0x0a, 0x6f, 0xd9, 0x47, 0xe8, 0xa6, 0x12, 0x49, 0x45,
	0xe1, 0x55, 0x4c, 0xbc, 0xff, 0xb3, 0x2e, 0xb4, 0xd3, 0x00, 0x3e, 0x53, 0xfe, 0x39, 0xca, 0xa3,
	0x68, 0x84, 0x77, 0xb6, 0x3b, 0x27, 0x26, 0xee, 0xb2, 0x5d, 0x39, 0x22, 0x79, 0x6a, 0x4d, 0xc8,
	0xbe, 0x8f, 0xef, 0x58, 0x1b, 0xea, 0xa9, 0xe4, 0xb8, 0x5a, 0xd2, 0x3c, 0x31, 0x23, 0xd8, 0x13,
	0x3a, 0xae, 0xf6, 0x03, 0x92, 0x42, 0x0e, 0xf0, 0xbd, 0x9d, 0x2b, 0xa4, 0xa9, 0x1e, 0x13, 0xff,
	0x4a, 0x49, 0x1f, 0x15, 0xf6, 0x08, 0x0d, 0xcb, 0x72, 0x55, 0xee, 0x59, 0x1d, 0x1e, 0x32, 0x55,
	0x5e, 0x14, 0x56, 0xed, 0x0a, 0x26, 0xd3, 0x88, 0xcb, 0x80, 0x0b, 0x7c, 0xb0, 0xd7, 0xe1, 0xb8,
	0x29, 0x46, 0x1b, 0x9b, 0xf0, 0xa9, 0x30, 0xb8, 0xc6, 0x1a, 0x80, 0x56, 0xf4, 0x89, 0xb2, 0x7c,
	0xee, 0x01, 0xc9, 0xd7, 0xe1, 0x0c, 0xeb, 0xb9, 0xdc, 0x67, 0xdc, 0xb0, 0xaf, 0xd7, 0x84, 0xf7,
	0xf9, 0x54, 0xe8, 0x97, 0x19, 0x05, 0x42, 0x62, 0xd3, 0xde, 0xa9, 0xe3, 0x66, 0xb4, 0x96, 0x7d,
	0x1b, 0x8e, 0x7b, 0xc6, 0x9a, 0xa4, 0x87, 0xed, 0x7c, 0x0b, 0xcf, 0x82, 0x9b, 0x5a, 0x8f, 0xb9,
	0x16, 0xce, 0xb8, 0x93, 0x1f, 0xcf, 0x49, 0x06, 0x79, 0xca, 0x8d, 0x77, 0xa2, 0xdd, 0xdc, 0xd6,
	0x92, 0x97, 0x66, 0x8a, 0x7e, 0xe8, 0x7d, 0x82, 0x8e, 0xf9, 0x09, 0xf1, 0xfa, 0xb8, 0xd2, 0xdb,
	0xdf, 0x8b, 0xb7, 0x78, 0x37, 0x58, 0xac, 0xe2, 0xe8, 0xb0, 0x38, 0x1c, 0xf7, 0xac, 0x02, 0x65,
	0xe3, 0x52, 0x34, 0x09, 0xf0, 0xa6, 0xf7, 0x05, 0xd8, 0x25, 0x74, 0x19, 0xef, 0x7f, 0x99, 0xa0,
	0x98, 0x31, 0xa8, 0x26, 0x0f, 0x47, 0x44, 0x26, 0xad, 0x59, 0x79, 0x81, 0x21, 0x54, 0x52, 0x36,
	0xe7, 0x42, 0x61, 0xb1, 0x27, 0x2d, 0xef, 0x71, 0x1f, 0xef, 0x2e, 0xde, 0x59, 0x44, 0x61, 0xc6,
	0x5b, 0x83, 0xfb, 0x94, 0x45, 0x8a, 0x4b, 0x0f, 0x8b, 0x19, 0x24, 0x14, 0x79, 0x58, 0xea, 0xcd,
	0xa1, 0x76, 0xc9, 0xb7, 0xdb, 0x6c, 0x56, 0xea, 0x6d, 0x1b, 0x9b, 0xd1, 0x93, 0x6f, 0xa0, 0xbe,
	0x05, 0xa4, 0xfb, 0xa1, 0x30, 0x13, 0x16, 0xb2, 0x54, 0xfa, 0xe1, 0x84, 0x8f, 0x4f, 0xff, 0xf1,
	0x4a, 0x43, 0xf2, 0x9c, 0xb1, 0x79, 0xf9, 0xa5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0x52,
	0xc7, 0xdd, 0x18, 0x04, 0x00, 0x00,
}

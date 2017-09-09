// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bainiu_play.proto

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

// Ignoring public import of common_req_kickout from common_client.proto

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

// Ignoring public import of common_bc_leaveTimeout from common_client.proto

// Ignoring public import of common_desk_by_agent from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of bainiu_client_poker from bainiu_base.proto

// Ignoring public import of bainiu_enum_protoid from bainiu_base.proto

// Ignoring public import of bainiu_area_name from bainiu_base.proto

// Ignoring public import of bainiu_enum_PokerType from bainiu_base.proto

// ==============================百人牛牛==================================
type BainiuDeskStatus int32

const (
	BainiuDeskStatus_BAINIU_WAIT_YAZHU BainiuDeskStatus = 1
	BainiuDeskStatus_BAINIU_WAIT_FAPAI BainiuDeskStatus = 2
	BainiuDeskStatus_BAINIU_WAIT_BIPAI BainiuDeskStatus = 3
)

var BainiuDeskStatus_name = map[int32]string{
	1: "BAINIU_WAIT_YAZHU",
	2: "BAINIU_WAIT_FAPAI",
	3: "BAINIU_WAIT_BIPAI",
}
var BainiuDeskStatus_value = map[string]int32{
	"BAINIU_WAIT_YAZHU": 1,
	"BAINIU_WAIT_FAPAI": 2,
	"BAINIU_WAIT_BIPAI": 3,
}

func (x BainiuDeskStatus) Enum() *BainiuDeskStatus {
	p := new(BainiuDeskStatus)
	*p = x
	return p
}
func (x BainiuDeskStatus) String() string {
	return proto.EnumName(BainiuDeskStatus_name, int32(x))
}
func (x *BainiuDeskStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BainiuDeskStatus_value, data, "BainiuDeskStatus")
	if err != nil {
		return err
	}
	*x = BainiuDeskStatus(value)
	return nil
}
func (BainiuDeskStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// ===============================比牌结果================================
// 比牌类型
type BainiuWinType int32

const (
	BainiuWinType_BAINIU_BIPAI_PINGJU  BainiuWinType = 1
	BainiuWinType_BAINIU_BIPAI_SHENGLI BainiuWinType = 2
	BainiuWinType_BAINIU_BIPAI_SHIBAI  BainiuWinType = 3
	BainiuWinType_BAINIU_BIPAI_ZJTC    BainiuWinType = 4
	BainiuWinType_BAINIU_BIPAI_ZJTP    BainiuWinType = 5
)

var BainiuWinType_name = map[int32]string{
	1: "BAINIU_BIPAI_PINGJU",
	2: "BAINIU_BIPAI_SHENGLI",
	3: "BAINIU_BIPAI_SHIBAI",
	4: "BAINIU_BIPAI_ZJTC",
	5: "BAINIU_BIPAI_ZJTP",
}
var BainiuWinType_value = map[string]int32{
	"BAINIU_BIPAI_PINGJU":  1,
	"BAINIU_BIPAI_SHENGLI": 2,
	"BAINIU_BIPAI_SHIBAI":  3,
	"BAINIU_BIPAI_ZJTC":    4,
	"BAINIU_BIPAI_ZJTP":    5,
}

func (x BainiuWinType) Enum() *BainiuWinType {
	p := new(BainiuWinType)
	*p = x
	return p
}
func (x BainiuWinType) String() string {
	return proto.EnumName(BainiuWinType_name, int32(x))
}
func (x *BainiuWinType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BainiuWinType_value, data, "BainiuWinType")
	if err != nil {
		return err
	}
	*x = BainiuWinType(value)
	return nil
}
func (BainiuWinType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// 押注区域
type BainiuAreaInfo struct {
	Name             *BainiuAreaName    `protobuf:"varint,1,opt,name=name,enum=ddproto.BainiuAreaName" json:"name,omitempty"`
	Poker            *BainiuClientPoker `protobuf:"bytes,2,opt,name=poker" json:"poker,omitempty"`
	AllChips         *int64             `protobuf:"varint,3,opt,name=all_chips,json=allChips" json:"all_chips,omitempty"`
	MyChips          *int64             `protobuf:"varint,4,opt,name=my_chips,json=myChips" json:"my_chips,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *BainiuAreaInfo) Reset()                    { *m = BainiuAreaInfo{} }
func (m *BainiuAreaInfo) String() string            { return proto.CompactTextString(m) }
func (*BainiuAreaInfo) ProtoMessage()               {}
func (*BainiuAreaInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BainiuAreaInfo) GetName() BainiuAreaName {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuAreaInfo) GetPoker() *BainiuClientPoker {
	if m != nil {
		return m.Poker
	}
	return nil
}

func (m *BainiuAreaInfo) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuAreaInfo) GetMyChips() int64 {
	if m != nil && m.MyChips != nil {
		return *m.MyChips
	}
	return 0
}

// 房间状态
type BainiuClientDesk struct {
	Status           *BainiuDeskStatus   `protobuf:"varint,1,opt,name=status,enum=ddproto.BainiuDeskStatus" json:"status,omitempty"`
	Areas            []*BainiuAreaInfo   `protobuf:"bytes,2,rep,name=areas" json:"areas,omitempty"`
	SurplusTime      *int64              `protobuf:"varint,3,opt,name=surplus_time,json=surplusTime" json:"surplus_time,omitempty"`
	Users            []*BainiuClientUser `protobuf:"bytes,4,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *BainiuClientDesk) Reset()                    { *m = BainiuClientDesk{} }
func (m *BainiuClientDesk) String() string            { return proto.CompactTextString(m) }
func (*BainiuClientDesk) ProtoMessage()               {}
func (*BainiuClientDesk) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *BainiuClientDesk) GetStatus() BainiuDeskStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return BainiuDeskStatus_BAINIU_WAIT_YAZHU
}

func (m *BainiuClientDesk) GetAreas() []*BainiuAreaInfo {
	if m != nil {
		return m.Areas
	}
	return nil
}

func (m *BainiuClientDesk) GetSurplusTime() int64 {
	if m != nil && m.SurplusTime != nil {
		return *m.SurplusTime
	}
	return 0
}

func (m *BainiuClientDesk) GetUsers() []*BainiuClientUser {
	if m != nil {
		return m.Users
	}
	return nil
}

// 用户状态
type BainiuClientUser struct {
	UserId           *uint32     `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Image            *string     `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	NickName         *string     `protobuf:"bytes,3,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Coin             *int64      `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	WxInfo           *WeixinInfo `protobuf:"bytes,5,opt,name=wx_info,json=wxInfo" json:"wx_info,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *BainiuClientUser) Reset()                    { *m = BainiuClientUser{} }
func (m *BainiuClientUser) String() string            { return proto.CompactTextString(m) }
func (*BainiuClientUser) ProtoMessage()               {}
func (*BainiuClientUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *BainiuClientUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuClientUser) GetImage() string {
	if m != nil && m.Image != nil {
		return *m.Image
	}
	return ""
}

func (m *BainiuClientUser) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *BainiuClientUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *BainiuClientUser) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

// ===============================进房===================================
// 进房请求
type BainiuEnterDeskReq struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BainiuEnterDeskReq) Reset()                    { *m = BainiuEnterDeskReq{} }
func (m *BainiuEnterDeskReq) String() string            { return proto.CompactTextString(m) }
func (*BainiuEnterDeskReq) ProtoMessage()               {}
func (*BainiuEnterDeskReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BainiuEnterDeskReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 进房回复
type BainiuEnterDeskAck struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskState        *BainiuClientDesk `protobuf:"bytes,2,opt,name=desk_state,json=deskState" json:"desk_state,omitempty"`
	UserInfo         *BainiuClientUser `protobuf:"bytes,3,opt,name=user_info,json=userInfo" json:"user_info,omitempty"`
	GamerSum         *int32            `protobuf:"varint,4,opt,name=gamer_sum,json=gamerSum" json:"gamer_sum,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *BainiuEnterDeskAck) Reset()                    { *m = BainiuEnterDeskAck{} }
func (m *BainiuEnterDeskAck) String() string            { return proto.CompactTextString(m) }
func (*BainiuEnterDeskAck) ProtoMessage()               {}
func (*BainiuEnterDeskAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *BainiuEnterDeskAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BainiuEnterDeskAck) GetDeskState() *BainiuClientDesk {
	if m != nil {
		return m.DeskState
	}
	return nil
}

func (m *BainiuEnterDeskAck) GetUserInfo() *BainiuClientUser {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *BainiuEnterDeskAck) GetGamerSum() int32 {
	if m != nil && m.GamerSum != nil {
		return *m.GamerSum
	}
	return 0
}

// ===============================押注===================================
// 押注overturn
type BainiuYazhuOt struct {
	DeskTime         *int64              `protobuf:"varint,1,opt,name=desk_time,json=deskTime" json:"desk_time,omitempty"`
	Users            []*BainiuClientUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	GamerSum         *int32              `protobuf:"varint,3,opt,name=gamer_sum,json=gamerSum" json:"gamer_sum,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *BainiuYazhuOt) Reset()                    { *m = BainiuYazhuOt{} }
func (m *BainiuYazhuOt) String() string            { return proto.CompactTextString(m) }
func (*BainiuYazhuOt) ProtoMessage()               {}
func (*BainiuYazhuOt) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *BainiuYazhuOt) GetDeskTime() int64 {
	if m != nil && m.DeskTime != nil {
		return *m.DeskTime
	}
	return 0
}

func (m *BainiuYazhuOt) GetUsers() []*BainiuClientUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BainiuYazhuOt) GetGamerSum() int32 {
	if m != nil && m.GamerSum != nil {
		return *m.GamerSum
	}
	return 0
}

// 押注请求
type BainiuYazhuReq struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	AreaName         *BainiuAreaName `protobuf:"varint,2,opt,name=area_name,json=areaName,enum=ddproto.BainiuAreaName" json:"area_name,omitempty"`
	Chips            *uint32         `protobuf:"varint,3,opt,name=chips" json:"chips,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BainiuYazhuReq) Reset()                    { *m = BainiuYazhuReq{} }
func (m *BainiuYazhuReq) String() string            { return proto.CompactTextString(m) }
func (*BainiuYazhuReq) ProtoMessage()               {}
func (*BainiuYazhuReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *BainiuYazhuReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuYazhuReq) GetAreaName() BainiuAreaName {
	if m != nil && m.AreaName != nil {
		return *m.AreaName
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuYazhuReq) GetChips() uint32 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

// 押注回复
type BainiuYazhuAck struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32         `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	AreaName         *BainiuAreaName `protobuf:"varint,3,opt,name=area_name,json=areaName,enum=ddproto.BainiuAreaName" json:"area_name,omitempty"`
	Chips            *int64          `protobuf:"varint,4,opt,name=chips" json:"chips,omitempty"`
	AllChips         *int64          `protobuf:"varint,5,opt,name=all_chips,json=allChips" json:"all_chips,omitempty"`
	SurplusCoin      *int64          `protobuf:"varint,6,opt,name=surplus_coin,json=surplusCoin" json:"surplus_coin,omitempty"`
	MyChips          *int64          `protobuf:"varint,7,opt,name=my_chips,json=myChips" json:"my_chips,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BainiuYazhuAck) Reset()                    { *m = BainiuYazhuAck{} }
func (m *BainiuYazhuAck) String() string            { return proto.CompactTextString(m) }
func (*BainiuYazhuAck) ProtoMessage()               {}
func (*BainiuYazhuAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *BainiuYazhuAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BainiuYazhuAck) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuYazhuAck) GetAreaName() BainiuAreaName {
	if m != nil && m.AreaName != nil {
		return *m.AreaName
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuYazhuAck) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

func (m *BainiuYazhuAck) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuYazhuAck) GetSurplusCoin() int64 {
	if m != nil && m.SurplusCoin != nil {
		return *m.SurplusCoin
	}
	return 0
}

func (m *BainiuYazhuAck) GetMyChips() int64 {
	if m != nil && m.MyChips != nil {
		return *m.MyChips
	}
	return 0
}

// 押注广播
type BainiuYazhuBc struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	AreaName         *BainiuAreaName `protobuf:"varint,2,opt,name=area_name,json=areaName,enum=ddproto.BainiuAreaName" json:"area_name,omitempty"`
	Chips            *int64          `protobuf:"varint,3,opt,name=chips" json:"chips,omitempty"`
	AllChips         *int64          `protobuf:"varint,4,opt,name=all_chips,json=allChips" json:"all_chips,omitempty"`
	SurplusCoin      *int64          `protobuf:"varint,5,opt,name=surplus_coin,json=surplusCoin" json:"surplus_coin,omitempty"`
	MyChips          *int64          `protobuf:"varint,6,opt,name=my_chips,json=myChips" json:"my_chips,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BainiuYazhuBc) Reset()                    { *m = BainiuYazhuBc{} }
func (m *BainiuYazhuBc) String() string            { return proto.CompactTextString(m) }
func (*BainiuYazhuBc) ProtoMessage()               {}
func (*BainiuYazhuBc) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *BainiuYazhuBc) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuYazhuBc) GetAreaName() BainiuAreaName {
	if m != nil && m.AreaName != nil {
		return *m.AreaName
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuYazhuBc) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

func (m *BainiuYazhuBc) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuYazhuBc) GetSurplusCoin() int64 {
	if m != nil && m.SurplusCoin != nil {
		return *m.SurplusCoin
	}
	return 0
}

func (m *BainiuYazhuBc) GetMyChips() int64 {
	if m != nil && m.MyChips != nil {
		return *m.MyChips
	}
	return 0
}

// ===============================发牌=================================
// 发牌广播
type BainiuFapaiBc struct {
	Poker            []*BainiuClientPoker `protobuf:"bytes,1,rep,name=poker" json:"poker,omitempty"`
	DeskTime         *int64               `protobuf:"varint,2,opt,name=desk_time,json=deskTime" json:"desk_time,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *BainiuFapaiBc) Reset()                    { *m = BainiuFapaiBc{} }
func (m *BainiuFapaiBc) String() string            { return proto.CompactTextString(m) }
func (*BainiuFapaiBc) ProtoMessage()               {}
func (*BainiuFapaiBc) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *BainiuFapaiBc) GetPoker() []*BainiuClientPoker {
	if m != nil {
		return m.Poker
	}
	return nil
}

func (m *BainiuFapaiBc) GetDeskTime() int64 {
	if m != nil && m.DeskTime != nil {
		return *m.DeskTime
	}
	return 0
}

// 输赢列表
type BainiuWinUserItem struct {
	NickName         *string `protobuf:"bytes,1,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Bill             *int64  `protobuf:"varint,2,opt,name=bill" json:"bill,omitempty"`
	UserId           *uint32 `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	SurplusCoin      *int64  `protobuf:"varint,4,opt,name=surplus_coin,json=surplusCoin" json:"surplus_coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BainiuWinUserItem) Reset()                    { *m = BainiuWinUserItem{} }
func (m *BainiuWinUserItem) String() string            { return proto.CompactTextString(m) }
func (*BainiuWinUserItem) ProtoMessage()               {}
func (*BainiuWinUserItem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *BainiuWinUserItem) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *BainiuWinUserItem) GetBill() int64 {
	if m != nil && m.Bill != nil {
		return *m.Bill
	}
	return 0
}

func (m *BainiuWinUserItem) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuWinUserItem) GetSurplusCoin() int64 {
	if m != nil && m.SurplusCoin != nil {
		return *m.SurplusCoin
	}
	return 0
}

// 每个区域输赢多少钱
type BainiuAreaBillItem struct {
	AreaName         *BainiuAreaName      `protobuf:"varint,1,opt,name=area_name,json=areaName,enum=ddproto.BainiuAreaName" json:"area_name,omitempty"`
	AllChips         *int64               `protobuf:"varint,2,opt,name=all_chips,json=allChips" json:"all_chips,omitempty"`
	Bill             *int64               `protobuf:"varint,3,opt,name=bill" json:"bill,omitempty"`
	WinUsers         []*BainiuWinUserItem `protobuf:"bytes,4,rep,name=win_users,json=winUsers" json:"win_users,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *BainiuAreaBillItem) Reset()                    { *m = BainiuAreaBillItem{} }
func (m *BainiuAreaBillItem) String() string            { return proto.CompactTextString(m) }
func (*BainiuAreaBillItem) ProtoMessage()               {}
func (*BainiuAreaBillItem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *BainiuAreaBillItem) GetAreaName() BainiuAreaName {
	if m != nil && m.AreaName != nil {
		return *m.AreaName
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuAreaBillItem) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuAreaBillItem) GetBill() int64 {
	if m != nil && m.Bill != nil {
		return *m.Bill
	}
	return 0
}

func (m *BainiuAreaBillItem) GetWinUsers() []*BainiuWinUserItem {
	if m != nil {
		return m.WinUsers
	}
	return nil
}

// 比牌结果广播
type BainiuBipaiBc struct {
	DeskTime         *int64                `protobuf:"varint,1,opt,name=desk_time,json=deskTime" json:"desk_time,omitempty"`
	SurplusCoin      *int64                `protobuf:"varint,2,opt,name=surplus_coin,json=surplusCoin" json:"surplus_coin,omitempty"`
	Type             *BainiuWinType        `protobuf:"varint,3,opt,name=type,enum=ddproto.BainiuWinType" json:"type,omitempty"`
	MyChips          *int64                `protobuf:"varint,4,opt,name=my_chips,json=myChips" json:"my_chips,omitempty"`
	MyBill           *int64                `protobuf:"varint,5,opt,name=my_bill,json=myBill" json:"my_bill,omitempty"`
	AllChips         *int64                `protobuf:"varint,6,opt,name=all_chips,json=allChips" json:"all_chips,omitempty"`
	BankerBill       *int64                `protobuf:"varint,7,opt,name=banker_bill,json=bankerBill" json:"banker_bill,omitempty"`
	Winer            []*BainiuWinUserItem  `protobuf:"bytes,8,rep,name=winer" json:"winer,omitempty"`
	AreaBill         []*BainiuAreaBillItem `protobuf:"bytes,9,rep,name=area_bill,json=areaBill" json:"area_bill,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *BainiuBipaiBc) Reset()                    { *m = BainiuBipaiBc{} }
func (m *BainiuBipaiBc) String() string            { return proto.CompactTextString(m) }
func (*BainiuBipaiBc) ProtoMessage()               {}
func (*BainiuBipaiBc) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *BainiuBipaiBc) GetDeskTime() int64 {
	if m != nil && m.DeskTime != nil {
		return *m.DeskTime
	}
	return 0
}

func (m *BainiuBipaiBc) GetSurplusCoin() int64 {
	if m != nil && m.SurplusCoin != nil {
		return *m.SurplusCoin
	}
	return 0
}

func (m *BainiuBipaiBc) GetType() BainiuWinType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return BainiuWinType_BAINIU_BIPAI_PINGJU
}

func (m *BainiuBipaiBc) GetMyChips() int64 {
	if m != nil && m.MyChips != nil {
		return *m.MyChips
	}
	return 0
}

func (m *BainiuBipaiBc) GetMyBill() int64 {
	if m != nil && m.MyBill != nil {
		return *m.MyBill
	}
	return 0
}

func (m *BainiuBipaiBc) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuBipaiBc) GetBankerBill() int64 {
	if m != nil && m.BankerBill != nil {
		return *m.BankerBill
	}
	return 0
}

func (m *BainiuBipaiBc) GetWiner() []*BainiuWinUserItem {
	if m != nil {
		return m.Winer
	}
	return nil
}

func (m *BainiuBipaiBc) GetAreaBill() []*BainiuAreaBillItem {
	if m != nil {
		return m.AreaBill
	}
	return nil
}

// =============================胜负走势==============================
type BainiuWincountReq struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BainiuWincountReq) Reset()                    { *m = BainiuWincountReq{} }
func (m *BainiuWincountReq) String() string            { return proto.CompactTextString(m) }
func (*BainiuWincountReq) ProtoMessage()               {}
func (*BainiuWincountReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *BainiuWincountReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type BainiuWincountItem struct {
	BankerBill       *int64 `protobuf:"varint,1,opt,name=banker_bill,json=bankerBill" json:"banker_bill,omitempty"`
	TianBill         *int64 `protobuf:"varint,2,opt,name=tian_bill,json=tianBill" json:"tian_bill,omitempty"`
	DiBill           *int64 `protobuf:"varint,3,opt,name=di_bill,json=diBill" json:"di_bill,omitempty"`
	XuanBill         *int64 `protobuf:"varint,4,opt,name=xuan_bill,json=xuanBill" json:"xuan_bill,omitempty"`
	HuangBill        *int64 `protobuf:"varint,5,opt,name=huang_bill,json=huangBill" json:"huang_bill,omitempty"`
	Time             *int64 `protobuf:"varint,6,opt,name=time" json:"time,omitempty"`
	TianIswin        *bool  `protobuf:"varint,7,opt,name=tian_iswin,json=tianIswin" json:"tian_iswin,omitempty"`
	DiIswin          *bool  `protobuf:"varint,8,opt,name=di_iswin,json=diIswin" json:"di_iswin,omitempty"`
	XuanIswin        *bool  `protobuf:"varint,9,opt,name=xuan_iswin,json=xuanIswin" json:"xuan_iswin,omitempty"`
	HuangIswin       *bool  `protobuf:"varint,10,opt,name=huang_iswin,json=huangIswin" json:"huang_iswin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BainiuWincountItem) Reset()                    { *m = BainiuWincountItem{} }
func (m *BainiuWincountItem) String() string            { return proto.CompactTextString(m) }
func (*BainiuWincountItem) ProtoMessage()               {}
func (*BainiuWincountItem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *BainiuWincountItem) GetBankerBill() int64 {
	if m != nil && m.BankerBill != nil {
		return *m.BankerBill
	}
	return 0
}

func (m *BainiuWincountItem) GetTianBill() int64 {
	if m != nil && m.TianBill != nil {
		return *m.TianBill
	}
	return 0
}

func (m *BainiuWincountItem) GetDiBill() int64 {
	if m != nil && m.DiBill != nil {
		return *m.DiBill
	}
	return 0
}

func (m *BainiuWincountItem) GetXuanBill() int64 {
	if m != nil && m.XuanBill != nil {
		return *m.XuanBill
	}
	return 0
}

func (m *BainiuWincountItem) GetHuangBill() int64 {
	if m != nil && m.HuangBill != nil {
		return *m.HuangBill
	}
	return 0
}

func (m *BainiuWincountItem) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *BainiuWincountItem) GetTianIswin() bool {
	if m != nil && m.TianIswin != nil {
		return *m.TianIswin
	}
	return false
}

func (m *BainiuWincountItem) GetDiIswin() bool {
	if m != nil && m.DiIswin != nil {
		return *m.DiIswin
	}
	return false
}

func (m *BainiuWincountItem) GetXuanIswin() bool {
	if m != nil && m.XuanIswin != nil {
		return *m.XuanIswin
	}
	return false
}

func (m *BainiuWincountItem) GetHuangIswin() bool {
	if m != nil && m.HuangIswin != nil {
		return *m.HuangIswin
	}
	return false
}

type BainiuWincountAck struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Rows             []*BainiuWincountItem `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *BainiuWincountAck) Reset()                    { *m = BainiuWincountAck{} }
func (m *BainiuWincountAck) String() string            { return proto.CompactTextString(m) }
func (*BainiuWincountAck) ProtoMessage()               {}
func (*BainiuWincountAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *BainiuWincountAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BainiuWincountAck) GetRows() []*BainiuWincountItem {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*BainiuAreaInfo)(nil), "ddproto.bainiu_area_info")
	proto.RegisterType((*BainiuClientDesk)(nil), "ddproto.bainiu_client_desk")
	proto.RegisterType((*BainiuClientUser)(nil), "ddproto.bainiu_client_user")
	proto.RegisterType((*BainiuEnterDeskReq)(nil), "ddproto.bainiu_enter_desk_req")
	proto.RegisterType((*BainiuEnterDeskAck)(nil), "ddproto.bainiu_enter_desk_ack")
	proto.RegisterType((*BainiuYazhuOt)(nil), "ddproto.bainiu_yazhu_ot")
	proto.RegisterType((*BainiuYazhuReq)(nil), "ddproto.bainiu_yazhu_req")
	proto.RegisterType((*BainiuYazhuAck)(nil), "ddproto.bainiu_yazhu_ack")
	proto.RegisterType((*BainiuYazhuBc)(nil), "ddproto.bainiu_yazhu_bc")
	proto.RegisterType((*BainiuFapaiBc)(nil), "ddproto.bainiu_fapai_bc")
	proto.RegisterType((*BainiuWinUserItem)(nil), "ddproto.bainiu_win_user_item")
	proto.RegisterType((*BainiuAreaBillItem)(nil), "ddproto.bainiu_area_bill_item")
	proto.RegisterType((*BainiuBipaiBc)(nil), "ddproto.bainiu_bipai_bc")
	proto.RegisterType((*BainiuWincountReq)(nil), "ddproto.bainiu_wincount_req")
	proto.RegisterType((*BainiuWincountItem)(nil), "ddproto.bainiu_wincount_item")
	proto.RegisterType((*BainiuWincountAck)(nil), "ddproto.bainiu_wincount_ack")
	proto.RegisterEnum("ddproto.BainiuDeskStatus", BainiuDeskStatus_name, BainiuDeskStatus_value)
	proto.RegisterEnum("ddproto.BainiuWinType", BainiuWinType_name, BainiuWinType_value)
}

func init() { proto.RegisterFile("bainiu_play.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1075 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0x29, 0x51, 0x22, 0x47, 0x4d, 0xeb, 0xac, 0x9d, 0x5a, 0x89, 0xeb, 0xc6, 0xe5, 0xc9,
	0x08, 0x1c, 0xb7, 0x71, 0x80, 0xa2, 0x48, 0x4f, 0xb6, 0xd1, 0xc6, 0x0a, 0x0a, 0x43, 0xa0, 0x6d,
	0x04, 0xc9, 0x85, 0xa0, 0xc4, 0x8d, 0xbd, 0xb0, 0x48, 0xaa, 0xfc, 0xa9, 0xac, 0x1e, 0x7a, 0xc8,
	0xb5, 0x6f, 0xd1, 0x7b, 0xcf, 0xbd, 0xf4, 0x21, 0x0a, 0xf4, 0xd2, 0x7b, 0x5f, 0xa4, 0x98, 0xd9,
	0x25, 0x45, 0x52, 0x3f, 0x75, 0x0a, 0xe4, 0x22, 0x91, 0x33, 0xdf, 0xec, 0x7e, 0xf3, 0xcd, 0xcc,
	0x2e, 0xe1, 0xee, 0xc0, 0x13, 0xa1, 0xc8, 0xdc, 0xf1, 0xc8, 0x9b, 0xee, 0x8f, 0xe3, 0x28, 0x8d,
	0x58, 0xdb, 0xf7, 0xe9, 0xe1, 0xc1, 0xfa, 0x30, 0x0a, 0x82, 0x28, 0x74, 0x87, 0x23, 0xc1, 0xc3,
	0x54, 0x7a, 0x1f, 0xe4, 0x01, 0x03, 0x2f, 0xe1, 0xd2, 0x64, 0xff, 0xa6, 0xc1, 0x9a, 0xb2, 0x7a,
	0x31, 0xf7, 0x5c, 0x11, 0xbe, 0x89, 0xd8, 0x63, 0x68, 0x86, 0x5e, 0xc0, 0xbb, 0xda, 0x8e, 0xb6,
	0xfb, 0xd1, 0xc1, 0xfd, 0x7d, 0xb5, 0xe8, 0x7e, 0x19, 0x88, 0x00, 0x87, 0x60, 0xec, 0x00, 0x8c,
	0x71, 0x74, 0xcd, 0xe3, 0xae, 0xbe, 0xa3, 0xed, 0x76, 0x0e, 0x3e, 0xad, 0xe3, 0x25, 0x07, 0x97,
	0x30, 0x8e, 0x84, 0xb2, 0x2d, 0xb0, 0xbc, 0xd1, 0xc8, 0x1d, 0x5e, 0x89, 0x71, 0xd2, 0x6d, 0xec,
	0x68, 0xbb, 0x0d, 0xc7, 0xf4, 0x46, 0xa3, 0x63, 0x7c, 0x67, 0xf7, 0xc1, 0x0c, 0xa6, 0xca, 0xd7,
	0x24, 0x5f, 0x3b, 0x98, 0x92, 0xcb, 0xfe, 0x53, 0x03, 0x56, 0x5d, 0xd6, 0xe7, 0xc9, 0x35, 0x7b,
	0x0a, 0xad, 0x24, 0xf5, 0xd2, 0x2c, 0x51, 0x9c, 0xb7, 0xea, 0x1c, 0x10, 0xe5, 0x4a, 0x88, 0xa3,
	0xa0, 0xec, 0x0b, 0x30, 0x30, 0x95, 0xa4, 0xab, 0xef, 0x34, 0x76, 0x3b, 0x4b, 0xf2, 0x44, 0x41,
	0x1c, 0x89, 0x63, 0x9f, 0xc3, 0x87, 0x49, 0x16, 0x8f, 0x47, 0x59, 0xe2, 0xa6, 0x22, 0xe0, 0x8a,
	0x77, 0x47, 0xd9, 0xce, 0x45, 0xc0, 0xd9, 0x13, 0x30, 0xb2, 0x84, 0xc7, 0xc8, 0x1b, 0xd7, 0xdc,
	0x5a, 0xa2, 0x05, 0x62, 0x1c, 0x89, 0xb4, 0x7f, 0x9d, 0x4b, 0x09, 0xed, 0x6c, 0x13, 0xda, 0xf8,
	0xef, 0x0a, 0x9f, 0x72, 0xba, 0xe3, 0xb4, 0xf0, 0xb5, 0xe7, 0xb3, 0x0d, 0x30, 0x44, 0xe0, 0x5d,
	0x72, 0x92, 0xdb, 0x72, 0xe4, 0x0b, 0x0a, 0x1a, 0x8a, 0xe1, 0x35, 0xd5, 0x85, 0x88, 0x59, 0x8e,
	0x89, 0x86, 0x53, 0xac, 0x10, 0x83, 0xe6, 0x30, 0x12, 0xa1, 0x12, 0x93, 0x9e, 0xd9, 0x1e, 0xb4,
	0x27, 0x37, 0x94, 0x5e, 0xd7, 0xa0, 0xba, 0xad, 0x17, 0x5c, 0x5f, 0x72, 0x71, 0x23, 0xc2, 0x1e,
	0x66, 0xde, 0x9a, 0xdc, 0xe0, 0xbf, 0xfd, 0x25, 0xdc, 0x53, 0x1c, 0x79, 0x98, 0xf2, 0x58, 0xea,
	0x19, 0xf3, 0x1f, 0x96, 0xd2, 0xb4, 0xff, 0xd6, 0x16, 0x85, 0x78, 0xc3, 0x6b, 0xb6, 0x07, 0xad,
	0x2b, 0xee, 0xf9, 0x3c, 0xa6, 0x88, 0xce, 0xc1, 0x46, 0xb1, 0x71, 0x1f, 0x7f, 0x4f, 0xc8, 0xe7,
	0x28, 0x0c, 0x7b, 0x06, 0x50, 0x14, 0x8f, 0xab, 0x16, 0x5b, 0x26, 0x2b, 0x02, 0x1d, 0x0b, 0x7f,
	0xcf, 0x10, 0xcd, 0xbe, 0x06, 0x4b, 0x92, 0xc3, 0x2c, 0x1b, 0x2b, 0x43, 0xa9, 0x22, 0x26, 0x71,
	0xc7, 0x11, 0xd8, 0x02, 0xeb, 0xd2, 0x0b, 0x78, 0xec, 0x26, 0x59, 0x40, 0xb2, 0x19, 0x8e, 0x49,
	0x86, 0xb3, 0x2c, 0xb0, 0x7f, 0x86, 0x8f, 0x55, 0xf0, 0xd4, 0xfb, 0xe9, 0x2a, 0x73, 0xa3, 0x14,
	0xf1, 0xc4, 0x92, 0xfa, 0x42, 0x93, 0xfd, 0x8c, 0x86, 0x6a, 0x53, 0xe8, 0xb7, 0x6d, 0x8a, 0xea,
	0xfe, 0x8d, 0xda, 0xfe, 0xd3, 0x62, 0x66, 0xe5, 0xfe, 0xab, 0xea, 0xc0, 0xbe, 0x02, 0xab, 0x18,
	0x58, 0x92, 0x6f, 0xe5, 0x44, 0x9b, 0xf8, 0x48, 0x3d, 0xb3, 0x01, 0xc6, 0x6c, 0x3a, 0xef, 0x38,
	0xf2, 0xc5, 0x7e, 0xab, 0xd7, 0xf6, 0x7e, 0xf7, 0x82, 0x96, 0x98, 0xea, 0xcb, 0x99, 0x36, 0xfe,
	0x07, 0x53, 0xd9, 0xde, 0xf2, 0xa5, 0x7a, 0xc2, 0x18, 0xb5, 0x13, 0xa6, 0x34, 0xc9, 0x34, 0x18,
	0xad, 0xca, 0x24, 0x1f, 0xe3, 0x7c, 0x94, 0x0f, 0xa1, 0x76, 0xf5, 0x10, 0xfa, 0x4b, 0xab, 0x35,
	0xc0, 0x60, 0xf8, 0x9e, 0xf5, 0x5f, 0x9c, 0x55, 0xf3, 0x3f, 0xb2, 0x32, 0x56, 0x67, 0xd5, 0xaa,
	0x66, 0x35, 0x28, 0x92, 0x7a, 0xe3, 0x8d, 0x3d, 0x81, 0x49, 0x15, 0x27, 0xbb, 0x46, 0x8d, 0x7b,
	0xdb, 0x93, 0x7d, 0x36, 0x09, 0x7a, 0x75, 0x12, 0xec, 0xb7, 0x1a, 0x6c, 0xa8, 0xd8, 0x89, 0x08,
	0x5d, 0xa9, 0x58, 0xca, 0x83, 0xea, 0xf1, 0xa5, 0xcd, 0x1f, 0x5f, 0x03, 0x31, 0x1a, 0xa9, 0xd5,
	0xe8, 0x99, 0x7d, 0x02, 0x4a, 0x60, 0xd5, 0x9f, 0xb9, 0xdc, 0x75, 0x0d, 0x9a, 0x73, 0x1a, 0xd8,
	0x7f, 0xcc, 0x4e, 0x26, 0x12, 0x1e, 0xd7, 0x93, 0x2c, 0x2a, 0xb5, 0xd2, 0x6e, 0x5f, 0xab, 0x4a,
	0x55, 0xf4, 0x5a, 0x55, 0x72, 0xf6, 0x8d, 0x12, 0xfb, 0x67, 0x60, 0xe5, 0xf9, 0xe7, 0x57, 0xc5,
	0x76, 0x7d, 0xa3, 0x8a, 0x40, 0x8e, 0x39, 0x11, 0xe1, 0x05, 0xdd, 0x17, 0xff, 0xe8, 0x45, 0xa1,
	0x06, 0x42, 0x15, 0x6a, 0xe5, 0xf1, 0x53, 0x97, 0x44, 0x9f, 0x6f, 0x8b, 0x3d, 0x68, 0xa6, 0xd3,
	0x71, 0x3e, 0x75, 0xdd, 0x45, 0x54, 0xd0, 0xef, 0x10, 0x6a, 0xc5, 0xfd, 0x8c, 0x63, 0x10, 0x4c,
	0x49, 0x51, 0xd5, 0x7d, 0xad, 0x60, 0x7a, 0x84, 0x19, 0x57, 0x24, 0x6a, 0xd5, 0x24, 0x7a, 0x08,
	0x9d, 0x81, 0x17, 0x5e, 0xf3, 0x58, 0x46, 0xca, 0x71, 0x03, 0x69, 0xa2, 0xe8, 0xa7, 0x60, 0x4c,
	0x44, 0xc8, 0xe3, 0xae, 0x79, 0x1b, 0xad, 0x24, 0x96, 0x7d, 0xa3, 0xaa, 0x49, 0x6b, 0x5a, 0x14,
	0xf8, 0xd9, 0xc2, 0x6a, 0x16, 0x0d, 0x20, 0x4b, 0x8a, 0x3b, 0xda, 0x8f, 0x61, 0x7d, 0xb6, 0xf6,
	0x30, 0xca, 0xc2, 0x94, 0x8e, 0xd9, 0x59, 0xdb, 0x55, 0x6f, 0xbb, 0xdf, 0xf5, 0x72, 0x63, 0x4b,
	0x3c, 0xb5, 0x54, 0x2d, 0x35, 0x6d, 0x2e, 0xb5, 0x2d, 0xb0, 0x52, 0xe1, 0x85, 0x6e, 0xa9, 0xc3,
	0x4d, 0x34, 0x90, 0x73, 0x13, 0xda, 0xbe, 0x70, 0x4b, 0xed, 0xd3, 0xf2, 0x45, 0x1e, 0x75, 0x93,
	0xe5, 0x51, 0xea, 0x1c, 0x40, 0x03, 0x39, 0xb7, 0x01, 0xae, 0x32, 0x2f, 0xbc, 0x2c, 0xd7, 0xc1,
	0x22, 0x0b, 0xb9, 0x19, 0x34, 0xa9, 0x4f, 0x64, 0x15, 0xe8, 0x19, 0x43, 0x88, 0x85, 0x48, 0x26,
	0x22, 0xa4, 0x02, 0x98, 0x0e, 0xf1, 0xea, 0xa1, 0x01, 0x2b, 0xee, 0x0b, 0xe5, 0x34, 0xc9, 0xd9,
	0xf6, 0x85, 0x74, 0x6d, 0x03, 0x10, 0x13, 0xe9, 0xb4, 0x64, 0x24, 0x5a, 0xa4, 0xfb, 0x21, 0x74,
	0x24, 0x17, 0xe9, 0x07, 0xf2, 0x4b, 0x7a, 0x04, 0xb0, 0x7f, 0x9c, 0x17, 0xfa, 0xdd, 0xef, 0x94,
	0x27, 0xd0, 0x8c, 0xa3, 0x49, 0x7e, 0xc1, 0x2e, 0x6a, 0x8f, 0x59, 0x49, 0x1c, 0x82, 0x3e, 0x7a,
	0x55, 0x7c, 0x75, 0x95, 0xbe, 0x0d, 0xd9, 0x3d, 0xb8, 0x7b, 0x74, 0xd8, 0x3b, 0xed, 0x5d, 0xb8,
	0x2f, 0x0f, 0x7b, 0xe7, 0xee, 0xab, 0xc3, 0xd7, 0x27, 0x17, 0x6b, 0x5a, 0xdd, 0xfc, 0xdd, 0x61,
	0xff, 0xb0, 0xb7, 0xa6, 0xd7, 0xcd, 0x47, 0x3d, 0x34, 0x37, 0x1e, 0xfd, 0x32, 0xbb, 0x1f, 0xf2,
	0xc9, 0x61, 0x9b, 0xb0, 0xae, 0xa0, 0x84, 0x72, 0xfb, 0xbd, 0xd3, 0xe7, 0x2f, 0x70, 0xe9, 0x2e,
	0x6c, 0x54, 0x1c, 0x67, 0x27, 0xdf, 0x9e, 0x3e, 0xff, 0x1e, 0x57, 0xaf, 0x87, 0x9c, 0x9d, 0xf4,
	0x8e, 0x70, 0xfd, 0xd2, 0xb6, 0xd2, 0xf1, 0xfa, 0xc5, 0xf9, 0xf1, 0x5a, 0x73, 0x91, 0xb9, 0xbf,
	0x66, 0xf4, 0x3f, 0xe8, 0x6b, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xd4, 0x75, 0x11, 0x2b,
	0x0c, 0x00, 0x00,
}

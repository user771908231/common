// Code generated by protoc-gen-go.
// source: zjh_play.proto
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

// Ignoring public import of EProtoId from zjh_base.proto

// Ignoring public import of zjh_enum_playerGameStatus from zjh_base.proto

// Ignoring public import of zjh_enum_deskState from zjh_base.proto

// Ignoring public import of zjh_enum_userState from zjh_base.proto

// Ignoring public import of zjh_enum_roomType from zjh_base.proto

// Ignoring public import of zjh_req_getRoomList from zjh_desk.proto

// Ignoring public import of zjh_base_roomInfo from zjh_desk.proto

// Ignoring public import of zjh_ack_roomList from zjh_desk.proto

// Ignoring public import of zjh_req_enterDesk from zjh_desk.proto

// Ignoring public import of zjh_base_userInfo from zjh_desk.proto

// Ignoring public import of ZjhDeskStateAck from zjh_desk.proto

// 游戏信息(广播)（接收服务端消息）(别的玩家已看牌是独立协议更好吧? 否则通过此协议下发是否过于冗余?)
type ZjhBcGameInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       []*ZjhBasePlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	ZjhDeskInfo      *ZjhBaseDeskInfo     `protobuf:"bytes,3,opt,name=zjhDeskInfo" json:"zjhDeskInfo,omitempty"`
	SenderUserId     *uint32              `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *int32               `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhBcGameInfo) Reset()                    { *m = ZjhBcGameInfo{} }
func (m *ZjhBcGameInfo) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcGameInfo) ProtoMessage()               {}
func (*ZjhBcGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *ZjhBcGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcGameInfo) GetPlayerInfo() []*ZjhBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *ZjhBcGameInfo) GetZjhDeskInfo() *ZjhBaseDeskInfo {
	if m != nil {
		return m.ZjhDeskInfo
	}
	return nil
}

func (m *ZjhBcGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *ZjhBcGameInfo) GetIsReconnect() int32 {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return 0
}

// 新加入玩家后 桌子广播需要的协议
type ZjhBaseDeskInfo struct {
	GameStatus       *int32           `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	RoomInfo         *ZjhBaseRoomInfo `protobuf:"bytes,2,opt,name=roomInfo" json:"roomInfo,omitempty"`
	PlayerNum        *int32           `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveUserId     *uint32          `protobuf:"varint,4,opt,name=activeUserId" json:"activeUserId,omitempty"`
	ActionTime       *int32           `protobuf:"varint,5,opt,name=actionTime" json:"actionTime,omitempty"`
	NInitActionTime  *int32           `protobuf:"varint,6,opt,name=nInitActionTime" json:"nInitActionTime,omitempty"`
	InitRoomCoin     *int64           `protobuf:"varint,7,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32           `protobuf:"varint,8,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32           `protobuf:"varint,9,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	RoomNumber       *string          `protobuf:"bytes,10,opt,name=roomNumber" json:"roomNumber,omitempty"`
	RoomOwnerUserId  *uint32          `protobuf:"varint,11,opt,name=roomOwnerUserId" json:"roomOwnerUserId,omitempty"`
	PlayRate         *int32           `protobuf:"varint,12,opt,name=playRate" json:"playRate,omitempty"`
	CurrRoundCount   *int32           `protobuf:"varint,13,opt,name=currRoundCount" json:"currRoundCount,omitempty"`
	TotalRoundCount  *int32           `protobuf:"varint,14,opt,name=totalRoundCount" json:"totalRoundCount,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ZjhBaseDeskInfo) Reset()                    { *m = ZjhBaseDeskInfo{} }
func (m *ZjhBaseDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*ZjhBaseDeskInfo) ProtoMessage()               {}
func (*ZjhBaseDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *ZjhBaseDeskInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetRoomInfo() *ZjhBaseRoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

func (m *ZjhBaseDeskInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetActionTime() int32 {
	if m != nil && m.ActionTime != nil {
		return *m.ActionTime
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetNInitActionTime() int32 {
	if m != nil && m.NInitActionTime != nil {
		return *m.NInitActionTime
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *ZjhBaseDeskInfo) GetRoomOwnerUserId() uint32 {
	if m != nil && m.RoomOwnerUserId != nil {
		return *m.RoomOwnerUserId
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetCurrRoundCount() int32 {
	if m != nil && m.CurrRoundCount != nil {
		return *m.CurrRoundCount
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetTotalRoundCount() int32 {
	if m != nil && m.TotalRoundCount != nil {
		return *m.TotalRoundCount
	}
	return 0
}

// 有玩家进入房间(服务器广播)
type ZjhBcNewPlayerEnter struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       *ZjhBasePlayerInfo `protobuf:"bytes,2,opt,name=playerInfo" json:"playerInfo,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhBcNewPlayerEnter) Reset()                    { *m = ZjhBcNewPlayerEnter{} }
func (m *ZjhBcNewPlayerEnter) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcNewPlayerEnter) ProtoMessage()               {}
func (*ZjhBcNewPlayerEnter) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *ZjhBcNewPlayerEnter) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcNewPlayerEnter) GetPlayerInfo() *ZjhBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

// 离开房间
type ZjhReqLeave struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhReqLeave) Reset()                    { *m = ZjhReqLeave{} }
func (m *ZjhReqLeave) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqLeave) ProtoMessage()               {}
func (*ZjhReqLeave) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *ZjhReqLeave) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 离开房间的广播(广播给所有人, 让所有人知道)
type ZjhBcLeave struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       *ZjhBasePlayerInfo `protobuf:"bytes,2,opt,name=playerInfo" json:"playerInfo,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhBcLeave) Reset()                    { *m = ZjhBcLeave{} }
func (m *ZjhBcLeave) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcLeave) ProtoMessage()               {}
func (*ZjhBcLeave) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *ZjhBcLeave) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcLeave) GetPlayerInfo() *ZjhBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

// 单个玩家信息
type ZjhBasePlayerInfo struct {
	IsFirst          *bool                    `protobuf:"varint,1,opt,name=isFirst" json:"isFirst,omitempty"`
	PlayerPokers     []*ClientBasePoker       `protobuf:"bytes,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	Coin             *int64                   `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	NickName         *string                  `protobuf:"bytes,4,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32                   `protobuf:"varint,5,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32                  `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	BReady           *int32                   `protobuf:"varint,8,opt,name=bReady" json:"bReady,omitempty"`
	Status           *ZjhEnumPlayerGameStatus `protobuf:"varint,9,opt,name=status,enum=ddproto.ZjhEnumPlayerGameStatus" json:"status,omitempty"`
	WxInfo           *WeixinInfo              `protobuf:"bytes,10,opt,name=wxInfo" json:"wxInfo,omitempty"`
	OnlineStatus     *int32                   `protobuf:"varint,11,opt,name=onlineStatus" json:"onlineStatus,omitempty"`
	IsCheckedPokers  *bool                    `protobuf:"varint,12,opt,name=isCheckedPokers" json:"isCheckedPokers,omitempty"`
	SeatIndex        *int32                   `protobuf:"varint,13,opt,name=seatIndex" json:"seatIndex,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ZjhBasePlayerInfo) Reset()                    { *m = ZjhBasePlayerInfo{} }
func (m *ZjhBasePlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*ZjhBasePlayerInfo) ProtoMessage()               {}
func (*ZjhBasePlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *ZjhBasePlayerInfo) GetIsFirst() bool {
	if m != nil && m.IsFirst != nil {
		return *m.IsFirst
	}
	return false
}

func (m *ZjhBasePlayerInfo) GetPlayerPokers() []*ClientBasePoker {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

func (m *ZjhBasePlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *ZjhBasePlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetStatus() ZjhEnumPlayerGameStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ZjhEnumPlayerGameStatus_ZJH_TEMP
}

func (m *ZjhBasePlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *ZjhBasePlayerInfo) GetOnlineStatus() int32 {
	if m != nil && m.OnlineStatus != nil {
		return *m.OnlineStatus
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetIsCheckedPokers() bool {
	if m != nil && m.IsCheckedPokers != nil {
		return *m.IsCheckedPokers
	}
	return false
}

func (m *ZjhBasePlayerInfo) GetSeatIndex() int32 {
	if m != nil && m.SeatIndex != nil {
		return *m.SeatIndex
	}
	return 0
}

// 开局（接收服务端消息）: 客户端收到消息后就开始【霸底】、【发牌】的动画, 然后处理第一个 OverTurn
type ZjhBcOpening struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfoList   []*ZjhBasePlayerInfo `protobuf:"bytes,2,rep,name=playerInfoList" json:"playerInfoList,omitempty"`
	BaseAnte         *int64               `protobuf:"varint,3,opt,name=baseAnte" json:"baseAnte,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhBcOpening) Reset()                    { *m = ZjhBcOpening{} }
func (m *ZjhBcOpening) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcOpening) ProtoMessage()               {}
func (*ZjhBcOpening) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *ZjhBcOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcOpening) GetPlayerInfoList() []*ZjhBasePlayerInfo {
	if m != nil {
		return m.PlayerInfoList
	}
	return nil
}

func (m *ZjhBcOpening) GetBaseAnte() int64 {
	if m != nil && m.BaseAnte != nil {
		return *m.BaseAnte
	}
	return 0
}

// 轮到谁操作（接收服务端消息）
type ZjhBcOverTurn struct {
	Header            *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CurOprtPlayerInfo *ZjhBasePlayerInfo `protobuf:"bytes,2,opt,name=curOprtPlayerInfo" json:"curOprtPlayerInfo,omitempty"`
	IsXuepin          *bool              `protobuf:"varint,3,opt,name=isXuepin" json:"isXuepin,omitempty"`
	Time              *int32             `protobuf:"varint,4,opt,name=time" json:"time,omitempty"`
	TotalTime         *int32             `protobuf:"varint,5,opt,name=totalTime" json:"totalTime,omitempty"`
	XXX_unrecognized  []byte             `json:"-"`
}

func (m *ZjhBcOverTurn) Reset()                    { *m = ZjhBcOverTurn{} }
func (m *ZjhBcOverTurn) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcOverTurn) ProtoMessage()               {}
func (*ZjhBcOverTurn) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *ZjhBcOverTurn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcOverTurn) GetCurOprtPlayerInfo() *ZjhBasePlayerInfo {
	if m != nil {
		return m.CurOprtPlayerInfo
	}
	return nil
}

func (m *ZjhBcOverTurn) GetIsXuepin() bool {
	if m != nil && m.IsXuepin != nil {
		return *m.IsXuepin
	}
	return false
}

func (m *ZjhBcOverTurn) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *ZjhBcOverTurn) GetTotalTime() int32 {
	if m != nil && m.TotalTime != nil {
		return *m.TotalTime
	}
	return 0
}

// 玩家点击看牌
type ZjhReqCheckCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhReqCheckCards) Reset()                    { *m = ZjhReqCheckCards{} }
func (m *ZjhReqCheckCards) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqCheckCards) ProtoMessage()               {}
func (*ZjhReqCheckCards) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *ZjhReqCheckCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 玩家点击看牌 广播 bc
type ZjhBcCheckCards struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	HandPokers       []*ClientBasePoker `protobuf:"bytes,2,rep,name=handPokers" json:"handPokers,omitempty"`
	HandPokerType    *string            `protobuf:"bytes,3,opt,name=handPokerType" json:"handPokerType,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhBcCheckCards) Reset()                    { *m = ZjhBcCheckCards{} }
func (m *ZjhBcCheckCards) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcCheckCards) ProtoMessage()               {}
func (*ZjhBcCheckCards) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *ZjhBcCheckCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcCheckCards) GetHandPokers() []*ClientBasePoker {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *ZjhBcCheckCards) GetHandPokerType() string {
	if m != nil && m.HandPokerType != nil {
		return *m.HandPokerType
	}
	return ""
}

// =========================弃牌=========================
// 玩家弃牌
type ZjhReqFold struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhReqFold) Reset()                    { *m = ZjhReqFold{} }
func (m *ZjhReqFold) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqFold) ProtoMessage()               {}
func (*ZjhReqFold) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *ZjhReqFold) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 玩家弃牌 广播 bc
type ZjhBcFold struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhBcFold) Reset()                    { *m = ZjhBcFold{} }
func (m *ZjhBcFold) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcFold) ProtoMessage()               {}
func (*ZjhBcFold) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{11} }

func (m *ZjhBcFold) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// ======================跟注=======================
// 玩家跟注
type ZjhReqCall struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhReqCall) Reset()                    { *m = ZjhReqCall{} }
func (m *ZjhReqCall) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqCall) ProtoMessage()               {}
func (*ZjhReqCall) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{12} }

func (m *ZjhReqCall) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 玩家跟注 广播 bc
type ZjhBcCall struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Chips            *int64       `protobuf:"varint,2,opt,name=chips" json:"chips,omitempty"`
	IsXuepin         *bool        `protobuf:"varint,3,opt,name=isXuepin" json:"isXuepin,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhBcCall) Reset()                    { *m = ZjhBcCall{} }
func (m *ZjhBcCall) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcCall) ProtoMessage()               {}
func (*ZjhBcCall) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{13} }

func (m *ZjhBcCall) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcCall) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

func (m *ZjhBcCall) GetIsXuepin() bool {
	if m != nil && m.IsXuepin != nil {
		return *m.IsXuepin
	}
	return false
}

// =====================血拼===========================
// 发起血拼
type ZjhReqBloodFight struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhReqBloodFight) Reset()                    { *m = ZjhReqBloodFight{} }
func (m *ZjhReqBloodFight) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqBloodFight) ProtoMessage()               {}
func (*ZjhReqBloodFight) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{14} }

func (m *ZjhReqBloodFight) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 玩家发起血拼 广播 bc  ，如果一人发起血拼则说明血拼开始，直到血拼结束
type ZjhBcBloodFight struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Chips            *int64       `protobuf:"varint,2,opt,name=chips" json:"chips,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhBcBloodFight) Reset()                    { *m = ZjhBcBloodFight{} }
func (m *ZjhBcBloodFight) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcBloodFight) ProtoMessage()               {}
func (*ZjhBcBloodFight) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{15} }

func (m *ZjhBcBloodFight) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcBloodFight) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

// 血拼结束广播
type ZjhBcBloodEnd struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhBcBloodEnd) Reset()                    { *m = ZjhBcBloodEnd{} }
func (m *ZjhBcBloodEnd) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcBloodEnd) ProtoMessage()               {}
func (*ZjhBcBloodEnd) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{16} }

func (m *ZjhBcBloodEnd) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// ====================加注========================
// 玩家加注
type ZjhReqRaiseFight struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Chips            *int64       `protobuf:"varint,2,opt,name=chips" json:"chips,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhReqRaiseFight) Reset()                    { *m = ZjhReqRaiseFight{} }
func (m *ZjhReqRaiseFight) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqRaiseFight) ProtoMessage()               {}
func (*ZjhReqRaiseFight) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{17} }

func (m *ZjhReqRaiseFight) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhReqRaiseFight) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

// 玩家加注 广播 bc
type ZjhBcRaiseAck struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Chips            *int64       `protobuf:"varint,2,opt,name=chips" json:"chips,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhBcRaiseAck) Reset()                    { *m = ZjhBcRaiseAck{} }
func (m *ZjhBcRaiseAck) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcRaiseAck) ProtoMessage()               {}
func (*ZjhBcRaiseAck) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{18} }

func (m *ZjhBcRaiseAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcRaiseAck) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

// ===================比牌、孤注一掷===================
// 玩家比牌
type ZjhReqCompare struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ComparedUserId   *uint32      `protobuf:"varint,2,opt,name=comparedUserId" json:"comparedUserId,omitempty"`
	IsAllIn          *bool        `protobuf:"varint,3,opt,name=isAllIn" json:"isAllIn,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhReqCompare) Reset()                    { *m = ZjhReqCompare{} }
func (m *ZjhReqCompare) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqCompare) ProtoMessage()               {}
func (*ZjhReqCompare) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{19} }

func (m *ZjhReqCompare) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhReqCompare) GetComparedUserId() uint32 {
	if m != nil && m.ComparedUserId != nil {
		return *m.ComparedUserId
	}
	return 0
}

func (m *ZjhReqCompare) GetIsAllIn() bool {
	if m != nil && m.IsAllIn != nil {
		return *m.IsAllIn
	}
	return false
}

// 玩家比牌 广播 bc
type ZjhBcCompare struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ComparedUserId   *uint32      `protobuf:"varint,2,opt,name=comparedUserId" json:"comparedUserId,omitempty"`
	IsCompareWin     *bool        `protobuf:"varint,3,opt,name=isCompareWin" json:"isCompareWin,omitempty"`
	IsAllIn          *bool        `protobuf:"varint,4,opt,name=isAllIn" json:"isAllIn,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhBcCompare) Reset()                    { *m = ZjhBcCompare{} }
func (m *ZjhBcCompare) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcCompare) ProtoMessage()               {}
func (*ZjhBcCompare) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{20} }

func (m *ZjhBcCompare) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcCompare) GetComparedUserId() uint32 {
	if m != nil && m.ComparedUserId != nil {
		return *m.ComparedUserId
	}
	return 0
}

func (m *ZjhBcCompare) GetIsCompareWin() bool {
	if m != nil && m.IsCompareWin != nil {
		return *m.IsCompareWin
	}
	return false
}

func (m *ZjhBcCompare) GetIsAllIn() bool {
	if m != nil && m.IsAllIn != nil {
		return *m.IsAllIn
	}
	return false
}

// ==================牌局结束==============================
// 牌局结束信息
type ZjhBcGameEnd struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Winer            *uint32      `protobuf:"varint,2,opt,name=winer" json:"winer,omitempty"`
	Chip             *int64       `protobuf:"varint,3,opt,name=chip" json:"chip,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhBcGameEnd) Reset()                    { *m = ZjhBcGameEnd{} }
func (m *ZjhBcGameEnd) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcGameEnd) ProtoMessage()               {}
func (*ZjhBcGameEnd) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{21} }

func (m *ZjhBcGameEnd) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcGameEnd) GetWiner() uint32 {
	if m != nil && m.Winer != nil {
		return *m.Winer
	}
	return 0
}

func (m *ZjhBcGameEnd) GetChip() int64 {
	if m != nil && m.Chip != nil {
		return *m.Chip
	}
	return 0
}

func init() {
	proto.RegisterType((*ZjhBcGameInfo)(nil), "ddproto.zjh_bc_gameInfo")
	proto.RegisterType((*ZjhBaseDeskInfo)(nil), "ddproto.zjh_base_deskInfo")
	proto.RegisterType((*ZjhBcNewPlayerEnter)(nil), "ddproto.zjh_bc_newPlayerEnter")
	proto.RegisterType((*ZjhReqLeave)(nil), "ddproto.zjh_req_leave")
	proto.RegisterType((*ZjhBcLeave)(nil), "ddproto.zjh_bc_leave")
	proto.RegisterType((*ZjhBasePlayerInfo)(nil), "ddproto.zjh_base_playerInfo")
	proto.RegisterType((*ZjhBcOpening)(nil), "ddproto.zjh_bc_opening")
	proto.RegisterType((*ZjhBcOverTurn)(nil), "ddproto.zjh_bc_overTurn")
	proto.RegisterType((*ZjhReqCheckCards)(nil), "ddproto.zjh_req_checkCards")
	proto.RegisterType((*ZjhBcCheckCards)(nil), "ddproto.zjh_bc_checkCards")
	proto.RegisterType((*ZjhReqFold)(nil), "ddproto.zjh_req_fold")
	proto.RegisterType((*ZjhBcFold)(nil), "ddproto.zjh_bc_fold")
	proto.RegisterType((*ZjhReqCall)(nil), "ddproto.zjh_req_call")
	proto.RegisterType((*ZjhBcCall)(nil), "ddproto.zjh_bc_call")
	proto.RegisterType((*ZjhReqBloodFight)(nil), "ddproto.zjh_req_bloodFight")
	proto.RegisterType((*ZjhBcBloodFight)(nil), "ddproto.zjh_bc_bloodFight")
	proto.RegisterType((*ZjhBcBloodEnd)(nil), "ddproto.zjh_bc_bloodEnd")
	proto.RegisterType((*ZjhReqRaiseFight)(nil), "ddproto.zjh_req_raiseFight")
	proto.RegisterType((*ZjhBcRaiseAck)(nil), "ddproto.zjh_bc_raiseAck")
	proto.RegisterType((*ZjhReqCompare)(nil), "ddproto.zjh_req_compare")
	proto.RegisterType((*ZjhBcCompare)(nil), "ddproto.zjh_bc_compare")
	proto.RegisterType((*ZjhBcGameEnd)(nil), "ddproto.zjh_bc_gameEnd")
}

var fileDescriptor14 = []byte{
	// 1009 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x66, 0xe3, 0xd8, 0xb1, 0xc7, 0x8e, 0x4b, 0x27, 0x05, 0x59, 0x51, 0x85, 0xac, 0x15, 0x42,
	0x96, 0x40, 0xb9, 0xc8, 0x05, 0x17, 0x25, 0x80, 0x82, 0xdb, 0x82, 0x11, 0x4a, 0xad, 0x21, 0x55,
	0x7a, 0x67, 0x8d, 0x77, 0x4f, 0xb3, 0x53, 0xef, 0xce, 0x2c, 0x33, 0xb3, 0xf9, 0xbb, 0xab, 0x78,
	0x11, 0x10, 0x3c, 0x03, 0xef, 0xc2, 0xdb, 0xa0, 0x99, 0x1d, 0xef, 0x8f, 0x53, 0x50, 0x37, 0x2d,
	0x37, 0xd1, 0x9e, 0xcf, 0xdf, 0x99, 0xf9, 0xce, 0x99, 0x73, 0xce, 0x4c, 0xd0, 0xf0, 0xe6, 0x55,
	0xb4, 0x48, 0x63, 0x7a, 0x7d, 0x90, 0x4a, 0xa1, 0x05, 0xde, 0x09, 0x43, 0xfb, 0xb1, 0xbf, 0x17,
	0x88, 0x24, 0x11, 0x7c, 0x11, 0xc4, 0x0c, 0xb8, 0xce, 0x7f, 0xdd, 0xb7, 0xec, 0x25, 0x55, 0x50,
	0xb5, 0x43, 0x50, 0xab, 0xdc, 0xf6, 0x5f, 0x6f, 0xa1, 0x7b, 0x96, 0x12, 0x2c, 0xce, 0x69, 0x02,
	0x33, 0xfe, 0x52, 0xe0, 0x2f, 0x50, 0x27, 0x02, 0x1a, 0x82, 0x1c, 0x79, 0x63, 0x6f, 0xd2, 0x3f,
	0x7c, 0x70, 0xe0, 0xb6, 0x38, 0x98, 0x9b, 0xbf, 0x3f, 0xd8, 0xdf, 0x88, 0xe3, 0xe0, 0x23, 0x84,
	0x8c, 0x1a, 0x90, 0xc6, 0x77, 0xb4, 0x35, 0x6e, 0x4d, 0xfa, 0x87, 0x0f, 0x0b, 0x8f, 0xf5, 0xf6,
	0x8b, 0x92, 0x43, 0x2a, 0x7c, 0x7c, 0x84, 0xfa, 0x37, 0xaf, 0xa2, 0xc7, 0xa0, 0x56, 0xd6, 0xbd,
	0x65, 0x37, 0xdc, 0xbf, 0xed, 0x1e, 0x3a, 0x06, 0xa9, 0xd2, 0xb1, 0x8f, 0x06, 0x0a, 0x78, 0x08,
	0xf2, 0xb9, 0x02, 0x39, 0x0b, 0x47, 0xdb, 0x63, 0x6f, 0xb2, 0x4b, 0x6a, 0x18, 0x1e, 0xa3, 0x3e,
	0x53, 0x04, 0x02, 0xc1, 0x39, 0x04, 0x7a, 0xd4, 0x1e, 0x7b, 0x93, 0x36, 0xa9, 0x42, 0xfe, 0xef,
	0xdb, 0xe8, 0xfe, 0xad, 0x8d, 0xf0, 0x27, 0x08, 0x7d, 0x4f, 0x13, 0xf8, 0x59, 0x53, 0x9d, 0x29,
	0x9b, 0x89, 0x36, 0xa9, 0x20, 0xf8, 0x4b, 0xd4, 0x95, 0x42, 0x24, 0x2e, 0xea, 0x7f, 0x91, 0xbd,
	0x66, 0x90, 0x82, 0x8b, 0x1f, 0xa2, 0x5e, 0x1e, 0xff, 0x49, 0x96, 0xd8, 0x78, 0xdb, 0xa4, 0x04,
	0x4c, 0x44, 0x34, 0xd0, 0xec, 0x02, 0xea, 0x11, 0x55, 0x31, 0xa3, 0xcc, 0xd8, 0x82, 0x9f, 0xb2,
	0x04, 0x5c, 0x40, 0x15, 0x04, 0x4f, 0xd0, 0x3d, 0x3e, 0xe3, 0x4c, 0x1f, 0x97, 0xa4, 0x8e, 0x25,
	0x6d, 0xc2, 0x66, 0x37, 0xc6, 0x99, 0x26, 0x42, 0x24, 0x53, 0xc1, 0xf8, 0x68, 0x67, 0xec, 0x4d,
	0x5a, 0xa4, 0x86, 0xe1, 0x4f, 0xd1, 0x6e, 0x90, 0x49, 0x39, 0x8f, 0xe9, 0xf5, 0x54, 0x64, 0x5c,
	0x8f, 0xba, 0x76, 0xad, 0x3a, 0x88, 0x3f, 0x43, 0x43, 0x2d, 0x34, 0x8d, 0x4b, 0x5a, 0xcf, 0xd2,
	0x36, 0x50, 0xa3, 0xdd, 0x64, 0xe2, 0x24, 0x4b, 0x96, 0x20, 0x47, 0x68, 0xec, 0x4d, 0x7a, 0xa4,
	0x82, 0x18, 0xed, 0xc6, 0x7a, 0x76, 0xc9, 0x8b, 0x43, 0xed, 0xdb, 0x14, 0x6c, 0xc2, 0x78, 0x1f,
	0x75, 0x4d, 0xda, 0x08, 0xd5, 0x30, 0x1a, 0xd8, 0xbd, 0x0a, 0xdb, 0xa8, 0x31, 0xf2, 0x88, 0xc8,
	0x78, 0x98, 0xab, 0xd9, 0xcd, 0xd5, 0xd4, 0x51, 0xb3, 0x9b, 0xd5, 0x57, 0x21, 0x0e, 0xf3, 0x4c,
	0x6d, 0xc0, 0xfe, 0xaf, 0x1e, 0xfa, 0xc8, 0xf5, 0x09, 0x87, 0xcb, 0xb9, 0x3d, 0xaf, 0x27, 0x5c,
	0x83, 0x7c, 0xc7, 0x6e, 0xf1, 0x9a, 0x74, 0x8b, 0xff, 0x35, 0xda, 0x35, 0x14, 0x09, 0xbf, 0x2c,
	0x62, 0xa0, 0x17, 0xd0, 0x6c, 0x73, 0xff, 0x06, 0x0d, 0x5c, 0x0c, 0x77, 0xf0, 0x7e, 0x47, 0xe9,
	0x7f, 0xb5, 0xd0, 0xde, 0x1b, 0x38, 0x78, 0x84, 0x76, 0x98, 0x7a, 0xca, 0xa4, 0xd2, 0x56, 0x44,
	0x97, 0xac, 0x4d, 0xfc, 0x0d, 0x1a, 0xe4, 0xbc, 0xb9, 0x58, 0x81, 0x54, 0x6e, 0xb4, 0x94, 0x4d,
	0x96, 0xcf, 0x39, 0xb7, 0xa0, 0xa1, 0x90, 0x1a, 0x1f, 0x63, 0xb4, 0x1d, 0x98, 0xa2, 0x6e, 0xd9,
	0xa2, 0xb6, 0xdf, 0xa6, 0x68, 0x38, 0x0b, 0x56, 0x27, 0x34, 0x01, 0xdb, 0x5a, 0x3d, 0x52, 0xd8,
	0xf8, 0x43, 0xd4, 0x52, 0x70, 0xe5, 0xfa, 0xc9, 0x7c, 0xe2, 0x8f, 0x51, 0x27, 0xcb, 0x6b, 0xb0,
	0x63, 0x6b, 0xd0, 0x59, 0x06, 0x5f, 0x12, 0xa0, 0xe1, 0xb5, 0xeb, 0x05, 0x67, 0xe1, 0x47, 0xa8,
	0xa3, 0xf2, 0x71, 0x61, 0x8a, 0x7f, 0x78, 0xe8, 0xd7, 0xb2, 0x03, 0x3c, 0x4b, 0x5c, 0xe4, 0xe5,
	0x18, 0x21, 0xce, 0x03, 0x7f, 0x8e, 0x3a, 0x97, 0x57, 0x36, 0xb3, 0xc8, 0x66, 0x76, 0xaf, 0xf0,
	0x3d, 0x03, 0x76, 0xc5, 0xb8, 0x4d, 0xa8, 0xa3, 0x98, 0xbe, 0x15, 0x3c, 0x66, 0x7c, 0x3d, 0x9d,
	0xfa, 0x56, 0x46, 0x0d, 0x33, 0xb5, 0xcd, 0xd4, 0x34, 0x82, 0x60, 0x05, 0xa1, 0xcb, 0xe0, 0xc0,
	0x26, 0x78, 0x13, 0x36, 0x13, 0x49, 0x01, 0xd5, 0x33, 0x1e, 0xc2, 0x95, 0x6b, 0x94, 0x12, 0xf0,
	0x7f, 0xf3, 0xf2, 0x2b, 0x67, 0x19, 0x2c, 0x44, 0x0a, 0x9c, 0xf1, 0xf3, 0x86, 0x75, 0xf3, 0x18,
	0x0d, 0xcb, 0xf3, 0xfe, 0x89, 0x29, 0xfd, 0x56, 0x97, 0xc4, 0x86, 0x8f, 0x39, 0x39, 0x43, 0x39,
	0xe6, 0x1a, 0xdc, 0x89, 0x16, 0xb6, 0xff, 0xb7, 0x57, 0x5c, 0x62, 0xe2, 0x02, 0xe4, 0x69, 0x26,
	0x79, 0x43, 0x8d, 0x3f, 0xa2, 0xfb, 0x41, 0x26, 0x9f, 0xa5, 0x52, 0xcf, 0x9b, 0x95, 0xf8, 0x6d,
	0x37, 0xa3, 0x94, 0xa9, 0x17, 0x19, 0xa4, 0xae, 0xf6, 0xba, 0xa4, 0xb0, 0x4d, 0x4d, 0x6a, 0xe6,
	0x6a, 0xaf, 0x4d, 0xec, 0xb7, 0x49, 0xbf, 0x9d, 0x36, 0x95, 0x69, 0x5e, 0x02, 0xfe, 0x77, 0x08,
	0xaf, 0x5b, 0x3e, 0x30, 0xa7, 0x36, 0xa5, 0x32, 0x54, 0x0d, 0xfb, 0xfe, 0x0f, 0xcf, 0x5d, 0x70,
	0xc1, 0x9d, 0xd7, 0xc0, 0x8f, 0x10, 0x8a, 0x28, 0x0f, 0xdf, 0xba, 0x17, 0x2b, 0x6c, 0x73, 0x85,
	0x14, 0xd6, 0xe9, 0x75, 0x9a, 0x1f, 0x60, 0x8f, 0xd4, 0x41, 0xff, 0x28, 0x9f, 0x4e, 0x26, 0xd2,
	0x97, 0x22, 0x0e, 0x1b, 0xc6, 0xf8, 0x95, 0x7d, 0x48, 0x98, 0x10, 0xef, 0xe0, 0x5c, 0xd9, 0x3a,
	0xa0, 0x71, 0xdc, 0xd0, 0x3b, 0x29, 0xb6, 0x6e, 0xee, 0x8c, 0x1f, 0xa0, 0x76, 0x10, 0xb1, 0x54,
	0xd9, 0x6a, 0x6b, 0x91, 0xdc, 0xf8, 0xaf, 0x1a, 0xaa, 0x56, 0xc4, 0x32, 0x16, 0x22, 0x7c, 0xca,
	0xce, 0x23, 0xdd, 0x50, 0xf2, 0x59, 0x51, 0x10, 0x77, 0x5d, 0xe2, 0xcd, 0xc2, 0xfd, 0x6f, 0x8b,
	0x4e, 0xb4, 0x0b, 0x3f, 0xe1, 0x4d, 0x8f, 0xe2, 0x45, 0x19, 0x9d, 0xa4, 0x4c, 0xc1, 0xfb, 0x93,
	0xf6, 0xbc, 0x90, 0x66, 0x17, 0x3e, 0x0e, 0x56, 0xef, 0x65, 0xd9, 0xd7, 0x6e, 0xf8, 0xd8, 0xe2,
	0x11, 0x49, 0x4a, 0x65, 0xd3, 0x8b, 0xd5, 0xbc, 0x56, 0x72, 0xc7, 0xd0, 0x3d, 0x79, 0xb6, 0xec,
	0x75, 0xb3, 0x81, 0xe6, 0x57, 0xe5, 0x71, 0x1c, 0xcf, 0xd6, 0x35, 0xb1, 0x36, 0xfd, 0x3f, 0xcb,
	0x19, 0xfd, 0xff, 0x4a, 0x30, 0x0f, 0x46, 0x35, 0xcd, 0xb1, 0xb3, 0xa2, 0x36, 0x6b, 0x58, 0x55,
	0xe6, 0x76, 0x5d, 0x66, 0x54, 0xa8, 0x34, 0xff, 0x6b, 0x34, 0xae, 0x0d, 0x73, 0x00, 0x97, 0x8c,
	0x83, 0x74, 0xe2, 0x72, 0xc3, 0xde, 0xf3, 0x11, 0x4b, 0x8b, 0x7b, 0x3e, 0x62, 0xe9, 0xfc, 0x83,
	0xb9, 0x37, 0xdf, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x19, 0xef, 0x96, 0x2b, 0x0d, 0x00,
	0x00,
}
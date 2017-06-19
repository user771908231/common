// Code generated by protoc-gen-go.
// source: hall_playback_pdk.proto
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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of pdk_base_roomTypeInfo from common_pdk.proto

// Ignoring public import of pdk_base_timerInfo from common_pdk.proto

// Ignoring public import of pdk_enum_paiType from common_pdk.proto

// Ignoring public import of pdk_enum_playerStatus from common_pdk.proto

// Ignoring public import of pdk_enum_roomType from common_pdk.proto

// Ignoring public import of pdk_enum_enterType from common_pdk.proto

// Ignoring public import of pdk_enum_coinRoomLevel from common_pdk.proto

// Ignoring public import of pdk_enum_deskGameStatus from common_pdk.proto

// 跑得快的操作类型
type PlaybackPdkActType int32

const (
	PlaybackPdkActType_PPAT_GUO   PlaybackPdkActType = 1
	PlaybackPdkActType_PPAT_DAPAI PlaybackPdkActType = 2
)

var PlaybackPdkActType_name = map[int32]string{
	1: "PPAT_GUO",
	2: "PPAT_DAPAI",
}
var PlaybackPdkActType_value = map[string]int32{
	"PPAT_GUO":   1,
	"PPAT_DAPAI": 2,
}

func (x PlaybackPdkActType) Enum() *PlaybackPdkActType {
	p := new(PlaybackPdkActType)
	*p = x
	return p
}
func (x PlaybackPdkActType) String() string {
	return proto.EnumName(PlaybackPdkActType_name, int32(x))
}
func (x *PlaybackPdkActType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlaybackPdkActType_value, data, "PlaybackPdkActType")
	if err != nil {
		return err
	}
	*x = PlaybackPdkActType(value)
	return nil
}
func (PlaybackPdkActType) EnumDescriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

type PlaybackPdkAckPage struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 分页信息
	Length *int32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Page   *int32 `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	Total  *int32 `protobuf:"varint,4,opt,name=total" json:"total,omitempty"`
	// 快照信息
	PlaybackSnapshots []*PdkPlaybackSnapshot `protobuf:"bytes,5,rep,name=playbackSnapshots" json:"playbackSnapshots,omitempty"`
	XXX_unrecognized  []byte                 `json:"-"`
}

func (m *PlaybackPdkAckPage) Reset()                    { *m = PlaybackPdkAckPage{} }
func (m *PlaybackPdkAckPage) String() string            { return proto.CompactTextString(m) }
func (*PlaybackPdkAckPage) ProtoMessage()               {}
func (*PlaybackPdkAckPage) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *PlaybackPdkAckPage) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PlaybackPdkAckPage) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetPage() int32 {
	if m != nil && m.Page != nil {
		return *m.Page
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetTotal() int32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetPlaybackSnapshots() []*PdkPlaybackSnapshot {
	if m != nil {
		return m.PlaybackSnapshots
	}
	return nil
}

// 玩家的信息
type PdkPlayerInfo struct {
	PlayerPokers     []*ClientBasePoker   `protobuf:"bytes,1,rep,name=playerPokers" json:"playerPokers,omitempty"`
	Coin             *int64               `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	NickName         *string              `protobuf:"bytes,3,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32               `protobuf:"varint,4,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32              `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool                `protobuf:"varint,6,opt,name=isOwner" json:"isOwner,omitempty"`
	BReady           *int32               `protobuf:"varint,7,opt,name=bReady" json:"bReady,omitempty"`
	Status           *PdkEnumPlayerStatus `protobuf:"varint,8,opt,name=status,enum=ddproto.PdkEnumPlayerStatus" json:"status,omitempty"`
	WxInfo           *WeixinInfo          `protobuf:"bytes,9,opt,name=wxInfo" json:"wxInfo,omitempty"`
	OnlineStatus     *int32               `protobuf:"varint,10,opt,name=onlineStatus" json:"onlineStatus,omitempty"`
	RemainPaiCount   *int32               `protobuf:"varint,11,opt,name=remainPaiCount" json:"remainPaiCount,omitempty"`
	OutPais          []*ClientBasePoker   `protobuf:"bytes,12,rep,name=outPais" json:"outPais,omitempty"`
	OutPaiType       *PdkEnumPaiType      `protobuf:"varint,13,opt,name=outPaiType,enum=ddproto.PdkEnumPaiType" json:"outPaiType,omitempty"`
	RoomCard         *int64               `protobuf:"varint,14,opt,name=roomCard" json:"roomCard,omitempty"`
	IsBanker         *bool                `protobuf:"varint,15,opt,name=isBanker" json:"isBanker,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkPlayerInfo) Reset()                    { *m = PdkPlayerInfo{} }
func (m *PdkPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkPlayerInfo) ProtoMessage()               {}
func (*PdkPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

func (m *PdkPlayerInfo) GetPlayerPokers() []*ClientBasePoker {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

func (m *PdkPlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PdkPlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PdkPlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *PdkPlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkPlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PdkPlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *PdkPlayerInfo) GetStatus() PdkEnumPlayerStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return PdkEnumPlayerStatus_PDK_NORMAL_DEFAULT
}

func (m *PdkPlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PdkPlayerInfo) GetOnlineStatus() int32 {
	if m != nil && m.OnlineStatus != nil {
		return *m.OnlineStatus
	}
	return 0
}

func (m *PdkPlayerInfo) GetRemainPaiCount() int32 {
	if m != nil && m.RemainPaiCount != nil {
		return *m.RemainPaiCount
	}
	return 0
}

func (m *PdkPlayerInfo) GetOutPais() []*ClientBasePoker {
	if m != nil {
		return m.OutPais
	}
	return nil
}

func (m *PdkPlayerInfo) GetOutPaiType() PdkEnumPaiType {
	if m != nil && m.OutPaiType != nil {
		return *m.OutPaiType
	}
	return PdkEnumPaiType_PDK_ERRORCARD
}

func (m *PdkPlayerInfo) GetRoomCard() int64 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *PdkPlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

// 回放快照信息（包含玩家信息、桌面信息、房间信息、右上角时间信息等），另外包含协议信息（PID 和 二进制数据）
type PdkPlaybackSnapshot struct {
	PlayerInfo       []*PdkPlayerInfo     `protobuf:"bytes,1,rep,name=PlayerInfo" json:"PlayerInfo,omitempty"`
	PlayBackDeskInfo *PdkPlaybackDeskInfo `protobuf:"bytes,2,opt,name=PlayBackDeskInfo" json:"PlayBackDeskInfo,omitempty"`
	RoomTypeInfo     *PdkBaseRoomTypeInfo `protobuf:"bytes,3,opt,name=RoomTypeInfo" json:"RoomTypeInfo,omitempty"`
	Time             *string              `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	ActType          *PlaybackPdkActType  `protobuf:"varint,5,opt,name=actType,enum=ddproto.PlaybackPdkActType" json:"actType,omitempty"`
	Userid           *uint32              `protobuf:"varint,6,opt,name=userid" json:"userid,omitempty"`
	OutPais          []*ClientBasePoker   `protobuf:"bytes,7,rep,name=outPais" json:"outPais,omitempty"`
	OutPaiType       *PdkEnumPaiType      `protobuf:"varint,8,opt,name=outPaiType,enum=ddproto.PdkEnumPaiType" json:"outPaiType,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkPlaybackSnapshot) Reset()                    { *m = PdkPlaybackSnapshot{} }
func (m *PdkPlaybackSnapshot) String() string            { return proto.CompactTextString(m) }
func (*PdkPlaybackSnapshot) ProtoMessage()               {}
func (*PdkPlaybackSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{2} }

func (m *PdkPlaybackSnapshot) GetPlayerInfo() []*PdkPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetPlayBackDeskInfo() *PdkPlaybackDeskInfo {
	if m != nil {
		return m.PlayBackDeskInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetRoomTypeInfo() *PdkBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

func (m *PdkPlaybackSnapshot) GetActType() PlaybackPdkActType {
	if m != nil && m.ActType != nil {
		return *m.ActType
	}
	return PlaybackPdkActType_PPAT_GUO
}

func (m *PdkPlaybackSnapshot) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *PdkPlaybackSnapshot) GetOutPais() []*ClientBasePoker {
	if m != nil {
		return m.OutPais
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetOutPaiType() PdkEnumPaiType {
	if m != nil && m.OutPaiType != nil {
		return *m.OutPaiType
	}
	return PdkEnumPaiType_PDK_ERRORCARD
}

// 回放桌面信息
type PdkPlaybackDeskInfo struct {
	GameStatus       *int32            `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	PlayerNum        *int32            `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveUserId     *uint32           `protobuf:"varint,4,opt,name=activeUserId" json:"activeUserId,omitempty"`
	CurrPlayCount    *int32            `protobuf:"varint,10,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32            `protobuf:"varint,11,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	RoomNumber       *string           `protobuf:"bytes,12,opt,name=roomNumber" json:"roomNumber,omitempty"`
	BankerId         *uint32           `protobuf:"varint,14,opt,name=bankerId" json:"bankerId,omitempty"`
	EnterType        *PdkEnumEnterType `protobuf:"varint,18,opt,name=enterType,enum=ddproto.PdkEnumEnterType" json:"enterType,omitempty"`
	CoinTicket       *int64            `protobuf:"varint,19,opt,name=coinTicket" json:"coinTicket,omitempty"`
	TimerInfo        *PdkBaseTimerInfo `protobuf:"bytes,20,opt,name=timerInfo" json:"timerInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *PdkPlaybackDeskInfo) Reset()                    { *m = PdkPlaybackDeskInfo{} }
func (m *PdkPlaybackDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkPlaybackDeskInfo) ProtoMessage()               {}
func (*PdkPlaybackDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{3} }

func (m *PdkPlaybackDeskInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *PdkPlaybackDeskInfo) GetBankerId() uint32 {
	if m != nil && m.BankerId != nil {
		return *m.BankerId
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetEnterType() PdkEnumEnterType {
	if m != nil && m.EnterType != nil {
		return *m.EnterType
	}
	return PdkEnumEnterType_PDK_FRIEND
}

func (m *PdkPlaybackDeskInfo) GetCoinTicket() int64 {
	if m != nil && m.CoinTicket != nil {
		return *m.CoinTicket
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetTimerInfo() *PdkBaseTimerInfo {
	if m != nil {
		return m.TimerInfo
	}
	return nil
}

type PdkDataRecoverPlayerInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	HandPokers       []int32 `protobuf:"varint,2,rep,name=HandPokers" json:"HandPokers,omitempty"`
	UserGameStatus   *int32  `protobuf:"varint,3,opt,name=UserGameStatus" json:"UserGameStatus,omitempty"`
	DissolveStatus   *int32  `protobuf:"varint,4,opt,name=DissolveStatus" json:"DissolveStatus,omitempty"`
	OutPaiList       []int32 `protobuf:"varint,5,rep,name=OutPaiList" json:"OutPaiList,omitempty"`
	OnlineStatus     *bool   `protobuf:"varint,6,opt,name=OnlineStatus" json:"OnlineStatus,omitempty"`
	IsBigWinner      *bool   `protobuf:"varint,7,opt,name=isBigWinner" json:"isBigWinner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,8,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64  `protobuf:"varint,9,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PdkDataRecoverPlayerInfo) Reset()                    { *m = PdkDataRecoverPlayerInfo{} }
func (m *PdkDataRecoverPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkDataRecoverPlayerInfo) ProtoMessage()               {}
func (*PdkDataRecoverPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{4} }

func (m *PdkDataRecoverPlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetHandPokers() []int32 {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *PdkDataRecoverPlayerInfo) GetUserGameStatus() int32 {
	if m != nil && m.UserGameStatus != nil {
		return *m.UserGameStatus
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetDissolveStatus() int32 {
	if m != nil && m.DissolveStatus != nil {
		return *m.DissolveStatus
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetOutPaiList() []int32 {
	if m != nil {
		return m.OutPaiList
	}
	return nil
}

func (m *PdkDataRecoverPlayerInfo) GetOnlineStatus() bool {
	if m != nil && m.OnlineStatus != nil {
		return *m.OnlineStatus
	}
	return false
}

func (m *PdkDataRecoverPlayerInfo) GetIsBigWinner() bool {
	if m != nil && m.IsBigWinner != nil {
		return *m.IsBigWinner
	}
	return false
}

func (m *PdkDataRecoverPlayerInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

type PdkDataRecoverDeskInfo struct {
	DeskGameStatus   *int32            `protobuf:"varint,1,opt,name=DeskGameStatus" json:"DeskGameStatus,omitempty"`
	PlayerNum        *int32            `protobuf:"varint,2,opt,name=PlayerNum" json:"PlayerNum,omitempty"`
	ActiveUserId     *uint32           `protobuf:"varint,3,opt,name=ActiveUserId" json:"ActiveUserId,omitempty"`
	IsFollowPai      *bool             `protobuf:"varint,4,opt,name=isFollowPai" json:"isFollowPai,omitempty"`
	CurrentRound     *int32            `protobuf:"varint,5,opt,name=CurrentRound" json:"CurrentRound,omitempty"`
	TotalRound       *int32            `protobuf:"varint,6,opt,name=TotalRound" json:"TotalRound,omitempty"`
	TimerInfo        *PdkBaseTimerInfo `protobuf:"bytes,7,opt,name=timerInfo" json:"timerInfo,omitempty"`
	LatestOutPokers  []int32           `protobuf:"varint,8,rep,name=LatestOutPokers" json:"LatestOutPokers,omitempty"`
	RoomNumber       *string           `protobuf:"bytes,9,opt,name=RoomNumber" json:"RoomNumber,omitempty"`
	Owner            *uint32           `protobuf:"varint,10,opt,name=owner" json:"owner,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *PdkDataRecoverDeskInfo) Reset()                    { *m = PdkDataRecoverDeskInfo{} }
func (m *PdkDataRecoverDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkDataRecoverDeskInfo) ProtoMessage()               {}
func (*PdkDataRecoverDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{5} }

func (m *PdkDataRecoverDeskInfo) GetDeskGameStatus() int32 {
	if m != nil && m.DeskGameStatus != nil {
		return *m.DeskGameStatus
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetIsFollowPai() bool {
	if m != nil && m.IsFollowPai != nil {
		return *m.IsFollowPai
	}
	return false
}

func (m *PdkDataRecoverDeskInfo) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetTotalRound() int32 {
	if m != nil && m.TotalRound != nil {
		return *m.TotalRound
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetTimerInfo() *PdkBaseTimerInfo {
	if m != nil {
		return m.TimerInfo
	}
	return nil
}

func (m *PdkDataRecoverDeskInfo) GetLatestOutPokers() []int32 {
	if m != nil {
		return m.LatestOutPokers
	}
	return nil
}

func (m *PdkDataRecoverDeskInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *PdkDataRecoverDeskInfo) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

type PdkDataRecover struct {
	PlayerInfo       []*PdkDataRecoverPlayerInfo `protobuf:"bytes,1,rep,name=playerInfo" json:"playerInfo,omitempty"`
	DeskInfo         *PdkDataRecoverDeskInfo     `protobuf:"bytes,2,opt,name=deskInfo" json:"deskInfo,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *PdkDataRecover) Reset()                    { *m = PdkDataRecover{} }
func (m *PdkDataRecover) String() string            { return proto.CompactTextString(m) }
func (*PdkDataRecover) ProtoMessage()               {}
func (*PdkDataRecover) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{6} }

func (m *PdkDataRecover) GetPlayerInfo() []*PdkDataRecoverPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *PdkDataRecover) GetDeskInfo() *PdkDataRecoverDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

type PdkDataRecoverDeskIdList struct {
	DeskId           []int32 `protobuf:"varint,1,rep,name=DeskId" json:"DeskId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PdkDataRecoverDeskIdList) Reset()                    { *m = PdkDataRecoverDeskIdList{} }
func (m *PdkDataRecoverDeskIdList) String() string            { return proto.CompactTextString(m) }
func (*PdkDataRecoverDeskIdList) ProtoMessage()               {}
func (*PdkDataRecoverDeskIdList) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{7} }

func (m *PdkDataRecoverDeskIdList) GetDeskId() []int32 {
	if m != nil {
		return m.DeskId
	}
	return nil
}

func init() {
	proto.RegisterType((*PlaybackPdkAckPage)(nil), "ddproto.playback_pdk_ack_page")
	proto.RegisterType((*PdkPlayerInfo)(nil), "ddproto.PdkPlayerInfo")
	proto.RegisterType((*PdkPlaybackSnapshot)(nil), "ddproto.PdkPlaybackSnapshot")
	proto.RegisterType((*PdkPlaybackDeskInfo)(nil), "ddproto.PdkPlaybackDeskInfo")
	proto.RegisterType((*PdkDataRecoverPlayerInfo)(nil), "ddproto.PdkDataRecoverPlayerInfo")
	proto.RegisterType((*PdkDataRecoverDeskInfo)(nil), "ddproto.PdkDataRecoverDeskInfo")
	proto.RegisterType((*PdkDataRecover)(nil), "ddproto.PdkDataRecover")
	proto.RegisterType((*PdkDataRecoverDeskIdList)(nil), "ddproto.PdkDataRecoverDeskIdList")
	proto.RegisterEnum("ddproto.PlaybackPdkActType", PlaybackPdkActType_name, PlaybackPdkActType_value)
}

var fileDescriptor30 = []byte{
	// 1065 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0x24, 0x35,
	0x13, 0xfe, 0x3b, 0x93, 0x39, 0x55, 0x26, 0xd9, 0xfc, 0xce, 0x12, 0x4c, 0x76, 0xb5, 0x0c, 0x23,
	0x84, 0x46, 0x80, 0x72, 0x11, 0xc1, 0x8a, 0x15, 0x12, 0x52, 0x0e, 0x62, 0x13, 0xb4, 0x4a, 0x5a,
	0xde, 0xac, 0xf6, 0x72, 0xe4, 0x74, 0x9b, 0xc4, 0xea, 0x69, 0x7b, 0xd4, 0xed, 0xc9, 0xe1, 0x9a,
	0x17, 0xe0, 0x8e, 0xf7, 0x40, 0x82, 0x17, 0xe1, 0x85, 0x50, 0x95, 0x7b, 0xfa, 0x30, 0x13, 0x84,
	0x96, 0x9b, 0xc4, 0xf5, 0xf9, 0x73, 0xb5, 0xab, 0xea, 0xab, 0xf2, 0xc0, 0xc7, 0x37, 0x72, 0x3a,
	0x9d, 0xcc, 0xa6, 0xf2, 0xe1, 0x4a, 0x46, 0xc9, 0x64, 0x16, 0x27, 0xfb, 0xb3, 0xcc, 0x3a, 0xcb,
	0xba, 0x71, 0x4c, 0x8b, 0xbd, 0x9d, 0xc8, 0xa6, 0xa9, 0x35, 0x93, 0x68, 0xaa, 0x95, 0x71, 0x7e,
	0x77, 0x6f, 0xbb, 0x00, 0x4b, 0xfe, 0xe8, 0xaf, 0x00, 0x3e, 0xaa, 0xbb, 0x99, 0xd0, 0x7f, 0x79,
	0xad, 0xd8, 0xd7, 0xd0, 0xb9, 0x51, 0x32, 0x56, 0x19, 0x0f, 0x86, 0xc1, 0x78, 0xe3, 0xe0, 0xe9,
	0x7e, 0xe1, 0x7a, 0x3f, 0xc4, 0xbf, 0xa7, 0xb4, 0x27, 0x0a, 0x0e, 0xdb, 0x85, 0xce, 0x54, 0x99,
	0x6b, 0x77, 0xc3, 0xd7, 0x86, 0xc1, 0xb8, 0x2d, 0x0a, 0x8b, 0x31, 0x58, 0x47, 0x6f, 0xbc, 0x45,
	0x28, 0xad, 0xd9, 0x53, 0x68, 0x3b, 0xeb, 0xe4, 0x94, 0xaf, 0x13, 0xe8, 0x0d, 0xf6, 0x13, 0xfc,
	0x7f, 0x71, 0x91, 0xb7, 0x46, 0xce, 0xf2, 0x1b, 0xeb, 0x72, 0xde, 0x1e, 0xb6, 0xc6, 0x1b, 0x07,
	0xcf, 0xab, 0x4f, 0xc7, 0x49, 0xb8, 0x44, 0x12, 0xab, 0xc7, 0x46, 0x7f, 0xac, 0xc3, 0x66, 0x41,
	0x55, 0xd9, 0x99, 0xf9, 0xd9, 0xb2, 0x1f, 0x60, 0x30, 0x23, 0x2b, 0xb4, 0x89, 0xca, 0x72, 0x1e,
	0x90, 0xe3, 0xbd, 0xd2, 0xb1, 0x4f, 0xd3, 0xe4, 0x4a, 0xe6, 0x6a, 0x32, 0x43, 0x8a, 0x68, 0xf0,
	0x31, 0x8e, 0xc8, 0x6a, 0x43, 0xd1, 0xb5, 0x04, 0xad, 0xd9, 0x1e, 0xf4, 0x8c, 0x8e, 0x92, 0x73,
	0x99, 0xfa, 0xf8, 0xfa, 0xa2, 0xb4, 0xd9, 0x36, 0xb4, 0x72, 0x75, 0x5f, 0x44, 0x88, 0x4b, 0xcc,
	0xd0, 0x3c, 0x57, 0xd9, 0x59, 0xcc, 0xdb, 0xc3, 0x60, 0xbc, 0x29, 0x0a, 0x8b, 0x71, 0xe8, 0xea,
	0xfc, 0xe2, 0xce, 0xa8, 0x8c, 0x77, 0x86, 0xc1, 0xb8, 0x27, 0x16, 0x26, 0x9e, 0xb8, 0x12, 0x4a,
	0xc6, 0x0f, 0xbc, 0xeb, 0x73, 0xea, 0x2d, 0xf6, 0x12, 0x3a, 0xb9, 0x93, 0x6e, 0x9e, 0xf3, 0xde,
	0x30, 0x18, 0x6f, 0x1d, 0xbc, 0x28, 0xa3, 0xc0, 0x02, 0x2a, 0x33, 0x4f, 0x27, 0xfe, 0xee, 0x6f,
	0x89, 0x25, 0x0a, 0x36, 0xfb, 0x0a, 0x3a, 0x77, 0xf7, 0x98, 0x0d, 0xde, 0xa7, 0x8a, 0xee, 0x94,
	0xe7, 0xde, 0x2b, 0x7d, 0xaf, 0x0d, 0x6e, 0x89, 0x82, 0xc2, 0x46, 0x30, 0xb0, 0x66, 0xaa, 0x8d,
	0xf2, 0x4e, 0x38, 0xd0, 0x15, 0x1a, 0x18, 0xfb, 0x02, 0xb6, 0x32, 0x95, 0x4a, 0x6d, 0x42, 0xa9,
	0x8f, 0xed, 0xdc, 0x38, 0xbe, 0x41, 0xac, 0x25, 0x94, 0x7d, 0x03, 0x5d, 0x3b, 0x77, 0xa1, 0xd4,
	0x39, 0x1f, 0xfc, 0x6b, 0xde, 0x17, 0x54, 0xf6, 0x0a, 0xc0, 0x2f, 0x2f, 0x1f, 0x66, 0x8a, 0x6f,
	0x52, 0xa8, 0x9f, 0x3c, 0x12, 0xaa, 0x27, 0x88, 0x1a, 0x19, 0x2b, 0x93, 0x59, 0x9b, 0x1e, 0xcb,
	0x2c, 0xe6, 0x5b, 0x54, 0xb1, 0xd2, 0xc6, 0x3d, 0x9d, 0x1f, 0x49, 0x93, 0xa8, 0x8c, 0x3f, 0xa1,
	0x84, 0x97, 0xf6, 0xe8, 0xf7, 0x16, 0xec, 0x3c, 0x22, 0x31, 0xf6, 0x12, 0xa0, 0xd2, 0x52, 0xa1,
	0x9d, 0xdd, 0x65, 0x51, 0xfa, 0x5d, 0x51, 0x63, 0xb2, 0x53, 0xd8, 0x46, 0xeb, 0x48, 0x46, 0xc9,
	0x89, 0xca, 0x13, 0x3a, 0xbd, 0x46, 0xb9, 0x7f, 0x54, 0xd2, 0x0b, 0x8e, 0x58, 0x39, 0xc5, 0x8e,
	0x60, 0x20, 0xac, 0x4d, 0x31, 0x3a, 0xf2, 0xd2, 0x22, 0x2f, 0xcd, 0xca, 0x53, 0x12, 0xb3, 0x1a,
	0x4b, 0x34, 0xce, 0xa0, 0x86, 0x9d, 0x4e, 0x15, 0x89, 0xb2, 0x2f, 0x68, 0xcd, 0xbe, 0x83, 0xae,
	0x8c, 0x1c, 0x65, 0xb8, 0xbd, 0x2c, 0xa6, 0xe6, 0x58, 0x70, 0x13, 0x87, 0x69, 0x5e, 0xd0, 0x17,
	0x7a, 0xd6, 0x31, 0xc9, 0xb6, 0xd0, 0xb3, 0x8e, 0xeb, 0xc5, 0xee, 0xfe, 0xd7, 0x62, 0xf7, 0x3e,
	0xa0, 0xd8, 0xa3, 0xdf, 0x9a, 0x45, 0x2b, 0x53, 0xf6, 0x02, 0xe0, 0xb5, 0x4c, 0x17, 0xfa, 0x0d,
	0x48, 0x99, 0x35, 0x84, 0x3d, 0x87, 0xbe, 0x6f, 0x93, 0xf3, 0x79, 0x5a, 0xcc, 0xa7, 0x0a, 0x40,
	0xfd, 0xcb, 0xc8, 0xe9, 0x5b, 0xf5, 0xce, 0x37, 0xed, 0x3a, 0x05, 0xd9, 0xc0, 0xd8, 0xe7, 0xb0,
	0x19, 0xcd, 0xb3, 0x0c, 0xbf, 0xec, 0xe5, 0xef, 0x9b, 0xa4, 0x09, 0x62, 0x97, 0xd0, 0x84, 0xab,
	0x68, 0x45, 0x97, 0x34, 0x51, 0xbc, 0x2f, 0x16, 0xef, 0x7c, 0x9e, 0x5e, 0xa9, 0x8c, 0x0f, 0xa8,
	0x48, 0x35, 0x04, 0x85, 0x7b, 0x45, 0x32, 0x3d, 0xf3, 0xa2, 0xde, 0x14, 0xa5, 0xcd, 0x5e, 0x41,
	0x5f, 0x19, 0xa7, 0x32, 0xca, 0x1e, 0xa3, 0xec, 0x3d, 0x5b, 0xcd, 0x5e, 0x49, 0x11, 0x15, 0x1b,
	0x3f, 0x8b, 0xd3, 0xec, 0x52, 0x47, 0x89, 0x72, 0x7c, 0x87, 0xba, 0xa5, 0x86, 0xa0, 0x6b, 0x54,
	0x8a, 0x97, 0xfe, 0x53, 0x92, 0xdd, 0xb3, 0x55, 0xd9, 0x95, 0x14, 0x51, 0xb1, 0x47, 0x7f, 0xae,
	0x01, 0x0f, 0xe3, 0xe4, 0x44, 0x3a, 0x29, 0x54, 0x64, 0x6f, 0x55, 0x56, 0xeb, 0x8d, 0x6a, 0x1e,
	0x06, 0x8d, 0x79, 0xf8, 0x02, 0xe0, 0x54, 0x9a, 0xb8, 0x98, 0xd3, 0x6b, 0xc3, 0x16, 0x96, 0xad,
	0x42, 0x30, 0x9d, 0x98, 0xfe, 0x5a, 0x69, 0x7d, 0xed, 0x96, 0x50, 0xe4, 0x9d, 0xe8, 0x3c, 0xb7,
	0xd3, 0xdb, 0x05, 0xcf, 0x0f, 0xe3, 0x25, 0x14, 0xbf, 0x77, 0x41, 0x62, 0x7a, 0xa3, 0x73, 0x47,
	0x0f, 0x4e, 0x5b, 0xd4, 0x10, 0x14, 0xc2, 0x45, 0x7d, 0x10, 0xfa, 0x21, 0xdd, 0xc0, 0xd8, 0x10,
	0x36, 0x74, 0x7e, 0xa4, 0xaf, 0xdf, 0x6b, 0x83, 0x73, 0xbc, 0x4b, 0x94, 0x3a, 0x84, 0x53, 0xfe,
	0x4e, 0x9b, 0x63, 0x7c, 0x42, 0x7a, 0x94, 0xe2, 0x85, 0x59, 0xbe, 0x2c, 0xfd, 0xea, 0x65, 0x19,
	0xfd, 0xd2, 0x82, 0xdd, 0x66, 0xe2, 0x4a, 0x55, 0x63, 0x58, 0x2a, 0x4f, 0x56, 0x94, 0xbd, 0x84,
	0xa2, 0xba, 0xc3, 0x52, 0xdd, 0xfe, 0x4d, 0xae, 0x00, 0x0c, 0xea, 0xb0, 0xae, 0xee, 0x96, 0x57,
	0x77, 0x1d, 0xf3, 0x41, 0xfd, 0x68, 0xa7, 0x53, 0x7b, 0x17, 0x4a, 0x4d, 0xd9, 0xa3, 0xa0, 0x4a,
	0x08, 0xbd, 0x1c, 0xcf, 0xb3, 0x4c, 0x19, 0x27, 0xec, 0xdc, 0xf8, 0x87, 0xad, 0x2d, 0x1a, 0x18,
	0xa6, 0xf7, 0x12, 0x75, 0xee, 0x19, 0x1d, 0xdf, 0x85, 0x15, 0xd2, 0x94, 0x57, 0xf7, 0x43, 0xe4,
	0xc5, 0xc6, 0xf0, 0xe4, 0x8d, 0x74, 0x2a, 0x77, 0x58, 0x2d, 0x2f, 0x97, 0x1e, 0x95, 0x6f, 0x19,
	0xc6, 0x4b, 0x88, 0xaa, 0xb5, 0xfa, 0xbe, 0xb5, 0x2a, 0x04, 0x7f, 0x91, 0x58, 0x7a, 0x81, 0x81,
	0xf2, 0xe0, 0x8d, 0xd1, 0xaf, 0x01, 0x6c, 0x35, 0xab, 0xc0, 0x0e, 0x01, 0x66, 0xcb, 0x0f, 0xc1,
	0x67, 0xf5, 0x51, 0xfe, 0xa8, 0xd6, 0x45, 0xed, 0x10, 0xfb, 0x1e, 0x7a, 0x71, 0xf3, 0x2d, 0xf8,
	0xf4, 0x1f, 0x1c, 0x94, 0xcf, 0x41, 0x79, 0x60, 0x74, 0xb0, 0xdc, 0x50, 0xc4, 0x89, 0x49, 0xa8,
	0xbb, 0xd0, 0xf1, 0x16, 0xdd, 0xab, 0x2d, 0x0a, 0xeb, 0xcb, 0x6f, 0x57, 0x7e, 0xe1, 0xf9, 0x51,
	0xce, 0x06, 0xd0, 0x0b, 0xc3, 0xc3, 0xcb, 0xc9, 0xeb, 0x77, 0x17, 0xdb, 0x01, 0xdb, 0x02, 0x20,
	0xeb, 0xe4, 0x30, 0x3c, 0x3c, 0xdb, 0x5e, 0x0b, 0xff, 0x17, 0x06, 0x7f, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x0d, 0x85, 0x9b, 0x4f, 0x67, 0x0a, 0x00, 0x00,
}

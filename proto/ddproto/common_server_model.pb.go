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

// 用户的资
type User struct {
	RoomCard            *int64   `protobuf:"varint,1,opt,name=RoomCard" json:"RoomCard,omitempty"`
	Pwd                 *string  `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Coin                *int64   `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                  *uint32  `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName            *string  `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	Scores              *int32   `protobuf:"varint,6,opt,name=scores" json:"scores,omitempty"`
	LastDrawLotteryTime *string  `protobuf:"bytes,7,opt,name=lastDrawLotteryTime" json:"lastDrawLotteryTime,omitempty"`
	LastSignTime        *string  `protobuf:"bytes,8,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignTotalDays       *int32   `protobuf:"varint,9,opt,name=signTotalDays" json:"signTotalDays,omitempty"`
	SignContinuousDays  *int32   `protobuf:"varint,10,opt,name=signContinuousDays" json:"signContinuousDays,omitempty"`
	Diamond             *int64   `protobuf:"varint,11,opt,name=Diamond" json:"Diamond,omitempty"`
	Diamond2            *int64   `protobuf:"varint,12,opt,name=Diamond2" json:"Diamond2,omitempty"`
	OpenId              *string  `protobuf:"bytes,13,opt,name=openId" json:"openId,omitempty"`
	UnionId             *string  `protobuf:"bytes,14,opt,name=UnionId" json:"UnionId,omitempty"`
	HeadUrl             *string  `protobuf:"bytes,15,opt,name=headUrl" json:"headUrl,omitempty"`
	City                *string  `protobuf:"bytes,16,opt,name=city" json:"city,omitempty"`
	Sex                 *int32   `protobuf:"varint,17,opt,name=sex" json:"sex,omitempty"`
	RobotType           *int32   `protobuf:"varint,18,opt,name=robotType" json:"robotType,omitempty"`
	Ticket              *int32   `protobuf:"varint,19,opt,name=ticket" json:"ticket,omitempty"`
	Bonus               *float64 `protobuf:"fixed64,20,opt,name=bonus" json:"bonus,omitempty"`
	RegTime             *int64   `protobuf:"varint,21,opt,name=regTime" json:"regTime,omitempty"`
	RegChannel          *string  `protobuf:"bytes,22,opt,name=regChannel" json:"regChannel,omitempty"`
	AgentId             *uint32  `protobuf:"varint,23,opt,name=agentId" json:"agentId,omitempty"`
	LastIp              *string  `protobuf:"bytes,24,opt,name=lastIp" json:"lastIp,omitempty"`
	LastTime            *int64   `protobuf:"varint,25,opt,name=lastTime" json:"lastTime,omitempty"`
	// 用户兑换信息
	RealName         *string  `protobuf:"bytes,26,opt,name=realName" json:"realName,omitempty"`
	PhoneNumber      *string  `protobuf:"bytes,27,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	WxNumber         *string  `protobuf:"bytes,28,opt,name=wxNumber" json:"wxNumber,omitempty"`
	RealAddress      *string  `protobuf:"bytes,29,opt,name=realAddress" json:"realAddress,omitempty"`
	ChannelId        *int32   `protobuf:"varint,30,opt,name=channelId" json:"channelId,omitempty"`
	NewUserAward     *bool    `protobuf:"varint,31,opt,name=newUserAward" json:"newUserAward,omitempty"`
	RegIp            *string  `protobuf:"bytes,32,opt,name=regIp" json:"regIp,omitempty"`
	Longitude        *float32 `protobuf:"fixed32,33,opt,name=longitude" json:"longitude,omitempty"`
	Latitude         *float32 `protobuf:"fixed32,34,opt,name=latitude" json:"latitude,omitempty"`
	Location         *string  `protobuf:"bytes,35,opt,name=location" json:"location,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

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

func (m *User) GetScores() int32 {
	if m != nil && m.Scores != nil {
		return *m.Scores
	}
	return 0
}

func (m *User) GetLastDrawLotteryTime() string {
	if m != nil && m.LastDrawLotteryTime != nil {
		return *m.LastDrawLotteryTime
	}
	return ""
}

func (m *User) GetLastSignTime() string {
	if m != nil && m.LastSignTime != nil {
		return *m.LastSignTime
	}
	return ""
}

func (m *User) GetSignTotalDays() int32 {
	if m != nil && m.SignTotalDays != nil {
		return *m.SignTotalDays
	}
	return 0
}

func (m *User) GetSignContinuousDays() int32 {
	if m != nil && m.SignContinuousDays != nil {
		return *m.SignContinuousDays
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

func (m *User) GetRobotType() int32 {
	if m != nil && m.RobotType != nil {
		return *m.RobotType
	}
	return 0
}

func (m *User) GetTicket() int32 {
	if m != nil && m.Ticket != nil {
		return *m.Ticket
	}
	return 0
}

func (m *User) GetBonus() float64 {
	if m != nil && m.Bonus != nil {
		return *m.Bonus
	}
	return 0
}

func (m *User) GetRegTime() int64 {
	if m != nil && m.RegTime != nil {
		return *m.RegTime
	}
	return 0
}

func (m *User) GetRegChannel() string {
	if m != nil && m.RegChannel != nil {
		return *m.RegChannel
	}
	return ""
}

func (m *User) GetAgentId() uint32 {
	if m != nil && m.AgentId != nil {
		return *m.AgentId
	}
	return 0
}

func (m *User) GetLastIp() string {
	if m != nil && m.LastIp != nil {
		return *m.LastIp
	}
	return ""
}

func (m *User) GetLastTime() int64 {
	if m != nil && m.LastTime != nil {
		return *m.LastTime
	}
	return 0
}

func (m *User) GetRealName() string {
	if m != nil && m.RealName != nil {
		return *m.RealName
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil && m.PhoneNumber != nil {
		return *m.PhoneNumber
	}
	return ""
}

func (m *User) GetWxNumber() string {
	if m != nil && m.WxNumber != nil {
		return *m.WxNumber
	}
	return ""
}

func (m *User) GetRealAddress() string {
	if m != nil && m.RealAddress != nil {
		return *m.RealAddress
	}
	return ""
}

func (m *User) GetChannelId() int32 {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return 0
}

func (m *User) GetNewUserAward() bool {
	if m != nil && m.NewUserAward != nil {
		return *m.NewUserAward
	}
	return false
}

func (m *User) GetRegIp() string {
	if m != nil && m.RegIp != nil {
		return *m.RegIp
	}
	return ""
}

func (m *User) GetLongitude() float32 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *User) GetLatitude() float32 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *User) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

// 通知公告
type TNotice struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NoticeType       *int32   `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string  `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string  `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string  `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Noticefileds     []string `protobuf:"bytes,6,rep,name=noticefileds" json:"noticefileds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TNotice) Reset()                    { *m = TNotice{} }
func (m *TNotice) String() string            { return proto.CompactTextString(m) }
func (*TNotice) ProtoMessage()               {}
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

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

func (m *TNotice) GetNoticefileds() []string {
	if m != nil {
		return m.Noticefileds
	}
	return nil
}

// 玩家游戏场次统计
type TGameCounts struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DdzCount         *int32  `protobuf:"varint,3,opt,name=ddzCount" json:"ddzCount,omitempty"`
	DdzWinCount      *int32  `protobuf:"varint,4,opt,name=ddzWinCount" json:"ddzWinCount,omitempty"`
	MjCount          *int32  `protobuf:"varint,5,opt,name=mjCount" json:"mjCount,omitempty"`
	MjWinCount       *int32  `protobuf:"varint,6,opt,name=mjWinCount" json:"mjWinCount,omitempty"`
	ThCount          *int32  `protobuf:"varint,7,opt,name=thCount" json:"thCount,omitempty"`
	ThWinWinCount    *int32  `protobuf:"varint,8,opt,name=thWinWinCount" json:"thWinWinCount,omitempty"`
	ZjhCount         *int32  `protobuf:"varint,9,opt,name=zjhCount" json:"zjhCount,omitempty"`
	ZjhWinCount      *int32  `protobuf:"varint,10,opt,name=zjhWinCount" json:"zjhWinCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TGameCounts) Reset()                    { *m = TGameCounts{} }
func (m *TGameCounts) String() string            { return proto.CompactTextString(m) }
func (*TGameCounts) ProtoMessage()               {}
func (*TGameCounts) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *TGameCounts) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TGameCounts) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TGameCounts) GetDdzCount() int32 {
	if m != nil && m.DdzCount != nil {
		return *m.DdzCount
	}
	return 0
}

func (m *TGameCounts) GetDdzWinCount() int32 {
	if m != nil && m.DdzWinCount != nil {
		return *m.DdzWinCount
	}
	return 0
}

func (m *TGameCounts) GetMjCount() int32 {
	if m != nil && m.MjCount != nil {
		return *m.MjCount
	}
	return 0
}

func (m *TGameCounts) GetMjWinCount() int32 {
	if m != nil && m.MjWinCount != nil {
		return *m.MjWinCount
	}
	return 0
}

func (m *TGameCounts) GetThCount() int32 {
	if m != nil && m.ThCount != nil {
		return *m.ThCount
	}
	return 0
}

func (m *TGameCounts) GetThWinWinCount() int32 {
	if m != nil && m.ThWinWinCount != nil {
		return *m.ThWinWinCount
	}
	return 0
}

func (m *TGameCounts) GetZjhCount() int32 {
	if m != nil && m.ZjhCount != nil {
		return *m.ZjhCount
	}
	return 0
}

func (m *TGameCounts) GetZjhWinCount() int32 {
	if m != nil && m.ZjhWinCount != nil {
		return *m.ZjhWinCount
	}
	return 0
}

// 玩家的任务
type TUserTask struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	TaskId           *int32  `protobuf:"varint,3,opt,name=taskId" json:"taskId,omitempty"`
	AwardId          *int32  `protobuf:"varint,4,opt,name=awardId" json:"awardId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TUserTask) Reset()                    { *m = TUserTask{} }
func (m *TUserTask) String() string            { return proto.CompactTextString(m) }
func (*TUserTask) ProtoMessage()               {}
func (*TUserTask) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *TUserTask) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TUserTask) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TUserTask) GetTaskId() int32 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func (m *TUserTask) GetAwardId() int32 {
	if m != nil && m.AwardId != nil {
		return *m.AwardId
	}
	return 0
}

// 推送协议
type Push struct {
	Id               *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Push) Reset()                    { *m = Push{} }
func (m *Push) String() string            { return proto.CompactTextString(m) }
func (*Push) ProtoMessage()               {}
func (*Push) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *Push) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Push) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "ddproto.User")
	proto.RegisterType((*TNotice)(nil), "ddproto.TNotice")
	proto.RegisterType((*TGameCounts)(nil), "ddproto.TGameCounts")
	proto.RegisterType((*TUserTask)(nil), "ddproto.TUserTask")
	proto.RegisterType((*Push)(nil), "ddproto.Push")
}

var fileDescriptor13 = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x8e, 0xdb, 0x36,
	0x10, 0x85, 0xbc, 0xf6, 0xda, 0xe6, 0xae, 0xd3, 0x94, 0x49, 0xb7, 0x4c, 0x9a, 0x6e, 0x55, 0xb7,
	0x1f, 0x46, 0x3f, 0x82, 0xa2, 0x37, 0x08, 0xbc, 0x40, 0x21, 0xa0, 0x5d, 0x14, 0xaa, 0x17, 0xf9,
	0x0c, 0xb8, 0xe2, 0xd4, 0xa6, 0x2d, 0x91, 0x06, 0x49, 0xd5, 0xb1, 0xaf, 0xd4, 0x73, 0xf4, 0x1a,
	0x3d, 0x4b, 0x31, 0x43, 0x49, 0x2b, 0xa3, 0xf9, 0xe8, 0xdf, 0xbc, 0xf7, 0x38, 0xd4, 0xcc, 0x13,
	0x67, 0xd8, 0xab, 0xc2, 0x56, 0x95, 0x35, 0x1f, 0x3c, 0xb8, 0x3f, 0xc1, 0x7d, 0xa8, 0xac, 0x82,
	0xf2, 0xed, 0xde, 0xd9, 0x60, 0xf9, 0x58, 0x29, 0x0a, 0xe6, 0xff, 0x8c, 0xd9, 0xf0, 0xc1, 0x83,
	0xe3, 0xaf, 0xd9, 0x24, 0xb7, 0xb6, 0x5a, 0x4a, 0xa7, 0x44, 0x92, 0x26, 0x8b, 0x8b, 0xbc, 0xc3,
	0xfc, 0x39, 0xbb, 0xd8, 0x1f, 0x94, 0x18, 0xa4, 0xc9, 0x62, 0x9a, 0x63, 0xc8, 0x39, 0x1b, 0x16,
	0x56, 0x1b, 0x71, 0x41, 0x27, 0x29, 0xe6, 0xcf, 0xd8, 0x40, 0x2b, 0x31, 0x4c, 0x93, 0xc5, 0x2c,
	0x1f, 0x68, 0x85, 0x37, 0x1a, 0x5d, 0xec, 0xee, 0x65, 0x05, 0x62, 0x44, 0xa9, 0x1d, 0xe6, 0x37,
	0xec, 0xd2, 0x17, 0xd6, 0x81, 0x17, 0x97, 0x69, 0xb2, 0x18, 0xe5, 0x0d, 0xe2, 0x3f, 0xb2, 0x17,
	0xa5, 0xf4, 0xe1, 0xce, 0xc9, 0xc3, 0x2f, 0x36, 0x04, 0x70, 0xc7, 0x95, 0xae, 0x40, 0x8c, 0x29,
	0xfd, 0x53, 0x12, 0x9f, 0xb3, 0x6b, 0xa4, 0x7f, 0xd7, 0x6b, 0x43, 0x47, 0x27, 0x74, 0xf4, 0x8c,
	0xe3, 0xdf, 0xb3, 0x99, 0xc7, 0xd8, 0x06, 0x59, 0xde, 0xc9, 0xa3, 0x17, 0x53, 0xfa, 0xe8, 0x39,
	0xc9, 0xdf, 0x32, 0x8e, 0xc4, 0xd2, 0x9a, 0xa0, 0x4d, 0x6d, 0x6b, 0x4f, 0x47, 0x19, 0x1d, 0xfd,
	0x84, 0xc2, 0x05, 0x1b, 0xdf, 0x69, 0x59, 0x59, 0xa3, 0xc4, 0x15, 0xd9, 0xd0, 0x42, 0xec, 0xbc,
	0x09, 0x7f, 0x12, 0xd7, 0xd1, 0xcb, 0x16, 0x63, 0xe7, 0x76, 0x0f, 0x26, 0x53, 0x62, 0x46, 0x95,
	0x36, 0x08, 0x6f, 0x7b, 0x30, 0xda, 0xa2, 0xf0, 0x8c, 0x84, 0x16, 0xa2, 0xb2, 0x01, 0xa9, 0x1e,
	0x5c, 0x29, 0x3e, 0x8b, 0x4a, 0x03, 0xe9, 0x2f, 0xe8, 0x70, 0x14, 0xcf, 0x89, 0xa6, 0x18, 0xff,
	0x95, 0x87, 0x8f, 0xe2, 0x73, 0x2a, 0x1b, 0x43, 0xfe, 0x86, 0x4d, 0x9d, 0x7d, 0xb4, 0x61, 0x75,
	0xdc, 0x83, 0xe0, 0xc4, 0x3f, 0x11, 0x58, 0x4f, 0xd0, 0xc5, 0x0e, 0x82, 0x78, 0x11, 0xff, 0x44,
	0x44, 0xfc, 0x25, 0x1b, 0x3d, 0x5a, 0x53, 0x7b, 0xf1, 0x32, 0x4d, 0x16, 0x49, 0x1e, 0x01, 0xd6,
	0xe2, 0x60, 0x4d, 0x46, 0x7f, 0x11, 0x7b, 0x6e, 0x20, 0xbf, 0x65, 0xcc, 0xc1, 0x7a, 0xb9, 0x91,
	0xc6, 0x40, 0x29, 0x6e, 0xa8, 0xa2, 0x1e, 0x83, 0x99, 0x72, 0x0d, 0x26, 0x64, 0x4a, 0x7c, 0x49,
	0x4f, 0xa4, 0x85, 0x58, 0x01, 0xfe, 0xad, 0x6c, 0x2f, 0x44, 0x74, 0x24, 0x22, 0x74, 0x11, 0x23,
	0xfa, 0xd8, 0xab, 0xe8, 0x62, 0x8b, 0x51, 0x73, 0x20, 0x4b, 0x7a, 0x5b, 0xaf, 0xe3, 0xdb, 0x6a,
	0x31, 0x4f, 0xd9, 0xd5, 0x7e, 0x63, 0x0d, 0xdc, 0xd7, 0xd5, 0x23, 0x38, 0xf1, 0x15, 0xc9, 0x7d,
	0x0a, 0xb3, 0x0f, 0x1f, 0x1b, 0xf9, 0x4d, 0xcc, 0x6e, 0x31, 0x66, 0xe3, 0x4d, 0xef, 0x94, 0x72,
	0xe0, 0xbd, 0xf8, 0x3a, 0x66, 0xf7, 0x28, 0xf4, 0xb3, 0x88, 0x4d, 0x65, 0x4a, 0xdc, 0x46, 0x3f,
	0x3b, 0x02, 0xdf, 0xa3, 0x81, 0x03, 0x8e, 0xd4, 0xbb, 0x03, 0xce, 0xd2, 0x37, 0x69, 0xb2, 0x98,
	0xe4, 0x67, 0x1c, 0x7a, 0xeb, 0x60, 0x9d, 0xed, 0x45, 0x4a, 0xb7, 0x47, 0x80, 0xf7, 0x96, 0xd6,
	0xac, 0x75, 0xa8, 0x15, 0x88, 0x6f, 0xd3, 0x64, 0x31, 0xc8, 0x9f, 0x88, 0xe8, 0x46, 0x88, 0xe2,
	0x9c, 0xc4, 0x0e, 0x93, 0x66, 0x0b, 0x19, 0xb4, 0x35, 0xe2, 0xbb, 0xd8, 0x4f, 0x8b, 0xe7, 0x7f,
	0x27, 0x6c, 0xbc, 0xba, 0xb7, 0x41, 0x17, 0xd0, 0x4c, 0x68, 0x42, 0x25, 0xe3, 0x84, 0xde, 0x32,
	0x66, 0x48, 0xa1, 0xa7, 0x31, 0x20, 0xbe, 0xc7, 0xa0, 0x17, 0x0d, 0xd2, 0xa1, 0x04, 0x1a, 0xf6,
	0x69, 0xde, 0xa7, 0x70, 0xb2, 0x22, 0xc4, 0xd9, 0x00, 0x13, 0x68, 0xfc, 0xa7, 0xf9, 0x39, 0xf9,
	0xf4, 0x9d, 0x5f, 0xa1, 0xb2, 0xcd, 0x2e, 0xe8, 0x31, 0xe4, 0x19, 0xa1, 0x3f, 0x74, 0x09, 0x0a,
	0x77, 0xc2, 0x05, 0xce, 0x70, 0x9f, 0x9b, 0xff, 0x35, 0x60, 0x57, 0xab, 0x9f, 0x65, 0x05, 0x4b,
	0x5b, 0x9b, 0xe0, 0xff, 0xd3, 0xcb, 0x0d, 0xbb, 0xac, 0x3d, 0xb8, 0x2c, 0xae, 0xa9, 0x59, 0xde,
	0x20, 0xf4, 0x46, 0xa9, 0x13, 0x25, 0x51, 0x03, 0xa3, 0xbc, 0xc3, 0xd8, 0x9f, 0x52, 0xa7, 0xf7,
	0xda, 0x44, 0x79, 0x48, 0x72, 0x9f, 0xc2, 0x57, 0x5b, 0x6d, 0xa3, 0x3a, 0x22, 0xb5, 0x85, 0xd8,
	0x53, 0xb5, 0xed, 0x52, 0xe3, 0x16, 0xeb, 0x31, 0x98, 0x19, 0x36, 0x51, 0x1c, 0xc7, 0xcc, 0x06,
	0xa2, 0x67, 0x61, 0xf3, 0x5e, 0x9b, 0x2e, 0x79, 0x12, 0xb7, 0xd1, 0x19, 0x89, 0x75, 0x9f, 0xb6,
	0xcd, 0x05, 0x71, 0x5d, 0x75, 0x18, 0xeb, 0x3e, 0x6d, 0x37, 0x5d, 0x7e, 0x5c, 0x51, 0x7d, 0x6a,
	0x0e, 0x6c, 0xba, 0xc2, 0xf7, 0xb6, 0x92, 0x7e, 0xf7, 0xbf, 0xad, 0xc2, 0x55, 0x20, 0xfd, 0x2e,
	0x53, 0x8d, 0x51, 0x0d, 0xa2, 0xd1, 0xc5, 0x77, 0x9b, 0xa9, 0xc6, 0xa2, 0x16, 0xce, 0x7f, 0x60,
	0xc3, 0xdf, 0x6a, 0xbf, 0xe9, 0x7d, 0x21, 0xae, 0x7e, 0xce, 0x86, 0x4a, 0x06, 0x49, 0xf7, 0x5f,
	0xe7, 0x14, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xb2, 0xce, 0xd2, 0x8e, 0x06, 0x00, 0x00,
}

package model

import (
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"time"
	"casino_common/utils/timeUtils"
	"casino_common/utils/numUtils"
)

type MjRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

func (b MjRecordBean) TransBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{
		NickName:  proto.String(b.NickName),
		UserId:    proto.Uint32(b.UserId),
		WinAmount: proto.Int64(b.WinAmount),
	}
	return result
}

//一把结束,战绩可以通过这个表来查询
type T_mj_desk_round struct {
	DeskId       int32
	GameNumber   int32
	UserIds      string
	BeginTime    time.Time
	EndTime      time.Time
	Records      []MjRecordBean
	FriendPlay   bool   //是否是朋友桌
	PassWord     string //朋友桌的房间号
	PlayBackData []*ddproto.PlaybackSnapshot
}

func (t T_mj_desk_round) TransRecord() *ddproto.BeanGameRecord {
	result := &ddproto.BeanGameRecord{
		BeginTime: proto.String(timeUtils.Format(t.EndTime)),
		DeskId:    proto.Int32(int32(numUtils.String2Int(t.PassWord))),
		Id:        proto.Int32(t.GameNumber), //战绩id就是 游戏编号
	}

	for _, bean := range t.Records {
		b := bean.TransBeanUserRecord()
		result.Users = append(result.Users, b)
	}
	return result
}

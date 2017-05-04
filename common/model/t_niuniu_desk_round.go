package model

import (
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"time"
	"casino_common/utils/timeUtils"
)

//一把结束,战绩可以通过这个表来查询
type T_niuniu_desk_round struct {
	DeskId     int32  //房间号
	GameNumber int32  //一局游戏编号
	UserIds    string
	BeginTime  time.Time
	EndTime    time.Time
	Records    []NiuRecordBean
}

func (t T_niuniu_desk_round) TransRecord() *ddproto.BeanGameRecord {
	result := &ddproto.BeanGameRecord{
		BeginTime: proto.String(timeUtils.Format(t.BeginTime)),
		DeskId:    proto.Int32(t.DeskId),
		Id:        proto.Int32(t.GameNumber),
	}
	for _, bean := range t.Records {
		b := bean.TransBeanUserRecord()
		result.Users = append(result.Users, b)
	}
	return result
}

type NiuRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

func (b NiuRecordBean) TransBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{
		NickName:  proto.String(b.NickName),
		UserId:    proto.Uint32(b.UserId),
		WinAmount: proto.Int64(b.WinAmount),
	}
	return result
}

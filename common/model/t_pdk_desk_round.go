package model

import (
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"time"
	"casino_common/utils/timeUtils"
)

//一把结束,战绩可以通过这个表来查询
type T_pdk_desk_round struct {
	DeskId     int32
	GameNumber int32
	UserIds    string
	BeginTime  time.Time
	EndTime    time.Time
	Records    []PdkRecordBean
}

func (t T_pdk_desk_round) TransRecord() *ddproto.BeanGameRecord {
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

type PdkRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

func (b PdkRecordBean) TransBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{
		NickName:  proto.String(b.NickName),
		UserId:    proto.Uint32(b.UserId),
		WinAmount: proto.Int64(b.WinAmount),
	}
	return result
}

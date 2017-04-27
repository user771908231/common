package model

import (
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"time"
	"casino_common/utils/timeUtils"
)

//一把结束,战绩可以通过这个表来查询
type T_ddz_desk_round struct {
	DeskId     int32
	GameNumber int32
	UserIds    string
	BeginTime  time.Time
	EndTime    time.Time
	Records    []DdzRecordBean
}

func (t T_ddz_desk_round) TransRecord() *ddproto.BeanGameRecord {
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

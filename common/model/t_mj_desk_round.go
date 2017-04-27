package model

import (
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
)

type DdzRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

func (b DdzRecordBean) TransBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{
		NickName:  proto.String(b.NickName),
		UserId:    proto.Uint32(b.UserId),
		WinAmount: proto.Int64(b.WinAmount),
	}
	return result
}

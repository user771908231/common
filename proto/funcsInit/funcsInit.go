package commonNewPorot

import (
	"casino_common/proto"
)

func NewHeader() *casinoCommonProto.ProtoHeader {
	ret := new(casinoCommonProto.ProtoHeader)
	ret.UserId = new(uint32)
	ret.Code = new(int32)
	ret.Error = new(string)
	return ret
}

func NewHeartbeat() *casinoCommonProto.Heartbeat {
	ret := new(casinoCommonProto.Heartbeat)
	ret.Header = NewHeader()
	return ret
}


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

func NewGameSession() *casinoCommonProto.GameSession {
	ret := new(casinoCommonProto.GameSession)
	ret.UserId = new(uint32)
	ret.GameId = new(int32)
	ret.DeskId = new(int32)
	ret.RoomId = new(int32)
	ret.GameStatus = new(int32)
	ret.GameNumber = new(int32)
	ret.GameCustomStatus = new(int32)
	return ret
}

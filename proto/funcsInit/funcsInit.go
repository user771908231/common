package commonNewPorot

import "casino_common/proto/ddproto"

func NewHeader() *ddproto.ProtoHeader {
	ret := new(ddproto.ProtoHeader)
	ret.UserId = new(uint32)
	ret.Code = new(int32)
	ret.Error = new(string)
	return ret
}

func NewHeartbeat() *ddproto.Heartbeat {
	ret := new(ddproto.Heartbeat)
	ret.Header = NewHeader()
	return ret
}

func NewGameSession() *ddproto.GameSession {
	ret := new(ddproto.GameSession)
	ret.UserId = new(uint32)
	ret.GameId = new(int32)
	ret.DeskId = new(int32)
	ret.RoomId = new(int32)
	ret.GameStatus = new(int32)
	ret.GameNumber = new(int32)
	ret.GameCustomStatus = new(int32)
	ret.IsBreak = new(bool)
	ret.IsLeave = new(bool)
	return ret
}

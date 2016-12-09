package commonNewPorot

import (
	"casino_common/proto/ddproto"
	"casino_common/common/game"
	"casino_majiang/msg/funcsInit"
)

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

func NewGameDesk() *game.GameDesk {
	ret := new(game.GameDesk)
	return ret
}

func NewCommonSrvGameUser() *ddproto.CommonSrvGameUser {
	ret := new(ddproto.CommonSrvGameUser)
	ret.UserId = new(uint32)
	ret.NickName = new(string)
	ret.Coin = new(int64)
	ret.Status = new(int32)
	ret.IsBreak = new(bool)
	ret.IsLeave = new(bool)
	ret.DeskId = new(int32)
	ret.RoomId = new(int32)
	ret.WaitTime = new(int64)
	return ret
}

func NewGameUser() *game.GameUser {
	ret := new(game.GameUser)
	ret.CommonSrvGameUser = NewCommonSrvGameUser()
	return ret
}

func NewWeixinInfo() *ddproto.WeixinInfo {
	ret := new(ddproto.WeixinInfo)
	ret.City = new(string)
	ret.HeadUrl = new(string)
	ret.NickName = new(string)
	ret.OpenId = new(string)
	ret.Sex = new(int32)
	ret.UnionId = new(string)
	return ret
}

func NewCommonReqGameLogin() *ddproto.CommonReqGameLogin {
	ret := new(ddproto.CommonReqGameLogin)
	ret.Header = NewHeader()
	ret.ProtoVersion = new(int32)
	ret.UserId = new(uint32)
	ret.WxInfo = NewWeixinInfo()
	return ret
}

func NewCommonAckGameLogin() *ddproto.CommonAckGameLogin {
	ret := new(ddproto.CommonAckGameLogin)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	ret.NickName = new(string)
	ret.RoomPassword = new(string)
	ret.CostCreateRoom = new(int64)
	ret.CostRebuy = new(int64)
	ret.Championship = new(bool)
	ret.Chip = new(int64)
	ret.MailCount = new(int32)
	ret.Notice = new(string)
	ret.GameStatus = new(int32)
	return ret
}




























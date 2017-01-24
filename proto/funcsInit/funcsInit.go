package commonNewPorot

import (
	"casino_common/proto/ddproto"
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

func NewServerInfo() *ddproto.ServerInfo {
	ret := new(ddproto.ServerInfo)
	ret.Ip = new(string)
	ret.Port = new(int32)
	ret.Status = new(int32)
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
	ret.RoomType = new(int32)
	ret.RoomPassword = new(string)
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
	ret.GameNumber = new(int32)
	ret.RoomType = new(int32)
	ret.RoomPassword = new(string)
	ret.IsRobot = new(bool)
	return ret
}

//服务端通过的桌子的信息,桌子公用的一些属性
func NewCommonSrvGameDesk() *ddproto.CommonSrvGameDesk {
	ret := new(ddproto.CommonSrvGameDesk)
	ret.DeskId = new(int32)
	ret.GameNumber = new(int32)
	ret.RoomId = new(int32)
	ret.Round = new(int32)
	ret.Password = new(string)
	ret.Owner = new(uint32)
	ret.Banker = new(uint32)
	ret.CreateFee = new(int64)
	ret.Status = new(int32)
	ret.BoardsCout = new(int32)
	ret.BaseValue = new(int64)
	ret.BeginTime = new(string)
	ret.EndTime = new(string)
	ret.NInitActionTime = new(int32)
	ret.UserCountLimit = new(int32)
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
	ret.Chip = new(int64)
	ret.Coin = new(int64)
	return ret
}

func NewCommonBcMessage() *ddproto.CommonBcMessage {
	result := new(ddproto.CommonBcMessage)
	result.Header = NewHeader()
	result.Id = new(int32)
	result.Msg = new(string)
	result.MsgType = new(int32)
	result.UserId = new(uint32)
	result.ToUserId = new(uint32)
	return result
}

//回复公告的通用new方法
func NewCommonAckNotice() *ddproto.CommonAckNotice {
	ret := new(ddproto.CommonAckNotice)
	ret.Id = new(int32)
	ret.NoticeContent = new(string)
	ret.NoticeMemo = new(string)
	ret.NoticeTitle = new(string)
	ret.NoticeType = new(int32)
	return ret
}

//红包回复
func NewAckRedbag() *ddproto.RedbagAck {
	ret := new(ddproto.RedbagAck)
	ret.Cash = new(string)
	ret.Num2Win4OpenRB = new(int32)
	ret.TotalCash = new(string)
	return ret
}

func NewAckEnterAgentMode() *ddproto.CommonAckEnterAgentMode {
	ret := new(ddproto.CommonAckEnterAgentMode)
	ret.UserId = new(uint32)
	return ret
}

func NewAckQuitAgentMode() *ddproto.CommonAckQuitAgentMode {
	ret := new(ddproto.CommonAckQuitAgentMode)
	ret.UserId = new(uint32)
	return ret
}

func NewAckLeaveDesk() *ddproto.CommonAckLeaveDesk {
	ret := new(ddproto.CommonAckLeaveDesk)
	ret.Header = NewHeader()
	ret.IsExchange = new(bool)
	ret.UserId = new(uint32)
	return ret
}

func NewAckAllowance() *ddproto.CommonAckAllowance {
	ret := new(ddproto.CommonAckAllowance)
	ret.Header = NewHeader()
	ret.AllowanceCoin = new(int64)
	ret.IsSuccessed = new(bool)
	ret.Times2Get = new(int32)
	ret.UserId = new(uint32)
	ret.UserTotalCoin = new(int64)
	return ret
}


func NewBCKickOut() *ddproto.CommonBcKickout {
	ret := new(ddproto.CommonBcKickout)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	ret.Msg = new(string)
	ret.Type = new(ddproto.COMMON_ENUM_KICKOUT)
	return ret
}

func NewBcApplyDissolve() *ddproto.CommonBcApplyDissolve {
	ret := new(ddproto.CommonBcApplyDissolve)
	ret.Header = NewHeader()
	ret.Result = new(bool)
	ret.UserId = new(uint32)
	return ret
}

func NewAckApplyDissolveBack() *ddproto.CommonAckApplyDissolveBack {
	ret := new(ddproto.CommonAckApplyDissolveBack)
	ret.Header = NewHeader()
	ret.Agree = new(bool)
	ret.UserId = new(uint32)
	return ret
}
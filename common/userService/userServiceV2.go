package userService

import (
	"casino_common/proto/ddproto"
)

type U_REDIS uint32

func (s U_REDIS) GetMainSession() *ddproto.GameSession {
	//return sessionService.GetSessionAuto(uint32(s))
	return nil
}

func (s U_REDIS) GetNickName() string {
	return ""
}

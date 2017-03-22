package userService

import (
	"casino_common/proto/ddproto"
)

type UREDIS uint32

func (s UREDIS) GetMainSession() *ddproto.GameSession {
	//return sessionService.GetSessionAuto(uint32(s))
	return nil
}

func (s UREDIS) GetNickName() string {
	return ""
}

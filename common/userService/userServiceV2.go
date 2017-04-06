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
	user := GetUserById(uint32(s))
	return user.GetNickName()
}

func (s U_REDIS) GetSex() int32 {
	user := GetUserById(uint32(s))
	return user.GetSex()
}

func (u U_REDIS) GetHeadUrl() string {
	user := GetUserById(uint32(u))
	return user.GetHeadUrl()
}

func (u U_REDIS) GetOpenId() string {
	user := GetUserById(uint32(u))
	return user.GetOpenId()
}

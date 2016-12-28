package handlers

import (
	"github.com/name5566/leaf/gate"
	"casino_common/common/log"
	"casino_common/common/service/signService"
)

func HandlerSignLottery(userId uint32, a gate.Agent) error {
	log.T("u.id[%v]请求签到领奖", userId)

	error := signService.DoSignLottery(userId)
	if error != nil {
		log.T("签到失败, 错误信息[%v]", error)
		return error
	}

	return nil
}
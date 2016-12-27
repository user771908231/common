package handlers

import (
	"casino_common/common/userService"
	"casino_common/common/Error"
	"github.com/name5566/leaf/gate"
	"casino_common/common/log"
	"casino_common/common/service/signService"
	"casino_common/proto/funcsInit"
)

func HandlerSignLottery(userId uint32, a gate.Agent) error {
	log.T("u.id[%v]请求签到领奖", userId)

	user := userService.GetUserById(userId)
	if user == nil {
		log.T("存储中找不到u.id为[%v]的用户", userId)
		return Error.NewError(-1, "找不到用户个人信息")
	}

	error := signService.DoSign(user)
	if error != nil {
		log.T("签到失败, 错误信息[%v]", error)
		return error
	}

	lotteryId := signService.DoSignLottery(user)

	ack := commonNewPorot.NewAckSignLottery()
	*ack.LotteryId = lotteryId
	a.WriteMsg(ack)

	return nil
}
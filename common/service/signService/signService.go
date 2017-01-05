package signService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/Error"
	"sync/atomic"
	"time"
	"casino_common/common/log"
	"github.com/golang/protobuf/proto"
	"casino_common/common/userService"
	"casino_common/utils/timeUtils"
	"fmt"
)


//用户签到
func DoSign(user *ddproto.User) error {

	log.T("[%v]开始签到", user.GetId())

	//验证用户能否签到
	if !isAbleToSign(user) {
		return Error.NewError(-1, fmt.Sprintf("用户[%v]不能签到,签到失败", user.GetId()))
	}

	//不是连续签到 清空连续签到计数
	if !isContinuousSign(user) {
		clearContinuousSignCount(user)
	}

	addSignCount(user)
	user.LastSignTime = proto.String(time.Now().Format("2006-01-02 15:04:05"))

	log.T("[%v]签到成功", user.GetId())

	return nil;
}

func isContinuousSign(user *ddproto.User) bool {
	if user.GetLastSignTime() == "" {
		return false
	}
	timeNow := time.Now()
	lastSignTime, _ := time.Parse(user.GetLastSignTime(), "2006-01-02 15:04:05")

	//上次签到＋1天 是否等于今天
	if !timeUtils.EqualDate(lastSignTime.AddDate(0, 0, 1), timeNow) {
		log.T("上次签到日期[%v]+1天不等于当前日期[%v]", user.GetLastSignTime(), timeNow.Format("2006-01-02 15:04:05"))
		return false
	}
	return true
}

func isAbleToSign(user *ddproto.User) bool {
	//验证用户能否签到
	if user.GetLastSignTime() != "" {
		timeNow := time.Now()
		lastSignTime, _ := time.Parse(user.GetLastSignTime(), "2006-01-02 15:04:05")
		//上次签到时间在当前时间之后
		if lastSignTime.After(timeNow) {
			log.T("上次签到时间[%v]在当前时间[%v]之后", user.GetLastSignTime(), timeNow.Format("2006-01-02 15:04:05"))
			return false
		}

		// 每天只能签一次
		//上次签到的日期(天)等于当前日期(天)
		if timeUtils.EqualDate(lastSignTime, timeNow) {
			log.T("上次签到日期[%v]等于当前日期[%v]", user.GetLastSignTime(), timeNow.Format("2006-01-02 15:04:05"))
			return false
		}
	}
	return true
}

func addSignCount(user *ddproto.User){
	if user.SignContinuousDays == nil {
		user.SignContinuousDays = proto.Int32(0) //addSignCount
	}
	if user.SignTotalDays == nil {
		user.SignTotalDays = proto.Int32(0) //addSignCount
	}
	atomic.AddInt32(user.SignContinuousDays, 1)
	atomic.AddInt32(user.SignTotalDays, 1)
}

func clearContinuousSignCount(user *ddproto.User){
	user.SignContinuousDays = proto.Int32(0) //clearContinuousSignCount
}


func DoSignLottery(userId uint32) error {
	log.T("[%v]开始签到领取奖励", userId)

	user := userService.GetUserById(userId)
	if user == nil {
		log.T("存储中找不到u.id为[%v]的用户", userId)
		return Error.NewError(-1, "找不到用户个人信息")
	}

	err1 := DoSign(user)
	if err1 != nil {
		log.T("[%v]签到领取奖励失败, 错误[%v]", user.GetId(), err1)
		return err1
	}

	err2 := deliveryUserSignLottery(user)
	if err2 != nil {
		log.T("[%v]签到领取奖励失败, 错误[%v]", user.GetId(), err2)
		return err2
	}

	log.T("[%v]签到领取奖励成功", user.GetId())
	return nil
}

func getUserLotteryId(user *ddproto.User) int32 {
	//todo 从数据库中取出连续签到天数对应奖励物品id
	m := make(map[int32]int32) //map[signContinuousDays]lotteryId
	m[2] = 1
	m[5] = 2
	m[7] = 3
	m[10] = 4
	m[15] = 5
	m[20] = 6
	m[50] = 7
	return m[user.GetSignCount()]
}

//根据签到情况获得可领的奖品id
func deliveryUserSignLottery(user *ddproto.User) error {
	//todo 从数据库根据奖励id中获得奖励类型和数量 设置玩家数据
	switch getUserLotteryId(user) {
	case 1:
		userService.INCRUserCOIN(user.GetId(), 50)
	case 2:
		userService.INCRUserCOIN(user.GetId(), 100)
	case 3:
		userService.INCRUserCOIN(user.GetId(), 200)
	case 4:
		userService.INCRUserCOIN(user.GetId(), 500)
	case 5:
		userService.INCRUserCOIN(user.GetId(), 1000)
	case 6:
		userService.INCRUserCOIN(user.GetId(), 2000)
	case 7:
		userService.INCRUserCOIN(user.GetId(), 5000)

	default:
	}
	return nil
}


func IsUserSignedToday(u *ddproto.User) bool {
	timeNow := time.Now()
	lastSignTime, _ := time.Parse(u.GetLastSignTime(), "2006-01-02 15:04:05")
	return timeUtils.EqualDate(lastSignTime, timeNow)
}
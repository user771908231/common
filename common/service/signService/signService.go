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
	"casino_common/common/model/userSignDao"
	"casino_hall/service/task"
	"errors"
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
	lastSignTime := timeUtils.String2YYYYMMDDHHMMSS(user.GetLastSignTime())

	//上次签到＋1天 是否等于今天
	if !timeUtils.EqualDate(lastSignTime.AddDate(0, 0, 1), timeNow) {
		log.T("上次签到日期[%v]+1天不等于当前日期[%v]", user.GetLastSignTime(), timeUtils.Format(timeNow))
		return false
	}
	return true
}

func isAbleToSign(user *ddproto.User) bool {
	//验证用户能否签到
	if user.GetLastSignTime() != "" {
		timeNow := time.Now()
		lastSignTime := timeUtils.String2YYYYMMDDHHMMSS(user.GetLastSignTime())
		//上次签到时间在当前时间之后
		if lastSignTime.After(timeNow) {
			log.T("上次签到时间[%v]在当前时间[%v]之后", user.GetLastSignTime(), timeUtils.Format(timeNow))
			return false
		}

		// 每天只能签一次
		//上次签到的日期(天)等于当前日期(天)
		if timeUtils.EqualDate(lastSignTime, timeNow) {
			log.T("上次签到日期[%v]等于当前日期[%v]", user.GetLastSignTime(), timeUtils.Format(timeNow))
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

//签到
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

//领取签到奖品
func deliveryUserSignLottery(user *ddproto.User) error {
	reward := userSignDao.FindSignRewardByDay(user.GetSignContinuousDays())
	if reward == nil {
		return errors.New("领取奖品失败")
	}
	//发放奖品
	newReward := []*ddproto.HallBagItem{
		&ddproto.HallBagItem{
			Type:ddproto.HallEnumTradeType(reward.Type).Enum(),
			Amount:&reward.Amount,
		},
	}
	task.CheckReward(user.GetId(), newReward)
	return nil
}


func IsUserSignedToday(u *ddproto.User) bool {
	timeNow := time.Now()
	lastSignTime := timeUtils.String2YYYYMMDDHHMMSS(u.GetLastSignTime())
	return timeUtils.EqualDate(lastSignTime, timeNow)
}
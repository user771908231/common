package signService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/Error"
	"time"
	"casino_common/common/log"
	"casino_common/common/userService"
	"casino_common/utils/timeUtils"
	"fmt"
	"casino_common/common/model/userSignDao"
	task "casino_common/common/service/taskService"
	"errors"
	"casino_common/common/model/userAttachDao"
	"casino_common/common/model"
)


//用户签到
func DoSign(user *ddproto.User) error {

	log.T("[%v]开始签到", user.GetId())

	userAttach := userAttachDao.FindUserAttachByUserId(user.GetId())

	//验证用户能否签到
	if !isAbleToSign(userAttach) {
		return Error.NewError(-1, fmt.Sprintf("用户[%v]不能签到,签到失败", user.GetId()))
	}

	//不是连续签到 清空连续签到计数
	if !isContinuousSign(userAttach) {
		clearContinuousSignCount(userAttach)
	}

	addSignCount(userAttach)
	userAttach.LastSignTime = timeUtils.Format(time.Now())
	userAttachDao.UpdateUserAttachByModel(userAttach)

	log.T("[%v]签到成功", user.GetId())

	return nil;
}

func isContinuousSign(userAttach *model.T_user_attach) bool {
	if userAttach.LastSignTime == "" {
		return false
	}
	timeNow := time.Now()
	lastSignTime := timeUtils.String2YYYYMMDDHHMMSS(userAttach.LastSignTime)

	//上次签到＋1天 是否等于今天
	if !timeUtils.EqualDate(lastSignTime.AddDate(0, 0, 1), timeNow) {
		log.T("上次签到日期[%v]+1天不等于当前日期[%v]", userAttach.LastSignTime, timeUtils.Format(timeNow))
		return false
	}
	return true
}

func isAbleToSign(userAttach *model.T_user_attach) bool {
	//验证用户能否签到
	if userAttach.LastSignTime != "" {
		timeNow := time.Now()
		lastSignTime := timeUtils.String2YYYYMMDDHHMMSS(userAttach.LastSignTime)
		//上次签到时间在当前时间之后
		if lastSignTime.After(timeNow) {
			log.T("上次签到时间[%v]在当前时间[%v]之后", userAttach.LastSignTime, timeUtils.Format(timeNow))
			return false
		}

		// 每天只能签一次
		//上次签到的日期(天)等于当前日期(天)
		if timeUtils.EqualDate(lastSignTime, timeNow) {
			log.T("上次签到日期[%v]等于当前日期[%v]", userAttach.LastSignTime, timeUtils.Format(timeNow))
			return false
		}
	}
	return true
}

func addSignCount(userAttach *model.T_user_attach){
	userAttach.SignContinuousDays++
	userAttach.SignTotalDays++
	userAttachDao.UpdateUserAttachByModel(userAttach)
}

func clearContinuousSignCount(userAttach *model.T_user_attach){
	userAttach.SignContinuousDays = 0 //clearContinuousSignCount
	userAttachDao.UpdateUserAttachByModel(userAttach)
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

	log.T("[%v]签到领取奖励成功，签到%v天。", user.GetId(), user.GetSignContinuousDays())
	return nil
}

//领取签到奖品
func deliveryUserSignLottery(user *ddproto.User) error {
	userAttach := userAttachDao.FindUserAttachByUserId(user.GetId())
	reward := userSignDao.FindSignRewardByDay(userAttach.SignContinuousDays)
	if reward == nil {
		return errors.New("没有找到对应的签到奖励！")
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


func IsUserSignedToday(userAttach *model.T_user_attach) bool {
	timeNow := time.Now()
	lastSignTime := timeUtils.String2YYYYMMDDHHMMSS(userAttach.LastSignTime)
	return timeUtils.EqualDate(lastSignTime, timeNow)
}
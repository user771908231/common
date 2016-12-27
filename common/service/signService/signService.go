package signService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/Error"
	"sync/atomic"
	"time"
	"casino_common/common/log"
	"github.com/golang/protobuf/proto"
)


//用户签到
func DoSign(user *ddproto.User) error {

	//验证用户能否签到
	if user.GetLastSignTime() != "" {
		timeNow := time.Now()
		lastSignTime, _ := time.Parse(user.GetLastSignTime(), "2006-01-02 15:04:05")
		//上次签到时间在当前时间之后
		if lastSignTime.After(timeNow) {
			log.T("上次签到时间[%v]在当前时间[%v]之后", user.GetLastSignTime(), timeNow.Format("2006-01-02 15:04:05"))
			return Error.NewError(-1, "上次签到时间在当前时间之后, 签到失败")
		}

		// 每天只能签一次
		//上次签到的日期(天)等于当前日期(天)
		if IsSameDate(lastSignTime, timeNow) {
			log.T("上次签到日期[%v]等于当前日期[%v]", user.GetLastSignTime(), timeNow.Format("2006-01-02 15:04:05"))
			return Error.NewError(-1, "上次签到日期等于当前日期, 签到失败")
		}
	}

	addSignCount(user)
	user.LastSignTime = proto.String(time.Now().Format("2006-01-02 15:04:05"))

	return nil;
}

func addSignCount(user *ddproto.User){
	if user.SignCount == nil {
		user.SignCount = proto.Int32(1)
	}else{
		atomic.AddInt32(user.SignCount, 1)
	}
}


func DoSignLottery(user *ddproto.User) int32 {
	//todo
	return 2
}



func IsUserSignedToday(u *ddproto.User) bool {
	timeNow := time.Now()
	lastSignTime, _ := time.Parse(u.GetLastSignTime(), "2006-01-02 15:04:05")
	return IsSameDate(lastSignTime, timeNow)
}

//判断两个日期是否是同一天
func IsSameDate(time1 time.Time, time2 time.Time) bool {
	if time1.Year() == time2.Year() &&
		time1.Month() == time2.Month() &&
		time1.Day() == time2.Day() {
		return true
	}
	return false
}

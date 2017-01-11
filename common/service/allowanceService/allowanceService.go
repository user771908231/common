package allowanceService

import (
	"github.com/name5566/leaf/gate"
	"casino_common/common/model/userAttachDao"
	"casino_common/proto/funcsInit"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/timeUtils"
	"casino_common/common/Error"
	"casino_common/common/log"
	"errors"
	"time"
	"casino_common/common/model"
)

var TIEMES2GETADAY int32 = 2 //用户一天能领多少次补助

func DoAllowance(uid uint32, a gate.Agent) error {
	log.T("开始领取补助流程")
	user := userService.GetUserById(uid)
	if user == nil {
		return errors.New("找不到用户")
	}


	//检查今日领取补助的次数
	times := GetTimes4UserReceiveAllowanceToday(user)
	log.T("用户[%v]今日领取补助的次数[%v]", uid, times)
	if times >= TIEMES2GETADAY {
		//无法领取
		ack := commonNewPorot.NewAckAllowance()
		*ack.UserId = uid
		*ack.IsSuccessed = false
		a.WriteMsg(ack)
		return Error.NewError(-1, "今日领取限额已满")
	}


	//领取补助
	//更新redis和数据库金币
	userService.INCRUserCOIN(user.GetId(), 1000) //配置？

	//更新内存
	*user.Coin += 1000

	//更新时间与次数
	UpdateUserAllowanceInfo(user)

	//ack
	ack := commonNewPorot.NewAckAllowance()
	*ack.UserId = uid
	*ack.IsSuccessed = true
	*ack.UserTotalCoin = user.GetCoin()
	*ack.Times2Get = times
	*ack.AllowanceCoin = 1000
	a.WriteMsg(ack)
	return nil
}


//返回今天领取补助的次数
func GetTimes4UserReceiveAllowanceToday(user *ddproto.User) int32 {
	//用户今天没领取补助
	if !IsUserReceiveAllowanceToday(user) {
		return 0
	}

	userAttach := userAttachDao.FindUserAttachByUserId(user.GetId())
	return userAttach.AllowanceTimes
}


//用户今天是否领取补助
func IsUserReceiveAllowanceToday(user *ddproto.User) bool {
	userAttach := userAttachDao.FindUserAttachByUserId(user.GetId())
	if userAttach == nil || userAttach.LastAllowanceTime == "" {
		//log.T("IsUserReceiveAllowanceToday user[%v] userAttach is nil", user.GetId())
		return false
	}

	lastAllowanceTime := timeUtils.String2YYYYMMDDHHMMSS(userAttach.LastAllowanceTime)
	//log.T("IsUserReceiveAllowanceToday lastAllowanceTime[%v]", userAttach.LastAllowanceTime)
	return timeUtils.EqualDate(time.Now(), lastAllowanceTime)
}


//更新领取记录 时间和次数
func UpdateUserAllowanceInfo(user *ddproto.User) {
	userAttach := userAttachDao.FindUserAttachByUserId(user.GetId())

	needInsert := false
	if userAttach == nil { //为空需要插入一条新数据
		//log.T("UpdateUserAllowanceInfo userAttach is nil")
		userAttach = &model.T_user_attach{}
		userAttach.UserId = user.GetId()
		needInsert = true
	}

	//今天没有领取补助 清空计数
	if !IsUserReceiveAllowanceToday(user) {
		userAttach.AllowanceTimes = 0
	}

	userAttach.AllowanceTimes++

	userAttach.LastAllowanceTime = timeUtils.Format(time.Now())
	if needInsert {
		userAttachDao.InsertUserAttachByModel(userAttach)
	}
	userAttachDao.UpdateUserAttachByModel(userAttach)
}

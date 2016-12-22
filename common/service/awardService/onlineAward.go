package awardService

import (
	"casino_common/proto/ddproto"
	"casino_common/utils/redisUtils"
	"github.com/golang/protobuf/proto"
	"casino_common/utils/timeUtils"
	"time"
	"casino_common/common/log"
	"github.com/name5566/leaf/gate"
	"casino_common/common/userService"
	"errors"
)

const DRUATION_ONLINE int64 = 10 * 60

//在线奖励
func InitOnlineAward(userId uint32) (*ddproto.AwardMOnline, error) {
	//ddproto.AwardMOnline{} 初始化
	d := getOnlineData(userId)
	//初始化之后的数据保存到redis
	oninit(d)
	setOnlienData(d)
	return d, nil
}

func oninit(d *ddproto.AwardMOnline) {
	tnow := time.Now()
	d.BeginTime = proto.String(timeUtils.Format(tnow))
	d.DurationSec = proto.Int64(DRUATION_ONLINE) //现在是默认10分钟，之后修改成可以配置的
	d.Capital = proto.Int64(0)
	d.Rate = proto.Int64(3)
}

//获取在线奖励的model
func getOnlineData(userId uint32) *ddproto.AwardMOnline {
	d := redisUtils.GetObj(redisUtils.K(AWARD_ONLINE_READIS_KEY, userId), &ddproto.AwardMOnline{})
	if d == nil {
		d := new(ddproto.AwardMOnline)
		d.UserId = proto.Uint32(userId)
		log.T("没有找到玩家[%v]在线奖励的数据（第一次登录的时候没有）...")
		setOnlienData(d)
		return d
	} else {
		return d.(*ddproto.AwardMOnline)
	}
}

func setOnlienData(d *ddproto.AwardMOnline) {
	redisUtils.SetObj(redisUtils.K(AWARD_ONLINE_READIS_KEY, d.GetUserId()), d)
}

//处理奖励
func DoAwardOnline(uid uint32, a gate.Agent) error {
	//加锁来处理
	l := userService.UserLockPools.GetUserLockByUserId(uid)
	l.Lock()
	defer l.Unlock()

	//处理在线奖励
	d := getOnlineData(uid)
	tnow := time.Now()
	if timeUtils.String2YYYYMMDDHHMMSS(d.GetBeginTime()).Add(time.Second * time.Duration(d.GetDurationSec())).Unix() > tnow.Unix() {
		return errors.New("时间不够，不能获得奖励")
	}
	//开始发送奖励
	award := d.GetCapital() * d.GetRate() % 100
	balance, _ := userService.INCRUserCOIN(uid, award) //增加用户的金币

	ack := new(ddproto.WardAckOnline)
	ack.Coin = proto.Int64(balance)
	ack.Duration = proto.Int64(DRUATION_ONLINE) // 默认值
	ack.ChangeCoin = proto.Int64(award)

	//再次初始化，并且放置在数据库中...
	oninit(d)
	setOnlienData(d)
	//发送给app
	a.WriteMsg(ack)
	return nil
}

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

const DRUATION_ONLINE int64 = 15 //第一次是15秒

//在线奖励
func InitOnlineAward(userId uint32) (*ddproto.AwardMOnline, error) {
	//ddproto.AwardMOnline{} 初始化
	d := getOnlineData(userId)
	//初始化之后的数据保存到redis
	oninit(d)
	return d, nil
}

//登录的时候初始化在线奖励,
/**
	1，原来的登录奖励记录是当天，需要重新开始计时，但是时间等级一致
	2，原来的登录奖励是今天之前，那么全部从头开始计算
	3，初始化之后保存到redis
 */
func oninit(d *ddproto.AwardMOnline) {
	//需要判断是不是同一天
	tnow := time.Now()
	dbeginTime := timeUtils.String2YYYYMMDDHHMMSS(d.GetBeginTime())
	d.Rate = proto.Int64(OnlineConfig.Rate)
	d.BeginTime = proto.String(timeUtils.Format(tnow))

	log.T("oninit tnow:%v", tnow)
	if timeUtils.EqualDate(tnow, dbeginTime) {
		//表示是同一天
		log.T("玩家[%v]同一天重复登录，设置在线奖励...", d.GetUserId())
		if d.GetLevel() < OnlineConfig.MaxLevel {
			d.Level = proto.Int32(d.GetLevel() + 1) //表示是第一季
		}
	} else {
		log.T("玩家[%v]新的一天登录，设置在线奖励...", d.GetUserId())
		//表示不是同一天
		d.Level = proto.Int32(OnlineConfig.LeastLevel) //表示是第一季
	}
	d.DurationSec = proto.Int64(OnlineConfig.TimeConfig[d.GetLevel()]) //现在是默认10分钟，之后修改成可以配置的
	d.Capital = proto.Int64(0)
	log.T("初始化玩家[%v]在线奖励的详情[%v]", d.GetUserId(), d)
	setOnlineData(d)

}

//获取在线奖励的model
func getOnlineData(userId uint32) *ddproto.AwardMOnline {
	d := redisUtils.GetObj(redisUtils.K(AWARD_ONLINE_READIS_KEY, userId), &ddproto.AwardMOnline{})
	if d == nil {
		d := new(ddproto.AwardMOnline)
		d.UserId = proto.Uint32(userId)
		log.T("没有找到玩家[%v]在线奖励的数据（第一次登录的时候没有）...")
		setOnlineData(d)
		return d
	} else {
		log.T("得到用户的在线奖励detail %v ", d)
		return d.(*ddproto.AwardMOnline)
	}
}

func setOnlineData(d *ddproto.AwardMOnline) {
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
	tend := timeUtils.String2YYYYMMDDHHMMSS(d.GetBeginTime()).Add(time.Second * time.Duration(d.GetDurationSec()))
	if timeUtils.DiffSec(tend, tnow) < 0 {
		log.T("时间不够begin: %v ,duration:%v，[%v]不能获得奖励", d.GetBeginTime(), d.GetDurationSec(), uid)
		return errors.New("时间不够，不能获得奖励")
	}
	//开始发送奖励
	award := d.GetCapital() * d.GetRate() % 100
	balance, _ := userService.INCRUserCOIN(uid, award) //增加用户的金币//是否需要增加金币的订单

	ack := new(ddproto.WardAckOnline)
	ack.Coin = proto.Int64(balance)
	ack.Duration = proto.Int64(DRUATION_ONLINE) // 默认值
	ack.ChangeCoin = proto.Int64(award)

	//再次初始化，并且放置在数据库中...
	oninit(d)
	//发送给app
	a.WriteMsg(ack)
	return nil
}

//获取奖励时间
func GetOnlineInfo(userId uint32, a gate.Agent) error {
	//获取在线奖励详情
	d := getOnlineData(10213)
	//计算时间差
	dura := timeUtils.DiffSec(time.Now(), timeUtils.String2YYYYMMDDHHMMSS(d.GetBeginTime()).Add(time.Second * time.Duration(d.GetDurationSec())))

	//回复消息
	ack := new(ddproto.AwardAckOnlineInfo)
	ack.Duration = proto.Int64(dura)
	a.WriteMsg(ack)
	return nil
}

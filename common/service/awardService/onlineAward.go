package awardService

import (
	"casino_common/proto/ddproto"
	"casino_common/utils/redisUtils"
	"github.com/golang/protobuf/proto"
	"casino_common/utils/timeUtils"
	"time"
	"casino_server/common/log"
	"github.com/name5566/leaf/gate"
)

//在线奖励
func InitOnlineAward(userId uint32) (*ddproto.AwardMOnline, error) {
	//ddproto.AwardMOnline{} 初始化
	d := getOnlineData(userId)
	d.UserId = proto.Uint32(userId)
	tnow := time.Now()
	d.BeginTime = proto.String(timeUtils.Format(tnow))
	d.DurationSec = proto.Int64(60 * 10) //现在是默认10分钟，之后修改成可以配置的
	d.Capital = proto.Int64(0)
	d.Rate = proto.Int64(3)
	//初始化之后的数据保存到redis
	setOnlienData(d)
	return d, nil
}

//获取在线奖励的model
func getOnlineData(userId uint32) *ddproto.AwardMOnline {
	d := redisUtils.GetObj(redisUtils.K(AWARD_ONLINE_READIS_KEY, userId), &ddproto.AwardMOnline{})
	if d == nil {
		d := new(ddproto.AwardMOnline)
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
func DoAwardOnline(userId uint32, a gate.Agent) error {
	return nil


}

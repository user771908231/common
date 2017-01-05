package awardService

import (
	"testing"
	"casino_common/common/sys"
	"casino_common/utils/timeUtils"
	"time"
)

func init() {
	sys.InitRedis("127.0.0.1:6379", "test")
}

func TestGetOnlineInfo(t *testing.T) {
	d := getOnlineData(10213)
	timeBegin := timeUtils.String2YYYYMMDDHHMMSS(d.GetBeginTime())
	timeEnd := timeBegin.Add(time.Second * time.Duration(d.GetDurationSec()))
	dura := timeUtils.DiffSec(timeEnd, time.Now())
	t.Logf("dura:%v \n", dura)
}

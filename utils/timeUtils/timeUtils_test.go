package timeUtils

import (
	"casino_server/utils/timeUtils"
	"testing"
	"time"
)

func TestString2YYYYMMDDHHMMSS(t *testing.T) {
	tn := time.Now()
	t.Logf("当前的时间:%v", tn)
	stn := timeUtils.Format(tn)
	t.Logf("format之后:%v", stn)
	tn2 := timeUtils.String2YYYYMMDDHHMMSS(stn)
	t.Logf("tn2:%v", tn2)
}

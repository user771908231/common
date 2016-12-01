package jobUtils

import (
	"time"
	"casino_common/common/Error"
)

//定义一个异步任务
func DoAsynJob(d time.Duration, f func() bool) {
	ticker := time.NewTicker(d)
	go func() {
		defer Error.ErrorRecovery("定时任务中")
		for _ = range ticker.C {
			if f() {
				break;
			}
		}
	}()
}
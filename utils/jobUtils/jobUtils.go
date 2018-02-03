package jobUtils

import (
	"casino_common/common/Error"
	"time"
)

//定义一个异步任务 间隔 d time.Duration 执行一次函数f
func DoAsynJob(d time.Duration, f func() bool) {
	ticker := time.NewTicker(d)
	go func() { //定时任务中
		defer Error.ErrorRecovery("定时任务中")
		for _ = range ticker.C {
			if f() {
				break
			}
		}
	}()
}

package redisUtils

import (
	"testing"
	"fmt"
	"casino_common/utils/redis"
	"casino_common/common/log"
)

func init() {
	data.InitRedis("192.168.199.200:6379", "test")
	log.InitLogger("", "") //初始化日志处理
}

func TestData_Open(t *testing.T) {
	for i := 1; i < 100000; i++ {
		fmt.Printf("%v ...........\n", i)
		go func() {
			r := GetConn()
			fmt.Println("处理...")
			r.Close()
		}()
	}
}

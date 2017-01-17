package redisUtils

import (
	"testing"
	"fmt"
	"casino_common/utils/redis"
	"casino_common/common/log"
	"sync/atomic"
)

func init() {
	data.InitRedis("192.168.199.200:6379", "test")
	log.InitLogger("", "") //初始化日志处理
}

var count int32 = 0

func TestData_Open(t *testing.T) {

	done := make(chan bool)
	max := int32(100000)
	for i := 1; i < int(max); i++ {
		fmt.Printf("%v ...........\n", i)
		go func() {
			r := GetConn()
			atomic.AddInt32(&count, 1)
			fmt.Printf("处理...%v \n", count)
			r.Close()
			if count > (max - 2 ) {
				done <- true
			}
		}()
	}

	<-done
	fmt.Println("最后的结果...", count)

}

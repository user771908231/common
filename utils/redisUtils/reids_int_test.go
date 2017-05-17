package redisUtils

import (
	"casino_common/common/consts"
	"casino_common/common/log"
	"casino_common/utils/redis"
	"fmt"
	"sync/atomic"
	"testing"
)

func init() {
	data.InitRedis("192.168.2.145:6379", "test")
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

func TestSADD(t *testing.T) {
	SADD(consts.RKEY_ALL_ONLINE_USERS, "user198")
}

func TestSCARD(t *testing.T) {
	ret := SCARD(consts.RKEY_ALL_ONLINE_USERS)
	t.Logf("得到的数据是:%v", ret)
}

func TestSMEMBERS(t *testing.T) {
	ret := SMEMBERS(consts.RKEY_ALL_ONLINE_USERS)
	t.Logf("得到的数据是:%v", ret)
}

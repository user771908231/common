package iputils

import (
	"testing"
	"time"
	"casino_common/common/log"
	"casino_common/common/sys"
)

func init() {
	sys.InitRedis("localhost", "test")
}

func TestIP2site(t *testing.T) {
	add := iP2site("118.112.57.224")
	t.Logf("得到的地址:%v", add)
}

func TestIP2siteAsyn(t *testing.T) {
	log.Printf("开始获取地址\n")
	iP2siteAsyn("118.112.57.224", func(add string) {
		log.Printf("这里是异步调用:%v \n ", add)
	})
	log.Printf("开始等待\n")
	time.Sleep(time.Second * 10)
	log.Printf("end\n")
}

func TestGetIPSite(t *testing.T) {
	log.Printf("%v \n", GetIPSite("118.112.57.224"))
}

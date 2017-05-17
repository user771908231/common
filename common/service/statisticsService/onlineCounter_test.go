package statisticsService

import (
	"casino_common/common/log"
	"casino_common/utils/redis"
	"testing"
)

func init() {
	data.InitRedis("192.168.2.145:6379", "test")
	log.InitLogger("", "") //初始化日志处理
}

func TestOnlineCountAll(t *testing.T) {
	count := OnlineCountAll()
	t.Logf("总共的数量:%v", count)
	ret2 := OnlineCountUsers()
	t.Logf("总共的数量:%v", ret2)
}

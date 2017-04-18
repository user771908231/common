package utils

import (
	"testing"
	"time"
)

func TestDingMessage_Send(t *testing.T) {
	for i := 0; i < 100; i++ {
		DingMsger.Send("测试1、\"")
		DingMsger.Send("测试2'''sfsfwe")
		DingMsger.Send("测试3[desk-哈哈---]sfasfafa''' sfsaf   ")
		DingMsger.Send("测试4\tfsahflalaf")
		DingMsger.Send("测试5")
		DingMsger.Send("测试6")
	}
	time.Sleep(time.Second * 10)
}

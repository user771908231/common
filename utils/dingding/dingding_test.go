package utils

import (
	"testing"
	"time"
)

func TestDingMessage_Send(t *testing.T) {
	DingMsger.Send("测试1、\"")
	time.Sleep(time.Second * 10)
}

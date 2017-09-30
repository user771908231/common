package utils

import (
	"testing"
	"time"
	"fmt"
)

func TestDingMessage_Send(t *testing.T) {
	DingMsger.Send("测试1、\"")
	time.Sleep(time.Second * 10)
}

func TestSingleRobot_SendTxt(t *testing.T) {
	for i:=0;i<10;i++ {
		res := WxRobot.SendTxt(fmt.Sprintf("%d", i))
		t.Log(res)
	}
}

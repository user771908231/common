package testUtils

import (
	"testing"
)

func TestGetTestXiPai(t *testing.T) {
	ret := GetTestXiPai(1)
	t.Logf("最后的结果:%v", ret)
}

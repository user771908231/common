package loginService

import (
	"testing"
	"casino_common/test"
)

func init() {
	test.TestInit()
}
func TestTouristReg(t *testing.T) {
	ret := TouristReg("","")
	t.Logf("游客注册: %v", ret)
}

package userService

import (
	"casino_common/test"
	"testing"
)

func init() {
	test.TestInit()
}

func TestUSESSION_GetMainSession(t *testing.T) {
	var u U_REDIS = 10356
	s := u.GetMainSession()
	t.Logf("得到的session%v", s)
}

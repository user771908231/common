package sessionService

import (
	"casino_common/test"
	"testing"
)

func init() {
	test.TestInit()
}

func TestGetSession(t *testing.T) {
	userId := uint32(10404)
	//s := GetSession(10507, 2)
	fs := GetSession(userId, 1)
	t.Logf("得到的session： %v", fs)
	cs := GetSession(userId, 2)
	t.Logf("得到的session： %v", cs)
}

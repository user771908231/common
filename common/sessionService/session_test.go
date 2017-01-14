package sessionService

import (
	"casino_common/test"
	"testing"
)

func init() {
	test.TestInit()
}

func TestGetSession(t *testing.T) {
	//s := GetSession(10507, 2)
	s := GetSession(10244, 2)
	t.Logf("得到的session： %v", s)
}

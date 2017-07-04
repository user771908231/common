package roomAgent

import (
	"casino_common/test"
	"testing"
	"time"
)

func init() {
	test.TestInit()
}

func TestCreateDesk(t *testing.T) {
	t.Log(GetAgentRooms(1))
	err := CreateDesk(1, "12345", 1, 1, "qwe", time.Now().Unix())
	t.Log(err)
	t.Log(GetAgentRooms(1))
	DoAddUser(1, 1, 1, "james")
	DoAddUser(1, 1, 1, "james")
	DoAddUser(1, 1, 1, "smith")
	t.Log(GetAgentRooms(1))
	DoStart(1, 1, 1)
	t.Log(GetAgentRooms(1))
	err = DoEnd(1, 1, 1)
	t.Log(err)
	t.Log(GetAgentRooms(1))
}

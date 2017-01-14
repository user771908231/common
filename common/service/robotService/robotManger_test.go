package robotService

import (
	"testing"
	"casino_common/proto/ddproto"
	"casino_common/test"
)

func init() {
	test.TestInit()
}

func TestNewRobotManager(t *testing.T) {
	m := NewRobotManager(ddproto.CommonEnumGame_GID_MAHJONG)
	t.Logf("机器人总数量[%v]", len(m.robots))

	for _, r := range m.robots {
		t.Logf("机器人[%v]: %v", r.GetId(), r)

	}
}

func TestNewRobots(t *testing.T) {
	m := NewRobotManager(ddproto.CommonEnumGame_GID_MAHJONG)
	t.Logf("所有的机器人: %v", m.robots)

	for i := 0; i < 100; i++ {
		m.NewRobotAndSave()
		t.Logf("所有的机器人: %v", m.robots);
	}

}

func TestNewDdzRobots(t *testing.T) {
	m := NewRobotManager(ddproto.CommonEnumGame_GID_DDZ)
	t.Logf("所有的机器人: %v", m.robots)

	m.newRobotAndSave()
	t.Logf("所有的机器人: %v", m.robots)

}

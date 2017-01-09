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
	m.Oninit()
	t.Logf("所有的机器人: %v", m.robots)
}

func TestNewRobots(t *testing.T) {
	m := NewRobotManager(ddproto.CommonEnumGame_GID_MAHJONG)
	m.Oninit()
	t.Logf("所有的机器人: %v", m.robots)

	m.newRobotAndSave()
	t.Logf("所有的机器人: %v", m.robots)

}

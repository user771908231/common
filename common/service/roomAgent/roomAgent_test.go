package roomAgent

import (
	"casino_common/test"
	"testing"
	"time"
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
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

func TestHistoryList(t *testing.T) {
	ack := &ddproto.HallAckAgentRoomHistoryList{
		List: []*ddproto.CommonDeskByAgent{},
	}
	db.C(tableName.DBT_AGENT_CREATE_ROOM_RECORD).Page(bson.M{
		"creator": 11688,
	}, &ack.List, "-createtime", 1, 20)
	t.Log(*ack)
}
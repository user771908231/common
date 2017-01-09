package configDao_test

import (
	"testing"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"casino_common/common/model/configDao"
	"casino_common/test"
)

func init() {
	test.TestInit()
}

func TestInsertConfigSys(t *testing.T) {
	cfg := &ddproto.TConfigSys{
		Id:             proto.Int32(1),
		NewUserCoin:    proto.Int64(1000),
		NewUserRoomcard:proto.Int64(20),
	}
	configDao.UpsertConfigSys(cfg)
}

func TestGetConfigSys(t *testing.T) {
	ret := configDao.GetConfigSys()
	t.Logf("得到的结果是: %v", ret)
}

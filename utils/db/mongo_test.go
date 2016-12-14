package db_test

import (
	"testing"
	"casino_common/utils/db"
	"casino_common/common/model/userDao"
	"github.com/golang/protobuf/proto"
)

func init() {
	db.Oninit("127.0.0.1", 51668, "test", "id")
}

func TestMongo(t *testing.T) {
	t.Logf("开始查询")
	tuser := userDao.FindUserById(10164)
	tuser.Diamond2 = proto.Int64(998)
	userDao.UpdateUser2Mgo(tuser)
	t.Logf("查询到的结果:%v", tuser)
}
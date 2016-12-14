package db_test

import (
	"testing"
	"casino_common/utils/db"
	"github.com/golang/protobuf/proto"
)

func init() {
	db.Oninit("127.0.0.1", 51668, "test", "id")
}

type t_temp struct {
	Id       *uint32 //id
	Password *string
	T        *t_temp2
}

type t_temp2 struct {
	P2 *string
}

func TestMongo(t *testing.T) {
	t.Logf("开始测试...")
	tt := &t_temp{}
	tt.T = &t_temp2{}
	tt.T.P2 = proto.String("1010010")
	tt.Id = proto.Uint32(9232323)
	tt.Password = proto.String("myname888s")
	t.Logf("要保存的数据:%v", tt)
	err := db.InsertMgoData("t_temp", tt)
	t.Logf("保存之后的错误:%v", err)

}
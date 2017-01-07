package handlers

import (
	"casino_common/utils/db"
	"casino_common/common/sys"
	"testing"
	"casino_common/common/service/noticeServer"
)

func init() {
	db.Oninit("127.0.0.1", 51668, "test", "id")
	sys.InitRedis("127.0.0.1:6379", "test")
}

func TestHandlerGetCommonAckNotice(t *testing.T) {
	ret := noticeServer.GetCommonAckNotice(2)
	t.Logf("得到的公告信息..:%v", ret)
	t.Logf("得到的公告信息..fileds:%v", ret.GetFileds())

}

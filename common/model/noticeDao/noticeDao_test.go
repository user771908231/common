package noticeDao_test

import (
	"testing"
	"casino_common/utils/db"
	"casino_common/common/model/noticeDao"
)

func init() {
	db.Oninit("127.0.0.1", 51668, "test", "id")
}

func TestFindNoticeByType(t *testing.T) {
	notice := noticeDao.FindNoticeByType(2)
	t.Logf("查询到的结果:%v", notice)
}

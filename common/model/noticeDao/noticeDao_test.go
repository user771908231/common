package noticeDao_test

import (
	"testing"
	"casino_common/utils/db"
	"casino_common/common/model/noticeDao"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
)

func init() {
	db.Oninit("127.0.0.1", 51668, "test", "id")
}

func TestFindNoticeByType(t *testing.T) {
	notice := noticeDao.FindNoticeByType(2)
	t.Logf("查询到的结果:%v", notice)
}

func TestSaveNotice2Mgo(t *testing.T) {
	notice := new(ddproto.TNotice)
	notice.Id = proto.Int32(4)
	notice.NoticeTitle = proto.String("活动公告")
	notice.NoticeContent = proto.String("神经游戏诚招各地代理！申请代理门槛极低，利润丰厚，轻松月入万元。欢迎联系洽谈！")
	noticeDao.SaveNotice2Mgo(notice)
}

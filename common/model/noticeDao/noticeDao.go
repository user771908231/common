package noticeDao

import (
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
	"casino_common/proto/ddproto"
)

//得到一个notice
func FindNoticeByType(noticeType int32) *ddproto.TNotice {
	notice := new(ddproto.TNotice)
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_TH_NOTICE).Find(bson.M{"noticetype": noticeType}).One(notice)
	})
	if notice.GetId() > 0 {
		return notice
	} else {
		return nil
	}
}

//
func SaveNotice2Mgo(notice *ddproto.TNotice) error {
	return db.InsertMgoData(tableName.DBT_T_TH_NOTICE, notice)
}

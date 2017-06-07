package noticeDao

import (
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//得到一个notice
func FindNoticeByType(noticeType int32, channelId string) *ddproto.TNotice {
	notice := new(ddproto.TNotice)
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_T_TH_NOTICE).Find(bson.M{"noticetype": noticeType, "channelid": channelId}).One(notice)
	})
	if notice.GetId() > 0 {
		return notice
	} else {
		return nil
	}
}

//
func SaveNotice2Mgo(notice *ddproto.TNotice) error {
	return db.InsertMgoData("", tableName.DBT_T_TH_NOTICE, notice)
}

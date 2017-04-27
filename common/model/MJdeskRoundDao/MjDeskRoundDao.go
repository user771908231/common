package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/numUtils"
	"casino_common/utils/db"
	"casino_common/common/log"
	"casino_common/common/consts/tableName"
	"casino_hall/data/model"
	"casino_common/proto/ddproto"
)

//更具userId查询战绩
func GetMjDeskRoundByUserId(userId uint32) []model.T_mj_desk_round {
	var deskRecords []model.T_mj_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_MJ_DESK_ROUND).Find(bson.M{
			"userids":    bson.RegEx{querKey, "."},
			"friendplay": true}).Sort("-gamenumber").Limit(20).All(&deskRecords)
	})

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]麻将相关的战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}

func GetMjPlayBack(gamenumber int32) []*ddproto.PlaybackSnapshot {
	ret := &model.T_mj_desk_round{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_MJ_DESK_ROUND).Find(bson.M{"gamenumber": gamenumber}).One(ret)
	})
	if ret.DeskId > 0 {
		return ret.PlayBackData
	} else {
		return nil
	}
}

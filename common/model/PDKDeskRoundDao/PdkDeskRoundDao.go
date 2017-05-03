package PDKDeskRoundDao

import (
"gopkg.in/mgo.v2"
"casino_common/common/consts/tableName"
"gopkg.in/mgo.v2/bson"
"casino_common/common/model"
"casino_common/utils/numUtils"
"casino_common/utils/db"
"casino_common/common/log"
)

func GetPdkDeskRoundByUserId(userId uint32) []model.T_pdk_desk_round {
	var deskRecords []model.T_pdk_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_PDK_DESK_ROUND).Find(bson.M{"userids": bson.RegEx{querKey, "."}}).Sort("-deskid").Limit(20).All(&deskRecords)
	})

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]斗地主相关的战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}


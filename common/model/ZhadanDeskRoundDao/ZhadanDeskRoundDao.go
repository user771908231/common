package ZhadanDeskRoundDao

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/model"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"gopkg.in/mgo.v2/bson"
)

//查询总战绩
func GetZhadanDeskRoundByUserId(userId uint32) []model.T_zhadan_desk_round {
	var deskRecords []model.T_zhadan_desk_round
	querKey, _ := numUtils.Uint2String(userId)

	db.Log(tableName.DBT_ZHADAN_DESK_ROUND_ALL).Page(bson.M{"userids": bson.RegEx{querKey, "."}}, &deskRecords, "-endtime", 1, 20)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]炸弹相关的战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}

//查询俱乐部战绩
func GetZhadanDeskRoundByDeskIds(deskIds []int32) []model.T_zhadan_desk_round {
	var deskRecords []model.T_zhadan_desk_round

	db.Log(tableName.DBT_ZHADAN_DESK_ROUND_ALL).Page(bson.M{"deskid": bson.M{"$in": deskIds}}, &deskRecords, "-endtime", 1, 100)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到牌桌[%v]炸弹相关的战绩...", deskIds)
		return nil
	} else {
		return deskRecords
	}
}

//查询牌桌内战绩
func GetZhadanDeskRoundByDeskId(userId uint32, deskId int32) []model.T_zhadan_desk_round {
	var deskRecords []model.T_zhadan_desk_round
	querKey, _ := numUtils.Uint2String(userId)

	db.Log(tableName.DBT_ZHADAN_DESK_ROUND_ONE).Page(bson.M{
		"userids": bson.RegEx{querKey, "."},
		"deskid":  deskId,
	}, &deskRecords, "-gamenumber", 1, 20)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]炸弹相关的牌桌内战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}

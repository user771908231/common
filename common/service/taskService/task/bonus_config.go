package task

import (
	"gopkg.in/mgo.v2/bson"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
)

type T_Bonus_Config struct {
	Id bson.ObjectId `bson:"_id"`
	GameId ddproto.CommonEnumGame //所属游戏
	CoinPercent float64  //金币与红包的兑换比率
	BonusPercent float64  //红包兑换系数
	Min float64  //最小红包数
	Max float64 //最大红包数
}

//获取新版红包任务配置
func GetBonusTaskComputeConfigByGameId(gameid ddproto.CommonEnumGame) *T_Bonus_Config {
	row := new(T_Bonus_Config)
	err := db.C(tableName.DBT_T_BONUS_TASK_COMPUTE_CONFIG).Find(bson.M{
		"gameid": gameid,
	}, row)
	if err != nil {
		return nil
	}
	return row
}

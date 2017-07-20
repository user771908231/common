package hall

import (
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"casino_common/utils/timeUtils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//统计房卡的消耗
type T_statistics_roomcard struct {
	Time          time.Time //使用的时间
	UserId        uint32    //使用的玩家
	Gid           int32     //游戏的Id
	Memo          string    //备注
	RoomCardCount int64     //消耗房卡的数量，这里的roomCardCount是有正负值大小的
}

func (t T_statistics_roomcard) Insert() {
	db.InsertMgoData(tableName.DBT_STATISTICS_ROOMCARD, &t)
	//2，增加统计的数据
	T_statistics_roomcard_day_details{
		Time:          t.Time,
		Gid:           t.Gid,
		Day:           timeUtils.FormatYYYYMMDD(t.Time),
		RoomCardCount: t.RoomCardCount,
	}.UpSert()
}

//房卡的每日消耗记录
type T_statistics_roomcard_day_details struct {
	Time          time.Time //时间，按天来统计
	Day           string
	Gid           int32  //游戏id
	RoomCardCount int64  //房卡的数量
}

func (t T_statistics_roomcard_day_details) UpSert() {
	db.Query(func(db *mgo.Database) {
		//单个游戏
		db.C(tableName.DBT_STATISTICS_ROOMCARD_DAY_DETAILS).Upsert(
			bson.M{"day": t.Day, "gid": t.Gid},
			bson.M{"$inc": bson.M{"roomcardcount": t.RoomCardCount, }, "$set": bson.M{"time": t.Time}})
		//全部游戏
		if t.Gid != 0 {  //防止重复计数
			db.C(tableName.DBT_STATISTICS_ROOMCARD_DAY_DETAILS).Upsert(
				bson.M{"day": t.Day, "gid": 0},
				bson.M{"$inc": bson.M{"roomcardcount": t.RoomCardCount, }, "$set": bson.M{"time": t.Time}})
		}
	})
}

package hall

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/db"
	"casino_common/utils/timeUtils"
	"testing"
	"time"
)

func init() {
	db.InitMongoDb("192.168.2.145", "test", tableName.DB_ENSURECOUNTER_KEY, nil) //初始化mongo 的地址
}

func TestT_statistics_roomcard_day_details_UpSert(t *testing.T) {
	data := T_statistics_roomcard_day_details{
		Time:          timeUtils.FormatYYYYMMDD(time.Now()),
		Gid:           2,
		RoomCardCount: -1,
	}
	t.Logf("得到的数据:%+v", data)
	data.UpSert()
}

package pinerzhang

import (
	"time"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	"casino_common/common/Error"
)

type PEZRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

type T_pinerzhang_desk_round struct {
	DeskId     int32 //房间号
	GameNumber int32 //一局游戏编号
	UserIds    string
	BeginTime  time.Time
	EndTime    time.Time
	Records    []PEZRecordBean
}

//insert
func (t T_pinerzhang_desk_round) Insert() {
	//插入到数据库
	go func(d *T_pinerzhang_desk_round) {
		Error.ErrorRecovery("插入拼二张的战绩")
		db.InsertMgoData(tableName.DBT_PEZ_DESK_ROUND, d)
	}(&t)
}

//todo 找到拼二张的战绩列表
func ListTPEZDeskRound() []T_pinerzhang_desk_round {
	panic("实现超着拼二张")
}

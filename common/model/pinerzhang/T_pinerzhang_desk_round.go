package pinerzhang

import (
	"casino_common/common/Error"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"casino_common/utils/timeUtils"
	"github.com/golang/protobuf/proto"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type PEZRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

type T_pinerzhang_desk_round struct {
	DeskId     int32  //桌子ID
	Password   string //房间号
	GameNumber int32  //一局游戏编号
	UserIds    string
	BeginTime  time.Time
	EndTime    time.Time
	Records    []PEZRecordBean
	RoundStr   string //局数信息
}

func (b PEZRecordBean) TransBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{
		NickName:  proto.String(b.NickName),
		UserId:    proto.Uint32(b.UserId),
		WinAmount: proto.Int64(b.WinAmount),
	}
	return result
}

func (t T_pinerzhang_desk_round) TransRecord() *ddproto.BeanGameRecord {
	result := &ddproto.BeanGameRecord{
		BeginTime: proto.String(timeUtils.Format(t.BeginTime)),
		DeskId:    proto.Int32(t.DeskId),
		Password:  proto.String(t.Password),
		Id:        proto.Int32(t.GameNumber),
		RoundStr:  proto.String(t.RoundStr),
	}
	for _, bean := range t.Records {
		b := bean.TransBeanUserRecord()
		result.Users = append(result.Users, b)
	}
	return result
}

//insert
func (t T_pinerzhang_desk_round) Insert() {
	//插入到数据库
	go func(d *T_pinerzhang_desk_round) {
		Error.ErrorRecovery("插入拼二张的战绩")
		db.InsertMgoData(tableName.DBT_PEZ_DESK_ROUND, d)
	}(&t)
}

//查询总战绩
func GetPEZDeskRoundByUserId(userId uint32) []T_pinerzhang_desk_round {
	var deskRecords []T_pinerzhang_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	db.Log(tableName.DBT_PEZ_DESK_ROUND_ALL).Page(bson.M{
		"userids": bson.RegEx{querKey, "."},
	}, &deskRecords, "-gamenumber", 1, 20)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]拼二张相关的战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}

//查询牌桌内战绩
func GetPEZDeskRoundByDeskId(userId uint32, deskId int32) []T_pinerzhang_desk_round {
	var deskRecords []T_pinerzhang_desk_round
	querKey, _ := numUtils.Uint2String(userId)

	db.Log(tableName.DBT_PEZ_DESK_ROUND).Page(bson.M{
		"userids": bson.RegEx{querKey, "."},
		"deskid":  deskId,
	}, &deskRecords, "-gamenumber", 1, 20)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]拼二张相关的牌桌内战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}

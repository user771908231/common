package paodekuai

import (
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

//一把结束,战绩可以通过这个表来查询
type T_pdk_desk_round struct {
	DeskId     int32
	Password   string //房间号
	GameNumber int32
	UserIds    string
	BeginTime  time.Time
	EndTime    time.Time
	RoundStr   string //局数信息
	Records    []PdkRecordBean
}

func (t T_pdk_desk_round) TransRecord() *ddproto.BeanGameRecord {
	result := &ddproto.BeanGameRecord{
		BeginTime: proto.String(timeUtils.Format(t.BeginTime)),
		DeskId:    proto.Int32(t.DeskId), //这里实际返回的是房间号
		Password:  proto.String(t.Password),
		Id:        proto.Int32(t.GameNumber),
		RoundStr:  proto.String(t.RoundStr), //局数信息
	}
	for _, bean := range t.Records {
		b := bean.TransBeanUserRecord()
		result.Users = append(result.Users, b)
	}
	return result
}

type PdkRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

func (b PdkRecordBean) TransBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{
		NickName:  proto.String(b.NickName),
		UserId:    proto.Uint32(b.UserId),
		WinAmount: proto.Int64(b.WinAmount),
	}
	return result
}

//DAO
func GetPdkDeskRoundByUserId(userId uint32) []T_pdk_desk_round {
	var deskRecords []T_pdk_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	db.Log(tableName.DBT_PDK_DESK_ROUND_ALL).Page(bson.M{
		"userids": bson.RegEx{querKey, "."},
	}, &deskRecords, "-deskid", 1, 20)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]跑得快相关的战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}

//查询俱乐部战绩
func GetPdkDeskRoundByDeskIds(deskIds []int32) []T_pdk_desk_round {
	var deskRecords []T_pdk_desk_round
	db.Log(tableName.DBT_PDK_DESK_ROUND_ALL).Page(bson.M{
		"deskid": bson.M{"$in": deskIds},
	}, &deskRecords, "-deskid", 1, 100)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到桌子[%v]跑得快相关的战绩...", deskIds)
		return nil
	} else {
		return deskRecords
	}
}

//查询牌桌内战绩
func GetPdkDeskRoundByDeskId(userId uint32, deskId int32) []T_pdk_desk_round {
	var deskRecords []T_pdk_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	db.Log(tableName.DBT_PDK_DESK_ROUND).Page(bson.M{
		"userids": bson.RegEx{querKey, "."},
		"deskid":  deskId,
	}, &deskRecords, "-gamenumber", 1, 20)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]跑得快相关的牌桌内战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}

package majiang

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"casino_common/utils/timeUtils"
	"github.com/golang/protobuf/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MjRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

func (b MjRecordBean) TransBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{
		NickName:  proto.String(b.NickName),
		UserId:    proto.Uint32(b.UserId),
		WinAmount: proto.Int64(b.WinAmount),
	}
	return result
}

//一把结束,战绩可以通过这个表来查询
type T_mj_desk_round struct {
	DeskId       int32
	GameNumber   int32
	UserIds      string
	BeginTime    time.Time
	EndTime      time.Time
	Records      []MjRecordBean
	FriendPlay   bool   //是否是朋友桌
	PassWord     string //朋友桌的房间号
	PlayBackData []*ddproto.PlaybackSnapshot
	RoundStr     string //局数信息
}

func (t T_mj_desk_round) TransRecord() *ddproto.BeanGameRecord {
	result := &ddproto.BeanGameRecord{
		BeginTime: proto.String(timeUtils.Format(t.EndTime)),
		DeskId:    proto.Int32(t.DeskId),
		Password:  proto.String(t.PassWord),
		Id:        proto.Int32(t.GameNumber), //战绩id就是 游戏编号
		RoundStr:  proto.String(t.RoundStr),  //局数信息
	}

	for _, bean := range t.Records {
		b := bean.TransBeanUserRecord()
		result.Users = append(result.Users, b)
	}
	return result
}

//更具userId查询战绩
func GetMjDeskRoundByUserId(userId uint32) []T_mj_desk_round {
	var deskRecords []T_mj_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_MJ_DESK_ROUND_ALL).Find(bson.M{
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

//查询牌桌内战绩
func GetMjDeskRoundByDeskId(userId uint32, passWord int32) []T_mj_desk_round {
	var deskRecords []T_mj_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	pwd := numUtils.Int2String2(passWord)
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_MJ_DESK_ROUND).Find(bson.M{
			"userids":    bson.RegEx{querKey, "."},
			"friendplay": true,
			"password":   pwd,
		}).Sort("-deskid").Limit(20).All(&deskRecords)
	})

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]麻将相关的牌桌[%v]内战绩...", userId, passWord)
		return nil
	} else {
		return deskRecords
	}
}

func GetMjPlayBack(gamenumber int32) []*ddproto.PlaybackSnapshot {
	ret := &T_mj_desk_round{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_MJ_DESK_ROUND).Find(bson.M{"gamenumber": gamenumber}).One(ret)
	})
	if ret.DeskId > 0 {
		return ret.PlayBackData
	} else {
		return nil
	}
}

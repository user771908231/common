package majiang

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"casino_common/utils/timeUtils"
	"github.com/golang/protobuf/proto"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"time"
)

func init() {
	//初始化缓存相关
	PlayBackStack = map[int32][]*ddproto.PlaybackSnapshot{}
	PlayBackNumbers = []int32{}
}

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

//
type T_mj_desk_round_playback struct {
	DeskId       int32
	GameNumber   int32
	PlayBackData []byte
	EndTime      time.Time
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

func (t T_mj_desk_round) TransRecordCoin() *ddproto.BeanGameRecordCoin {
	result := &ddproto.BeanGameRecordCoin{
		BeginTime: proto.String(timeUtils.Format(t.EndTime)),
		DeskId:    proto.Int32(t.DeskId),
		Password:  proto.String(t.PassWord),
		Id:        proto.Int32(t.GameNumber), //战绩id就是 游戏编号
	}

	for _, bean := range t.Records {
		b := bean.TransBeanUserRecord()
		result.Users = append(result.Users, b)
	}
	return result
}

//根据userId查询全局战绩
func GetMjDeskRoundByUserId(userId uint32, gid, roomType int32) []T_mj_desk_round {
	var deskRecords []T_mj_desk_round
	querKey, _ := numUtils.Uint2String(userId)

	tbName := tableName.DBT_MJ_DESK_ROUND_ALL
	switch gid {
	case int32(ddproto.CommonEnumGame_GID_ZXZ):
		tbName = tableName.DBT_MJ_ZXZ_DESK_ROUND_ALL

	case int32(ddproto.CommonEnumGame_GID_MJBAISHAN):
		tbName = tableName.DBT_MJ_BS_DESK_ROUND_ALL

	case int32(ddproto.CommonEnumGame_GID_MJ_SONGJIANGHE):
		tbName = tableName.DBT_MJ_SJH_DESK_ROUND_ALL

	case int32(ddproto.CommonEnumGame_GID_ZHUANZHUAN):
		tbName = tableName.DBT_MJ_ZHZH_DESK_ROUND_ALL
		if roomType == int32(ddproto.MJRoomType_roomType_mj_hongzhong) {
			tbName = tableName.DBT_MJ_HZH_DESK_ROUND_ALL
		}
	case int32(ddproto.CommonEnumGame_GID_MJ_SHENQI):
		tbName = tableName.DBT_MJ_SHENQI_DESK_ROUND_ALL
	default:
	}

	log.T("userId%v gameId%v roomType%v tbName[%v]", userId, gid, roomType, tbName)
	db.Log(tbName).Page(bson.M{
		"userids":    bson.RegEx{querKey, "."},
		"friendplay": true,
	}, &deskRecords, "-deskid", 1, 20)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]麻将相关的战绩... gid[%v] roomType[%v]", userId, gid, roomType)
		return nil
	} else {
		return deskRecords
	}
}

//查询俱乐部战绩
func GetMjDeskRoundByDeskIds(DeskIds []int32, gid, roomType int32) []T_mj_desk_round {
	var deskRecords []T_mj_desk_round

	tbName := tableName.DBT_MJ_DESK_ROUND_ALL
	switch gid {
	case int32(ddproto.CommonEnumGame_GID_ZXZ):
		tbName = tableName.DBT_MJ_ZXZ_DESK_ROUND_ALL

	case int32(ddproto.CommonEnumGame_GID_MJBAISHAN):
		tbName = tableName.DBT_MJ_BS_DESK_ROUND_ALL

	case int32(ddproto.CommonEnumGame_GID_MJ_SONGJIANGHE):
		tbName = tableName.DBT_MJ_SJH_DESK_ROUND_ALL

	case int32(ddproto.CommonEnumGame_GID_ZHUANZHUAN):
		tbName = tableName.DBT_MJ_ZHZH_DESK_ROUND_ALL
		if roomType == int32(ddproto.MJRoomType_roomType_mj_hongzhong) {
			tbName = tableName.DBT_MJ_HZH_DESK_ROUND_ALL
		}
	case int32(ddproto.CommonEnumGame_GID_MJ_SHENQI):
		tbName = tableName.DBT_MJ_SHENQI_DESK_ROUND_ALL
	default:
	}

	log.T("deskids%v gameId%v roomType%v tbName[%v]", DeskIds, gid, roomType, tbName)
	db.Log(tbName).Page(bson.M{
		"deskid":     bson.M{"$in": DeskIds},
		"friendplay": true,
	}, &deskRecords, "-deskid", 1, 100)

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到desk[%v]麻将相关的战绩... gid[%v] roomType[%v]", DeskIds, gid, roomType)
		return nil
	} else {
		return deskRecords
	}
}

//根据userId查询金币场每局的战绩
func GetMjCoinDeskRoundByUserId(userId uint32, gid, roomType int32) []T_mj_desk_round {
	var deskRecords []T_mj_desk_round
	querKey, _ := numUtils.Uint2String(userId)

	tbName := ""
	switch gid {
	case int32(ddproto.CommonEnumGame_GID_MJ_ZIGONG):
		tbName = tableName.DBT_MJ_ZIGONG_DESK_ROUND_COIN
	default:
	}

	db.Log(tbName).Page(bson.M{
		"userids":    bson.RegEx{querKey, "."},
		"friendplay": false,
	}, &deskRecords, "-gamenumber", 1, 20)
	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]麻将相关的牌桌内战绩... gid[%v] roomType[%v]", userId, gid, roomType)
		return nil
	} else {
		return deskRecords
	}
}

//根据deskId查询单局战绩
func GetMjDeskRoundByDeskId(userId uint32, deskId, gid, roomType int32) []T_mj_desk_round {
	var deskRecords []T_mj_desk_round
	querKey, _ := numUtils.Uint2String(userId)

	tbName := tableName.DBT_MJ_DESK_ROUND
	switch gid {
	case int32(ddproto.CommonEnumGame_GID_ZXZ):
		tbName = tableName.DBT_MJ_ZXZ_DESK_ROUND

	case int32(ddproto.CommonEnumGame_GID_MJBAISHAN):
		tbName = tableName.DBT_MJ_BS_DESK_ROUND

	case int32(ddproto.CommonEnumGame_GID_MJ_SONGJIANGHE):
		tbName = tableName.DBT_MJ_SJH_DESK_ROUND

	case int32(ddproto.CommonEnumGame_GID_ZHUANZHUAN):
		tbName = tableName.DBT_MJ_ZHZH_DESK_ROUND
		if roomType == int32(ddproto.MJRoomType_roomType_mj_hongzhong) {
			tbName = tableName.DBT_MJ_HZH_DESK_ROUND
		}
	case int32(ddproto.CommonEnumGame_GID_MJ_SHENQI):
		tbName = tableName.DBT_MJ_SHENQI_DESK_ROUND
	default:
	}

	db.Log(tbName).Page(bson.M{
		"userids":    bson.RegEx{querKey, "."},
		"friendplay": true,
		"deskid":     deskId,
	}, &deskRecords, "-gamenumber", 1, 20)
	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]麻将相关的牌桌内战绩... deskId[%v] gid[%v] roomType[%v]", userId, deskId, gid, roomType)
		return nil
	} else {
		return deskRecords
	}
}

/************************* 麻将回放 start ******************************/
func GetMjPlayBack(gamenumber int32) []*ddproto.PlaybackSnapshot {
	ret := &T_mj_desk_round{}
	db.Log(tableName.DBT_MJ_DESK_ROUND).Find(bson.M{
		"gamenumber": gamenumber,
	}, &ret)
	if ret.DeskId > 0 {
		return ret.PlayBackData
	} else {
		return nil
	}
}

func GetMjZXZPlayBack(gamenumber int32) []*ddproto.PlaybackSnapshot {
	ret := &T_mj_desk_round{}
	db.Log(tableName.DBT_MJ_ZXZ_DESK_ROUND).Find(bson.M{
		"gamenumber": gamenumber,
	}, &ret)
	if ret.DeskId > 0 {
		return ret.PlayBackData
	} else {
		return nil
	}
}

func GetMjBSPlayBack(gamenumber int32) []*ddproto.PlaybackSnapshot {
	ret := &T_mj_desk_round{}
	db.Log(tableName.DBT_MJ_BS_DESK_ROUND).Find(bson.M{
		"gamenumber": gamenumber,
	}, &ret)
	if ret.DeskId > 0 {
		return ret.PlayBackData
	} else {
		return nil
	}
}

func GetMjShenQiPlayBack(gamenumber int32) []*ddproto.PlaybackSnapshot {
	ret := &T_mj_desk_round{}
	db.Log(tableName.DBT_MJ_SHENQI_DESK_ROUND).Find(bson.M{
		"gamenumber": gamenumber,
	}, &ret)
	if ret.DeskId > 0 {
		return ret.PlayBackData
	} else {
		return nil
	}
}

func GetMjZHZHPlayBack(gamenumber, roomType int32) []*ddproto.PlaybackSnapshot {
	ret := &T_mj_desk_round_playback{}
	tbName := tableName.DBT_MJ_ZHZH_DESK_ROUND_PLAYBACK
	if roomType == int32(ddproto.MJRoomType_roomType_mj_hongzhong) {
		tbName = tableName.DBT_MJ_HZH_DESK_ROUND_PLAYBACK
	}
	db.Log(tbName).Find(bson.M{
		"gamenumber": gamenumber,
	}, &ret)
	if ret.DeskId > 0 {
		log.T("查询到转转/红中麻将回放 回放数据大小[%v]byte", len(ret.PlayBackData))
		playBackData := &ddproto.PlaybackData{}
		err := proto.Unmarshal(ret.PlayBackData, playBackData)
		if err != nil {
			log.E("查询转转/红中麻将回放 反序列化回放数据出错 err:%v", err)
		}
		return playBackData.PlaybackSnapshots
	} else {
		return nil
	}
}

/************************* 麻将回放 end ******************************/

//==================================回放缓存==================================
//回放缓存-全局变量
var PlayBackStack map[int32][]*ddproto.PlaybackSnapshot

//缓存编号列表
var PlayBackNumbers []int32

//为缓存写上锁-保证线程安全
var playBackWLock sync.Mutex

//从内存缓存中取出-回放数据
func GetMjPlayBackFromMemory(gamenumber, gid, roomType int32) []*ddproto.PlaybackSnapshot {
	//如果缓存中存在则直接从缓存中取数据
	if data, ok := PlayBackStack[gamenumber]; ok {
		log.T("从内存读取gid[%v] roomType[%v]的数据：gamenumber:%d 当前缓存条数：%d", gid, roomType, gamenumber, len(PlayBackNumbers))
		return data
	}

	//如果不存在则向缓存中写入一条
	data := []*ddproto.PlaybackSnapshot{}
	switch gid {
	case int32(ddproto.CommonEnumGame_GID_ZXZ):
		data = GetMjZXZPlayBack(gamenumber)
	case int32(ddproto.CommonEnumGame_GID_MJBAISHAN):
		data = GetMjBSPlayBack(gamenumber)
	case int32(ddproto.CommonEnumGame_GID_ZHUANZHUAN):
		data = GetMjZHZHPlayBack(gamenumber, roomType)
	case int32(ddproto.CommonEnumGame_GID_MJ_SHENQI):
		data = GetMjShenQiPlayBack(gamenumber)
	default:
		data = GetMjPlayBack(gamenumber)
	}

	if data == nil {
		return nil
	}

	PlayBackStack[gamenumber] = data

	//为写上锁，保证线程安全
	playBackWLock.Lock()
	defer playBackWLock.Unlock()

	PlayBackNumbers = append(PlayBackNumbers, gamenumber)

	//如果缓存条数超过100条则删除队列最前的一条
	cache_len := len(PlayBackNumbers)
	if cache_len > 50 {
		old_number := PlayBackNumbers[0]
		delete(PlayBackStack, old_number)
		PlayBackNumbers = PlayBackNumbers[1:]
		log.T("mj回放缓存队列已满，删除一条最前面的数据gamenumber:%d", old_number)
	}

	log.T("从数据库读取mj数据：gamenumber:%d 当前缓存条数：%d", gamenumber, cache_len)
	return data
}

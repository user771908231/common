package paodekuai

import (
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"casino_common/common/log"
)

func init() {
	//初始化缓存相关
	PlayBackStack = map[int32][]*ddproto.PdkPlaybackSnapshot{}
	PlayBackNumbers = []int32{}
}

//回放可以通过这个表来查询
type T_pdk_desk_round_playback struct {
	DeskId       int32
	GameNumber   int32
	PlayBackData []*ddproto.PdkPlaybackSnapshot
}

func GetPdkPlayBack(gamenumber int32) []*ddproto.PdkPlaybackSnapshot {
	ret := &T_pdk_desk_round_playback{}
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_PDK_DESK_ROUND_PLAYBACK).Find(bson.M{"gamenumber": gamenumber}).One(ret)
	})
	if ret.DeskId > 0 {
		return ret.PlayBackData
	} else {
		return nil
	}
}

//==================================回放缓存==================================
//回放缓存-全局变量
var PlayBackStack map[int32][]*ddproto.PdkPlaybackSnapshot
//缓存编号列表
var PlayBackNumbers []int32
//为缓存写上锁-保证线程安全
var playBackWLock sync.Mutex

//从内存缓存中取出-回放数据
func GetPdkPlayBackFromMemory(gamenumber int32) []*ddproto.PdkPlaybackSnapshot {
	//如果缓存中存在则直接从缓存中取数据
	if data,ok := PlayBackStack[gamenumber]; ok {
		log.T("从内存读取pdk数据：gamenumber:%d 当前缓存条数：%d", gamenumber, len(PlayBackNumbers))
		return data
	}

	//如果不存在则向缓存中写入一条
	data := GetPdkPlayBack(gamenumber)
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
		log.T("pdk回放缓存队列已满，删除一条最前面的数据gamenumber:%d", old_number)
	}

	log.T("从数据库读取pdk数据：gamenumber:%d 当前缓存条数：%d", gamenumber, cache_len)
	return data
}

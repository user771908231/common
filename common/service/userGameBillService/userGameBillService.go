package userGameBillService

import (
	"time"
	"github.com/name5566/leaf/util"
	"gopkg.in/mgo.v2/bson"
	"casino_common/common/log"
	"casino_common/utils/db"
	"errors"
	"fmt"
	"casino_common/common/consts/tableName"
	"casino_common/common/Error"
	"casino_common/utils/numUtils"
	"strings"
	"casino_common/common/consts"
	"casino_common/utils/redisUtils"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
)

//玩家游戏账单数据 没玩完一局游戏存储一次数据
// 1.展示玩家输赢走势 2.下一局输赢判定策略的依据

func init() {
	UserBill = new(util.Map)
}

func OnInit(gameId, limitLength int32) {
	Cfg.gameId = gameId
	Cfg.limitLength = limitLength
	Cfg.tbName = tableName.DBT_USER_GAME_BILL
}

//程序启动加载时初始化配置信息
var Cfg struct {
	gameId      int32  //游戏id
	limitLength int32  //内存中数据的条数
	tbName      string //表名
}

//游戏账单数据 数据库里一局游戏一条记录
type T_user_game_bill struct {
	UserId     uint32
	GameId     int32     //游戏id
	RoomType   int32     //游戏类型 gameid的补充
	DeskId     int32     //房间Id
	Password   string    //房间号
	GameNumber int32     //一局游戏编号
	WinAmount  int64     //输赢得分
	EndTime    time.Time //结束时间
}

//将struct转成protobuf
func (t T_user_game_bill) GetProto() *ddproto.UserGameBill {
	ret := &ddproto.UserGameBill{}
	ret.UserId = proto.Uint32(t.UserId)
	ret.WinAmount = proto.Int64(t.WinAmount)
	ret.Password = proto.String(t.Password)
	ret.GameNumber = proto.Int32(t.GameNumber)
	ret.RoomType = proto.Int32(t.RoomType)
	ret.DeskId = proto.Int32(t.DeskId)
	ret.EndTime = proto.Int64(t.EndTime.Unix())
	return ret
}

func Transform(bill *ddproto.UserGameBill) T_user_game_bill {
	ret := T_user_game_bill{}
	ret.EndTime = time.Unix(bill.GetEndTime(), 0)
	ret.DeskId = bill.GetDeskId()
	ret.GameNumber = bill.GetGameNumber()
	ret.Password = bill.GetPassword()
	ret.WinAmount = bill.GetWinAmount()
	ret.UserId = bill.GetUserId()
	ret.GameId = Cfg.gameId
	ret.RoomType = bill.GetRoomType()
	return ret
}

//每个游戏只会维护自己游戏的数据
var UserBill *util.Map

//根据userId从数据库查询玩家账单数据
func GetViaDB(userId uint32, roomType int32) []T_user_game_bill {
	var gameBills []T_user_game_bill

	log.T("GetViaDB userId%v gameId%v roomType%v tbName[%v] limitLength[%v]", userId, Cfg.gameId, roomType, Cfg.tbName, Cfg.limitLength)
	db.Log(Cfg.tbName).Page(bson.M{
		"userid":   userId,
		"gameid":   Cfg.gameId,
		"roomtype": roomType,
	}, &gameBills, "-endtime", 1, int(Cfg.limitLength))

	if gameBills == nil || len(gameBills) <= 0 {
		log.T("没有从数据库中找到玩家[%v]相关的游戏账单数据... gid[%v] roomType[%v]", userId, Cfg.gameId, roomType)
		return nil
	}
	return gameBills
}

//redis key 例如 rkey_user_game_bill_11290_2_0
func getRedisKey(userId uint32, gameId, roomType int32) string {
	idstr, _ := numUtils.Uint2String(userId)
	gameIdSrt, _ := numUtils.Int2String(gameId)
	roomTypeSrt, _ := numUtils.Int2String(roomType)
	ret := strings.Join([]string{consts.RKEY_USER_GAME_BILL, idstr, gameIdSrt, roomTypeSrt}, "_")
	return ret
}

//根据userId从redis查询玩家账单数据
func GetViaRedis(userId uint32, roomType int32) []*ddproto.UserGameBill {
	log.T("GetViaRedis userId%v gameId%v roomType%v tbName[%v] limitLength[%v]", userId, Cfg.gameId, roomType, Cfg.tbName, Cfg.limitLength)
	gameBills := &ddproto.RedisUserGameBill{}

	redisUtils.GetObj(getRedisKey(userId, Cfg.gameId, roomType), gameBills)
	if gameBills == nil || gameBills.GetData() == nil || len(gameBills.GetData()) <= 0 {
		log.T("没有从redis里找到玩家[%v]相关的游戏账单数据... gid[%v] roomType[%v]", userId, Cfg.gameId, roomType)
		return nil
	}

	bills := gameBills.GetData()
	//返回限定长度
	if len(bills) > int(Cfg.limitLength) {
		return bills[:Cfg.limitLength]
	}
	return bills
}

//找到玩家账单数据 先从内存 找不到再从数据库并缓存到内存
func GetViaCache(userId uint32, roomType int32) []*ddproto.UserGameBill {
	log.T("开始查找玩家[%v]的游戏账单数据...", userId)
	gets := UserBill.Get(userId)
	bills := []*ddproto.UserGameBill{}
	if gets != nil {
		//从内存中获取到了 且长度符合要求 直接返回
		bills = gets.([]*ddproto.UserGameBill)
		if bills != nil && len(bills) == int(Cfg.limitLength) {
			log.T("从内存中查询玩家[%v]的游戏账单数据, 返回[%v]", userId, bills)
			return bills
		}
	}

	//从数据库中查询
	bills = GetViaRedis(userId, roomType)
	if bills != nil && len(bills) > 0 {
		//查询到了 缓存到内存里
		UserBill.Set(userId, bills)
		log.T("从redis中查询玩家[%v]的游戏账单数据, 缓存到内存并返回[%v]", userId, bills)
		return bills
	}
	log.W("没有找到玩家[%v]的游戏账单数据", userId)
	return nil
}

func Insert(userId uint32, roomType, deskId, gameNumber int32, password string, winAmount int64) {
	log.T("开始添加玩家[%v]的游戏账单数据...", userId)
	bill := T_user_game_bill{}
	bill.RoomType = roomType
	bill.GameId = Cfg.gameId
	bill.UserId = userId
	bill.DeskId = deskId
	bill.GameNumber = gameNumber
	bill.Password = password
	bill.WinAmount = winAmount
	bill.EndTime = time.Now()

	//同步插入到内存
	insert2Memory(bill.GetProto())

	//同步插入到redis
	insert2Redis(bill.GetProto(), roomType)

	//异步插入数据库
	go func(data interface{}) {
		defer Error.ErrorRecovery("保存玩家游戏账单到mgo")
		log.T("开始异步添加玩家[%v]的游戏账单数据到数据库中... bill[%v]", userId, bill)
		db.Log(Cfg.tbName).Insert(data)
	}(&bill)
	log.T("添加玩家[%v]的游戏账单数据完毕", userId)
}

//将一条账单插入到内存中
func insert2Memory(b *ddproto.UserGameBill) {
	log.T("开始添加玩家[%v]的游戏账单数据到内存中... bill[%v]", b.GetUserId(), b)
	//将内存中的bills取出来 把新的bill追加进去 再将汇总的bills重新set回去
	gets := UserBill.Get(b.GetUserId())
	userBills := []*ddproto.UserGameBill{}
	if gets != nil {
		userBills = gets.([]*ddproto.UserGameBill)
	}

	//将这条数据追加到第一个位置
	newUserBills := []*ddproto.UserGameBill{b}

	if len(userBills) >= int(Cfg.limitLength) {
		//原始长度超出限制
		newUserBills = append(newUserBills, userBills[:Cfg.limitLength-1]...)
	} else {
		newUserBills = append(newUserBills, userBills...)
	}

	UserBill.Set(b.GetUserId(), newUserBills)
	log.T("添加玩家[%v]的游戏账单数据到内存中完毕", b.GetUserId())
}

func insert2Redis(b *ddproto.UserGameBill, roomType int32) {
	log.T("开始添加玩家[%v]的游戏账单数据到redis中... bill[%v]", b.GetUserId(), b)
	//将内存中的bills取出来 把新的bill追加进去 再将汇总的bills重新set回去
	redisUserGameBill := &ddproto.RedisUserGameBill{}
	redisUtils.GetObj(getRedisKey(b.GetUserId(), Cfg.gameId, roomType), redisUserGameBill)

	if redisUserGameBill.GetData() == nil {
		redisUserGameBill.Data = []*ddproto.UserGameBill{b}
		redisUtils.SetObj(getRedisKey(b.GetUserId(), Cfg.gameId, roomType), redisUserGameBill)
		log.T("redis中没有数据, 第一次添加玩家[%v]的游戏账单数据到redis中完毕", b.GetUserId())
		return
	}

	//将这条数据追加到第一个位置
	newUserBills := []*ddproto.UserGameBill{b}
	newUserBills = append(newUserBills, redisUserGameBill.GetData()...)

	redisUserGameBill.Data = newUserBills

	redisUtils.SetObj(getRedisKey(b.GetUserId(), Cfg.gameId, roomType), redisUserGameBill)
	log.T("添加玩家[%v]的游戏账单数据到redis中完毕", b.GetUserId())
}

//根据玩家游戏账单数据判定是否让他赢
func IsUserShouldWin(userId uint32, roomType int32, f func(b []*ddproto.UserGameBill) bool) (bool, error) {
	log.T("开始根据玩家[%v]游戏账单数据判定是否让他赢", userId)
	bills := GetViaCache(userId, roomType)

	if bills == nil || len(bills) <= 0 {
		log.W("找不到玩家[%v]的游戏账单数据, 无法判定IsUserShouldWin", userId)
		return false, errors.New(fmt.Sprintf("找不到玩家[%v]的游戏账单数据, 无法判定", userId))
	}

	return f(bills), nil
}

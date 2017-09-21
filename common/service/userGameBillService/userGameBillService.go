package userGameBillService

import (
	"time"
	"github.com/name5566/leaf/util"
	"gopkg.in/mgo.v2/bson"
	"casino_common/common/log"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"casino_common/common/Error"
	"casino_common/utils/numUtils"
	"strings"
	"casino_common/common/consts"
	"casino_common/utils/redisUtils"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"math"
)

//玩家游戏账单数据 没玩完一局游戏存储一次数据
// 1.展示玩家输赢走势 2.下一局输赢判定策略的依据

const (
	//连输的限制局数在这两个常量间取随机数
	CONTIUNELOSECOUNT_MAX int32 = 5 //连输最多不能超过的局数
	CONTIUNELOSECOUNT_MIN int32 = 3 //连输最少不能低于的局数
)

func init() {
	UserBill = new(util.Map)
}

func OnInit(gameId, limitLength int32, watchPoint float64) {
	Cfg.gameId = gameId
	Cfg.limitLength = limitLength
	Cfg.tbName = tableName.DBT_USER_GAME_BILL
	Cfg.watchPoint = watchPoint
}

//程序启动加载时初始化配置信息
var Cfg struct {
	gameId      int32   //游戏id
	limitLength int32   //内存中数据的条数
	tbName      string  //表名
	watchPoint  float64 //是否进入退出赢模式的点数
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
func GetViaRedis(userId uint32, roomType int32) *ddproto.RedisUserGameBill {
	log.T("GetViaRedis userId%v gameId%v roomType%v tbName[%v] limitLength[%v]", userId, Cfg.gameId, roomType, Cfg.tbName, Cfg.limitLength)
	gameBills := &ddproto.RedisUserGameBill{}

	redisUtils.GetObj(getRedisKey(userId, Cfg.gameId, roomType), gameBills)
	if gameBills == nil || gameBills.GetData() == nil || len(gameBills.GetData()) <= 0 {
		log.T("没有从redis里找到玩家[%v]相关的游戏账单数据... gid[%v] roomType[%v]", userId, Cfg.gameId, roomType)
		return nil
	}

	return gameBills
}

//找到玩家账单数据 先从内存 找不到再从数据库并缓存到内存
func GetViaCache(userId uint32, roomType int32) *ddproto.RedisUserGameBill {
	log.T("开始查找玩家[%v]的游戏账单数据...", userId)
	gets := UserBill.Get(userId)
	gameBill := &ddproto.RedisUserGameBill{}
	if gets != nil {
		//从内存中获取到了 且长度符合要求 直接返回
		gameBill = gets.(*ddproto.RedisUserGameBill)
		if gameBill != nil && gameBill.GetData() != nil {
			log.T("从内存中查询到玩家[%v]的游戏账单数据, 已返回", userId)
			return gameBill
		}
	}

	//从redis中查询
	gameBill = GetViaRedis(userId, roomType)
	if gameBill.GetData() != nil && len(gameBill.GetData()) > 0 {
		//查询到了 缓存到内存里
		UserBill.Set(userId, gameBill)
		log.T("从redis中查询玩家[%v]的游戏账单数据, 缓存到内存并返回", userId)
		return gameBill
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
	gameBill := &ddproto.RedisUserGameBill{}
	if gets != nil {
		gameBill = gets.(*ddproto.RedisUserGameBill)
	}

	//将这条数据追加到第一个位置
	newUserBills := []*ddproto.UserGameBill{b}

	if gameBill.GetData() != nil && len(gameBill.GetData()) >= int(Cfg.limitLength) {
		//原始长度超出限制
		newUserBills = append(newUserBills, gameBill.Data[:Cfg.limitLength - 1]...)
	} else {
		newUserBills = append(newUserBills, gameBill.Data...)
	}

	gameBill.Data = newUserBills
	UserBill.Set(b.GetUserId(), gameBill)
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

	lenth := len(newUserBills)
	//只保留限定条数的数据
	if len(newUserBills) > int(Cfg.limitLength) {
		lenth = int(Cfg.limitLength)
	}
	redisUserGameBill.Data = newUserBills[:lenth]

	redisUtils.SetObj(getRedisKey(b.GetUserId(), Cfg.gameId, roomType), redisUserGameBill)
	log.T("添加玩家[%v]的游戏账单数据到redis中完毕", b.GetUserId())
}

func updateUserGameBill(userId uint32, roomType int32, b *ddproto.RedisUserGameBill) {
	log.T("开始更新玩家[%v]redis和内存中的游戏账单数据... bill[%v]", userId, b)
	UserBill.Set(userId, b)
	redisUtils.SetObj(getRedisKey(userId, Cfg.gameId, roomType), b)
	log.T("完成更新玩家[%v]redis和内存中的游戏账单数据... bill[%v]", userId, b)
	return
}

//获取一组玩家中谁可以赢 赢的概率是多少 0~100
func GetWinUser(roomType int32, userIds []uint32) (winUserId uint32, winRandomRate int32) {
	log.T("GetWinUser 开始获取赢的玩家 userIds[%v] roomType[%v]", userIds, roomType)
	if len(userIds) <= 0 {
		log.T("GetWinUser 传入的userIds数组为空 返回0")
		return 0, 0
	}

	//首先更新玩家输赢模式
	updateWinMode(userIds, roomType)

	//赢的概率默认70
	winRandomRate = 70

	//todo 暂时注释连输的控制
	//continueLoseUsers := getContinueLoseUsers(userIds, roomType, rand.Rand(CONTIUNELOSECOUNT_MIN, CONTIUNELOSECOUNT_MAX))
	//if len(continueLoseUsers) == 1 {
	//	//只有一个玩家连输的时候直接返回
	//	winUserId = continueLoseUsers[0]
	//	winRandomRate = 100 //必赢
	//	log.T("GetWinUser 找到单个连输的玩家[%v] 让他必赢", winUserId)
	//	return
	//}
	//
	//if len(continueLoseUsers) <= 0 {
	//	//没有连输的玩家 设置为所有玩家
	//	continueLoseUsers = userIds
	//}

	//获取位于赢模式的玩家
	//winModeUsers := getWinModeUsers(continueLoseUsers, roomType)
	winModeUsers := getWinModeUsers(userIds, roomType)
	if len(winModeUsers) == 1 {
		//只有一个玩家处于赢模式的时候直接返回
		winUserId = winModeUsers[0]
		log.T("GetWinUser 找到单个位于赢模式的玩家[%v] 让他赢", winUserId)
		return
	}

	if len(winModeUsers) > 1 {
		//有多个处于赢模式的玩家 获取其中输分最多的玩家
		winUserId = getTOPDefeatedPointUser(winModeUsers, roomType)
		log.T("GetWinUser 有多个玩家处于赢的模式[%v] 找到最高输分玩家[%v] 让他赢", winModeUsers, winUserId)
		return
	}

	log.T("GetWinUser 没有处于赢模式的玩家, 返回0")
	return 0, 0
}

//从一组玩家中获取连输的玩家
func getContinueLoseUsers(userIds []uint32, roomType int32, watchContinueLoseCount int32) (continueLoseUsers []uint32) {
	for _, userId := range userIds {
		gameBill := GetViaCache(userId, roomType)
		if gameBill == nil {
			continue
		}

		//最近连输的场次统计
		continueLoseCount := int32(0)
		for _, b := range gameBill.GetData() {
			if b.GetWinAmount() > 0 {
				//有正的得分 就不是连输 退出循环
				break
			}
			if b.GetWinAmount() == 0 {
				//过滤掉0分
				continue
			}
			//累计输的局数
			continueLoseCount++
		}
		if continueLoseCount >= watchContinueLoseCount {
			log.T("getContinueLoseUsers 玩家[%v]最近已连输[%v]把 超过限制[%v]把", userId, continueLoseCount, watchContinueLoseCount)
			continueLoseUsers = append(continueLoseUsers, userId)
			continue
		}
	}
	return
}

//从一组玩家中获取处于赢模式的玩家
func getWinModeUsers(userIds []uint32, roomType int32) (winModeUsers []uint32) {
	for _, userId := range userIds {
		gameBill := GetViaCache(userId, roomType)
		if gameBill == nil {
			continue
		}

		if gameBill.GetIsWinMode() {
			winModeUsers = append(winModeUsers, userId)
		}
	}
	return
}

//从一组玩家中获取输分最多的一个玩家
func getTOPDefeatedPointUser(userIds []uint32, roomType int32) uint32 {
	topDefeatedPointUserId := uint32(0)
	topDefeatedPoint := float64(0)
	for _, userId := range userIds {
		gameBill := GetViaCache(userId, roomType)
		if gameBill == nil {
			continue
		}
		defeatedPoint, _, _ := getUserDefeatedPoint(gameBill.GetData())
		if topDefeatedPointUserId == uint32(0) {
			topDefeatedPointUserId = userId
			topDefeatedPoint = defeatedPoint
			continue
		}
		if defeatedPoint > topDefeatedPoint {
			topDefeatedPointUserId = userId
			topDefeatedPoint = defeatedPoint
		}
	}
	return topDefeatedPointUserId
}

//更新玩家输赢模式
func updateWinMode(userIds []uint32, roomType int32) {
	for _, userId := range userIds {
		gameBill := GetViaCache(userId, roomType)
		if gameBill == nil {
			continue
		}

		wonPoint, wonCount, totalWonAmount := getUserWonPoint(gameBill.GetData())
		defeatedPoint, loseCount, totalLoseAmount := getUserDefeatedPoint(gameBill.GetData())

		log.T("判断是否更新玩家输赢模式 玩家[%v]当前是否处于赢的模式[%v] 当前defeatedPoint[%v]-loseCount[%v]-totalLoseAmount[%v] wonPoint[%v]-wonCount[%v]-totalWonAmount[%v] watchPoint[%v]", userId, gameBill.GetIsWinMode(), defeatedPoint, loseCount, totalLoseAmount, wonPoint, wonCount, totalWonAmount, Cfg.watchPoint)

		if gameBill.GetIsWinMode() {
			//玩家当前在赢的模式中
			if wonPoint >= Cfg.watchPoint {
				//玩家超过赢的得分限制 退出赢的模式
				gameBill.IsWinMode = proto.Bool(false)
				updateUserGameBill(userId, roomType, gameBill)
				log.T("更新玩家输赢模式 玩家[%v]wonPoint[%v]满足条件 退出赢的模式 wonCount[%v]-totalWonAmount[%v]", userId, wonPoint, wonCount, totalWonAmount)
				continue
			}
			continue
		}
		//玩家没有处在赢的模式中 根据输的得分判定是否进入赢的模式
		if defeatedPoint >= Cfg.watchPoint {
			//玩家超过输的得分限制 进入赢的模式
			gameBill.IsWinMode = proto.Bool(true)
			updateUserGameBill(userId, roomType, gameBill)
			log.T("更新玩家输赢模式 玩家[%v]defeatedPoint[%v]满足条件 进入赢的模式 loseCount[%v]-totalLoseAmount[%v]", userId, defeatedPoint, loseCount, totalLoseAmount)
			continue
		}
	}
}

func getUserDefeatedPoint(bills []*ddproto.UserGameBill) (defeatedPoint, loseCount, totalLoseAmount float64) {
	for _, b := range bills {
		totalLoseAmount += float64(b.GetWinAmount())
		if b.GetWinAmount() < 0 {
			loseCount++
		}
	}

	if totalLoseAmount > 0 {
		totalLoseAmount = 0
	} else {
		totalLoseAmount = -totalLoseAmount
	}

	log.T("getUserDefeatedPoint loseCount:%v totalLoseAmount:%v", loseCount, totalLoseAmount)
	defeatedPoint = getPoint(loseCount, float64(len(bills)), totalLoseAmount)
	return
}

func getUserWonPoint(bills []*ddproto.UserGameBill) (winPoint, winCount, totalWinAmount float64) {
	for _, b := range bills {
		totalWinAmount += float64(b.GetWinAmount())
		if b.GetWinAmount() > 0 {
			winCount++
		}
	}
	if totalWinAmount <= 0 {
		totalWinAmount = 0
	}

	log.T("getUserWinPoint winCount:%v totalWinAmount:%v", winCount, totalWinAmount)
	winPoint = getPoint(winCount, float64(len(bills)), totalWinAmount)
	return
}

func getPoint(targetCount, totalCount, amount float64) (point float64) {
	ratio := targetCount / totalCount * amount
	if ratio <= 1 {
		point = 0
		return
	}
	point = math.Log2(ratio)
	return point
}

package cron

import (
	"casino_common/common/Error"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/model/userDao"
	"casino_common/common/service/userGameBillService"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/dbUtils"
	"fmt"
	"github.com/name5566/leaf/timer"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//加载定时任务
func OnInit() {
	//凌晨0点
	//清理每日任务记录
	Cron("0 0 0 * * *", func() {
		//用户任务状态表
		db.C(tableName.DBT_T_USER_TASK).Drop()
		//游戏计数表
		db.C(tableName.DBT_T_GAME_DAY_COUNT).Drop()
		//同步用户金币、房卡、钻石数据redis到mongo
		dbUtils.SaveAllRedisUserToMongo(true)
	}, "任务状态清理及玩家数据同步")

	//凌晨4点 同步玩家数据
	Cron("0 0 4 * * *", func() {
		//同步用户金币、房卡、钻石数据redis到mongo
		dbUtils.SaveAllRedisUserToMongo(true)
	}, "同步玩家redis数据到mongo")

	//凌晨10min 汇总每日机器人输赢金币及库存金币
	Cron("0 1 0 * * *", AmountRobotsBillAndBalance, "汇总机器人输赢金币及库存金币")
}

//开启一个定时任务 one cron per goroutine
func Cron(expr string, cronFunc func(), alias string) {
	go func() {
		defer Error.ErrorRecovery(fmt.Sprintf("Cron设置定时任务【%v】发生异常", alias))

		log.T("开启定时【%v】expr:%v", alias, expr)
		d := timer.NewDispatcher(10)

		// cron expr 每天凌晨两点整
		cronExpr, err := timer.NewCronExpr(expr)
		if err != nil {
			log.T("NewCronExpr err:%v", err)
			return
		}

		d.CronFunc(cronExpr, func() {
			log.T("正在执行定时任务【%v】，该任务下次执行时间为%s:", alias, cronExpr.Next(time.Now()).String())
			cronFunc()
		})

		// dispatch
		for {
			(<-d.ChanTimer).Cb()
		}
	}()
}

//用户使用房卡统计
type T_daily_bill_amount struct {
	Gid                float64   `bson:"gid"`                //游戏ID
	DailyWinAmount     int64     `bson:"dailywinamount"`     //单日输赢汇总
	DailyBalanceAmount int64     `bson:"dailybalanceamount"` //单日库存余额汇总
	Day                time.Time `bson:"day"`                //统计的哪一天的数据
}

//汇总机器人输赢金币及库存金币
func AmountRobotsBillAndBalance() {
	//需要统计的游戏
	gids := []ddproto.CommonEnumGame{
		ddproto.CommonEnumGame_GID_MAHJONG,        //四川麻将
		ddproto.CommonEnumGame_GID_DDZ,            //斗地主
		ddproto.CommonEnumGame_GID_ZHUANZHUAN,     //红中转转
		ddproto.CommonEnumGame_GID_ZJH,            //扎金花
		ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN, //经典牛牛
		ddproto.CommonEnumGame_GID_BAIRENNIUNIU,   //百人牛牛
		ddproto.CommonEnumGame_GID_PDK,            //跑得快
		ddproto.CommonEnumGame_GID_MJ_HAINAN,      //海南
	}

	//输赢金币从数据库的gameBill中查询
	query := bson.M{
		"$and": []bson.M{},
	}

	//获取昨天天的0点时间作为开始时间
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local)
	utcStart := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.UTC)
	//start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	//开始时间加一天作为结束时间
	end := start.AddDate(0, 0, 1)
	log.T("获得汇总机器人输赢金币和库存的查询时间 start:%v end:%v", start, end)
	group := struct {
		WinAmount int64
	}{}
	balanceAmount := int64(0)

	query["$and"] = append(query["$and"].([]bson.M), bson.M{
		//只限昨天一整天的记录
		"endtime": bson.M{
			"$gte": start,
			"$lt":  end,
		},
		//只找机器人
		"isrobot": true,
	})
	//求和每一条输赢金币
	query_win_amount := []bson.M{
		bson.M{"$match": query},
		bson.M{"$group": bson.M{
			"_id":       nil,
			"winamount": bson.M{"$sum": "$winamount"},
		}},
	}

	for _, gid := range gids {
		//汇总输赢金币
		log.T("开始尝试汇总机器人输赢金币及库存 game[%v]", gid.String())
		//查出金币场的机器人账单记录 根据记录统计出输赢金币数量
		tbName := userGameBillService.GetTableName(int32(gid), int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN))
		err := db.Log(tbName).Pipe(query_win_amount, &group)
		if err != nil {
			log.E("汇总机器人输赢金币时err game[%v] tbName[%v] err[%v] query:%v", gid.String(), tbName, err, query_win_amount)
		}

		//汇总库存金币
		//1.先从数据库查询到所有机器人
		//2.从redis里查询金币 累加起来
		users := userDao.FindUsersByKV("robottype", int32(gid))
		for _, u := range users {
			balanceAmount += userService.GetUserCoin(u.GetId())
		}

		log.T("汇总机器人输赢金币及库存 game[%v] 机器人数量[%v] 输赢分数[%v] 金币库存[%v]", gid.String(), len(users), group.WinAmount, balanceAmount)
		if len(users) > 0 {
			data := T_daily_bill_amount{
				Gid:                float64(gid),
				DailyWinAmount:     group.WinAmount,
				DailyBalanceAmount: balanceAmount,
				Day:                utcStart.Add(time.Minute * 1),
			}
			err := db.C(tableName.DBT_ROBOT_DAILY_BILL_AMOUNT).Insert(data)
			log.T("汇总机器人输赢金币及库存 game[%v] 插入数据[%v] err[%v]", gid.String(), data, err)
		}
		group.WinAmount = 0
		balanceAmount = 0
	}
}

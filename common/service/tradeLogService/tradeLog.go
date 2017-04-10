package tradeLogService

import (
	"gopkg.in/mgo.v2/bson"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"time"
	"sync"
	"github.com/name5566/leaf/timer"
	"casino_common/common/log"
	"casino_common/common/userService"
	"errors"
)

func init() {
	//首先初始化队列
	LogList = []TTradeLogRow{}
	//开始队列插入任务
	go insertTaskStart()
}

//货币流动记录表
type TTradeLogRow struct {
	Id bson.ObjectId `bson:"_id"`
	Uid uint32  //用户
	NickName string  //昵称
	Type ddproto.HallEnumTradeType  //货币类型
	Num float64  //数量
	Time time.Time
	Msg string  //备注
}

//队列
var LogList []TTradeLogRow
//对队列插入和清空加锁
var listLock sync.Mutex

/**
    四个参数分别为：用户id、货币类型、货币增减数目、备注（备注必须填写）
 */
func Add(user_id uint32, trade_type ddproto.HallEnumTradeType, num float64, msg string) error {
	listLock.Lock()
	defer listLock.Unlock()
	user_info := userService.GetUserById(user_id)
	if user_info == nil {
		return errors.New("user nil.")
	}

	LogList = append(LogList, TTradeLogRow{
		Id:bson.NewObjectId(),
		Uid: user_id,
		NickName: user_info.GetNickName(),
		Type:trade_type,
		Num:num,
		Msg:msg,
		Time:time.Now(),
	})

	return nil
}

//批量插入日志
func insertTaskStart() {
	d := timer.NewDispatcher(10)

	// 每1分钟执行一次日志批量插入任务
	cronExpr, err := timer.NewCronExpr("*/1 * * * *")
	if err != nil {
		log.E("cronExp build err.")
		log.Printf("cronExp build err.")
		return
	}

	d.CronFunc(cronExpr, func() {
		listLock.Lock()
		//用临时数组来存，防止插入数据时的锁导致队列插入被阻塞。
		docs := LogList
		//清空队列
		LogList = []TTradeLogRow{}
		listLock.Unlock()

		log.D("日志队列批量插入任务，该任务下次执行时间为%s:", cronExpr.Next(time.Now()).String())
		log.Printf("日志队列批量插入任务，该任务下次执行时间为%s:", cronExpr.Next(time.Now()).String())
		//批量插入
		db.C(tableName.DBT_T_TRADE_LOG).InsertAll(docs)
	})

	// dispatch
	for {
		(<-d.ChanTimer).Cb()
	}

}

package tradeLogService

import (
	"gopkg.in/mgo.v2/bson"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"time"
	"casino_common/common/log"
	"fmt"
)

func init() {
	//开始队列插入任务
	go insertTaskStart()
}

//货币流动记录表
type TTradeLogRow struct {
	Id bson.ObjectId `bson:"_id"`
	Uid uint32  //用户
	Type ddproto.HallEnumTradeType  //货币类型
	Num float64  //数量
	Amount float64  //操作后剩余数量
	Time time.Time
	Msg string  //备注
}

/**
    四个参数分别为：用户id、货币类型、货币增减数目、备注（备注必须填写）
 */
func Add(user_id uint32, trade_type ddproto.HallEnumTradeType, num float64, amount float64, msg string) error {
	chan_list <- TTradeLogRow{
		Id:bson.NewObjectId(),
		Uid: user_id,
		Type:trade_type,
		Num:num,
		Amount: amount,
		Msg:msg,
		Time:time.Now(),
	}

	return nil
}

//获取时间
func GetTableName(time time.Time) string {
	return fmt.Sprintf("%s_%s", tableName.DBT_T_ACCOUNT_LOG, time.Format("2006_01_02"))
}

//插入缓冲
var chan_list = make(chan TTradeLogRow, 1024)

//批量插入日志
func insertTaskStart() {
	for new_item := range chan_list {
		time_out := make(chan bool)
		go func(ch chan bool) {
			//插入记录
			db.Log(GetTableName(time.Now())).Insert(new_item)
			ch <- true
		}(time_out)

		//超时则pass,防止阻塞后面的插入任务
		select {
		case <-time_out:
		case <-time.After(5 * time.Second):
			log.E("插入交易日志超时，%v", new_item)
		}
	}
}

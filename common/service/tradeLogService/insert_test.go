package tradeLogService

import (
	"testing"
	"casino_common/proto/ddproto"
	"sync"
	"time"
	db_init "casino_common/common/db"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"casino_common/common/sys"
)

func init() {
	db_init.InitMongoDb("192.168.199.200", 27017, "test", "id",[]string{})
	sys.InitRedis("192.168.199.200:6379","test")
}

//测试：插入记录
func TestAdd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	time.AfterFunc(10*time.Minute, func() {
		wg.Done()
	})
	Add(10058, ddproto.HallEnumTradeType_TRADE_COIN, 10, "测试-插入10金币")
	wg.Wait()
}

func TestInsert(t *testing.T) {
	Add(10058, ddproto.HallEnumTradeType_TRADE_COIN, 10, "测试-插入10金币")
	Add(10058, ddproto.HallEnumTradeType_TRADE_COIN, 11, "测试-插入11金币")

	err,num := db.C(tableName.DBT_T_TRADE_LOG).InsertAll(LogList)
	t.Log(err, num)
}

func TestQueue(t *testing.T) {
	var wg sync.WaitGroup
	chan_a := make(chan int, 1000)
	wg.Add(1)
	go func() {
		chan_a <- 1
		//chan_a <- 1
		//chan_a <- 1
		//chan_a <- 1

		wg.Add(1)
	}()

	wg.Done()

	t.Log(<-chan_a)
	wg.Done()

	wg.Wait()
}
package tradeLogService

import (
	"testing"
	"casino_common/proto/ddproto"
	"sync"
	"time"
	"casino_common/test"
)

func init() {
	test.TestInit()
}

//测试：插入记录
func TestAdd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	time.AfterFunc(10*time.Minute, func() {
		wg.Done()
	})
	Add(10955, ddproto.HallEnumTradeType_TRADE_COIN, 10, 100, "测试-插入10金币 1")
	Add(10955, ddproto.HallEnumTradeType_TRADE_COIN, 10, 110, "测试-插入10金币 2")

	wg.Wait()
}

func TestInsert(t *testing.T) {
	//Add(10058, ddproto.HallEnumTradeType_TRADE_COIN, 10, "测试-插入10金币")
	//Add(10058, ddproto.HallEnumTradeType_TRADE_COIN, 11, "测试-插入11金币")
	//
	//err,num := db.C(tableName.DBT_T_TRADE_LOG).InsertAll(LogList)
	//t.Log(err, num)
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
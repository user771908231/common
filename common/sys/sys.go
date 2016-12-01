package sys

import (
	"fmt"
	"time"
	"casino_common/utils/redis"
	"math/rand"
	"runtime"
	"casino_common/common/log"
	"casino_common/common/db"
)

func SysInit(redisAddr string, redisName string, logPath string, logName string, mongoIp string, mongoPort int, mongoName string, mongoSeqKey string, mongoSeqTables []string) error {
	InitRedis(redisAddr, redisName)
	initRandSeed()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.InitLogger(logPath, logName)        //初始化日志处理
	db.InitMongoDb(mongoIp, mongoPort, mongoName, mongoSeqKey, mongoSeqTables)
	return nil
}

//
func InitRedis(addr, name string) error {
	fmt.Println("3，初始化redis...")
	data.InitRedis(addr, name)
	return nil
}

func initRandSeed() {
	fmt.Println("5，初始化随机数种子")
	s := time.Now().UTC().UnixNano()
	rand.Seed(s)
}


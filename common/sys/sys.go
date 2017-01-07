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

var GAMEENV struct {
	PRODMODE      bool
	DEVMODE       bool
	APPLEPAY_ONLY bool
}

func SysInit(releaseTag int32, prodMode bool, redisAddr string, redisName string, logPath string, logName string, mongoIp string, mongoPort int, mongoName string, mongoSeqKey string, mongoSeqTables []string) error {
	InitRedis(redisAddr, redisName)
	initRandSeed()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.InitLogger(logPath, logName) //初始化日志处理
	db.InitMongoDb(mongoIp, mongoPort, mongoName, mongoSeqKey, mongoSeqTables)

	//初始化游戏环境变量
	GAMEENV.PRODMODE = prodMode
	GAMEENV.DEVMODE = !prodMode
	GAMEENV.APPLEPAY_ONLY = false //上线的时候设置成true
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
